// Options - Tickers
// https://polygon.io/docs/options/get_v3_reference_tickers
// https://github.com/dhiaayachi/client-go/blob/master/rest/reference.go
package main

import (
	"context"
	"log"
	"os"

	polygon "github.com/dhiaayachi/client-go/rest"
	"github.com/dhiaayachi/client-go/rest/models"
)

func main() {

	// init client
	c := polygon.New(os.Getenv("POLYGON_API_KEY"))

	// set params
	params := models.ListTickersParams{}.
		WithType("CS").
		WithMarket(models.AssetStocks).
		WithExchange("XNAS").
		WithActive(true).
		WithSort(models.TickerSymbol).
		WithOrder(models.Asc).
		WithLimit(1000)

	// make request
	iter := c.ListTickers(context.Background(), params)

	// do something with the result
	for iter.Next() {
		log.Print(iter.Item())
	}
	if iter.Err() != nil {
		log.Fatal(iter.Err())
	}

}
