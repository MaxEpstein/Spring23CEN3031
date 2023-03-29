package main

import (
	"fmt"
	"math"
	"time"

	chart "github.com/piquette/finance-go/chart"
	"github.com/piquette/finance-go/datetime"
	"github.com/piquette/finance-go/quote"
)

type stock struct {
	symbol      string
	name        string
	data        map[int64]uint
	s_type      string
	recentPrice uint
	//@TODO any additional features needed add here
}

type data_list struct {
	data map[string]map[string]stock //For future scalability for etf's, stocks, crypto
}

func initializeWorkingList(s_type_name []string, s_type_sym []string, data_interval string, data_time_interval string) *data_list {
	main_working_list := new(data_list)
	main_working_list.data = make(map[string]map[string]stock)
	main_working_list.data["stock"] = make(map[string]stock)

	//data_time_interval = "15min" //remove in future

	if data_interval == "" || data_time_interval == "" {
		return main_working_list
	}
	for _, item := range s_type_sym {
		////@TODO any additional features needed add here
		////https://piquette.io/projects/finance-go/ website for full list of things
		////========================
		if checkIfStockExist(item) {
			main_working_list.data["stock"][item] = *getDataByTicker(item, "stock", data_interval, data_time_interval)
		}
	}

	return main_working_list
}

func addHistoricalData(temp_stock *stock, timeFrame string, chartInterval string) {
	timeFrameDate, timeInterval := getTimeFrame(timeFrame, chartInterval)
	p := &chart.Params{
		Symbol: temp_stock.symbol,

		Start: timeFrameDate,
		End: &datetime.Datetime{Month: int(time.Now().Month()),
			Day:  int(time.Now().Day() + 1),
			Year: int(time.Now().Year())},
		Interval: timeInterval, //@Todo might want to change this later
	}
	iter := chart.Get(p)

	// Iterate over results. Will exit upon any error.
	for iter.Next() { //
		b := iter.Bar()
		//RoundFloor or RoundUp
		open_price, _ := b.Open.Float64() //Open Price for that day
		//close_price, _ := b.Close.Float64()
		//Close Price for that day
		if open_price != 0 {
			temp_stock.data[int64(b.Timestamp)] = uint(math.Round(open_price * 100)) //Timestamp is for the days open  09:30:00 EST
		}
		//temp_stock.data[int64(b.Timestamp)+23400] = uint(math.Round(close_price * 100)) // Timestamp is for the days close at  16:00:00 EST
		//fmt.Println(b.Open) //b has Timestamp, Open, High, Low, Close, Volume, AdjClose
	}

}

func getTimeFrame(timeFrame string, chartIntervalString string) (*datetime.Datetime, datetime.Interval) {
	adjustedTime := time.Now()
	chartInterval := getChartInterval(chartIntervalString)
	switch choose := timeFrame; choose {
	case "1day":
		adjustedTime = passWeekends(1)
	case "5day":
		adjustedTime = passWeekends(5)
	case "1month":
		adjustedTime = skipWeekends(time.Now().AddDate(0, -1, 0))
	case "3month":
		adjustedTime = skipWeekends(time.Now().AddDate(0, -3, 0))
	case "6month":
		adjustedTime = skipWeekends(time.Now().AddDate(0, -6, 0))
	case "YTD":
		return &datetime.Datetime{Month: 1, Day: 1, Year: time.Now().Year()}, chartInterval
	case "1year":
		adjustedTime = skipWeekends(time.Now().AddDate(-1, 0, 0))
	}
	return &datetime.Datetime{Month: (int)(adjustedTime.Month()), Day: adjustedTime.Day(), Year: adjustedTime.Year()}, chartInterval
}

func passWeekends(numDays int) time.Time {
	i := 0
	adjustedTime := time.Now()
	for i < numDays {
		adjustedTime = adjustedTime.AddDate(0, 0, -1)
		if adjustedTime.Weekday() != 0 && adjustedTime.Weekday() != 6 {
			i++
			fmt.Println("hello")
		}
	}
	return adjustedTime
}

func skipWeekends(adjustedTime time.Time) time.Time {
	for adjustedTime.Weekday() == 0 || adjustedTime.Weekday() == 6 {
		adjustedTime = adjustedTime.AddDate(0, 0, 1)
	}
	return adjustedTime
}

func getChartInterval(chartIntervalString string) datetime.Interval {
	chartInterval := datetime.OneMin
	switch choose := chartIntervalString; choose {
	case "1min":
		chartInterval = datetime.OneMin
	case "5min":
		chartInterval = datetime.FiveMins
	case "15min":
		chartInterval = datetime.FifteenMins
	case "1hour":
		chartInterval = datetime.OneHour
	case "1day":
		chartInterval = datetime.OneDay
	case "1month":
		chartInterval = datetime.OneMonth
	case "3month":
		chartInterval = datetime.ThreeMonth
	case "1year":
		chartInterval = datetime.OneYear
	}
	return chartInterval
}

func getDataByTicker(ticker string, s_type string, data_interval string, data_time_interval string) *stock { //take ticker input
	//data_interval = Amount of info needed, 1 day etc
	//data_time_interval = Time intervals, 15 mins etc
	qt, err := quote.Get(ticker)
	if err != nil {
		panic(err)
	}
	//=========================
	temp_stock := new(stock)
	temp_stock.data = make(map[int64]uint)
	temp_stock.symbol = ticker
	temp_stock.name = qt.ShortName
	temp_stock.s_type = s_type
	addHistoricalData(temp_stock, data_interval, data_time_interval)
	//fmt.Println(len(temp_stock.data))
	return temp_stock

	//@TODO any additional features needed add here
	//https://piquette.io/projects/finance-go/ website for full list of things
	//========================

}

// get latest information for tickers within workinglist
func updateMainWorkingList(working_list *data_list) {
	for _, st_type := range working_list.data {

		for _, st_symb1 := range st_type {
			qt, err := quote.Get(st_symb1.symbol)
			if err != nil {
				panic(err)
			}
			st_symb1.data[time.Now().Unix()] = uint(qt.Ask * 100)
			st_symb1.recentPrice = uint(qt.Ask)
		}
	}
}

func addStockToMain(stockToAdd *stock, main_list *data_list) {
	main_list.data[stockToAdd.s_type][stockToAdd.symbol] = *stockToAdd
}

func checkIfStockExist(ticker string) bool {
	qt, err := quote.Get(ticker)
	if err != nil {
		panic(err)
	} else if qt != nil {
		return true
	}
	return false
}

/*
func mainForWorklistFuncs() { //used for testing various functions

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
	addStockToMain(getDataByTicker(s_type_sym_user, s_type_name_user, ), main_working_list)
	//==========================================================//

	//for {
	//updateMainWorkingList(main_working_list)
	//}
	//for future frequent updates of specific stock info

	fmt.Println(main_working_list)
}
*/
