// Forex - Market Holidays
// https://polygon.io/docs/forex/get_v1_marketstatus_upcoming
// https://github.com/dhiaayachi/client-go/blob/master/rest/reference.go
package main

import (
	"context"
	"log"
	"os"

	polygon "github.com/dhiaayachi/client-go/rest"
)

func main() {
	// init client
	c := polygon.New(os.Getenv("POLYGON_API_KEY"))

	// make request
	res, err := c.GetMarketHolidays(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// do something with the result
	log.Print(res)
}
