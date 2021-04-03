package news

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) handleHealth() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	}
}

func (s *Server) handleFetch() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Status(http.StatusOK)
	}
}
