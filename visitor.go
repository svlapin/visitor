package visitor

// PageLoader represents an interface to load HTML by URL
type PageLoader interface {
	LoadPage(url string) (statusCode int, body []byte, err error)
}

// ExternalScriptsExtractor represents an interface to extract external
// scripts from an HTML page
type ExternalScriptsExtractor interface {
	ExtractRefs(htmlBody []byte) (refs []string, err error)
}
