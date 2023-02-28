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

let sendMsg = (msg: string) => {
  console.log("sending msg: ", msg);
  socket.send(msg);
  socket.onmessage = (event) => {
    console.log(event.data);
    return event.data;
  }


};

export { connect, sendMsg };