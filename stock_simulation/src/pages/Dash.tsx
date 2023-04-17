import "./dash.css";
import {savedSearch} from "./Search";
import {Route, Link, Redirect} from "react-router-dom";
import React, { Component, useEffect } from 'react';
import {sendMsg} from "../server";
import { useState } from "react";
import { useRef } from "react";

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
import { mockComponent } from "react-dom/test-utils";

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


const delay = async (ms: number) => new Promise(
  resolve => setTimeout(resolve, ms)
)

export function Dash() {
  const dataFetchedRef = useRef(false);

   useEffect(() => {    
    if (dataFetchedRef.current) return;
    dataFetchedRef.current = true;
    updateSaved(["AAPL", "MSFT", "GOOG"]);
  })

  const [prices, setPrices] = useState<string[]>([]);

  const updateSaved = async(id: string[]) => {
    await delay(100)
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
            <h1>Saved Stocks</h1>
            <p> <>AAPL: {prices[0]}</></p>
            <p> <>MSFT: {prices[1]}</></p>
            <p> <>GOOG: {prices[2]}</></p>

          </div>
      </>
      
    );
  }