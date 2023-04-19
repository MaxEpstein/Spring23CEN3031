import "./dash.css";
import {savedSearch} from "./Search";
import {Route, Link, Redirect} from "react-router-dom";
import React, { Component } from 'react';
import {sendMsg} from "../server";
import { useState } from "react";
import { useRef } from "react";
import {useEffect} from "react";
import moment from "moment";


import {
  LineChart,
  Line,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  Legend
} from "recharts";
import {render} from "react-dom";

const data = [
  {
    name: "Page A",
    uv: 4000,
    pv: 2400,
    amt: 2400
  },
  {
    name: "Page B",
    uv: 3000,
    pv: 1398,
    amt: 2210
  },
  {
    name: "Page C",
    uv: 2000,
    pv: 9800,
    amt: 2290
  },
  {
    name: "Page D",
    uv: 2780,
    pv: 3908,
    amt: 2000
  },
  {
    name: "Page E",
    uv: 1890,
    pv: 4800,
    amt: 2181
  },
  {
    name: "Page F",
    uv: 2390,
    pv: 3800,
    amt: 2500
  },
  {
    name: "Page G",
    uv: 3490,
    pv: 4300,
    amt: 2100
  }
];


let pricesArr:string[] = [];
var format = 'hh:mm:ss';
let loggedIn: boolean = true;


const delay = async (ms: number) => new Promise(
  resolve => setTimeout(resolve, ms)
)



export function Dash() {
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
    if (loggedIn == true)
      updateSaved(["AAPL", "MSFT", "GOOG", "AAL", "META", "TSLA", "RCL"]);
  }

   useEffect(() => {    
    if (dataFetchedRef.current) return;

    else
    {
      logged();
      dataFetchedRef.current = true;
    }
    if (moment().isBetween(moment('9:30:00',format), moment('16:00:00', format))){
        const interval = setInterval(() => {
        updateSaved(["AAPL", "MSFT", "GOOG", "AAL", "META", "TSLA", "RCL"]);
      }, 10000);

      return ()=> clearInterval(interval);
  }
      
  })

  const [prices, setPrices] = useState<string[]>([]);

  const updateSaved = async(id: string[]) => {
    await delay(0)
    pricesArr = [];
    for (let tick of id){
        let incomming: string[] = [];

        console.log("initial Search: " + tick);
        incomming = await sendMsg(tick + ":now");


        let priceInt = parseInt(incomming[0].substr(incomming[0].indexOf(":") + 1));
        priceInt = priceInt / 100.00;
        
        console.log("Price: " + priceInt);
        pricesArr[pricesArr.length] = String(priceInt);

      }
      setPrices(pricesArr);
    };

    return(
        <>
        {!loggedIn ? (
           <Redirect to={"/login"} ></Redirect>
) : (
  <>
        <div className="heading">
            <h1>Dashboard</h1>
          <div className="stats">
            <ul>
              <li>Cash:</li>
              <li>YTD:</li>
              <li># of Stocks:</li>
            </ul>
          </div>
          
        </div>

        
        <div className="dashGraph">
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
                  <XAxis dataKey="name" />
                  <YAxis />
                  <Tooltip />
                  <Legend />
                  <Line
                    type="monotone"
                    dataKey="pv"
                    stroke="#8884d8"
                    activeDot={{ r: 8 }}
                  />
                  <Line type="monotone" dataKey="uv" stroke="#82ca9d" />
            </LineChart>
                </div>

          <div className = "savedStocks">
            <div className="SavedHeader">
              <h1>Saved Stocks</h1>
              <button className="Graph_button">Edit</button>
            </div>

            <p> <>AAPL: {prices[0]}</></p>
            <p> <>MSFT: {prices[1]}</></p>
            <p> <>GOOG: {prices[2]}</></p>
            <p> <>AAL: {prices[3]}</></p>
            <p> <>META: {prices[4]}</></p>
            <p> <>TSLA: {prices[5]}</></p>
            <p> <>RCL: {prices[6]}</></p>


          </div>
          </>
          )
            }
      </>
      
    );
  }