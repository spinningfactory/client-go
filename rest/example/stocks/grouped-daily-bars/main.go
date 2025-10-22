// Stocks - Grouped Daily (Bars)
// https://polygon.io/docs/stocks/get_v2_aggs_grouped_locale_us_market_stocks__date
// https://github.com/dhiaaayachi/client-go/blob/master/rest/aggs.go
package main

import (
	"context"
	"log"
	"os"
	"time"

	polygon "github.com/dhiaaayachi/client-go/rest"
	"github.com/dhiaaayachi/client-go/rest/models"
)

func main() {

	// init client
	c := polygon.New(os.Getenv("POLYGON_API_KEY"))

	// set params
	params := models.GetGroupedDailyAggsParams{
		Locale:     models.US,
		MarketType: models.Stocks,
		Date:       models.Date(time.Date(2023, 3, 8, 0, 0, 0, 0, time.Local)),
	}.WithAdjusted(true)

	// make request
	res, err := c.GetGroupedDailyAggs(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}

	// do something with the result
	log.Print(res)

}
