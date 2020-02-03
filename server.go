package main

import (
	"fmt"
	"net/http"
	"strings"
)

type Server struct {
	config *Config
}

func NewServer(config *Config) (*Server, error) {
	s := &Server{
		config: config,
	}
	return s, s.run()
}

func (s *Server) run() error {
	http.HandleFunc("/", s.handleHttp)
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", s.config.Listen, s.config.Port), nil)
	return err
}

func (s *Server) handleHttp(w http.ResponseWriter, r *http.Request) {
	destination, found := s.getDestination(r.URL.Path)
	if found {
		http.Redirect(w, r, destination, 301)
	} else {
		http.NotFound(w, r)
	}
}

func (s *Server) getDestination(path string) (string, bool) {
	// Remove slash
	d := strings.TrimPrefix(path, "/")
	d = strings.TrimSuffix(d, "/")
	destination, found := s.config.Destinations[d].(string)
	return destination, found
}
