import { useState } from 'react';
import "./Dash.tsx";
import { sendMsg } from '../server';
import {Login} from "./LogInPage"

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

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setMessage(event.target.value);
  };

  const handleKeyDown = (event: React.KeyboardEvent<HTMLInputElement> ) => {
    if (event.key === 'Enter') {
      handleClick();
    }
};
  const saveStock = () => {
    console.log("Saved Stock");

  };

  const  handleClick = async () => {
    data.splice(0);
    priceMax = 0;
    priceMin = 1000;

    let incomming:string[] = [];
      incomming = await sendMsg(message);
      console.log(incomming);
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

          setPrevMessage(message.toUpperCase() + "- $" + priceInt);
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
                <button className="submit" type="submit" onClick={handleClick}>Search</button>
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
                </div >

              }

          <button className="submit" type="submit" onClick={saveStock}> Save to Dashboard </button>

              <div/>
            </div>

  }
        </>
  );
}

