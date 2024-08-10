package database

import (
	"MOA/config"
	"MOA/infrastructure/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type TransactionRepository struct {
	conn *pgx.Conn
}

func NewTransactionRepository(cfg *config.Config) (*TransactionRepository, error) {
	conn, err := pgx.Connect(context.Background(), cfg.Database.Host)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	return &TransactionRepository{conn: conn}, nil
}

func (r *TransactionRepository) Create(transaction *models.Transaction) error {
	query := `INSERT INTO transactions (transaction_type, amount, category_id) VALUES ($1, $2, $3)`
	_, err := r.conn.Exec(context.Background(), query, transaction.TransactionType, transaction.Amount, transaction.CategoryId)
	if err != nil {
		return fmt.Errorf("failed to create transaction: %v", err)
	}

	return nil
}

func (r *TransactionRepository) Read(id int32) (*models.Transaction, error) {
	query := `SELECT id, transaction_type, amount, category_id FROM transactions WHERE id = $1`
	row := r.conn.QueryRow(context.Background(), query, id)

	var transaction models.Transaction
	err := row.Scan(&transaction.Id, &transaction.TransactionType, &transaction.Amount, &transaction.CategoryId)
	if err != nil {
		return nil, fmt.Errorf("failed to read transaction: %v", err)
	}

	return &transaction, nil
}

func (r *TransactionRepository) Update(id int32, transaction *models.Transaction) error {
	query := `UPDATE transactions SET transaction_type = $1, amount = $2, category_id = $3 WHERE id = $4`
	_, err := r.conn.Exec(context.Background(), query, transaction.TransactionType, transaction.Amount, transaction.CategoryId, id)
	if err != nil {
		return fmt.Errorf("failed to update transaction: %v", err)
	}

	return nil
}

func (r *TransactionRepository) ReadAllByUser(userId int64) ([]*models.Transaction, error) {
	query := `SELECT id, transaction_type, amount, category_id FROM transactions WHERE user_id = $1 AND is_deleted = false`
	rows, err := r.conn.Query(context.Background(), query, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to read transactions: %v", err)
	}
	defer rows.Close()

	var transactions []*models.Transaction
	for rows.Next() {
		var transaction models.Transaction
		if err := rows.Scan(&transaction.Id, &transaction.TransactionType, &transaction.Amount, &transaction.CategoryId); err != nil {
			return nil, fmt.Errorf("failed to scan transaction: %v", err)
		}
		transactions = append(transactions, &transaction)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to read transactions: %v", err)
	}
	return transactions, nil
}

func (r *TransactionRepository) Close() {
	r.conn.Close(context.Background())
}
