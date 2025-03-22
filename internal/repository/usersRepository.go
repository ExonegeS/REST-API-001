package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/ExonegeS/REST-API-001/internal/domain"
	"github.com/lib/pq"
)

type UsersRepository interface {
	GetUsersList(ctx context.Context, input domain.GetUsersInput) (*domain.List[domain.User], error)
	GetUsersOne(ctx context.Context, input domain.GetUserInput) (*domain.User, error)
	InsertUser(ctx context.Context, user *domain.User) (*domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error)
}

type usersRepository struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *usersRepository {
	return &usersRepository{
		db: db,
	}
}

func (u *usersRepository) GetUsersList(ctx context.Context, input domain.GetUsersInput) (*domain.List[domain.User], error) {
	rows, err := u.db.QueryContext(ctx, "select * from users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	usersList := domain.List[domain.User]{}

	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, fmt.Errorf("error occured while scanning rows in 'users': %s", err)
		}
		usersList.Elements = append(usersList.Elements, user)
	}

	err = u.db.QueryRow("SELECT COUNT(*) FROM users").Scan(&usersList.Total)
	if err != nil {
		return nil, err
	}

	return &usersList, nil
}

func (u *usersRepository) GetUsersOne(ctx context.Context, input domain.GetUserInput) (*domain.User, error) {
	query := "SELECT id, email, first_name, last_name, created_at, updated_at FROM users"

	var args []interface{}

	conditions := []string{}

	if input.ID != nil {
		conditions = append(conditions, "id = $"+strconv.Itoa(len(args)+1))
		args = append(args, *input.ID)
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	row := u.db.QueryRowContext(ctx, query, args...)

	var user domain.User
	if err := row.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.CreatedAt, &user.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("error occurred while scanning row in 'users': %s", err)
	}

	return &user, nil
}

func (u *usersRepository) InsertUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	var existingUser domain.User
	err := u.db.QueryRowContext(ctx, "SELECT id, email, first_name, last_name, created_at, updated_at FROM users WHERE email = $1", user.Email).Scan(
		&existingUser.ID, &existingUser.Email, &existingUser.FirstName, &existingUser.LastName, &existingUser.CreatedAt, &existingUser.UpdatedAt,
	)
	if err != nil && err != sql.ErrNoRows {
		return nil, fmt.Errorf("error checking for existing email: %v", err)
	}
	if err == nil {
		return nil, fmt.Errorf("email already in use")
	}

	query := `INSERT INTO users (id, email, first_name, last_name, created_at, updated_at) 
			  VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err = u.db.QueryRowContext(ctx, query, user.ID, user.Email, user.FirstName, user.LastName, user.CreatedAt, user.UpdatedAt).Scan(&user.ID)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return nil, fmt.Errorf("email already in use")
		}
		return nil, fmt.Errorf("error creating user: %v", err)
	}

	return user, nil
}

func (u *usersRepository) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	var existingUser domain.User
	err := u.db.QueryRowContext(ctx, "SELECT id, email, first_name, last_name, created_at, updated_at FROM users WHERE id = $1", user.ID).Scan(
		&existingUser.ID, &existingUser.Email, &existingUser.FirstName, &existingUser.LastName, &existingUser.CreatedAt, &existingUser.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with id %s not found", user.ID)
		}
		return nil, fmt.Errorf("error checking for existing user: %v", err)
	}

	query := `UPDATE users
			  SET email = $1, first_name = $2, last_name = $3, updated_at = $4
			  WHERE id = $5
			  RETURNING id, email, first_name, last_name, created_at, updated_at`

	err = u.db.QueryRowContext(ctx, query, user.Email, user.FirstName, user.LastName, user.UpdatedAt, user.ID).Scan(
		&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return nil, fmt.Errorf("the email address is already in use. Please use a different email.")
		}
		return nil, fmt.Errorf("error updating user: %v", err)
	}

	return user, nil
}
