package visitor

import "net/url"

func ResolveURL(stringURL string, baseURL *url.URL) (absoluteURL *url.URL, err error) {
	parsed, err := url.Parse(stringURL)
	if err != nil {
		return nil, err
	}
	return baseURL.ResolveReference(parsed), nil
}
