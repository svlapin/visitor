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
		log.Fatal(fmt.Errorf("%v: failed to get: %w", url, err))
	}

	log.Printf("%v: got response: %v", url, statusCode)

	externalExtractor := selectExternalScriptExtractor()
	refs, err := externalExtractor.ExtractRefs(body)

	if err != nil {
		log.Fatal(fmt.Errorf(" %v: failed to extract scripts: %w", url, err))
	}

	log.Printf("%v: external scripts: %v\n", url, refs)
}

func selectLoader() visitor.PageLoader {
	// TODO: make use of response buffer
	return visitor.FastHTTPPageLoader{Timeout: 10 * time.Second}
}

func selectExternalScriptExtractor() visitor.ExternalScriptsExtractor {
	return visitor.GoqueryParser{}
}
