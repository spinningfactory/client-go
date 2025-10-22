// Options - Daily Open/Close
// https://polygon.io/docs/options/get_v1_open-close__optionsticker___date
// https://github.com/dhiaayachi/client-go/blob/master/rest/aggs.go
package main

import (
	"context"
	"log"
	"os"
	"time"

	polygon "github.com/dhiaayachi/client-go/rest"
	"github.com/dhiaayachi/client-go/rest/models"
)

func main() {

	// init client
	c := polygon.New(os.Getenv("POLYGON_API_KEY"))

	// set params
	params := models.GetDailyOpenCloseAggParams{
		Ticker: "O:SPY251219C00650000",
		Date:   models.Date(time.Date(2023, 2, 22, 0, 0, 0, 0, time.Local)),
	}.WithAdjusted(true)

	// make request
	res, err := c.GetDailyOpenCloseAgg(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}

	// do something with the result
	log.Print(res)

}
