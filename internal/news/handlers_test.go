package news

import (
	"encoding/json"
	"github.com/petomalina/news-app/internal/news/feed"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap/zaptest"
	"net/http"
	"net/http/httptest"
	"testing"
)

type ServerHandlersSuite struct {
	suite.Suite
}

func (s *ServerHandlersSuite) TestHandleHealth() {
	newsServer := NewServer(zaptest.NewLogger(s.T()))
	srv := httptest.NewServer(newsServer.Routes())
	defer srv.Close()

	res, err := http.Get(srv.URL + "/health")
	s.NoError(err)
	s.NotNil(res)

	var jsonRes map[string]string
	err = json.NewDecoder(res.Body).Decode(&jsonRes)
	s.NoError(err)
	s.NoError(res.Body.Close())

	s.Equal(http.StatusOK, res.StatusCode)
	s.Equal("ok", jsonRes["status"])
}

type handleFetchCandidate struct {
	serverOpts       []ServerOpt
	queryString      string
	expectedStatus   int
	expectedArticles []*feed.Article
}

func (s *ServerHandlersSuite) TestHandleFetch() {
	candidates := []handleFetchCandidate{
		{
			serverOpts: []ServerOpt{
				WithProvider(
					feed.NewMockProvider(
						[]*feed.Article{
							{
								Title:   "Hello World",
								Content: "This is a test article",
							},
						},
						nil,
					),
				),
			},
			expectedStatus: http.StatusOK,
			expectedArticles: []*feed.Article{
				{
					Title:   "Hello World",
					Content: "This is a test article",
				},
			},
		},
	}
	for _, c := range candidates {
		s.Run(
			"", func() {
				newsServer := NewServer(zaptest.NewLogger(s.T()), c.serverOpts...)
				srv := httptest.NewServer(newsServer.Routes())
				defer srv.Close()

				res, err := http.Get(srv.URL + "/fetch" + c.queryString)
				s.NoError(err)
				s.NotNil(res)
				defer s.NoError(res.Body.Close())
				s.Equal(c.expectedStatus, res.StatusCode)

				var articles []*feed.Article
				err = json.NewDecoder(res.Body).Decode(&articles)
				s.NoError(err)

				s.ElementsMatch(c.expectedArticles, articles)
			},
		)
	}
}

func TestServerHandlerSuite(t *testing.T) {
	suite.Run(t, &ServerHandlersSuite{})
}
