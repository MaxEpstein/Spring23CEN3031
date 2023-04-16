package main

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"strings"
)

type user struct {
	username        string
	password        string
	watchlistStocks []string
	encryptedString string
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 15)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func createEncryptedInfo(username string, pw string) string {
	usernameHash, _ := HashPassword(username)
	pwHash, _ := HashPassword(pw)
	stringToAdd := usernameHash + ":" + pwHash + ":"
	return stringToAdd
}
func userNew(username string, pw string) *user {
	newUser := new(user)
	newUser.username = username
	newUser.password = pw
	newUser.encryptedString = createEncryptedInfo(username, pw)
	//tickers will be added else where
	return newUser
}

func addTicker(currentUser *user, ticker string) {
	currentUser.watchlistStocks = append(currentUser.watchlistStocks, ticker)
	currentUser.encryptedString += ticker + ":"
}
func addMultTicker(currentUser *user, ticker []string) {
	for _, element := range ticker {
		addTicker(currentUser, element)
	}
}

func initDatafile() {
	//Fix directory issues
	f, err := os.Create("info.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
} //
func grabLoginFile(currUser *user) {
	f, err := os.Open("info.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
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

func removeUser() {

}

func logout() {

}

func addStockToWatchlist() {

}

func removeStockFromWatchlist() {

}
