// Indices - Market Status
// https://polygon.io/docs/indices/get_v1_marketstatus_now
// https://github.com/spinningfactory/client-go/blob/master/rest/reference.go
package main

import (
	"context"
	"log"
	"os"

	polygon "github.com/spinningfactory/client-go/rest"
)

func main() {
	// init client
	c := polygon.New(os.Getenv("POLYGON_API_KEY"))

	// make request
	res, err := c.GetMarketStatus(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// do something with the result
	log.Print(res)
}
