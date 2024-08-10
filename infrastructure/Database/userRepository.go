package database

import (
	"MOA/config"
	"MOA/infrastructure/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	conn *pgx.Conn
}

func NewUserRepository(cfg *config.Config) (*UserRepository, error) {
	conn, err := pgx.Connect(context.Background(), cfg.Database.Host)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	return &UserRepository{conn: conn}, nil
}

func (r *UserRepository) Create(user *models.User) error {
	query := `INSERT INTO users (name) VALUES ($1)`
	_, err := r.conn.Exec(context.Background(), query, user.Name)
	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}

	return nil
}

func (r *UserRepository) Read(id int64) (*models.User, error) {
	query := `SELECT id, name FROM users WHERE id = $1`
	row := r.conn.QueryRow(context.Background(), query, id)

	var user models.User
	err := row.Scan(&user.Id, &user.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to read user: %v", err)
	}

	return &user, nil
}

func (r *UserRepository) Update(id int64, user *models.User) error {
	query := `UPDATE users SET name = $1 WHERE id = $3`
	_, err := r.conn.Exec(context.Background(), query, user.Name, id)
	if err != nil {
		return fmt.Errorf("failed to update user: %v", err)
	}

	return nil
}

func (r *UserRepository) Delete(id int64) error {
	query := `UPDATE users SET is_deleted = true WHERE id = $1`
	_, err := r.conn.Exec(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}

	return nil
}

func (r *UserRepository) ReadAll() ([]*models.User, error) {
	query := `SELECT id, name FROM users WHERE is_deleted = false`
	rows, err := r.conn.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("failed to read users: %v", err)
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Name); err != nil {
			return nil, fmt.Errorf("failed to scan user: %v", err)
		}
		users = append(users, &user)
	}

	return users, nil
}

func (r *UserRepository) Close() {
	r.conn.Close(context.Background())
}
