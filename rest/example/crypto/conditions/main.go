// Crypto - Conditions
// https://polygon.io/docs/crypto/get_v3_reference_conditions
// https://github.com/dhiaayachi/client-go/blob/master/rest/reference.go
package main

import (
	"context"
	"log"
	"os"

	polygon "github.com/dhiaayachi/client-go/rest"
	"github.com/dhiaayachi/client-go/rest/models"
)

func main() {
	// init client
	c := polygon.New(os.Getenv("POLYGON_API_KEY"))

	// set params
	params := models.ListConditionsParams{}.
		WithAssetClass(models.AssetCrypto).
		WithLimit(1000)

	// make request
	iter := c.ListConditions(context.Background(), params)

	// do something with the result
	for iter.Next() {
		log.Print(iter.Item())
	}
	if iter.Err() != nil {
		log.Fatal(iter.Err())
	}
}
