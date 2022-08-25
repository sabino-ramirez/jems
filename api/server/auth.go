package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/sabino-ramirez/jems/api/models"
	"golang.org/x/crypto/bcrypt"
)

func (s *Server) handleAuthSignup() http.HandlerFunc {
	log.Println("handleSignup invoked")
	var user models.User

	return func(w http.ResponseWriter, r *http.Request) {
		// actual handlerFunc
		// signup logic done here
		// call s.dao.NewUserRepo().CreateUser to add user to db once logic is done

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			// return 400 if request body is broken
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// use bcrypt to salt and hash password
		user.Password, err = hashPassword(user.Password)
		if err != nil {
			// hashing failed try again after some time
			// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
      fmt.Errorf("hashPassword:", err)
      return
		}

		// insert username and hashed password to the database
		user, err := s.dao.UserRepo(s.db).CreateUser(user)
		if err != nil {
			// logic to handle error for creating user in database
      
      return
		}
	}
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	return string(bytes), fmt.Errorf("hashPassword:", err)
}
