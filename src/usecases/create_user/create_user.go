package create_user

import (
	"context"
	"time"

	"minhex/src/domain/shared/ports"
	"minhex/src/domain/users/entities"
	"minhex/src/domain/users/errors"
	"minhex/src/domain/users/events"
	userPorts "minhex/src/domain/users/ports"
)

type CreateUserUseCase interface {
	Execute(request Request) (*Response, error)
}

type Request struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Response struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type useCase struct {
	userRepo       userPorts.UserRepository
	eventPublisher ports.EventPublisher
}

func New(userRepo userPorts.UserRepository, eventPublisher ports.EventPublisher) CreateUserUseCase {
	return &useCase{
		userRepo:       userRepo,
		eventPublisher: eventPublisher,
	}
}

func (uc *useCase) Execute(request Request) (*Response, error) {
	ctx := context.Background()

	if err := uc.validateRequest(request); err != nil {
		return nil, err
	}

	if existingUser, _ := uc.userRepo.FindByEmail(ctx, request.Email); existingUser != nil {
		return nil, errors.ErrUserAlreadyExists
	}

	user := entities.NewUser(request.Email, request.Name)

	if err := uc.userRepo.Save(ctx, user); err != nil {
		return nil, err
	}

	event := events.UserCreated{
		UserID:    user.ID,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: time.Now(),
	}

	if err := uc.eventPublisher.Publish(ctx, "user.created", event); err != nil {
		return nil, err
	}

	return &Response{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}, nil
}

func (uc *useCase) validateRequest(request Request) error {
	if request.Email == "" || len(request.Email) < 3 {
		return errors.ErrInvalidEmail
	}

	if request.Name == "" || len(request.Name) < 2 {
		return errors.ErrInvalidName
	}

	return nil
}
