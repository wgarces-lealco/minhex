package postgres

import (
	"context"
	"log"
	"minhex/src/domain/commerces/entities"
	commercePorts "minhex/src/domain/commerces/ports"
)

type CommerceRepository struct {
	config *Config
}

func NewCommerceRepository() commercePorts.CommerceRepository {
	cfg := &Config{
		DatabaseURL: getEnv("POSTGRES_URL", "postgres://user:pass@localhost/minhex"),
	}

	return &CommerceRepository{
		config: cfg,
	}
}

func (r *CommerceRepository) Save(ctx context.Context, commerce *entities.Commerce) error {
	log.Printf("[PostgreSQL] Saving commerce to %s", r.config.DatabaseURL)
	log.Printf("[PostgreSQL] Commerce: %+v", commerce)
	return nil
}

func (r *CommerceRepository) FindByID(ctx context.Context, id string) (*entities.Commerce, error) {
	log.Printf("[PostgreSQL] Finding commerce by ID from %s", r.config.DatabaseURL)
	log.Printf("[PostgreSQL] ID: %s", id)
	return nil, nil
}

func (r *CommerceRepository) FindByOwnerID(ctx context.Context, ownerID string) ([]*entities.Commerce, error) {
	log.Printf("[PostgreSQL] Finding commerces by owner from %s", r.config.DatabaseURL)
	log.Printf("[PostgreSQL] OwnerID: %s", ownerID)
	return nil, nil
}

func (r *CommerceRepository) Update(ctx context.Context, commerce *entities.Commerce) error {
	log.Printf("[PostgreSQL] Updating commerce in %s", r.config.DatabaseURL)
	log.Printf("[PostgreSQL] Commerce: %+v", commerce)
	return nil
}
