package main

import (
	"flag"
	"fmt"
	"github.com/piquette/finance-go/quote"
)

func main() {
	flag.Parse()

	var smbls []string
	smbls = append(smbls, "aapl")
	//smbls = append(smbls, "aapl")

	stock_data := make(map[string][]uint)
	sum := 1
	for sum < 100 {
		sum += 1
		iter := quote.List(smbls) // This creates a list that unless
		// you keep creating will return false and only run once
		for iter.Next() {
			q := iter.Quote()
			stock_data[q.ShortName] = append(stock_data[q.ShortName], uint(q.Ask*100))

		}
	}
	fmt.Println(stock_data)
}
