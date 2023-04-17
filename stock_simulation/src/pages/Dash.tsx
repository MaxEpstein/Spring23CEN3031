import "./dash.css";
import {savedSearch} from "./Search";
import {Route, Link, Redirect} from "react-router-dom";
import React, { Component } from 'react';
import {sendMsg} from "../server";

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

export function Dash() {
  

  function searchStock() {
    savedSearch("AMZN");
    }

  const updateSaved = async(id: string) => {
      let incomming: string[] = [];

      console.log("initial Search: " + id);
      incomming = await sendMsg(id + ":1day:5min");

      console.log("Value " + incomming[0]);

      let priceMax = 0;
      let priceMin = 1000;


      let date = parseInt(incomming[0].substr(0, incomming[0].indexOf(":")));

      let priceInt = parseInt(incomming[0].substr(incomming[0].indexOf(":") + 1));
      priceInt = priceInt / 100.00;

      if (priceInt > priceMax) {
        priceMax = priceInt;
      }

      if (priceInt < priceMin) {
        priceMin = priceInt;
      }
      console.log("Price: " + priceInt);

      return priceInt;
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
            <p> AAPL: <>{updateSaved("AAPL")}</></p>

          </div>
      </>
      
    );
  }