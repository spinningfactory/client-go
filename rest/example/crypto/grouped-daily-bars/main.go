// Crypto - Grouped Daily (Bars)
// https://polygon.io/docs/crypto/get_v2_aggs_grouped_locale_global_market_crypto__date
// https://github.com/spinningfactory/client-go/blob/master/rest/aggs.go
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
	params := models.GetGroupedDailyAggsParams{
		Locale:     models.Global,
		MarketType: models.Crypto,
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
