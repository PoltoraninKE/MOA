package database

import (
	"MOA/config"
	"MOA/infrastructure/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type CategoryRepository struct {
	conn *pgx.Conn
}

func NewCategoryRepository(cfg *config.Config) (*CategoryRepository, error) {
	conn, err := pgx.Connect(context.Background(), cfg.Database.Host)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	return &CategoryRepository{conn: conn}, nil
}

func (r *CategoryRepository) Create(category *models.Category) error {
	query := `INSERT INTO categories (name, user_id) VALUES ($1)`
	_, err := r.conn.Exec(context.Background(), query, category.Name, category.UserId)
	if err != nil {
		return fmt.Errorf("failed to create category: %v", err)
	}

	return nil
}

func (r *CategoryRepository) Read(id int64) (*models.Category, error) {
	query := `SELECT id, name, user_id FROM categories WHERE id = $1`
	row := r.conn.QueryRow(context.Background(), query, id)

	var category models.Category
	err := row.Scan(&category.Id, &category.Name, &category.UserId)
	if err != nil {
		return nil, fmt.Errorf("failed to read category: %v", err)
	}
	return &category, nil
}

func (r *CategoryRepository) Update(id int64, category *models.Category) error {
	query := `UPDATE categories SET name = $1 WHERE id = $2`
	_, err := r.conn.Exec(context.Background(), query, category.Name, id)
	if err != nil {
		return fmt.Errorf("failed to update category: %v", err)
	}

	return nil
}

func (r *CategoryRepository) Delete(id int64) error {
	query := `UPDATE categories SET is_deleted = true WHERE id = $1`
	_, err := r.conn.Exec(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("failed to delete category: %v", err)
	}

	return nil
}

func (r *CategoryRepository) ReadAllByUser(userId int64) ([]*models.Category, error) {
	query := `SELECT id, name, user_id FROM categories WHERE user_id = $1 AND is_deleted = false`
	rows, err := r.conn.Query(context.Background(), query, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to read categories: %v", err)
	}
	defer rows.Close()

	var categories []*models.Category
	for rows.Next() {
		var category models.Category
		if err := rows.Scan(&category.Id, &category.Name, &category.UserId); err != nil {
			return nil, fmt.Errorf("failed to scan category: %v", err)
		}
		categories = append(categories, &category)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to read categories: %v", err)
	}
	return categories, nil
}

func (r *CategoryRepository) Close() {
	r.conn.Close(context.Background())
}
