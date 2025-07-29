package entities

import "time"

type Commerce struct {
	ID          string
	Name        string
	Description string
	OwnerID     string
	Status      CommerceStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type CommerceStatus string

const (
	StatusPending   CommerceStatus = "pending"
	StatusActive    CommerceStatus = "active"
	StatusSuspended CommerceStatus = "suspended"
)

func NewCommerce(name, description, ownerID string) *Commerce {
	return &Commerce{
		ID:          generateCommerceID(),
		Name:        name,
		Description: description,
		OwnerID:     ownerID,
		Status:      StatusPending,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (c *Commerce) Activate() {
	c.Status = StatusActive
	c.UpdatedAt = time.Now()
}

func (c *Commerce) Suspend() {
	c.Status = StatusSuspended
	c.UpdatedAt = time.Now()
}

func generateCommerceID() string {
	return "commerce_" + time.Now().Format("20060102150405")
}
