package visitor

import (
	"time"

	"github.com/valyala/fasthttp"
)

// FastHTTPPageLoader represents a way to load HTML by URL via fasthttp package
type FastHTTPPageLoader struct {
	Timeout time.Duration
	Buf     []byte
}

func (l FastHTTPPageLoader) LoadPage(url string) (statusCode int, body []byte, error error) {
	return fasthttp.GetTimeout(l.Buf, url, l.Timeout)
}
