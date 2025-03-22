package service

import (
	"context"
	"fmt"
	"time"

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
		return nil, fmt.Errorf("input validation error: %v", err)
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
		return nil, fmt.Errorf("input validation error: %v", err)
	}

	user, err := u.UseCase.GetUsersOne(ctx, input)
	if err != nil {
		return nil, err
	}
	return &domain.GetUserResponse{
		User: *user,
	}, nil
}

func (u *UsersService) CreateUser(ctx context.Context, input domain.CreateUserInput) (*domain.GetUserResponse, error) {
	if err := input.Validate(); err != nil {
		return nil, fmt.Errorf("input validation error: %v", err)
	}

	user, err := u.UseCase.CreateUser(ctx, &domain.User{
		Email:     input.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return nil, err
	}

	return &domain.GetUserResponse{
		User: *user,
	}, nil
}

func (u *UsersService) UpdateUser(ctx context.Context, input domain.UpdateUserInput) (*domain.GetUserResponse, error) {
	if err := input.Validate(); err != nil {
		return nil, fmt.Errorf("input validation error: %v", err)
	}

	user, err := u.UseCase.GetUsersOne(ctx, domain.GetUserInput{
		ID: &input.ID,
	})
	if err != nil {
		return nil, err
	}

	now := time.Now()
	if input.Email != nil {
		if user.Email != *input.Email {
			user.UpdatedAt = now
		}
		user.Email = *input.Email
	}
	if input.FirstName != nil {
		if user.FirstName != *input.FirstName {
			user.UpdatedAt = now
		}
		user.FirstName = *input.FirstName
	}
	if input.LastName != nil {
		if user.LastName != *input.LastName {
			user.UpdatedAt = now
		}
		user.LastName = *input.LastName
	}

	user, err = u.UseCase.UpdateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &domain.GetUserResponse{
		User: *user,
	}, nil
}
