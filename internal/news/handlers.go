package news

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
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
		articles, err := provider.Fetch(c.Request.Context(), category)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "could not fetch articles for the given provider",
			})
			return
		}

		// support descending order of news in case the app doesn't want
		// to render them based on relevance
		if c.Query("sort") == "desc" {
			sort.Slice(articles, func(i, j int) bool {
				return articles[i].PublishedParsed.After(*articles[j].PublishedParsed)
			})
		}

		c.JSON(http.StatusOK, articles)
	}
}
