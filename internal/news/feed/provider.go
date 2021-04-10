package feed

import "context"

// Provider abstracts an Article provider for a feed
type Provider interface {
	Fetch(ctx context.Context, category string) ([]*Article, error)
}
