package domain

import (
	"context"
)

type List[T any] struct {
	Elements []T
	Total    int64
}

type (
	GetUsersInput struct {
		Query   *string
		OrderBy *string
		Limit   int
		Offset  int
	}
	GetUsersResponse struct {
		Users []User `json:"users"`
		Total int64  `json:"total"`
	}
	GetUserInput struct {
		ID    *string `json:"id"`
		Email *string `json:"email"`
	}
	GetUserResponse struct {
		User User `json:"user"`
	}
)

type Service interface {
	GetUsersMany(context.Context, GetUsersInput) (*GetUsersResponse, error)
	GetUsersOne(context.Context, GetUserInput) (*GetUserResponse, error)
}

func (v *GetUsersInput) Validate() error {
	return nil
}

func (v *GetUserInput) Validate() error {
	return nil
}
