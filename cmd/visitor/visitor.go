package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/svlapin/visitor"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("url required")
	}

	loader := selectLoader()

	url := os.Args[1]

	statusCode, body, err := loader.LoadPage(url)

	if err != nil {
		log.Fatal(fmt.Errorf("failed to get %v: %w", url, err))
	}

	log.Println("Got response: ", statusCode, string(body), err)
}

func selectLoader() visitor.PageLoader {
	// TODO: make use of response buffer
	return visitor.FastHTTPPageLoader{Timeout: 10 * time.Second}
}
