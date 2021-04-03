package news

import (
	"github.com/gin-gonic/gin"
	"github.com/petomalina/news-app/internal/news/feed"
	"go.uber.org/zap"
	"net/http"
	"os"
)

// Server handles requests to the News API
type Server struct {
	log       *zap.Logger
	providers map[string]feed.Provider
}

// ServerOpt is a type used to configure a Server instance
type ServerOpt = func(s *Server)

// WithProvider registers a new feed.Provider to the Server
func WithProvider(name string, p feed.Provider) ServerOpt {
	return func(s *Server) {
		s.providers[name] = p
	}
}

// NewServer instantiates a new Server instance for the News API
func NewServer(log *zap.Logger, opts ...ServerOpt) *Server {
	srv := &Server{
		log:       log,
		providers: map[string]feed.Provider{},
	}

	for _, opt := range opts {
		opt(srv)
	}

	return srv
}

// Routes returns a mux.Router with registered routes
func (s *Server) Routes() http.Handler {
	// make the ReleaseMode the new default instead of DebugMode
	if os.Getenv("GIN_MODE") == "" && gin.Mode() == gin.DebugMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.GET("/health", s.handleHealth())
	r.GET("/sources", s.handleSourceInfo())
	r.GET("/fetch", s.handleFetch())

	s.handleFetch()

	return r
}
