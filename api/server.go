package api

import (
	"log"
	"net/http"
	"webservermod/dao"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type server struct {
	addr   string
	mux    *chi.Mux
	userDB dao.UserDB
}

type Server interface {
	Start() error
	HandlePing(http.ResponseWriter, *http.Request)
	HandleIP(http.ResponseWriter, *http.Request)
	HandleGetByEmail(w http.ResponseWriter, r *http.Request)
}

func NewServer(addr string, userDB dao.UserDB) Server {
	mux := chi.NewRouter()

	if userDB == nil {
		log.Fatal("error conneting user db ")
		return nil
	}
	return &server{addr: addr, mux: mux, userDB: userDB}
}

func (s *server) Start() error {
	return http.ListenAndServe(s.addr, registerRoutes(s))
}

func registerRoutes(s *server) http.Handler {

	s.mux.Use(middleware.Recoverer)
	s.mux.Use(addIPToRequest)

	s.mux.Get("/ping", s.HandlePing)
	s.mux.Get("/ip", s.HandleIP)
	s.mux.Post("/user", s.HandleCreateUser)
	s.mux.Get("/user/{email}", s.HandleGetByEmail)
	return s.mux
}
