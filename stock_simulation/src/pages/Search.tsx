import { useState } from 'react';
import { FormatCodeSettings } from 'typescript';

import { sendMsg } from '../server';

import "./pageStyles.css";
import React from "react";
import {
  LineChart,
  Line,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  Legend
} from "recharts";

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


let userSearched = false;

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




  const handleClick = () => {
    let map = sendMsg(message);
    if (map == null){
      setPrevMessage("Invalid Stock Ticker");
    }
    else{
      setPrevMessage(message.toUpperCase() + "- $" + map);
      setMessage("");
      console.log(message.toUpperCase());

      userSearched = true;
    }
    
    
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
              </div>
}
        </>
  );
}