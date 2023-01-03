package infrastructure

import (
	"errors"
	"sync"

	"github.com/lkeix/dip-sandbox/anti-pattern/entity"
)

type InMemoryUserStore struct {
	Users map[int]*entity.User
	Mux   sync.RWMutex
}

func NewInMemoryUserStore() *InMemoryUserStore {
	return &InMemoryUserStore{
		Users: make(map[int]*entity.User),
		Mux:   sync.RWMutex{},
	}
}
func (i *InMemoryUserStore) UserByID(id int) (*entity.User, error) {
	i.Mux.RLock()
	defer i.Mux.RUnlock()

	user, ok := i.Users[id]

	if !ok {
		return nil, errors.New("failed to find user")
	}

	return user, nil
}

func (i *InMemoryUserStore) Create(user *entity.User) error {
	i.Mux.Lock()
	defer i.Mux.Unlock()
	_, ok := i.Users[int(user.ID)]
	if ok {
		return errors.New("user already exist")
	}

	i.Users[int(user.ID)] = user
	return nil
}
