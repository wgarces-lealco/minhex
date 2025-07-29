package postgres

import (
	"context"
	"log"
	"minhex/src/domain/users/entities"
	userPorts "minhex/src/domain/users/ports"
	"os"
)

type Config struct {
	DatabaseURL string
}

type UserRepository struct {
	config *Config
}

func NewUserRepository() userPorts.UserRepository {
	cfg := &Config{
		DatabaseURL: getEnv("POSTGRES_URL", "postgres://user:pass@localhost/minhex"),
	}

	return &UserRepository{
		config: cfg,
	}
}

func (r *UserRepository) Save(ctx context.Context, user *entities.User) error {
	log.Printf("[PostgreSQL] Saving user to %s", r.config.DatabaseURL)
	log.Printf("[PostgreSQL] User: %+v", user)
	return nil
}

func (r *UserRepository) FindByID(ctx context.Context, id string) (*entities.User, error) {
	log.Printf("[PostgreSQL] Finding user by ID from %s", r.config.DatabaseURL)
	log.Printf("[PostgreSQL] ID: %s", id)
	return nil, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*entities.User, error) {
	log.Printf("[PostgreSQL] Finding user by email from %s", r.config.DatabaseURL)
	log.Printf("[PostgreSQL] Email: %s", email)
	return nil, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
