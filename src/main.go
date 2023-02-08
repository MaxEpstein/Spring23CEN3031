package main

import (
	"fmt"
	"time"

	finance "github.com/piquette/finance-go"
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

func add_historic_data(qt *finance.Quote, temp_stock *stock) {
	//@TODO figure out pointer situaion and get maps to update accross
	//@TODO add historic data  anythin behind 01/01/1970 UTC midnight is a negative int
	p := &chart.Params{
		Symbol: temp_stock.symbol,
		Start:  &datetime.Datetime{Month: 1, Day: 1, Year: 1970},
		End: &datetime.Datetime{Month: int(time.Now().Month()),
			Day:  int(time.Now().Day()),
			Year: int(time.Now().Year())},
		Interval: datetime.OneDay, //@Todo might want to change this later
	}
	iter := chart.Get(p)

	// Iterate over results. Will exit upon any error.
	for iter.Next() {
		b := iter.Bar()
		fmt.Println(b)

		// Meta-data for the iterator - (*finance.ChartMeta).
		//fmt.Println(iter.Meta())
	}

	// Catch an error, if there was one.
	//if iter.Err() != nil {
	// Uh-oh!
	//	panic(err)
	//}
	//It has something to do with
	//func GetHistoricalQuote(symbol string, month int, day int, year int) (*finance.ChartBar, error)
}

func setup_main_working_list(s_type_name []string, s_type_sym []string) *data_list {
	s_types := make(map[string][]string)

	for i, item := range s_type_name {
		s_types[item] = append(s_types[item], s_type_sym[i])
	}

	main_working_list := new(data_list)
	main_working_list.data = make(map[string][]stock)

	for st_type, symb_arr := range s_types {
		for i, item := range symb_arr {
			//here you can add more attributes

			qt, err := quote.Get(item)
			if err != nil {
				panic(err)
			}
			//=========================
			temp_stock := new(stock)
			temp_stock.data = make(map[int64]uint)
			temp_stock.symbol = item
			temp_stock.name = qt.ShortName
			temp_stock.s_type = st_type
			add_historic_data(qt, temp_stock)
			//@TODO any additional features needed add here
			//https://piquette.io/projects/finance-go/ website for full list of things
			//========================
			//fmt.Println(temp_stock) // comment out when ready to continue
			main_working_list.data[st_type] = append(main_working_list.data[st_type], *temp_stock)
			i++ // to quite goLang

		}

	}
	return main_working_list
}

func update_data_list(working_list *data_list) {

	for x, st_type := range working_list.data {

		for i, st_symb1 := range st_type {
			qt, err := quote.Get(st_symb1.symbol)
			if err != nil {
				panic(err)
			}
			st_symb1.data[time.Now().Unix()] = uint(qt.Ask * 100)
			i++     //ignore
			x += "" //ignore
		}
	}
}
func getDataByTicker(ticker string) *stock { //take ticker input
	qt, err := quote.Get(ticker)
	if err != nil {
		panic(err)
	}
	//=========================
	temp_stock := new(stock)
	temp_stock.data = make(map[int64]uint)
	temp_stock.symbol = ticker
	temp_stock.name = qt.ShortName
	add_historic_data(qt, temp_stock)
	//@TODO any additional features needed add here
	//https://piquette.io/projects/finance-go/ website for full list of things
	//========================
	//fmt.Println(temp_stock) // comment out when ready to continue
	return temp_stock

	//temp as stock, find some way to get stock type, eft, crypto, etc
}

func addStockToMain(stockToAdd *stock, main_list *data_list) {
	main_list.data[stockToAdd.s_type] = append(main_list.data[stockToAdd.s_type], *stockToAdd)
}

func main() {

	//s_type_name := []string{"stock", "stock", "stock"} //stock, eft, crypto, etc
	//_type_sym := []string{"aapl", "aal", "intc"} //ticker
	//Way to specify what type and symbol
	//main_working_list := setup_main_working_list(s_type_name, s_type_sym)

	var s_type_name_user []string
	fmt.Scanln(&s_type_name_user)
	var s_type_sym_user []string
	fmt.Scanln(&s_type_sym_user)

	main_working_list := setup_main_working_list(s_type_name_user, s_type_sym_user)

	var singularStockUser string
	fmt.Scanln(&singularStockUser)

	addStockToMain(getDataByTicker(singularStockUser), main_working_list)

	//wrap in some kind of time based loop
	//for {
	//update_data_list(main_working_list)
	//}
	//for future frequent updates of specific stock info

	fmt.Println(main_working_list)
}
