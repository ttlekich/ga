package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	port string
}

func NewServer(port string) *Server {
	return &Server{
		port: port,
	}
}

func (s *Server) Start() error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", s.getRoot)
	r.Post("/game", s.startGame)
	return http.ListenAndServe(s.port, r)
}

func (s *Server) getRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
}

func (s *Server) startGame(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
}
