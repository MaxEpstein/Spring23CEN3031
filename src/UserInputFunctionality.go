package main

import (
	"strconv"
	//"math"
	//"time"
)

//get string val from frontend, decide what func to run or what to do

//func searchByString(input string, main_list *data_list) string {
//	var current = main_list.data["stock"]
//
//	////search through stock map
//	//searchKeyVal, doesExist := current
//	//if doesExist {
//	//	return returnCurrentPriceString(&searchKeyVal[input])
//	//	//p
//	//	//defaultDisplay(searchKeyVal)
//	//} else {
//	//	fmt.Println("Search query does not exist within the stock database.")
//	//} //if searchKeyVal is empty, then does not exist in database
//	//return "" //empty string, does not exist
//}

func defaultDisplay(current []stock) {

}

func showFiveDay() {

}

func returnCurrentPriceString(currentStock stock) string {
	return strconv.FormatUint(uint64(currentStock.recentPrice), 10)
}

func returnCurrentPriceUInt(currentStock stock) uint {
	return currentStock.recentPrice
}
