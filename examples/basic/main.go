package main

import (
	"context"
	"fmt"
	"time"

	easyhttp "github.com/lockp111/go-easyhttp"
)

func main() {
	client := easyhttp.NewClient(
		easyhttp.Config{Timeout: 5 * time.Second},
		easyhttp.WithMaxConns(100),
	)

	req, _ := easyhttp.NewGet("https://httpbin.org/get")
	req.AddQuery("q", "ping").SetHeader("Accept", "application/json")

	resp, err := client.Fetch(context.Background(), req)
	if err != nil {
		panic(err)
	}

	fmt.Println("status:", resp.Status)
	fmt.Println("body:", resp.GetBodyString())
}
