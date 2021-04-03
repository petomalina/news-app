package feed

import "github.com/mmcdole/gofeed"

// RSSProvider fetches RSS feeds
type RSSProvider struct {
	source string
}

// NewRSSProvider instantiates a new RSSProvider with the given source
func NewRSSProvider(source string) *RSSProvider {
	return &RSSProvider{
		source: source,
	}
}

// Fetch parses the given source and returns its articles
func (p *RSSProvider) Fetch() ([]*Article, error) {
	feed, err := gofeed.NewParser().ParseURL(p.source)
	if err != nil {
		return nil, err
	}

	return feed.Items, nil
}
