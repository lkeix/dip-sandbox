package repository

import (
	"github.com/lkeix/dip-sandbox/domain/model"
)

type User interface {
	Users() []model.User               // fetch all users
	UserByID(int) (*model.User, error) // fetch user by id
	Update(*model.User) error          // update user info by id
	Create(*model.User) error          // create new user
	Delete(int) error
}
