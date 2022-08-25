package server

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sabino-ramirez/jems/api/repo"
)

type Server struct {
	db     *sql.DB
	dao    repo.DAO
	router *mux.Router
}

func NewServer(db *sql.DB, dao repo.DAO) *Server {
	s := &Server{
		db:     db,
		dao:    dao,
		router: mux.NewRouter(),
	}

	s.routes()
	return s
}

// make the server a httpHandler by implementing this one method
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// all routes in one place to "map out" services
func (s *Server) routes() {
	s.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "available")
	})

	s.router.HandleFunc("/signup", s.handleAuthSignup())
}
