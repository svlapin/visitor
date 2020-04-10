package visitor

import (
	"bytes"
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

// GoqueryParser represetns an entity to parse HTML with goquery
type GoqueryParser struct{}

// ExtractRefs extracts an array of external scripts from HTML document
func (g GoqueryParser) ExtractRefs(htmlBody []byte) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(htmlBody))

	if err != nil {
		return nil, fmt.Errorf("goquery: %w", err)
	}

	refs := make([]string, 0)

	doc.Find("script").Each(func(_ int, s *goquery.Selection) {
		if src, exists := s.Attr("src"); exists {
			refs = append(refs, src)
		}
	})

	return refs, nil
}
