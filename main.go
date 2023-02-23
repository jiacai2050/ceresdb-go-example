package main

import (
	"context"
	"fmt"
	"time"

	"github.com/CeresDB/ceresdb-client-go/ceresdb"
	"github.com/CeresDB/ceresdb-client-go/types"
	"github.com/CeresDB/ceresdb-client-go/utils"
)

func currentMS() int64 {
	return time.Now().UnixMilli()
}

func main() {
	endpoint := "127.0.0.1:8831"
	client, err := ceresdb.NewClient(endpoint, types.Direct,
		ceresdb.WithDefaultDatabase("public"),
	)
	if err != nil {
		panic(err)
	}

	// write
	points := make([]types.Point, 0, 2)
	for i := 0; i < 2; i++ {
		point, err := ceresdb.NewPointBuilder("demo").
			SetTimestamp(utils.CurrentMS()).
			AddTag("name", types.NewStringValue("test_tag1")).
			AddField("value", types.NewDoubleValue(1.0*float64(i))).
			Build()
		if err != nil {
			panic(err)
		}
		points = append(points, point)
	}

	resp, err := client.Write(context.Background(), types.WriteRequest{
		Points: points,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Write resp = %+v\n", resp)

}
