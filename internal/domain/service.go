package domain

import (
	"context"
	"fmt"
	"regexp"

	"github.com/google/uuid"
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
		ID *string `json:"id"`
	}
	GetUserResponse struct {
		User User `json:"user"`
	}
	CreateUserInput struct {
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}
	UpdateUserInput struct {
		ID        string  `json"id"`
		Email     *string `json:"email"`
		FirstName *string `json:"first_name"`
		LastName  *string `json:"last_name"`
	}
)

type Service interface {
	GetUsersMany(context.Context, GetUsersInput) (*GetUsersResponse, error)
	GetUsersOne(context.Context, GetUserInput) (*GetUserResponse, error)
	CreateUser(context.Context, CreateUserInput) (*GetUserResponse, error)
	UpdateUser(context.Context, UpdateUserInput) (*GetUserResponse, error)
}

func (v *GetUsersInput) Validate() error {
	if v.Limit < 1 {
		return fmt.Errorf("limit must be greater than 0")
	}
	if v.Offset < 0 {
		return fmt.Errorf("offset must be greater than or equal to 0")
	}
	if v.Query != nil && *v.Query == "" {
		return fmt.Errorf("query cannot be an empty string")
	}
	if v.OrderBy != nil && *v.OrderBy == "" {
		return fmt.Errorf("order_by cannot be an empty string")
	}
	return nil
}

func (v *GetUserInput) Validate() error {
	if v.ID == nil {
		return fmt.Errorf("id must be provided")
	}

	if v.ID != nil {
		if len(*v.ID) != 36 {
			return fmt.Errorf("invalid id format, length must be 36 characters")
		}

		_, err := uuid.Parse(*v.ID)
		if err != nil {
			return fmt.Errorf("invalid id format: %v", err)
		}
	}

	return nil
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func (v *CreateUserInput) Validate() error {
	if v.Email == "" {
		return fmt.Errorf("email is required")
	}
	if !isValidEmail(v.Email) {
		return fmt.Errorf("invalid email format")
	}

	if v.FirstName == "" {
		return fmt.Errorf("first_name is required")
	}

	if v.LastName == "" {
		return fmt.Errorf("last_name is required")
	}

	return nil
}

func (v *UpdateUserInput) Validate() error {
	return nil
}
