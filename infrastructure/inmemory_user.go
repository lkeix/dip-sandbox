package infrastructure

import (
	"errors"
	"sync"

	"github.com/lkeix/dip-sandbox/domain/model"
	"github.com/lkeix/dip-sandbox/domain/repository"
)

type inMemoryUserAdapter struct {
	users map[int]*model.User
	mux   sync.RWMutex
}

func NewInmemoryUserAdapter() repository.User {
	return &inMemoryUserAdapter{
		users: make(map[int]*model.User),
		mux:   sync.RWMutex{},
	}
}

func (i *inMemoryUserAdapter) Users() []model.User {
	var users []model.User

	i.mux.RLock()
	for _, user := range i.users {
		users = append(users, *user)
	}
	i.mux.Unlock()

	return users
}

func (i *inMemoryUserAdapter) UserByID(id int) (*model.User, error) {
	i.mux.RLock()
	defer i.mux.RUnlock()

	user, ok := i.users[id]

	if !ok {
		return nil, errors.New("failed to find user")
	}

	return user, nil
}

func (i *inMemoryUserAdapter) Update(user *model.User) error {
	i.mux.Lock()
	defer i.mux.Unlock()
	_, ok := i.users[int(user.ID)]

	if !ok {
		return errors.New("user doesn't exist")
	}

	i.users[int(user.ID)] = user

	return nil
}

func (i *inMemoryUserAdapter) Create(user *model.User) error {
	i.mux.Lock()
	defer i.mux.Unlock()
	_, ok := i.users[int(user.ID)]
	if ok {
		return errors.New("user already exist")
	}

	i.users[int(user.ID)] = user
	return nil
}

func (i *inMemoryUserAdapter) Delete(id int) error {
	i.mux.Lock()
	defer i.mux.Unlock()
	_, ok := i.users[id]

	if !ok {
		return errors.New("user doesn't exist")
	}

	delete(i.users, id)

	return nil
}
