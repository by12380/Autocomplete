package routers

import (
	"net/http"

	"github.com/by12380/Autocomplete/services/trie"
	"github.com/gin-gonic/gin"
)

// InitAutocomplete - Setup routes under /autocomplete
func InitAutocomplete(r *gin.RouterGroup) {
	r.GET("", autocompleteSearch)
	r.PUT("", autocompleteUpdate)
}

func autocompleteUpdate(c *gin.Context) {
	test := c.Param("test")

	responseBody := map[string]interface{}{
		"test": test,
	}

	c.JSON(http.StatusOK, responseBody)
}

func autocompleteSearch(c *gin.Context) {
	results := []*trie.TextItem{}

	text := c.Query("q")

	if text != "" {
		t := trie.GetInstance()
		results = t.Search(text)
	}

	responseBody := map[string]interface{}{
		"results": results,
	}

	c.JSON(http.StatusOK, responseBody)
}
