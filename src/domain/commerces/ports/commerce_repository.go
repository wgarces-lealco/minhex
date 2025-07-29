package ports

import (
	"context"
	"minhex/src/domain/commerces/entities"
)

type CommerceRepository interface {
	Save(ctx context.Context, commerce *entities.Commerce) error
	FindByID(ctx context.Context, id string) (*entities.Commerce, error)
	FindByOwnerID(ctx context.Context, ownerID string) ([]*entities.Commerce, error)
	Update(ctx context.Context, commerce *entities.Commerce) error
}
