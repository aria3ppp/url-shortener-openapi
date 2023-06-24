package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/aria3ppp/url-shortener-openapi/pkg/client"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "error: no username and password provided.\n")
		fmt.Printf("usage: %s <username> <password>\n", os.Args[0])
		os.Exit(1)
		return
	}

	username, password := os.Args[1], os.Args[2]

	ctx := context.Background()

	c, err := client.NewClient("http://localhost:8080")
	if err != nil {
		panic(err)
	}

	resp, err := c.CreateUser(
		ctx,
		client.CreateUserJSONRequestBody{
			Username: username,
			Password: password,
		},
	)
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
