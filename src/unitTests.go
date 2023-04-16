package main

import (
	"fmt"
	"time"
)

//go run worklistFunctions.go main.go unitTests.go

func unitTests() { //pass in example/testing data to various functions
	fmt.Println("Sprint 1 and 2:\n")

	testTypeTickerArray := []string{"stock", "stock", "stock", "stock", "stock"}
	testTickerArray := []string{"aapl", "amzn", "intc", "nvda", "wmt"}
	testDataList := new(data_list)
	testDataList = testInitializeWorkingList(testTypeTickerArray, testTickerArray)

	testStock := testGetDataByTicker("amd", "stock")

	testDataList = testAddStockToMain(testStock, testDataList)

	testCheckIfStockExist("aapl")
	testCheckIfStockExist("zzzzz")
	fmt.Println()

	fmt.Println("Sprint 3:\n")

	testPassWeekends(5, "5")
	testSkipWeekends("1year")

	testGetTimeFrame("1day", "1min")

	//testGetTimeFrame("5day", "5min")
	//testGetTimeFrame("1month", "15min")
	//testGetTimeFrame("3month", "1day")
	//testGetTimeFrame("YTD", "1month")
	//testGetTimeFrame("1year", "1year")
  testGetTimeFrame("1month", "1day")

  
  testAddHistoricalData(testStock, "1day", "1hour")
	//testAddHistoricalData(testStock, "1year", "1year")
	//testAddHistoricalData(testStock, "1year")
	passwordHashing("helloWorld")
}

func testInitializeWorkingList(typeTickerArray []string, tickerArray []string) *data_list {
	testMainDataStorage := initializeWorkingList(typeTickerArray, tickerArray, "1day", "1min")
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
	testStock := getDataByTicker(testTicker, testStockType, "1day", "1min")
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
	fmt.Println("Does the ticker " + testTicker + " exist? True or False.")
	fmt.Println(checkIfStockExist(testTicker))
}

func testGetTimeFrame(timeFrame string, chartInterval string) {
	timeFrameDate, timeInterval := getTimeFrame(timeFrame, chartInterval)
	fmt.Println("Date of requested period of time: " + timeFrame)
	fmt.Println(timeFrameDate.Time().Date())
	fmt.Println("Time interval used in creating chart data from API:")
	fmt.Println(timeInterval)
	fmt.Println()
}

func testPassWeekends(numDays int, numdays string) {
	fmt.Println("For " + numdays + " days: Starting date to be used only counting weekdays, passing over weekends. For 5 day or smaller.")
	fmt.Println(passWeekends(numDays).Date())
	fmt.Println()
}

func testSkipWeekends(skip string) {
	fmt.Println("For " + skip + ": Starting date to be used only counting weekdays, skipping over weekends. For 1 month or larger.")
	fmt.Println(skipWeekends(time.Now().AddDate(-1, 0, 0)).Date())
	fmt.Println()
}

func testAddHistoricalData(temp_stock *stock, timeFrame string, chartInterval string) {
	addHistoricalData(temp_stock, timeFrame, chartInterval)
	fmt.Println("Datapoints for " + temp_stock.name + ":")
	fmt.Println("Timeframe: " + timeFrame + "\nChart Interval: " + chartInterval)
	fmt.Println(temp_stock.data)
	fmt.Println()
	//t := time.Unix(1679664600, 0)
	//fmt.Println(t)
}
func passwordHashing(password string) {
	hash, _ := HashPassword(password) // ignore error for the sake of simplicity

	fmt.Println("Password:", password)
	fmt.Println("Hash:    ", hash)

	match := CheckPasswordHash(password, hash)
	fmt.Println("Match:   ", match)
}
