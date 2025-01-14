package server

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
}

type WeatherResponse struct {
	Temperature float64 `json:"temperature"`
	Condition   string  `json:"condition"`
	City        string  `json:"city"`
	Country     string  `json:"country"`
}

func NewServer() *Server {
	s := &Server{
		router: mux.NewRouter(),
	}
	s.routes()
	return s
}

func (s *Server) Router() *mux.Router {
	return s.router
}

func (s *Server) routes() {
	s.router.HandleFunc("/api/weather", s.handleWeather()).Methods("GET")
}

func (s *Server) handleWeather() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get client IP
		ip := getClientIP(r)

		// TODO: Implement IP geolocation
		// TODO: Implement weather lookup

		// For now, return mock data
		weather := WeatherResponse{
			Temperature: 20.5,
			Condition:   "Sunny",
			City:        "London",
			Country:     "UK",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(weather)
	}
}

func getClientIP(r *http.Request) string {
	// Check X-Forwarded-For header
	ip := r.Header.Get("X-Forwarded-For")
	if ip != "" {
		return strings.Split(ip, ",")[0]
	}
	
	// Check X-Real-IP header
	ip = r.Header.Get("X-Real-IP")
	if ip != "" {
		return ip
	}
	
	// Fall back to RemoteAddr
	return strings.Split(r.RemoteAddr, ":")[0]
}