// Stocks - Simple Moving Average (SMA)
// https://polygon.io/docs/stocks/get_v1_indicators_sma__stockticker
// https://github.com/spinningfactory/client-go/blob/master/rest/indicators.go
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
	params := models.GetSMAParams{
		Ticker: "AAPL",
	}.WithLimit(1000)

	// make request
	res, err := c.GetSMA(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}

	// do something with the result
	log.Print(res)

}
