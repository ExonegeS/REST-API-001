package service

import (
	"context"
	"fmt"

	"github.com/ExonegeS/REST-API-001/internal/domain"
	"github.com/ExonegeS/REST-API-001/internal/usecase"
)

type UsersService struct {
	UseCase usecase.UsersUseCase
}

func NewUsersService(usecase usecase.UsersUseCase) domain.Service {
	return &UsersService{
		UseCase: usecase,
	}
}

func (u UsersService) GetUsersMany(ctx context.Context, input domain.GetUsersInput) (*domain.GetUsersResponse, error) {
	if err := input.Validate(); err != nil {
		return nil, fmt.Errorf("GetUsersInput validation error: %v", err)
	}

	usersList, err := u.UseCase.GetUsersList(ctx, input)
	if err != nil {
		return nil, err
	}
	return &domain.GetUsersResponse{
		Users: usersList.Elements,
		Total: usersList.Total,
	}, nil
}

func (u UsersService) GetUsersOne(ctx context.Context, input domain.GetUserInput) (*domain.GetUserResponse, error) {
	if err := input.Validate(); err != nil {
		return nil, fmt.Errorf("GetUserInput validation error: %v", err)
	}

	user, err := u.UseCase.GetUsersOne(ctx, input)
	if err != nil {
		return nil, err
	}
	return &domain.GetUserResponse{
		User: *user,
	}, nil
}
