package repo

import (
	"database/sql"

	"github.com/sabino-ramirez/jems/api/models"
)

type UserRepo interface {
	CreateUser(user models.User) (models.User, error)
  // other functions related to users
}

type userRepo struct{
  db *sql.DB
}

func (u *userRepo) CreateUser(user models.User) (models.User, error) {
	// logic to create user in  db
	// where we actually use the sqlite3 database driver

	return user, nil
}
