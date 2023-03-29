package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

// We'll need to define an Upgrader
// this will require a Read and Write buffer size
var upgrader1 = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	// We'll need to check the origin of our connection
	// this will allow us to make requests from our React
	// development server to here.
	// For now, we'll do no checking and just allow any connection
	CheckOrigin: func(r *http.Request) bool { return true },
}

func reader1(conn *websocket.Conn) {
	//This defined reader will listen for the messages in the front end.
	_, p, err := conn.ReadMessage()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(p)

	if err := conn.WriteMessage(1, []byte(nil)); err != nil {
		//Return nill to front end with not found.
		log.Println(err)
		return
	}
}

// define our WebSocket endpoint
func serveWs1(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	// upgrade this connection to a WebSocket
	// connection
	ws, err := upgrader1.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	// listen indefinitely for new messages coming
	// through on our WebSocket connection

	reader1(ws)
}

func setupRoutes1() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello")
	})
	// mape our `/ws` endpoint to the `serveWs` function
	http.HandleFunc("/ws", serveWs1)

}

/*
func main() {
	//unitTests()
	fmt.Println("Server Backend")
	setupRoutes1()
	http.ListenAndServe(":8081", nil)
}
*/
