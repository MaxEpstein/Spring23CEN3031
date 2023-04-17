package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
	"strings"
)

var conn *pgx.Conn

func main() {
	//Might want to use Database URL
	dsn := "postgresql://leandro:bDoDK7mXiGb_dcN_1Mi5mg@cloned-giant-10351.7tt.cockroachlabs.cloud:26257/mindmywallet_userdata?sslmode=verify-full"
	var err error
	conn, err = pgx.Connect(context.Background(), dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	deleteTable()
	createTable()
	addUser("leo023,0233454,aapl:aal,154.45")
	findUser("aapl:aal")

}
func createTable() {
	_, err := conn.Exec(context.Background(),
		"CREATE TABLE IF NOT EXISTS userData (id STRING NOT NULL PRIMARY KEY, pwd STRING NOT NULL,  tiker STRING, balance STRING)")
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
func addUser(userInfo string) {
	//Username,PW,tiker:tiker,balance
	userInfo_cont := strings.Split(userInfo, ",")
	// Create
	_, err := conn.Exec(context.Background(),
		"INSERT INTO userData (id,pwd,tiker,balance) VALUES ('Leandro','vnrnr','aapl:aal','userInfo_cont[3]')")
	if err != nil {
		panic(err)
	}
	fmt.Println(userInfo_cont)
}
func findUser(id_s string) {
	var someText string
	row := conn.QueryRow(context.Background(),
		"SELECT * FROM userData WHERE id = id_s LIMIT 1")
	if err := row.Scan(&someText); err != nil {
		panic(err)
	}
	fmt.Println(someText)
}
