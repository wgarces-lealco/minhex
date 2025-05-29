package create_user

import (
	"context"
	"minhex/src/domain/entities"
	"minhex/src/domain/ports"
)

type UseCase struct {
	userRepo ports.UserRepository
}

func New(userRepo ports.UserRepository) *UseCase {
	return &UseCase{
		userRepo: userRepo,
	}
}

type Request struct {
	Email string
	Name  string
}

type Response struct {
	UserID string
}

func (uc *UseCase) Execute(ctx context.Context, req Request) (*Response, error) {
	user := entities.NewUser(req.Email, req.Name)

	if err := uc.userRepo.Save(ctx, user); err != nil {
		return nil, err
	}

	return &Response{
		UserID: user.ID,
	}, nil
}
