package memory

import (
	"context"
	"minhex/src/domain/users/entities"
	userPorts "minhex/src/domain/users/ports"
	"sync"
)

type UserRepository struct {
	users map[string]*entities.User
	mu    sync.RWMutex
}

func NewUserRepository() userPorts.UserRepository {
	return &UserRepository{
		users: make(map[string]*entities.User),
	}
}

func (r *UserRepository) Save(ctx context.Context, user *entities.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.users[user.ID] = user
	return nil
}

func (r *UserRepository) FindByID(ctx context.Context, id string) (*entities.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if user, exists := r.users[id]; exists {
		return user, nil
	}
	return nil, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*entities.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, nil
}
