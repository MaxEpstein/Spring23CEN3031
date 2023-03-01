package main

import (
	"fmt"
)

func unitTests() { //pass in example/testing data to various functions
	typeTickerArray := []string{"stock", "stock", "stock", "stock", "stock"}
	tickerArray := []string{"aapl", "amzn", "intc", "nvda", "wmt"}
	testMainDataStorage := initializeWorkingList(typeTickerArray, tickerArray)

	fmt.Println("Print every ticker with its type and name for testMainDataStorage below: ")
	for ticker, i := range testMainDataStorage.data {
		fmt.Println("Name: " + i[ticker].name + " -  Ticker: " + i[ticker].symbol + " -  Type: " + i[ticker].s_type)
	}

	testMainDataStorage = updateMainWorkingList(testMainDataStorage) //get latest information on each ticker inside the main working list

	fmt.Println("Print every ticker with its latest price for testMainDataStorage below: ")
	for ticker, i := range testMainDataStorage.data {
		fmt.Println("Ticker: " + i[ticker].symbol + " -  Latest Price: " + returnCurrentPriceString(i[ticker]))
	}

	//add stock

	//reprint list

	//check if exists


	//split into functions



	//fmt.Println(testMainDataStorage.data[])

}

//add function that takes in mainlist from main.go from frontend

func testmaun() {

	var s_type_container []string
	var s_type_sym_container []string

	var s_type_name_user, s_type_sym_user string
	fmt.Println("Type in type and ticker and then stop stop to exit loop")
	for s_type_name_user != "stop" {

		_, err := fmt.Scanln(&s_type_name_user, &s_type_sym_user) // take in stock type and stock ticker
		if err != nil {
			panic(err)
		}
		if s_type_name_user != "stop" {
			s_type_container = append(s_type_container, s_type_name_user) //Add it to ness list
			s_type_sym_container = append(s_type_sym_container, s_type_sym_user)
		}
	}

	main_working_list := initializeWorkingList(s_type_container, s_type_sym_container)

	fmt.Println("Enter a new type and symbol, mainly used to demo appending a new stock to main list")
	//============================Demo Purpose ======================//
	_, err := fmt.Scanln(&s_type_name_user, &s_type_sym_user)
	if err != nil {
		panic(err)
	}
	addStockToMain(getDataByTicker(s_type_sym_user, s_type_name_user), main_working_list)
	//==========================================================//

	//for {
	updateMainWorkingList(main_working_list)
	//}
	//for future frequent updates of specific stock info

	fmt.Println(main_working_list)
}

func searchTest(input string, main_list *data_list) {

}

func passTest() {

}
