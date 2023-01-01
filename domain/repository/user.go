package repository

import (
	"github.com/lkeix/dip-sandbox/domain/model"
)

type User interface {
	Users()                 // fetch all users
	UserByID(int)           // fetch user by id
	Update(int, model.User) // update user info by id
	Create(model.User)      // create new user
	Delete(int)
}
