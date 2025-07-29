package memory

import (
	"context"
	"minhex/src/domain/commerces/entities"
	commercePorts "minhex/src/domain/commerces/ports"
	"sync"
)

type CommerceRepository struct {
	commerces map[string]*entities.Commerce
	mu        sync.RWMutex
}

func NewCommerceRepository() commercePorts.CommerceRepository {
	return &CommerceRepository{
		commerces: make(map[string]*entities.Commerce),
	}
}

func (r *CommerceRepository) Save(ctx context.Context, commerce *entities.Commerce) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.commerces[commerce.ID] = commerce
	return nil
}

func (r *CommerceRepository) FindByID(ctx context.Context, id string) (*entities.Commerce, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if commerce, exists := r.commerces[id]; exists {
		return commerce, nil
	}
	return nil, nil
}

func (r *CommerceRepository) FindByOwnerID(ctx context.Context, ownerID string) ([]*entities.Commerce, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []*entities.Commerce
	for _, commerce := range r.commerces {
		if commerce.OwnerID == ownerID {
			result = append(result, commerce)
		}
	}
	return result, nil
}

func (r *CommerceRepository) Update(ctx context.Context, commerce *entities.Commerce) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.commerces[commerce.ID]; !exists {
		return nil
	}
	r.commerces[commerce.ID] = commerce
	return nil
}
