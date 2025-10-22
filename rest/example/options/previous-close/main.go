// Options - Previous Close
// https://polygon.io/docs/options/get_v2_aggs_ticker__optionsticker__prev
// https://github.com/spinningfactory/client-go/blob/master/rest/aggs.go
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
	params := models.GetPreviousCloseAggParams{
		Ticker: "O:SPY251219C00650000",
	}.WithAdjusted(true)

	// make request
	res, err := c.GetPreviousCloseAgg(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}

	// do something with the result
	log.Print(res)

}
