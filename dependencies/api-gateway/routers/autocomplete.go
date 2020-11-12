package routers

import (
	"net/http"
	"net/http/httputil"
	"os"

	"github.com/gin-gonic/gin"
)

// InitAutocomplete - Setup routes under /autocomplete
func InitAutocomplete(r *gin.RouterGroup) {
	r.GET("", reverseProxy)
}

func reverseProxy(c *gin.Context) {
	targetHost := os.Getenv("AUTOCOMPLETE_SERVICE_HOST") + ":" + os.Getenv("AUTOCOMPLETE_SERVICE_PORT")

	director := func(req *http.Request) {
		req.URL.Scheme = "http"
		req.URL.Host = targetHost
		req.Header["my-header"] = []string{req.Header.Get("my-header")}
		// Golang camelcases headers
		delete(req.Header, "My-Header")
	}

	proxy := &httputil.ReverseProxy{Director: director}
	proxy.ServeHTTP(c.Writer, c.Request)
}
