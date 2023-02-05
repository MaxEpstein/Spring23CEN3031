package main

import (
	"fmt"

	"github.com/svarlamov/goyhfin"
)

func main() {
	resp, err := goyhfin.GetTickerData("AAPL", goyhfin.OneMonth, goyhfin.OneDay, false)
	if err != nil {
		// NOTE: For library-specific errors, you can check the err against the errors exposed in goyhfin/errors.go
		fmt.Println("Error fetching Yahoo Finance data:", err)
		panic(err)
	}
	for ind := range resp.Quotes {
		fmt.Println("The day's high was", resp.Quotes[ind].Open, "on the", resp.Quotes[ind].OpensAt.Day(), "day of", resp.Quotes[ind].OpensAt.Month(), "of", resp.Quotes[ind].OpensAt.Year())
	}
}
