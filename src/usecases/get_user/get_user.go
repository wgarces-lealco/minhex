package get_user

import (
	"context"

	"minhex/src/domain/users/errors"
	userPorts "minhex/src/domain/users/ports"
)

type GetUserUseCase interface {
	Execute(request Request) (*Response, error)
}

type Request struct {
	ID string `json:"id"`
}

type Response struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type useCase struct {
	userRepo userPorts.UserRepository
}

func New(userRepo userPorts.UserRepository) GetUserUseCase {
	return &useCase{
		userRepo: userRepo,
	}
}

func (uc *useCase) Execute(request Request) (*Response, error) {
	ctx := context.Background()

	user, err := uc.userRepo.FindByID(ctx, request.ID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.ErrUserNotFound
	}

	return &Response{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}, nil
}
