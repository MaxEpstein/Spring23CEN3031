import { ReactComponentElement, useState } from 'react';

import { sendMsg } from '../server';

import "./pageStyles.css";
import React from "react";
import {
  ResponsiveContainer,
  LineChart,
  Line,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  Legend
} from "recharts"; 

let userSearched = false;
let validStock = false;
let data = [{date: "3/26/2023 15:17",price:  280.95}]
let priceMin = 1000;
let priceMax= 0;

export function Search() {
    const [message, setMessage] = useState('');
    const [prevMessage, setPrevMessage] = useState('');
    const [prevTicker, setPrevTicker] = useState('');

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setMessage(event.target.value);
  };

  const handleKeyDown = (event: React.KeyboardEvent<HTMLInputElement> ) => {
    if (event.key === 'Enter') {
      //handleClick();
    }
};

  const  handleClick = async (id:string, e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
    console.log("Id: " + id);
    data.splice(0);
    priceMax = 0;
    priceMin = 1000;

    let incomming:string[] = [];
    //send message as stock:timePeriod 
      if (id == "Search"){
        console.log("initial Search: " + message);
        incomming = await sendMsg(message+":1day:15min");
      }
      else{
        console.log("Update Graph: " + prevTicker);
        incomming = await sendMsg(prevTicker+":"+id);
      }
        
      //console.log(incomming);
      incomming.sort();
      for (let s of incomming) {
        let price = s;
        let date = parseInt(s.substr(0,s.indexOf(":")));

        let dateDate = new Date(date * 1000); // convert to current time
        let dateStr =  (dateDate.getMonth() +1) + "/" + dateDate.getDate()+ "/"  + dateDate.getFullYear() +  " " + dateDate.getHours() + ":" + dateDate.getMinutes();

        let priceInt = parseInt(s.substr(s.indexOf(":")+1));
        priceInt = priceInt/100.00;

        if (priceInt > priceMax){
          priceMax = priceInt;
        }
          
        if (priceInt < priceMin){
          priceMin = priceInt;
        }

        console.log("Price sent: " + priceInt  + "  Date: " + date);
        if (Number.isNaN(priceInt)){
          console.log("Invalid Ticker");
          setPrevMessage("Invalid Stock Ticker");
          userSearched = true;
        }
        else{
          if (message != "" || id == "Search"){
            setPrevTicker(message);
          }
          setPrevMessage(prevTicker.toUpperCase() + "- $" + priceInt);
          setMessage("");
          console.log(message.toUpperCase());

          data.push({date: dateStr, price: priceInt});
          console.log(data);

          userSearched = true;
          validStock = true;
        }
    
      }

      console.log("Min: " + priceMin + "  Max: " + priceMax);
  };

    return  (
      <>        
      <div className="SearchTop">
                <input type="text" placeholder="Stock Ticker" onChange={handleChange} value={message} name="message" id="message" onKeyDown={handleKeyDown}/>
                <button className="submit" type="submit" onClick={(e) => handleClick("Search", e)}>Search</button>
        </div>
        {userSearched === true &&
            <div className="stockInfo" >
              <h1 style={{paddingLeft: "2%"}}>Stock: {prevMessage}</h1>
              {validStock === true &&
                <div className="graph">
                    <LineChart
                  width={500}
                  height={300}
                  data={data}
                  margin={{
                    top: 5,
                    right: 30,
                    left: 20,
                    bottom: 5
                  }}
                >
                  <CartesianGrid strokeDasharray="3 3" />
                  <XAxis dataKey="date" />
                  <YAxis type="number" domain={[Math.floor(priceMin*0.98), Math.ceil((priceMax*1.02))]}/>
                  <Tooltip />
                  <Legend />
                  <Line
                    type="monotone"
                    dataKey="price"
                    stroke="#8884d8"
                    activeDot={{ r: 8 }}
                  />
            </LineChart>
            
            <button className='Graph Button' key={"1Day"} onClick={(e) => handleClick("1day:15min", e)}>1 Day</button>
            <button className='Graph Button' key={"5Day"} onClick={(e) => handleClick("5day:1hour", e)}>5 Day</button>
            <button className='Graph Button' key={"1Month"} onClick={(e) => handleClick("1month:1day", e)}>1 Month</button>
            <button className='Graph Button' key={"3Month"} onClick={(e) => handleClick("3month:1day", e)}>3 Month</button>
            <button className='Graph Button' key={"6Month"} onClick={(e) => handleClick("6month:1day", e)}>6 Month</button>
            <button className='Graph Button' key={"1Year"} onClick={(e) => handleClick("1year:1day", e)}>1 Year</button>
            <button className='Graph Button' key={"YTD"} onClick={(e) => handleClick("YTD:1day", e)}>YTD</button>
            <button className='Graph Button' key={"All"} onClick={(e) => handleClick("All:1day", e)}>All</button>
                  </div>
}
              </div>
  }

        </>
  );
}