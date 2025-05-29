package ports

import (
	"context"
	"minhex/src/domain/entities"
)

type UserRepository interface {
	Save(ctx context.Context, user *entities.User) error
	FindByID(ctx context.Context, id string) (*entities.User, error)
	FindByEmail(ctx context.Context, email string) (*entities.User, error)
}
