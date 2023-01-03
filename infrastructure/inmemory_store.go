package infrastructure

import (
	"errors"
	"sync"

	"github.com/lkeix/dip-sandbox/entity"
)

type inMemoryUserAdapter struct {
	users map[int]*entity.User
	Mux   sync.RWMutex
}

type User interface {
	Users() []*entity.User              // fetch all users
	UserByID(int) (*entity.User, error) // fetch user by id
	Create(*entity.User) error          // create new user
}

func NewInmemoryUserAdapter() User {
	return &inMemoryUserAdapter{
		users: make(map[int]*entity.User),
		Mux:   sync.RWMutex{},
	}
}

func (i *inMemoryUserAdapter) Users() []*entity.User {
	var users []*entity.User

	i.Mux.RLock()
	for _, user := range i.users {
		users = append(users, user)
	}
	i.Mux.Unlock()

	return users
}

func (i *inMemoryUserAdapter) UserByID(id int) (*entity.User, error) {
	i.Mux.RLock()
	defer i.Mux.RUnlock()

	user, ok := i.users[id]

	if !ok {
		return nil, errors.New("failed to find user")
	}

	return user, nil
}

func (i *inMemoryUserAdapter) Create(user *entity.User) error {
	i.Mux.Lock()
	defer i.Mux.Unlock()
	_, ok := i.users[int(user.ID)]
	if ok {
		return errors.New("user already exist")
	}

	i.users[int(user.ID)] = user
	return nil
}
