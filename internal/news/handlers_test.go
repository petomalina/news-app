package news

import (
	"encoding/json"
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
	srv := makeNewsServer(s.T())
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
	queryString    string
	expectedStatus int
}

func (s *ServerHandlersSuite) TestHandleFetch() {
	candidates := []handleFetchCandidate{
		{
			expectedStatus: http.StatusOK,
		},
	}
	for _, c := range candidates {
		s.Run("", func() {
			srv := makeNewsServer(s.T())
			defer srv.Close()

			res, err := http.Get(srv.URL + "/fetch" + c.queryString)
			s.NoError(err)
			s.NotNil(res)
			s.Equal(c.expectedStatus, res.StatusCode)
		})
	}
}

func TestServerHandlerSuite(t *testing.T) {
	suite.Run(t, &ServerHandlersSuite{})
}

func makeNewsServer(t *testing.T) *httptest.Server {
	newsServer := NewServer(zaptest.NewLogger(t), []string{})
	return httptest.NewServer(newsServer.Routes())
}
