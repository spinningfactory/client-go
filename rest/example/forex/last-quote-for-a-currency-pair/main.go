// Forex - Last Quote for a Currency Pair
// https://polygon.io/docs/forex/get_v1_last_quote_currencies__from___to
// https://github.com/dhiaayachi/client-go/blob/master/rest/quotes.go
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
	params := &models.GetLastForexQuoteParams{
		From: "AUD",
		To:   "USD",
	}

	// make request
	res, err := c.GetLastForexQuote(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}

	// do something with the result
	log.Print(res)
}
