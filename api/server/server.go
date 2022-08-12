package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sabino-ramirez/jems/api/repo"
)

type Server struct {
	db     repo.DAO
	router *mux.Router
}

func NewServer(db repo.DAO) *Server {
  s := &Server{
    db: db,
    router: mux.NewRouter(),
  }

  s.routes()
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  s.router.ServeHTTP(w, r)
}

func (s *Server) routes () {
	s.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "available")
	})

	s.router.HandleFunc("/signup", s.handleSignup())
}

