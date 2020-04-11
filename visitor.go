package visitor

// PageLoader represents an interface to load HTML by URL
type PageLoader interface {
	LoadPage(url string) (statusCode int, finalURL string, body []byte, err error)
}

// ExternalScriptsExtractor represents an interface to extract external
// scripts from an HTML page
type ExternalScriptsExtractor interface {
	ExtractRefs(htmlBody []byte) (refs []string, err error)
}

// InlineScriptsExtractor represents an interface to extract inline
// scripts from an HTML page
type InlineScriptsExtractor interface {
	ExtractInlineScripts(htmlBody []byte) (refs []string, err error)
}
