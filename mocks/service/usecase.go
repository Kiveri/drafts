package service

import "mocks/user"

// создание интерфеса с методом из репозитория

type repo interface {
	CreateUser(user *user.User) (*user.User, error)
}

// подключение репозитория

type UserUseCase struct {
	repo repo
}

// функция конструктор для нового usecase

func NewUserUseCase(repo repo) *UserUseCase {
	return &UserUseCase{
		repo: repo,
	}
}
