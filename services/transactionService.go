package services

import (
	"fmt"
	"log/slog"

	"MOA/config"
	"MOA/infrastructure/database"
	"MOA/infrastructure/models"
)

type TransactionService struct {
	repo   *database.TransactionRepository
	logger *slog.Logger
}

func NewTransactionService(cfg *config.Config, logger *slog.Logger) (*TransactionService, error) {
	repo, err := database.NewTransactionRepository(cfg)
	if err != nil {
		logger.Error("failed to create transaction repository: %v", err)
		return nil, fmt.Errorf("failed to create transaction repository: %v", err)
	}

	return &TransactionService{repo: repo, logger: logger}, nil
}

func (s *TransactionService) Create(transaction *models.Transaction) error {
	return s.repo.Create(transaction)
}

func (s *TransactionService) Read(id int32) (*models.Transaction, error) {
	return s.repo.Read(id)
}

func (s *TransactionService) Update(id int32, transaction *models.Transaction) error {
	return s.repo.Update(id, transaction)
}

func (s *TransactionService) ReadAllByUser(userId int64) ([]*models.Transaction, error) {
	return s.repo.ReadAllByUser(userId)
}

func (s *TransactionService) Close() {
	s.repo.Close()
}
