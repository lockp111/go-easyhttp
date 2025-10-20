package main

import (
	"context"
	"fmt"
	"time"

	easyhttp "github.com/lockp111/go-easyhttp"
)

func main() {
	client := easyhttp.NewClient(
		easyhttp.Config{Timeout: 3 * time.Second},
		easyhttp.WithConnTimeout(500*time.Millisecond),
		easyhttp.WithResponseHeaderTimeout(2*time.Second),
		easyhttp.WithIdleConnTimeout(30*time.Second),
		easyhttp.WithMaxConns(50),
	)

	req, _ := easyhttp.NewGet("https://httpbin.org/delay/1")
	start := time.Now()
	resp, err := client.Fetch(context.Background(), req)
	elapsed := time.Since(start)

	fmt.Println("err:", err)
	if resp != nil {
		fmt.Println("status:", resp.Status)
	}
	fmt.Println("elapsed:", elapsed)
}
