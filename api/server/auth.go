package server

import (
	"fmt"
	"net/http"

)

func (s *Server) handleSignup () http.HandlerFunc  {
  // closure. do anything here 
  num := 5

  return func (w http.ResponseWriter, r *http.Request)  {
    // actuall handlerFunc
    // call s.db.NewUserQuery().CreateUser to add user to db once logic is done
    fmt.Println("num: ", num)
  }
}
