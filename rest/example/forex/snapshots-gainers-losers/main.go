// Forex - Snapshot Gainers/Losers
// https://polygon.io/docs/forex/get_v2_snapshot_locale_global_markets_forex__direction
// https://github.com/spinningfactory/client-go/blob/master/rest/snapshot.go
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
	params := &models.GetGainersLosersSnapshotParams{
		Locale:     models.Global,
		MarketType: models.Forex,
		Direction:  models.Gainers, // or models.Losers
	}

	// make request
	res, err := c.GetGainersLosersSnapshot(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}

	// do something with the result
	log.Print(res)
}
