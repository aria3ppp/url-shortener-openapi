package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/aria3ppp/url-shortener-openapi/pkg/client"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "error: no shortened string provided.\n")
		fmt.Printf("usage: %s <shortened string>\n", os.Args[0])
		os.Exit(1)
		return
	}

	shortenedString := os.Args[1]

	ctx := context.Background()

	client, err := client.NewClient("http://localhost:8080")
	if err != nil {
		panic(err)
	}

	resp, err := client.GetLink(ctx, shortenedString)
	if err != nil {
		panic(err)
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if err := resp.Body.Close(); err != nil {
		panic(err)
	}

	fmt.Printf("status code: %d\n", resp.StatusCode)
	fmt.Printf("body: %s\n", responseBody)
}
