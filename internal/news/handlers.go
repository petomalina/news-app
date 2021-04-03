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

func (s *Server) handleGetProviders() gin.HandlerFunc {
	return func(c *gin.Context) {
		var names []string
		for name, _ := range s.providers {
			names = append(names, name)
		}
		c.JSON(200, names)
	}
}

func (s *Server) handleFetch() gin.HandlerFunc {
	return func(c *gin.Context) {
		provider, ok := s.providers[c.Query("p")]
		if !ok {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "requested provider could not be found",
			})
			return
		}

		category := c.Query("c")
		articles, err := provider.Fetch(category)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "could not fetch articles for the given provider",
			})
			return
		}

		c.JSON(http.StatusOK, articles)
	}
}
