package feed

type Provider interface {
	Fetch(category string) ([]*Article, error)
}
