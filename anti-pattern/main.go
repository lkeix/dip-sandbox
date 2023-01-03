package main

import (
	"fmt"

	"github.com/lkeix/dip-sandbox/anti-pattern/entity"
	"github.com/lkeix/dip-sandbox/anti-pattern/infrastructure"
)

func main() {
	store := infrastructure.NewInMemoryUserStore()
	setup(store)

	user, _ := store.UserByID(1)
	fmt.Println(user)

	for _, u := range store.Users {
		fmt.Println(u)
	}
}

func setup(store *infrastructure.InMemoryUserStore) {
	users := []*entity.User{
		&entity.User{
			ID:      1,
			Name:    "hoge",
			Mail:    "hoge@hoge.hoge",
			Address: "hogehoge",
		},
		&entity.User{
			ID:      2,
			Name:    "fuga",
			Mail:    "fuga@fuga.fuga",
			Address: "fugafuga",
		},
		&entity.User{
			ID:      3,
			Name:    "piyo",
			Mail:    "piyo@piyo.piyo",
			Address: "piyopiyo",
		},
	}

	for _, user := range users {
		store.Create(user)
	}
}
