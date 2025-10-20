package main

import (
	"context"
	"fmt"
	"time"

	easyhttp "github.com/lockp111/go-easyhttp"
)

func main() {
	client := easyhttp.NewClient(easyhttp.Config{Timeout: 10 * time.Second})

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, _ := easyhttp.NewGet("https://httpbin.org/delay/2")
	_, err := client.Fetch(ctx, req)
	fmt.Println("err:", err)
}
