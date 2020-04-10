package visitor

import (
	"fmt"
	"time"

	"github.com/valyala/fasthttp"
)

func CreateLoader() {
	// return
}

// FastHTTPPageLoader represents a way to load HTML by URL via fasthttp package
type FastHTTPPageLoader struct {
	Timeout      time.Duration
	MaxRedirects int
	UserAgent    string
	Buf          []byte
}

func (l FastHTTPPageLoader) LoadPage(url string) (statusCode int, finalUrl string, body []byte, error error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(res)

	req.Header.Add("User-Agent", l.UserAgent)
	finalURL := url
	for i := 0; i < l.MaxRedirects; i++ {
		req.SetRequestURI(finalURL)
		err := fasthttp.DoTimeout(req, res, l.Timeout)
		if err != nil {
			return 0, finalURL, nil, fmt.Errorf("loadPage: %w", err)
		}
		statusCode := res.StatusCode()
		if statusCode >= 300 && statusCode < 400 && res.Header.Peek("location") != nil {
			finalURL = string(res.Header.Peek("location"))
			continue
		}

		copy := fasthttp.AcquireResponse()
		res.CopyTo(copy)
		return statusCode, finalURL, copy.Body(), nil
	}

	return 0, finalURL, nil, fmt.Errorf("loadPage: maxRedirects (%v) exceeded", l.MaxRedirects)
}
