package db

import "mocks/user"

func (r *Repo) CreateUser(user *user.User) (*user.User, error) {
	user.ID = r.getNextID()
	r.users[user.ID] = user

	return user, nil
}
