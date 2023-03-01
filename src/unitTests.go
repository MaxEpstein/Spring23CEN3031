package main

import (
	"fmt"
)

func unitTests() { //pass in example/testing data to various functions
	testTypeTickerArray := []string{"stock", "stock", "stock", "stock", "stock"}
	testTickerArray := []string{"aapl", "amzn", "intc", "nvda", "wmt"}
	testDataList := new(data_list)
	testDataList = testInitializeWorkingList(testTypeTickerArray, testTickerArray)

	testStock := testGetDataByTicker("amd", "stock")

	testDataList = testAddStockToMain(testStock, testDataList)

	testCheckIfStockExist("aapl")
	testCheckIfStockExist("zzzzz")
}

func testInitializeWorkingList(typeTickerArray []string, tickerArray []string) *data_list {
	testMainDataStorage := initializeWorkingList(typeTickerArray, tickerArray)
	//fmt.Println(testMainDataStorage)

	fmt.Println("Example input string ticker array: ")
	fmt.Println(tickerArray)

	fmt.Println("Print every stock in the test tickerArray with its name, ticker, and type below: ")
	for _, i := range testMainDataStorage.data["stock"] {
		fmt.Println("Name: " + i.name + " | Ticker: " + i.symbol + " | Type: " + i.s_type)
	}
	fmt.Println()
	return testMainDataStorage
}

func testGetDataByTicker(testTicker string, testStockType string) *stock {
	testStock := getDataByTicker(testTicker, testStockType)
	fmt.Println("Print the fetched data of the given stock: ")
	fmt.Println("Name: " + testStock.name + " | Ticker: " + testStock.symbol + " | Type: " + testStock.s_type + "\n")
	return testStock
}

func testAddStockToMain(testStock *stock, testMainDataStorage *data_list) *data_list {
	addStockToMain(testStock, testMainDataStorage)
	fmt.Println("Print every stock, including the newly added stock, in the test tickerArray with its name, ticker, and type below: ")
	for _, i := range testMainDataStorage.data["stock"] {
		fmt.Println("Name: " + i.name + " | Ticker: " + i.symbol + " | Type: " + i.s_type)
	}
	fmt.Println()
	return testMainDataStorage
}

func testCheckIfStockExist(testTicker string) {
	fmt.Println("Does the ticker " + testTicker + " exist?")
	fmt.Println(checkIfStockExist(testTicker))
}

// func testUpdateDataList(testMainDataStorage *data_list) *data_list {
// 	testMainDataStorage = updateMainWorkingList(testMainDataStorage)
// 	fmt.Println("Print every ticker with its latest price for testMainDataStorage below: ")
// 	for ticker, i := range testMainDataStorage.data {
// 		tempStock := i[ticker]
// 		fmt.Println("Ticker: " + i[ticker].symbol + " -  Latest Price: " + (string)(tempStock.recentPrice))
// 	}
// 	return testMainDataStorage
// }
