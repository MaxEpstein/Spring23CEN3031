import { waitFor } from "@testing-library/react";
import { stringify } from "querystring";

var socket = new WebSocket("ws://localhost:8080/ws");

let connect = () => {
  console.log("Attempting Connection...");

  socket.onopen = () => {
    console.log("Successfully Connected");
  };

  socket.onmessage = msg => {
    console.log(msg);
  };

  socket.onclose = event => {
    console.log("Socket Closed Connection: ", event);
  };

  socket.onerror = error => {
    console.log("Socket Error: ", error);
  };
};

const sleep = (ms:number) => new Promise(r => setTimeout(r, ms));

let sendMsg =  async (msg: string): Promise<any> => {
  console.log("sending msg: ", msg);
  socket.send(msg);

  let price:number|null = null;

   await socket.addEventListener('message', (event) => {
    console.log("Incomming message: " + event.data);
    price = parseInt(event.data.substr(event.data.indexOf(":")+1));
    let intPrice = price/100.00;
    price = intPrice;
    console.log("Price: " + intPrice);
  });

  return new Promise((resolve, reject) => {
    setTimeout(() => {
      if (price != null) {
        resolve(price);
      } else {
        reject(null);
      }
    }, 1000);
  });
};



export { connect, sendMsg };