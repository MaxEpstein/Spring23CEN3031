package main

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/jackc/pgx/v4"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
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

func userFinder(conn *websocket.Conn, msg_cont []string) {
	//Format for incoming string LG:1:USERNAME:PW:TIK,TIK:BALANCE
	//User info gets hashed

	//msg_cont[2], _ = HashPassword(msg_cont[2])
	//msg_cont[3], _ = HashPassword(msg_cont[3])

	command := msg_cont[1]
	username := msg_cont[2]
	pw := msg_cont[3]
	tikers := msg_cont[4]
	balance := msg_cont[5]
	switch command {
	case "0": // AddUser
		addUser(strings.Join(msg_cont[2:], ":"))
	case "1": //Remove user
		removeUser(username)
	case "2": //returnUserData
		msg := returnUserData(username)
		temp := strings.Split(msg, ":")[0]
		if msg == "NIL:1" { //Wrong USername
			//do nothing
		} else if pw != temp { //Wrong PW
			msg = "NIL:0"
		} else {
			msg = strings.Join(strings.Split(msg, ":")[1:], ":")
			fmt.Println(msg)
		}
		if err := conn.WriteMessage(1, []byte(msg)); err != nil {
			log.Println(err)
			return
		} //
	case "3": //Update favorite
		temp := username + "," + tikers
		updateFavorite(temp)
	case "4":
		temp := username + "," + balance
		updateBalance(temp)
	}

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
		if msg_cont[0] == "LG" {
			//Check if message should be for the user database
			userFinder(conn, msg_cont)

		} else {
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
	//Might want to use Database URL
	dsn := "postgresql://leandro:bDoDK7mXiGb_dcN_1Mi5mg@cloned-giant-10351.7tt.cockroachlabs.cloud:26257/mindmywallet_userdata?sslmode=verify-full"
	//dsn := "postgresql://braydenb176:pf07lHgqf9HEqtJ-qYLhZg@humble-pegasus-10349.7tt.cockroachlabs.cloud:26257/mindmywallet?sslmode=verify-full"
	var err error
	conn, err = pgx.Connect(context.Background(), dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	//For testing we will erase it every time, comment out later
	deleteTable()
	//Check if table has been created, create it if not
	createTable()
	//unitTests()
	fmt.Println("Big boy app 2.0")
	setupRoutes()
	http.ListenAndServe(":8080", nil)

}
