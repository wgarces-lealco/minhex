package activate_commerce

import (
	"context"
	"time"

	"minhex/src/domain/commerces/errors"
	"minhex/src/domain/commerces/events"
	commercePorts "minhex/src/domain/commerces/ports"
	"minhex/src/domain/shared/ports"
)

type ActivateCommerceUseCase interface {
	Execute(request Request) (*Response, error)
}

type Request struct {
	CommerceID string `json:"commerce_id"`
}

type Response struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

type useCase struct {
	commerceRepo   commercePorts.CommerceRepository
	eventPublisher ports.EventPublisher
}

func New(commerceRepo commercePorts.CommerceRepository, eventPublisher ports.EventPublisher) ActivateCommerceUseCase {
	return &useCase{
		commerceRepo:   commerceRepo,
		eventPublisher: eventPublisher,
	}
}

func (uc *useCase) Execute(request Request) (*Response, error) {
	ctx := context.Background()

	commerce, err := uc.commerceRepo.FindByID(ctx, request.CommerceID)
	if err != nil {
		return nil, err
	}

	if commerce == nil {
		return nil, errors.ErrCommerceNotFound
	}

	commerce.Activate()

	if err := uc.commerceRepo.Update(ctx, commerce); err != nil {
		return nil, err
	}

	event := events.CommerceActivated{
		CommerceID: commerce.ID,
		OwnerID:    commerce.OwnerID,
		CreatedAt:  time.Now(),
	}

	if err := uc.eventPublisher.Publish(ctx, "commerce.activated", event); err != nil {
		return nil, err
	}

	return &Response{
		ID:     commerce.ID,
		Status: string(commerce.Status),
	}, nil
}
