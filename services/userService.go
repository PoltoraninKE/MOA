package services

import (
	"fmt"
	"log/slog"

	"MOA/config"
	"MOA/infrastructure/database"
	"MOA/infrastructure/models"
)

type UserService struct {
	repo   *database.UserRepository
	logger *slog.Logger
}

func NewUserService(cfg *config.Config, logger *slog.Logger) (*UserService, error) {
	repo, err := database.NewUserRepository(cfg)
	if err != nil {
		logger.Error("failed to create user repository: %v", err)
		return nil, fmt.Errorf("failed to create user repository: %v", err)
	}

	return &UserService{repo: repo, logger: logger}, nil
}

func (s *UserService) Create(user *models.User) error {
	return s.repo.Create(user)
}

func (s *UserService) Read(id int64) (*models.User, error) {
	return s.repo.Read(id)
}

func (s *UserService) Update(id int64, user *models.User) error {
	return s.repo.Update(id, user)
}

func (s *UserService) Delete(id int64) error {
	return s.repo.Delete(id)
}

func (s *UserService) ReadAll() ([]*models.User, error) {
	return s.repo.ReadAll()
}

func (s *UserService) Close() {
	s.repo.Close()
}
