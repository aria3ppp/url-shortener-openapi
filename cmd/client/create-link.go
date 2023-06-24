package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/aria3ppp/url-shortener-openapi/pkg/client"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "error: no link provided.\n")
		fmt.Printf("usage: %s <link> [<shortened string>]\n", os.Args[0])
		os.Exit(1)
		return
	}

	var (
		link            string
		shortenedString *string
	)

	link = os.Args[1]

	if len(os.Args) > 2 {
		shortenedString = new(string)
		*shortenedString = os.Args[2]
	}

	basicauth := os.Getenv("BASICAUTH")
	if basicauth == "" || !strings.Contains(basicauth, ":") {
		fmt.Fprintf(os.Stderr, "error: invalid BASICAUTH environment value %q\n", basicauth)
		os.Exit(1)
		return
	}

	username, password, _ := strings.Cut(basicauth, ":")

	ctx := context.Background()

	c, err := client.NewClient("http://localhost:8080")
	if err != nil {
		panic(err)
	}

	resp, err := c.CreateLink(
		ctx,
		client.CreateLinkJSONRequestBody{
			Url:             link,
			ShortenedString: shortenedString,
		},
		func(ctx context.Context, req *http.Request) error {
			req.SetBasicAuth(username, password)
			return nil
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
