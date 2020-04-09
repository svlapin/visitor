package visitor

// PageLoader represents an interface to load HTML by URL
type PageLoader interface {
	LoadPage(url string) (statusCode int, body []byte, err error)
}
