package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/svlapin/visitor"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("url required")
	}

	loader := selectLoader()

	targetURL := os.Args[1]

	statusCode, finalURL, body, err := loader.LoadPage(targetURL)

	if err != nil {
		log.Fatal(fmt.Errorf("%v: failed to get: %w", targetURL, err))
	}

	if finalURL != targetURL {
		log.Printf("%v: redirected: %v", targetURL, finalURL)
	}
	log.Printf("%v: got response: %v", targetURL, statusCode)

	externalExtractor := selectExternalScriptExtractor()
	refs, err := externalExtractor.ExtractRefs(body)

	if err != nil {
		log.Fatal(fmt.Errorf("%v: failed to extract scripts: %w", targetURL, err))
	}

	abs := make([]*url.URL, 0, len(refs))
	f, err := url.Parse(finalURL)
	if err != nil {
		log.Fatal(fmt.Errorf("%v: failed to parse URL: %w", finalURL, err))
	}

	for _, u := range refs {
		if r, err := visitor.ResolveURL(u, f); err == nil {
			abs = append(abs, r)
		}
	}

	for _, u := range abs {
		statusCode, _, script, err := loader.LoadPage(u.String())

		if err != nil {
			log.Fatal(fmt.Errorf("failed to load script %v: %w", u, err))
		}

		log.Printf("loaded script %v: %v", u, statusCode)

		_, err = visitor.ExtractStrings(script)
		if err != nil {
			log.Fatal(fmt.Errorf("failed to extract strings %v: %w", u, err))
		}

		// log.Println(strings)
	}

	inlineExtractor := selectInlineScriptExtractor()
	inlineScripts, err := inlineExtractor.ExtractInlineScripts(body)
	if err != nil {
		log.Fatal(fmt.Errorf("%v: failed to extract inline scripts: %w", targetURL, err))
	}

	for _, scr := range inlineScripts {
		strings, err := visitor.ExtractStrings([]byte(scr))
		if err != nil {
			log.Fatal(fmt.Errorf("failed to extract strings from inline script %v: %w", scr, err))
		}
		log.Println(strings)
	}
}

func selectLoader() visitor.PageLoader {
	// TODO: make use of response buffer
	return visitor.FastHTTPPageLoader{Timeout: 10 * time.Second, MaxRedirects: 5}
}

func selectExternalScriptExtractor() visitor.ExternalScriptsExtractor {
	return visitor.GoqueryParser{}
}

func selectInlineScriptExtractor() visitor.InlineScriptsExtractor {
	return visitor.GoqueryParser{}
}
