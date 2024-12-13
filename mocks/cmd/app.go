package main

import (
	"fmt"
	"mocks/db"
	"mocks/service"
)

func main() {
	repo := db.NewRepo()
	usecase := service.NewUserUseCase(repo)
	user1, err := usecase.CreateUser(service.CreateUserRequest{
		Username: "Denis",
		Password: "123456",
		Phone:    "79995398037",
		Email:    "denis@gmail.com",
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(user1)
}
