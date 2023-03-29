package main

import (
	"strings"
)

type user struct {
	username        string
	password        string
	watchlistStocks []string
}

func initializeUserDatabase() {
	userDatabase := make(map[string]user)
	for k := range userDatabase {
		delete(userDatabase, k)
	}
}

func login(loginInfo string) {
	loginInfo = strings.TrimSpace(loginInfo)            //get rid of leading and trailing spaces
	loginInfoSeperated := strings.Split(loginInfo, ",") //replace with delimiter being used
	username, password := loginInfoSeperated[0], loginInfoSeperated[1]
	getAllUsers(username, password)
}

func getAllUsers(username string, password string) {

}

func addUser() {

}

func removeUser() {

}

func logout() {

}

func addStockToWatchlist() {

}

func removeStockFromWatchlist() {

}
