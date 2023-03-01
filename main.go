package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/CeresDB/ceresdb-client-go/ceresdb"
)

var endpoint = "127.0.0.1:8831"

const table = "godemo"

func init() {
	if v := os.Getenv("CERESDB_ADDR"); v != "" {
		endpoint = v
	}
}

func main() {
	fmt.Printf("connect to %s...\n", endpoint)

	client, err := ceresdb.NewClient(endpoint, ceresdb.Direct,
		ceresdb.WithDefaultDatabase("public"),
	)
	if err != nil {
		panic(err)
	}

	// write
	points := make([]ceresdb.Point, 0, 2)
	for i := 0; i < 2; i++ {
		point, err := ceresdb.NewPointBuilder(table).
			SetTimestamp(currentMS()).
			AddTag("name", ceresdb.NewStringValue("test_tag1")).
			AddField("value", ceresdb.NewDoubleValue(1.0*float64(i))).
			Build()
		if err != nil {
			panic(err)
		}
		points = append(points, point)
	}

	ctx := context.TODO()
	resp, err := client.Write(ctx, ceresdb.WriteRequest{
		Points: points,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Write resp = %+v\n", resp)

	resp2, err := client.SQLQuery(ctx, ceresdb.SQLQueryRequest{
		Tables: []string{table},
		SQL:    fmt.Sprintf("select * from %s", table),
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Query resp = %+v\n", resp2)
}

func currentMS() int64 {
	return time.Now().UnixMilli()
}
