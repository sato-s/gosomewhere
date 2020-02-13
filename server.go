package main

import (
	"fmt"
	"github.com/agnivade/levenshtein"
	"html/template"
	"log"
	"math"
	"net/http"
	"strings"
)

type Server struct {
	config    *Config
	templates *template.Template
}

func NewServer(config *Config) (*Server, error) {
	templates := template.Must(template.ParseGlob("templates/*"))
	s := &Server{
		config:    config,
		templates: templates,
	}
	return s, s.run()
}

func (s *Server) run() error {
	http.HandleFunc("/", s.handleHttp)
	log.Printf("running server.... port:%d listening:%s", s.config.Port, s.config.Listen)
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", s.config.Listen, s.config.Port), nil)
	return err
}

func (s *Server) handleHttp(w http.ResponseWriter, r *http.Request) {
	destination, found := s.getDestination(r.URL.Path)
	if found {
		log.Printf("Redirect %s to %s", r.URL.Path, destination)
		http.Redirect(w, r, destination, 307)
	} else {
		s.handleDestinationUncertain(w, r)
	}
}

// Serve default html page if we can't determine target
func (s *Server) handleDestinationUncertain(w http.ResponseWriter, r *http.Request) {
	if err := s.templates.ExecuteTemplate(w, "index.html", s.config); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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

	if minDistance < 2 {
		return destination, true
	} else {
		return "", false
	}
}
