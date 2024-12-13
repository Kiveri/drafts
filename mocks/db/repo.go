package db

import (
	"mocks/user"
)

// репозиторий в котором храним сущности

type Repo struct {
	users  map[int64]*user.User
	nextID int64
}

// функция конструктор для репозитория

func NewRepo() *Repo {
	return &Repo{
		users:  make(map[int64]*user.User),
		nextID: 1,
	}
}

// функция создания id для сущности

func (r *Repo) getNextID() int64 {
	nextID := r.nextID
	r.nextID++

	return nextID
}
