package server

import (
	"encoding/json"
	"hsl/internal/server/middlewares"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := mux.NewRouter()

	// Apply CORS middleware
	r.Use(middlewares.CorsMiddleware)

	// Home Handlers

	// Auth

	// Default handlers
	r.HandleFunc("/api/", s.HelloWorldHandler)
	r.HandleFunc("/api/health", s.healthHandler)
	// straming handlers
	// Need to be authed, so gotta add a authMiddlewware
	r.Use(middlewares.AuthMiddleware)

	return r
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.Health())

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
