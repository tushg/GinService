package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Server represents the HTTP server
type Server struct {
	*http.Server
}

// New creates a new HTTP server
func New(port string, router *gin.Engine) *Server {
	return &Server{
		Server: &http.Server{
			Addr:    ":" + port,
			Handler: router,
		},
	}
}
