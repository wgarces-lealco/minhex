package events

import "time"

type CommerceActivated struct {
	CommerceID string
	OwnerID    string
	CreatedAt  time.Time
}

func NewCommerceActivated(commerceID, ownerID string) *CommerceActivated {
	return &CommerceActivated{
		CommerceID: commerceID,
		OwnerID:    ownerID,
		CreatedAt:  time.Now(),
	}
}
