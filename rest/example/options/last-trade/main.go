// Options - Last Trade
// https://polygon.io/docs/options/get_v2_last_trade__optionsticker
// https://github.com/dhiaaayachi/client-go/blob/master/rest/trades.go
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
	params := &models.GetLastTradeParams{
		Ticker: "O:TSLA210903C00700000",
	}

	// make request
	res, err := c.GetLastTrade(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}

	// do something with the result
	log.Print(res)

}
