package main

import (
	_ "embed"
	"fmt"
	"github.com/agnivade/levenshtein"
	"log"
	"math"
	"net/http"
	"strings"
)

//go:embed index.html
var indexHTML string

type Server struct {
	config  *Config
	topPage string
}

func NewServer(config *Config) (*Server, error) {
	s := &Server{
		config:  config,
		topPage: indexHTML,
	}
	return s, s.run()
}

func (s *Server) run() error {
	http.HandleFunc("/", s.handleHttp)
	log.Printf("running server.... port:%d listening:%s", s.config.Port, s.config.Listen)
	if s.config.IsBasicauthEnabled() {
		log.Printf("basic auth is enabled")
	} else {
		log.Printf("basic auth is disabled")
	}
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", s.config.Listen, s.config.Port), nil)
	return err
}

func (s *Server) handleHttp(w http.ResponseWriter, r *http.Request) {
	if err := s.checkAuthHeader(r); err != nil {
		log.Printf("Basic auth failed %s", err)
		w.Header().Add("WWW-Authenticate", `Basic realm="SECRET AREA"`)
		http.Error(w, "Unauthorized.", http.StatusUnauthorized)
		return
	} else {
		log.Printf("Basic auth success")
	}

	// Render "/" page
	if r.URL.Path == "/" {
		s.handleTopPage(w, r)
		return
	}

	destination, found := s.getDestination(r.URL.Path)
	if found {
		log.Printf("Redirect %s to %s", r.URL.Path, destination)
		http.Redirect(w, r, destination, 307)
	} else {
		log.Printf("Could not find %s from destinations", r.URL.Path)
		http.NotFound(w, r)
	}
}

func (s *Server) checkAuthHeader(r *http.Request) error {
	username, password, ok := r.BasicAuth()
	if !ok {
		fmt.Errorf("Failed to parse basic auth header")
	}
	if username == s.config.Basicauth.Username &&
		password == s.config.Basicauth.Password {
		return nil
	}

	return fmt.Errorf("username or password doensn't match")
}

func (s *Server) handleTopPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, s.topPage)
}

func (s *Server) getDestination(path string) (string, bool) {
	// Remove slash
	d := strings.TrimPrefix(path, "/")
	d = strings.TrimSuffix(d, "/")
	// Find closest destination using edit distance (=levenshtein distance).
	var destination string
	minDistance := math.MaxInt32
	for k, v := range s.config.Destinations {
		distance := levenshtein.ComputeDistance(k, d)
		if distance < minDistance {
			minDistance = distance
			destination = v
		}
	}

	if minDistance < 5 {
		return destination, true
	} else {
		return "", false
	}
}
