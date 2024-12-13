package user

// сущность "user"

type User struct {
	ID       int64
	Username string
	Password string
	Phone    string
	Email    string
}

// функция конструктор для создания сущности

func NewUser(username string, password string, phone string, email string) *User {
	return &User{
		Username: username,
		Password: password,
		Phone:    phone,
		Email:    email,
	}
}
