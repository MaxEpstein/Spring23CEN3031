//import "./Dash.tsx";
import { ReactComponentElement, useState } from 'react';
import {Redirect} from 'react-router-dom';
import { useRef } from 'react';
import { sendMsg } from '../server';
import {Login} from "./LogInPage"
import { useEffect } from 'react';
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
let data = [{date: 9999,price:  999.99}]
let newData = [{date: "99/99/9999 15:17",price:  999.99}]
let priceMin = 1000;
let priceMax= 0;
let dashTicker: string;
let loggedIn = false;

export function savedSearch(name: string) {
  window.location.replace('/search');
  console.log(name);
  dashTicker = name;
  userSearched = true;
}

const delay = async (ms: number) => new Promise(
  resolve => setTimeout(resolve, ms)
)

export function Search() {
  console.log("User Redirect? " + userSearched + "  dashTicker: " +dashTicker);
  const [message, setMessage] = useState('');
  const [prevMessage, setPrevMessage] = useState('');
  const [prevTicker, setPrevTicker] = useState('');  
  const[forceRefresh, setForceRefresh] = useState('');
  const dataFetchedRef = useRef(false);

  const logged = async() => {
    await delay(100);
    let incomming = await sendMsg("LOG");
    console.log("Incomming about log " + incomming);
  
    if (incomming == 1){
      loggedIn = true;
    }
    else if (incomming == 0){ 
      loggedIn = false;
    }
  
    console.log("Loggedin: " + loggedIn);
  }

  useEffect(() => {    
    if (dataFetchedRef.current) return;

    else
    {
      logged();
      dataFetchedRef.current = true;
    }
}) ;


  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setMessage(event.target.value);
  };

  const handleKeyDown = (event: React.KeyboardEvent<HTMLInputElement>) => {
    if (event.key === 'Enter') {
      handleClick("Search");
    }
  };
  const saveStock = async (ticker: string)  => {
    console.log("Saved Stock: " + ticker);
    let incomming = await sendMsg("LG:3:::"+ticker+":");
    console.log("Successfully saved?: " + incomming);
  };


  const compareNumbers = (a:{date: number, price: number}, b:{date: number, price:number}):number => {
    return a.date - b.date;
  }

  const handleClick = async (id: string) => {

    //console.log("Id: " + id);
    data.splice(0);
    newData.splice(0);
    console.log("Cleared Data; " + newData);
    priceMax = 0;
    priceMin = 1000;

    let incomming: string[] = [];
    //send message as stock:timePeriod 
    if (id == "Search") {
      console.log("initial Search: " + message);
      incomming = await sendMsg(message + ":1day:5min");
    } else {
      console.log("Update Graph: " + prevTicker);
      incomming = await sendMsg(prevTicker + ":" + id);
    }

    incomming.sort();
    for (let s of incomming) {
      let date = parseInt(s.substr(0, s.indexOf(":")));

      let priceInt = parseInt(s.substr(s.indexOf(":") + 1));
      priceInt = priceInt / 100.00;

      if (priceInt > priceMax) {
        priceMax = priceInt;
      }

      if (priceInt < priceMin) {
        priceMin = priceInt;
      }


      //console.log("Price sent: " + priceInt  + "  Date: " + date);
      

        setMessage("");
        //console.log(message.toUpperCase());

        data.push({date: date, price: priceInt});
        //console.log(data);

        userSearched = true;
        validStock = true;
      }

    data.sort(compareNumbers);
    

    for (let i of data){
      let dateDate = new Date(i.date * 1000); // convert to current time
      let dateStr = (dateDate.getMonth() + 1) + "/" + dateDate.getDate() + "/" + dateDate.getFullYear() + " " + dateDate.getHours() + ":" + dateDate.getMinutes();
      newData.push({date: dateStr, price: i.price});
    }

    if (Number.isNaN(newData[newData.length-1].price)) {
      console.log("Invalid Ticker");
      setPrevMessage("Invalid Stock Ticker");
      userSearched = true;
    } 
    else{
      if (message != "" || id == "Search") {
        setPrevTicker(message);
        setPrevMessage(message.toUpperCase() + "- $" + newData[newData.length-1].price);
      } 
      else{
          incomming = await sendMsg(prevTicker + ":now");

          let priceInt = parseInt(incomming[0].substr(incomming[0].indexOf(":") + 1));
          priceInt = priceInt / 100.00;
          
          console.log("Message: " + prevTicker +" Price: " + priceInt);

          setPrevMessage(prevTicker.toUpperCase() + "- $" + priceInt);
        }
      }

    console.log(newData);
    
    console.log("Min: " + priceMin + "  Max: " + priceMax);
    setForceRefresh(String(newData.length));
  };

  return (
      <>
        <div className="SearchTop" style={{paddingTop: "2%"}}>
          <input type="text" placeholder="Stock Ticker" onChange={handleChange} value={message} name="message"
                 id="message" onKeyDown={handleKeyDown} />
          <button className="submit" type="submit" onClick={(e) => handleClick("Search")}>Search</button>
        </div>
        {userSearched === true &&
            <div className="stockInfo">
              <h1 style={{paddingLeft: "2%", scrollPaddingBottom: "2%", textAlign: "center"}}>Stock: {prevMessage}</h1>
              {validStock === true &&
                  <div>
                    <div className="graph" style={{paddingTop: "2%", marginTop : "7%"}}>
                        <LineChart
                            width={500}
                            height={300}
                            data={newData}
                            margin={{
                              top: 5,
                              right: 30,
                              left: 20,
                              bottom: 5
                            }}
                            key={newData.length}
                        >
                          <CartesianGrid strokeDasharray="3 3"/>
                          <XAxis dataKey="date" allowDataOverflow={false}/>
                          <YAxis type="number" domain={[Math.floor(priceMin * 0.98), Math.ceil((priceMax * 1.02))]}/>
                          <Tooltip/>
                          <Legend/>
                          <Line
                              type="monotone"
                              dataKey="price"
                              stroke="#8884d8"
                              dot = {false}
                              activeDot={{r: 4}}
                              key={newData.length}
                          />
                        </LineChart>

                    <button className='Graph_button' key={"1Day"} onClick={(e) => handleClick("1day:15min")} >1 Day</button>
                    <button className='Graph_button' key={"5Day"} onClick={(e) => handleClick("5day:1hour")} >5 Day</button>
                    <button className='Graph_button' key={"1Month"} onClick={(e) => handleClick("1month:1day")} >1 Month</button>
                    <button className='Graph_button' key={"3Month"} onClick={(e) => handleClick("3month:1day")} >3 Month</button>
                    <button className='Graph_button' key={"6Month"} onClick={(e) => handleClick("6month:1day")} >6 Month</button>
                    <button className='Graph_button' key={"1Year"} onClick={(e) => handleClick("1year:1day")} >1 Year</button>
                    <button className='Graph_button' key={"YTD"} onClick={(e) => handleClick("YTD:1day")} >YTD</button>
                    <button className='Graph_button' key={"All"} onClick={(e) => handleClick("all:1month")} >All</button>

                    {!loggedIn ? (
                      <></>
                    ) : (
                      <button className="submit" type="submit" onClick={(e) => saveStock(prevTicker)}> Save to Dashboard </button>
                    )}
                    </div>  
                  </div>
              }
            </div>
        }
      </>

  );
}



