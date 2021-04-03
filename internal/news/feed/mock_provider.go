package feed

// MockProvider provides articles that were given to it directly
type MockProvider struct {
	articles []*Article
	err      error
}

// NewMockProvider instantiates a new MockProvider
func NewMockProvider(articles []*Article, err error) *MockProvider {
	return &MockProvider{
		articles: articles,
		err:      err,
	}
}

// Fetch returns the mocked articles provided in the constructor
func (p *MockProvider) Fetch() ([]*Article, error) {
	return p.articles, p.err
}
