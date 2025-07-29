package events

import "time"

type CommerceCreated struct {
	CommerceID  string
	Name        string
	Description string
	OwnerID     string
	CreatedAt   time.Time
}

func NewCommerceCreated(commerceID, name, description, ownerID string) *CommerceCreated {
	return &CommerceCreated{
		CommerceID:  commerceID,
		Name:        name,
		Description: description,
		OwnerID:     ownerID,
		CreatedAt:   time.Now(),
	}
}
