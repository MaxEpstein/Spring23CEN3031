package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/websocket"
) //

// We'll need to define an Upgrader
// this will require a Read and Write buffer size
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	// We'll need to check the origin of our connection
	// this will allow us to make requests from our React
	// development server to here.
	// For now, we'll do no checking and just allow any connection
	CheckOrigin: func(r *http.Request) bool { return true },
}

func reader(conn *websocket.Conn) {
	//This defined reader will listen for the messages in the front end.
	for {
		// read in a message
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		ticker := string(p)
		//Expected message ticker:interval:time_interval
		msg_cont := strings.Split(ticker, ":")
		//Check if the stock being submitted in is real, otherwise continue listening for an input
		if checkIfStockExist(msg_cont[0]) != true {
			if err := conn.WriteMessage(1, []byte(nil)); err != nil {
				//Return nill to front end with not found.
				log.Println(err)
				return
			}
			continue
		}
		if msg_cont[1] == "now" {
			msgCurrentPrice := strconv.FormatUint(uint64(getCurrentPrice(msg_cont[0])), 10)
			if err := conn.WriteMessage(1, []byte(msgCurrentPrice)); err != nil {
				log.Println(err)
				return
			}
		} else {
			main_list := initializeWorkingList(nil, nil, "", "")
			addStockToMain(getDataByTicker(msg_cont[0], "stock", msg_cont[1], msg_cont[2]), main_list)
			//updateMainWorkingList(main_list) //take away later
			temp_stock := main_list.data["stock"][msg_cont[0]]
			msg := ""
			for key, element := range temp_stock.data {
				//Send all the data within the current map
				msg = strconv.FormatUint(uint64(key), 10) + ":" + strconv.FormatUint(uint64(element), 10)
				if err := conn.WriteMessage(1, []byte(msg)); err != nil {
					log.Println(err)
					return
				}
			}
		}

	}
}

// define our WebSocket endpoint
func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	// upgrade this connection to a WebSocket
	// connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	// listen indefinitely for new messages coming
	// through on our WebSocket connection

	reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello")
	})
	// mape our `/ws` endpoint to the `serveWs` function
	http.HandleFunc("/ws", serveWs)

}

func main() {
	unitTests()
	fmt.Println("Big boy app 2.0")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
