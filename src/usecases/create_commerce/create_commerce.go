package create_commerce

import (
	"context"
	"time"

	"minhex/src/domain/commerces/entities"
	"minhex/src/domain/commerces/errors"
	"minhex/src/domain/commerces/events"
	commercePorts "minhex/src/domain/commerces/ports"
	"minhex/src/domain/shared/ports"
	userErrors "minhex/src/domain/users/errors"
	userPorts "minhex/src/domain/users/ports"
)

type CreateCommerceUseCase interface {
	Execute(request Request) (*Response, error)
}

type Request struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	OwnerID     string `json:"owner_id"`
}

type Response struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	OwnerID     string `json:"owner_id"`
	Status      string `json:"status"`
}

type useCase struct {
	commerceRepo   commercePorts.CommerceRepository
	userRepo       userPorts.UserRepository
	eventPublisher ports.EventPublisher
}

func New(commerceRepo commercePorts.CommerceRepository, userRepo userPorts.UserRepository, eventPublisher ports.EventPublisher) CreateCommerceUseCase {
	return &useCase{
		commerceRepo:   commerceRepo,
		userRepo:       userRepo,
		eventPublisher: eventPublisher,
	}
}

func (uc *useCase) Execute(request Request) (*Response, error) {
	ctx := context.Background()

	if err := uc.validateRequest(request); err != nil {
		return nil, err
	}

	owner, err := uc.userRepo.FindByID(ctx, request.OwnerID)
	if err != nil {
		return nil, err
	}
	if owner == nil {
		return nil, userErrors.ErrUserNotFound
	}

	commerce := entities.NewCommerce(request.Name, request.Description, request.OwnerID)

	if err := uc.commerceRepo.Save(ctx, commerce); err != nil {
		return nil, err
	}

	event := events.CommerceCreated{
		CommerceID:  commerce.ID,
		Name:        commerce.Name,
		Description: commerce.Description,
		OwnerID:     commerce.OwnerID,
		CreatedAt:   time.Now(),
	}

	if err := uc.eventPublisher.Publish(ctx, "commerce.created", event); err != nil {
		return nil, err
	}

	return &Response{
		ID:          commerce.ID,
		Name:        commerce.Name,
		Description: commerce.Description,
		OwnerID:     commerce.OwnerID,
		Status:      string(commerce.Status),
	}, nil
}

func (uc *useCase) validateRequest(request Request) error {
	if request.Name == "" || len(request.Name) < 2 {
		return errors.ErrInvalidCommerceName
	}

	if request.OwnerID == "" {
		return errors.ErrInvalidOwner
	}

	return nil
}
