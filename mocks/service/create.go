package service

import (
	"fmt"
	"mocks/user"
)

// вспомогательная дто

type CreateUserRequest struct {
	Username string
	Password string
	Phone    string
	Email    string
}

func (u *UserUseCase) CreateUser(req CreateUserRequest) (*user.User, error) {
	newUser := user.NewUser(req.Username, req.Password, req.Phone, req.Email)
	newUser, err := u.repo.CreateUser(newUser)
	if err != nil {
		return nil, fmt.Errorf("repo.CreateUser: %w", err)
	}

	return newUser, nil
}
