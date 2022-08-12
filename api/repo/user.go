package repo

import (

	"github.com/sabino-ramirez/jems/api/models"
)

type UserRepo interface{
  CreateUser()
  GetUserByID(id string) *models.User
}

type userRepo struct{}

// func NewUserQuery() UserQuery {
//   return &userQuery{}
// }

func (u *userRepo) CreateUser() {
  // logic to create user in  db
  // where we actually use the sqlite3 database driver

}

func (u *userRepo) GetUserByID(id string) *models.User {
  return nil
}
