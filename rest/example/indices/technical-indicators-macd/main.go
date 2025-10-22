// Indices - Moving Average Convergence/Divergence (MACD)
// https://polygon.io/docs/indices/get_v1_indicators_macd__indicesticker
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
	params := models.GetMACDParams{
		Ticker: "I:SPX",
	}.WithShortWindow(12).
		WithLongWindow(26).
		WithSignalWindow(9).
		WithOrder(models.Desc)

	// make request
	res, err := c.GetMACD(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}

	// do something with the result
	log.Print(res)
}
