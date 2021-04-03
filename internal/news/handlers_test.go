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
	newsServer := NewServer(zaptest.NewLogger(s.T()), []string{})
	srv := httptest.NewServer(newsServer.Routes())
	defer srv.Close()

	res, err := http.Get(srv.URL + "/health")
	s.NoError(err)
	s.NotNil(res, "response should not be nil")

	var jsonRes map[string]string
	err = json.NewDecoder(res.Body).Decode(&jsonRes)
	s.NoError(err)
	s.NoError(res.Body.Close())

	s.Equal(http.StatusOK, res.StatusCode)
	s.Equal("ok", jsonRes["status"])
}

func TestServerHandlerSuite(t *testing.T) {
	suite.Run(t, &ServerHandlersSuite{})
}
