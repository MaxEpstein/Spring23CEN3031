package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

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

func reader(conn *websocket.Conn, main_list *data_list) {
	//This defined reader will listen for the messages in the front end.
	for {
		// read in a message
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		ticker := string(p)
		//Check if the stock being submitted in is real, otherwise continue listening for an input
		if checkIfStockExist(ticker) != true {
			if err := conn.WriteMessage(1, []byte(nil)); err != nil {
				//Return nill to front end with not found.
				log.Println(err)
				return
			}
			continue
		}
		addStockToMain(getDataByTicker(ticker, "stock"), main_list)
		update_data_list(main_list) //take away later
		temp_stock := main_list.data["stock"][ticker]
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

//for {
//	// read in a message
//	_, p, err := conn.ReadMessage()
//
//	if err != nil {
//		log.Println(err)
//		return
//	}
//
//	r := regexp.MustCompile("[^\\s]+")
//	inputArray := r.FindAllString(string(p), -1)
//	inputType := inputArray[0]

//var current []stock
//
//switch choose := inputType; choose { //depending on button/passed in input type, do required functionality
//case "search":
//	current = searchByString(inputArray[1], &data_list{})
//case "display": // "display 1day", ie "inputType, button input function"
//	switch displayChoose := inputArray[1]; displayChoose {
//	case "1day":
//
//	case "5day":
//
//	case "10day":
//
//	case "1month": //etc
//
//	}
//case "home": //return to home page?
//
//case "other": //other button functionality

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
	main_working_list := initializeWorkingList(nil, nil)

	reader(ws, main_working_list)
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
	fmt.Println("Chat App v0.01")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
