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
	symbol string
	name   string
	data   map[int64]uint
	s_type string
	//@TODO any additional features needed add here
}

type data_list struct {
	data map[string][]stock //For future scalability for etf's, stocks, crypto
}

func add_historic_data(temp_stock *stock) {
	//@TODO figure out pointer situaion and get maps to update accross
	p := &chart.Params{
		Symbol: temp_stock.symbol,
		Start:  &datetime.Datetime{Month: 5, Day: 2, Year: 1792},
		End: &datetime.Datetime{Month: int(time.Now().Month()),
			Day:  int(time.Now().Day()),
			Year: int(time.Now().Year())},
		Interval: datetime.OneDay, //@Todo might want to change this later
	}
	iter := chart.Get(p)

	// Iterate over results. Will exit upon any error.
	for iter.Next() { //
		b := iter.Bar()
		//RoundFloor or RoundUp
		open_price, _ := b.Open.Float64()                                               //Open Price for that day
		close_price, _ := b.Close.Float64()                                             //Close Price for that day
		temp_stock.data[int64(b.Timestamp)] = uint(math.Round(open_price * 100))        //Timestamp is for the days open  09:30:00 EST
		temp_stock.data[int64(b.Timestamp)+23400] = uint(math.Round(close_price * 100)) // Timestamp is for the days close at  16:00:00 EST
		//fmt.Println(b.Open) //b has Timestamp, Open, High, Low, Close, Volume, AdjClose
	}
}

func setup_main_working_list(s_type_name []string, s_type_sym []string) *data_list {
	s_types := make(map[string][]string)

	for i, item := range s_type_name {
		s_types[item] = append(s_types[item], s_type_sym[i])
	}

	main_working_list := new(data_list)
	main_working_list.data = make(map[string][]stock)

	for st_type, symb_arr := range s_types {
		for _, item := range symb_arr {
			////@TODO any additional features needed add here
			////https://piquette.io/projects/finance-go/ website for full list of things
			////========================

			if checkIfStockExist(item) {
				main_working_list.data[st_type] = append(main_working_list.data[st_type], *getDataByTicker(item, st_type))
			}
		}

	}
	return main_working_list
}

func update_data_list(working_list *data_list) {

	for _, st_type := range working_list.data {

		for _, st_symb1 := range st_type {
			qt, err := quote.Get(st_symb1.symbol)
			if err != nil {
				panic(err)
			}
			st_symb1.data[time.Now().Unix()] = uint(qt.Ask * 100)
		}
	}
}
func getDataByTicker(ticker string, s_type string) *stock { //take ticker input
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
	add_historic_data(temp_stock)
	//@TODO any additional features needed add here
	//https://piquette.io/projects/finance-go/ website for full list of things
	//========================
	return temp_stock

	//temp as stock, find some way to get stock type, eft, crypto, etc
}

func addStockToMain(stockToAdd *stock, main_list *data_list) {
	main_list.data[stockToAdd.s_type] = append(main_list.data[stockToAdd.s_type], *stockToAdd)
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

func main() {

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

	main_working_list := setup_main_working_list(s_type_container, s_type_sym_container)

	fmt.Println("Enter a new type and symbol, mainly used to demo appending a new stock to main list")
	//============================Demo Purpose ======================//
	_, err := fmt.Scanln(&s_type_name_user, &s_type_sym_user)
	if err != nil {
		panic(err)
	}
	addStockToMain(getDataByTicker(s_type_sym_user, s_type_name_user), main_working_list)
	//==========================================================//

	//for {
	update_data_list(main_working_list)
	//}
	//for future frequent updates of specific stock info

	fmt.Println(main_working_list)
}
