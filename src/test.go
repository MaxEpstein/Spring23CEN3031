package main

import (
	"context"
	"fmt"
	"strings"

	_ "database/sql"

	"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
)

var conn *pgx.Conn

/*
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

		deleteTable()
		createTable()
		addUser("leo023,0233454,aapl:aal,154.45")
		returnUserData("leo023")
		updateFavorite("leo023,aapl")
		updateBalance("leo023,140")
		returnUserData("leo023")
		addUser("bry,123456,amd,100")
		returnUserData("bry")
	}
*/
func createTable() {
	_, err := conn.Exec(context.Background(),
		"CREATE TABLE IF NOT EXISTS userData (Username STRING NOT NULL UNIQUE PRIMARY KEY, Password STRING NOT NULL, Favorites STRING, Balance STRING)")
	if err != nil {
		panic(err)
	}
}
func deleteTable() {
	_, err := conn.Exec(context.Background(),
		"DROP TABLE IF EXISTS userData")
	if err != nil {
		panic(err)
	}
}

func addUser(userData string) {
	//Username,Password,Favorites,Balance
	// 	           Ticker:Ticker:Ticker...
	userInfo := strings.Split(userData, ",")
	add := "INSERT INTO userData (Username,Password,Favorites,Balance) VALUES ($1,$2,$3,$4)"
	_, err := conn.Exec(context.Background(), add, userInfo[0], userInfo[1], userInfo[2], userInfo[3])
	if err != nil {
		panic(err)
	}
}

func returnUserData(inputUsername string) { //turn into return for both vars
	var Favorites string
	var Balance string
	query := "SELECT Favorites, Balance FROM userData WHERE Username = $1"
	row := conn.QueryRow(context.Background(), query, inputUsername)
	switch err := row.Scan(&Favorites, &Balance); err {
	case pgx.ErrNoRows:
		fmt.Println("Error: No rows")
		panic(err)
	case nil:
		fmt.Println(Favorites, Balance)
	default:
		panic(err)
	}
}

func updateFavorite(userData string) { //pass in new string with removed or added tickers
	userInfo := strings.Split(userData, ",")
	update := "UPDATE userData SET Username = $1, Favorites = $2"
	_, err := conn.Exec(context.Background(), update, userInfo[0], userInfo[1])
	if err != nil {
		panic(err)
	}
}

func updateBalance(userData string) { //pass in new string with removed or added tickers
	userInfo := strings.Split(userData, ",")
	update := "UPDATE userData SET Username = $1, Balance = $2"
	_, err := conn.Exec(context.Background(), update, userInfo[0], userInfo[1])
	if err != nil {
		panic(err)
	}
}
