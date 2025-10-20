package main

import (
	"context"
	"os"
	"time"

	easyhttp "github.com/lockp111/go-easyhttp"
)

func main() {
	client := easyhttp.NewClient(easyhttp.Config{Timeout: 60 * time.Second})

	f, err := os.Open("./bigfile.bin")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	req, _ := easyhttp.NewPost("https://httpbin.org/post")
	req.SetHeader("Content-Type", "application/octet-stream")
	req.SetBodyReader(f)

	if _, err := client.Fetch(context.Background(), req); err != nil {
		panic(err)
	}
}
