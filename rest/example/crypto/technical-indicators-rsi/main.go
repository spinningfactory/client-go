// Crypto - Relative Strength Index (RSI)
// https://polygon.io/docs/crypto/get_v1_indicators_rsi__cryptoticker
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
	params := &models.GetRSIParams{
		Ticker: "X:BTCUSD",
	}

	// make request
	res, err := c.GetRSI(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}

	// do something with the result
	log.Print(res)
}
