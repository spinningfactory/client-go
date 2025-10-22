// Stocks - Last Quote (NBBO)
// https://polygon.io/docs/stocks/get_v2_last_nbbo__stocksticker
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
	params := &models.GetLastQuoteParams{
		Ticker: "AAPL",
	}

	// make request
	res, err := c.GetLastQuote(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}

	// do something with the result
	log.Print(res)

}
