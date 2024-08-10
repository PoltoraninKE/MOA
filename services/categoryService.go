package services

import (
	"fmt"
	"log/slog"

	"MOA/config"
	"MOA/infrastructure/database"
	"MOA/infrastructure/models"
)

type CategoryService struct {
	repo   *database.CategoryRepository
	logger *slog.Logger
}

func NewCategoryService(cfg *config.Config, logger *slog.Logger) (*CategoryService, error) {
	repo, err := database.NewCategoryRepository(cfg)
	if err != nil {
		logger.Error("failed to create category repository: %v", err)
		return nil, fmt.Errorf("failed to create category repository: %v", err)
	}

	return &CategoryService{repo: repo, logger: logger}, nil
}

func (s *CategoryService) Create(category *models.Category) error {
	return s.repo.Create(category)
}

func (s *CategoryService) Read(id int64) (*models.Category, error) {
	return s.repo.Read(id)
}

func (s *CategoryService) Update(id int64, category *models.Category) error {
	return s.repo.Update(id, category)
}

func (s *CategoryService) Delete(id int64) error {
	return s.repo.Delete(id)
}

func (s *CategoryService) ReadAllByUser(userId int64) ([]*models.Category, error) {
	return s.repo.ReadAllByUser(userId)
}

func (s *CategoryService) Close() {
	s.repo.Close()
}
