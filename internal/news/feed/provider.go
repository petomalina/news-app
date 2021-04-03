package feed

type Provider interface {
	Fetch() ([]*Article, error)
}
