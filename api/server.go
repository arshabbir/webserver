package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type server struct {
	addr string
	mux  *chi.Mux
}

type Server interface {
	Start() error
	HandlePing(http.ResponseWriter, *http.Request)
	HandleIP(http.ResponseWriter, *http.Request)
}

func NewServer(addr string) Server {
	mux := chi.NewRouter()
	return &server{addr: addr, mux: mux}
}

func (s *server) Start() error {
	return http.ListenAndServe(s.addr, registerRoutes(s))
}

func registerRoutes(s *server) http.Handler {

	s.mux.Use(middleware.Recoverer)
	s.mux.Use(addIPToRequest)

	s.mux.Get("/ping", s.HandlePing)
	s.mux.Get("/ip", s.HandleIP)
	return s.mux
}
