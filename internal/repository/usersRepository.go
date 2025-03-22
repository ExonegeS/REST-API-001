package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"strconv"
	"strings"

	"github.com/ExonegeS/REST-API-001/internal/domain"
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
	// Start building the base SQL query
	query := "SELECT id, email, first_name, last_name, created_at, updated_at FROM users"

	// To hold query parameters (values)
	var args []interface{}

	// Build WHERE clauses based on non-nil fields in the input
	conditions := []string{}

	if input.ID != nil {
		conditions = append(conditions, "id = $"+strconv.Itoa(len(args)+1))
		args = append(args, *input.ID)
	}

	// If any conditions are added, append them to the query
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	// Log the final query and args
	slog.Info("Running query", "query", query, "args", args)

	// Execute the query
	row := u.db.QueryRowContext(ctx, query, args...)

	var user domain.User
	if err := row.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.CreatedAt, &user.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			// No user found, log the ID and return nil
			slog.Info("No user found with ID", "id", *input.ID)
			return nil, nil
		}
		return nil, fmt.Errorf("error occurred while scanning row in 'users': %s", err)
	}

	// Return the user if found
	return &user, nil
}

func (u *usersRepository) InsertUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	query := `
        INSERT INTO users (email, first_name, last_name, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id, email, first_name, last_name, created_at, updated_at
    `

	err := u.db.QueryRowContext(ctx, query, user.Email, user.FirstName, user.LastName, user.CreatedAt, user.UpdatedAt).
		Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("error occurred while inserting new user: %w", err)
	}

	return user, nil
}

func (u *usersRepository) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	return user, nil
}
