// Stocks - Quotes (NBBO)
// https://polygon.io/docs/stocks/get_v3_quotes__stockticker
// https://github.com/spinningfactory/client-go/blob/master/rest/quotes.go
package main

import (
	"context"
	"log"
	"os"
	"time"

	polygon "github.com/spinningfactory/client-go/rest"
	"github.com/spinningfactory/client-go/rest/models"
)

func main() {

	// init client
	c := polygon.New(os.Getenv("POLYGON_API_KEY"))

	// set params
	params := models.ListQuotesParams{
		Ticker: "AAPL",
	}.WithTimestamp(models.EQ, models.Nanos(time.Date(2021, 7, 22, 0, 0, 0, 0, time.UTC))).
		WithSort(models.Timestamp).
		WithOrder(models.Asc).
		WithLimit(50000)

	// make request
	iter := c.ListQuotes(context.Background(), params)

	// do something with the result
	for iter.Next() {
		log.Print(iter.Item())
	}
	if iter.Err() != nil {
		log.Fatal(iter.Err())
	}

}
