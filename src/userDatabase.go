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
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func initDatafile() {
	//Fix directory issues
	f, err := os.Create("pw.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
}
func grabLoginFile() {
	f, err := os.Open("pw.txt")

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

func createEncryptedInfo(username string, pw string, tiker []string) string {
	stringToAdd := username
	pwHash, _ := HashPassword(pw)
	stringToAdd += pwHash
	for _, element := range tiker {
		stringToAdd += element
	}

	return stringToAdd
}

func removeUser() {

}

func logout() {

}

func addStockToWatchlist() {

}

func removeStockFromWatchlist() {

}
