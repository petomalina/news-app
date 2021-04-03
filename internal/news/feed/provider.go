package feed

// Provider abstracts an Article provider for a feed
type Provider interface {
	Fetch(category string) ([]*Article, error)
}
