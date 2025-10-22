// Indices - Ticker Types
// https://polygon.io/docs/indices/get_v3_reference_tickers_types
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
	params := models.GetTickerTypesParams{}.WithAssetClass("indices")

	// make request
	res, err := c.GetTickerTypes(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}

	// do something with the result
	log.Print(res)
}
