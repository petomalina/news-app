package feed

import (
	"fmt"
	"github.com/mmcdole/gofeed"
)

// RSSProvider fetches RSS feeds
type RSSProvider struct {
	// sourceFmt is a string that contains a single %s argument
	// to be replaced by the category provided to the Fetch method
	sourceFmt string
}

// NewRSSProvider instantiates a new RSSProvider with the given sourceFmt
func NewRSSProvider(sourceFmt string) *RSSProvider {
	return &RSSProvider{
		sourceFmt: sourceFmt,
	}
}

// Fetch parses the given sourceFmt and returns its articles
func (p *RSSProvider) Fetch(category string) ([]*Article, error) {
	feed, err := gofeed.NewParser().ParseURL(fmt.Sprintf(p.sourceFmt, category))
	if err != nil {
		return nil, err
	}

	return feed.Items, nil
}
