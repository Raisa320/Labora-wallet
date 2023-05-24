package config

import (
	"fmt"
	"net/http"
)

type Server struct {
	listenAddr string
	handler    http.Handler
}

func NewServer(listenAddr string, router http.Handler) *Server {

	return &Server{
		listenAddr: listenAddr,
		handler:    router,
	}
}

func (s *Server) Start() error {
	fmt.Println("Server running on port", s.listenAddr)
	return http.ListenAndServe(s.listenAddr, s.handler)
}
