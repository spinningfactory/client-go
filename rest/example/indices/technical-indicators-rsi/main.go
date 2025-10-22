// Indices - Relative Strength Index (RSI)
// https://polygon.io/docs/indices/get_v1_indicators_rsi__indicesticker
// https://github.com/dhiaaayachi/client-go/blob/master/rest/indicators.go
package main

import (
	"context"
	"log"
	"os"

	polygon "github.com/dhiaaayachi/client-go/rest"
	"github.com/dhiaaayachi/client-go/rest/models"
)

func main() {
	// init client
	c := polygon.New(os.Getenv("POLYGON_API_KEY"))

	// set params
	params := &models.GetRSIParams{
		Ticker: "I:SPX",
	}

	// make request
	res, err := c.GetRSI(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}

	// do something with the result
	log.Print(res)
}
