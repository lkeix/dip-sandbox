package main

import (
	"fmt"

	"github.com/lkeix/dip-sandbox/domain/model"
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
	users := []*model.User{
		&model.User{
			ID:      1,
			Name:    "hoge",
			Mail:    "hoge@hoge.hoge",
			Address: "hogehoge",
		},
		&model.User{
			ID:      2,
			Name:    "fuga",
			Mail:    "fuga@fuga.fuga",
			Address: "fugafuga",
		},
		&model.User{
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
