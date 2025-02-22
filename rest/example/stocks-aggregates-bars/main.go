// Stocks - Aggregates (Bars)
// https://polygon.io/docs/stocks/get_v2_aggs_ticker__stocksticker__range__multiplier___timespan___from___to
// https://github.com/polygon-io/client-go/blob/master/rest/aggs.go
package main

import (
	"context"
	"time"
	"log"
	"os"

	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
)

func main() {

	// init client
	c := polygon.New(os.Getenv("POLYGON_API_KEY"))

	// set params
	params := models.ListAggsParams{
		Ticker:     "AAPL",
		Multiplier: 1,
		Timespan:   "day",
		From:       models.Millis(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)),
		To:         models.Millis(time.Date(2023, 3, 9, 0, 0, 0, 0, time.UTC)),
	}.WithOrder(models.Desc).WithLimit(50000).WithAdjusted(true)

	// make request
	iter := c.ListAggs(context.Background(), params)

	// do something with the result
	for iter.Next() {
		log.Print(iter.Item()) 
	}
	if iter.Err() != nil {
		log.Fatal(iter.Err())
	}

}
