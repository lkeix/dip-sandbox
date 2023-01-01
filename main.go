package main

import (
	"fmt"

	"github.com/lkeix/dip-sandbox/domain/model"
	"github.com/lkeix/dip-sandbox/domain/repository"
	"github.com/lkeix/dip-sandbox/infrastructure"
)

func main() {
	inmemoryAdapter := infrastructure.NewInmemoryUserAdapter()
	setup(inmemoryAdapter)

	userID1, err := inmemoryAdapter.UserByID(1)

	if err != nil {
		panic(err)
	}

	fmt.Println(userID1)
}

func setup(adapter repository.User) {
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
