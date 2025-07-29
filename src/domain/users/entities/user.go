package entities

import "time"

type User struct {
	ID        string
	Email     string
	Name      string
	CreatedAt time.Time
}

func NewUser(email, name string) *User {
	return &User{
		ID:        generateID(),
		Email:     email,
		Name:      name,
		CreatedAt: time.Now(),
	}
}

func generateID() string {
	return "user_" + time.Now().Format("20060102150405")
}
