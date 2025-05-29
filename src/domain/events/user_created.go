package events

import "time"

type UserCreated struct {
	UserID    string
	Email     string
	Name      string
	CreatedAt time.Time
}

func NewUserCreated(userID, email, name string) *UserCreated {
	return &UserCreated{
		UserID:    userID,
		Email:     email,
		Name:      name,
		CreatedAt: time.Now(),
	}
}
