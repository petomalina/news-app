package news

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"os"
)

// Server handles requests to the News API
type Server struct {
	log     *zap.Logger
	sources []string
}

// NewServer instantiates a new Server instance for the News API
func NewServer(log *zap.Logger, sources []string) *Server {
	return &Server{
		log:     log,
		sources: sources,
	}
}

// Routes returns a mux.Router with registered routes
func (s *Server) Routes() http.Handler {
	// make the ReleaseMode the new default instead of DebugMode
	if os.Getenv("GIN_MODE") == "" && gin.Mode() == gin.DebugMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.GET("/health", s.handleHealth())
	r.GET("/fetch", s.handleFetch())

	s.handleFetch()

	return r
}
