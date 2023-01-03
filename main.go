package main

import (
	"fmt"

	"github.com/lkeix/dip-sandbox/entity"
	"github.com/lkeix/dip-sandbox/infrastructure"
)

func main() {
	inmemoryAdapter := infrastructure.NewInmemoryUserAdapter()
	setup(inmemoryAdapter)

	user, _ := inmemoryAdapter.UserByID(1)

	fmt.Println(user)

	users := inmemoryAdapter.Users()

	for _, u := range users {
		fmt.Println(u)
	}
}

func setup(adapter infrastructure.User) {
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
		adapter.Create(user)
	}
}
