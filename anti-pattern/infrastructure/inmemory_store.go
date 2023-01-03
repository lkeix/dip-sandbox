package infrastructure

import (
	"errors"
	"sync"

	"github.com/lkeix/dip-sandbox/anti-pattern/entity"
)

type InMemoryUserStore struct {
	users map[int]*entity.User
	mux   sync.RWMutex
}

func NewInMemoryUserStore() *InMemoryUserStore {
	return &InMemoryUserStore{
		users: make(map[int]*entity.User),
		mux:   sync.RWMutex{},
	}
}
func (i *InMemoryUserStore) UserByID(id int) (*entity.User, error) {
	i.mux.RLock()
	defer i.mux.RUnlock()

	user, ok := i.users[id]

	if !ok {
		return nil, errors.New("failed to find user")
	}

	return user, nil
}

func (i *InMemoryUserStore) Create(user *entity.User) error {
	i.mux.Lock()
	defer i.mux.Unlock()
	_, ok := i.users[int(user.ID)]
	if ok {
		return errors.New("user already exist")
	}

	i.users[int(user.ID)] = user
	return nil
}
