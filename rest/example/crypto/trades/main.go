// Crypto - Trades
// https://polygon.io/docs/crypto/get_v3_trades__cryptoticker
// https://github.com/spinningfactory/client-go/blob/master/rest/trades.go
package main

import (
	"context"
	"log"
	"os"

	polygon "github.com/spinningfactory/client-go/rest"
	"github.com/spinningfactory/client-go/rest/models"
)

func main() {
	// init client
	c := polygon.New(os.Getenv("POLYGON_API_KEY"))

	// set params
	params := models.ListTradesParams{
		Ticker: "X:BTC-USD",
	}.WithDay(2023, 4, 13).WithLimit(50000).WithOrder(models.Asc)

	// make request
	iter := c.ListTrades(context.Background(), params)

	// do something with the result
	for iter.Next() {
		log.Print(iter.Item())
	}
	if iter.Err() != nil {
		log.Fatal(iter.Err())
	}
}
