package main

import (
	"fmt"
	"time"
)

//go run main_worklist_funcs.go main.go unitTests.go

func unitTests() { //pass in example/testing data to various functions
	//testTypeTickerArray := []string{"stock", "stock", "stock", "stock", "stock"}
	//testTickerArray := []string{"aapl", "amzn", "intc", "nvda", "wmt"}
	//testDataList := new(data_list)
	//testDataList = testInitializeWorkingList(testTypeTickerArray, testTickerArray)

	//testStock := testGetDataByTicker("amd", "stock")

	//testDataList = testAddStockToMain(testStock, testDataList)

	//testCheckIfStockExist("aapl")
	//testCheckIfStockExist("zzzzz")
	//fmt.Println()
	testGetTimeFrame("1day")
	testGetTimeFrame("5day")
	testGetTimeFrame("1month")
	testGetTimeFrame("3month")
	testGetTimeFrame("YTD")
	testGetTimeFrame("1year")
	//testAddHistoricalData(testStock, "1year")
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

func testGetTimeFrame(timeFrame string) {
	timeFrameDate, timeInterval := getTimeFrame(timeFrame)
	fmt.Println("Day, Month, Year of starting date of requested period of time:")
	fmt.Println(timeFrameDate)
	fmt.Println("\n" + "Time interval used in creating chart data from API:")
	fmt.Println(timeInterval)
	fmt.Println()
}

func testAddHistoricalData(temp_stock *stock, timeFrame string) {
	addHistoricalData(temp_stock, timeFrame)
	fmt.Println(temp_stock.data)
	t := time.Unix(1679664600, 0)
	fmt.Println(t)
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

//go run main_worklist_funcs.go main.go unitTests.go
