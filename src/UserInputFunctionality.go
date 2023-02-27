package main

import (
	"fmt"
	//"math"
	//"time"
)

//get string val from frontend, decide what func to run or what to do

func searchByString(input string, main_list *data_list) []stock] {
	searchKeyVal, doesExist := main_list.data["stock"]
	if doesExist {
		defaultDisplay(searchKeyVal)
		return []stock
	} else {
		fmt.Println("Search query does not exist within the stock database.")
	}
	return 
}

func defaultDisplay(current []stock) {

}

func showFiveDay() {

}

func returnCurrentPrice(current []stock) {

}
