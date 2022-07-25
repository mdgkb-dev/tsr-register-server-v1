package search

import (
	handler "mdgkb/tsr-tegister-server-v1/handlers/search"

	"github.com/gin-gonic/gin"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("/main", h.SearchMain)
	r.GET("/v1", h.ElasticSearch)
	r.GET("/", h.Search)
	r.GET("/search-groups", h.SearchGroups)
	r.GET("/search-by-group", h.SearchGroups)
}
