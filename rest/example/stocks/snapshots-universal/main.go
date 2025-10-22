package main

import (
	"context"
	"log"
	"os"

	polygon "github.com/dhiaaayachi/client-go/rest"
	"github.com/dhiaaayachi/client-go/rest/models"
)

func main() {
	// Init client
	c := polygon.New(os.Getenv("POLYGON_API_KEY"))

	// Set parameters
	params := models.ListUniversalSnapshotsParams{}.WithTickerAnyOf("AAPL,META,F")

	// Make request
	iter := c.ListUniversalSnapshots(context.Background(), params)

	// do something with the result
	for iter.Next() {
		log.Println(iter.Item())
	}
	if iter.Err() != nil {
		log.Fatal(iter.Err())
	}
}
