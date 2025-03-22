package usecase

import (
	"context"

	"github.com/ExonegeS/REST-API-001/internal/domain"
	"github.com/ExonegeS/REST-API-001/internal/repository"
)

type UsersUseCase interface {
	GetUsersList(ctx context.Context, input domain.GetUsersInput) (*domain.List[domain.User], error)
	GetUsersOne(ctx context.Context, input domain.GetUserInput) (*domain.User, error)
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error)
}

type usersUseCase struct {
	usersRepository repository.UsersRepository
}

func NewUsersUseCase(repository repository.UsersRepository) *usersUseCase {
	return &usersUseCase{
		usersRepository: repository,
	}
}

func (u *usersUseCase) GetUsersList(ctx context.Context, input domain.GetUsersInput) (*domain.List[domain.User], error) {
	return u.usersRepository.GetUsersList(ctx, input)
}

func (u *usersUseCase) GetUsersOne(ctx context.Context, input domain.GetUserInput) (*domain.User, error) {
	return u.usersRepository.GetUsersOne(ctx, input)
}

func (u *usersUseCase) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	return u.usersRepository.InsertUser(ctx, user)
}

func (u *usersUseCase) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	return u.usersRepository.UpdateUser(ctx, user)
}
