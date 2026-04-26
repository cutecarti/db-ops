package service

import (
	"context"
	"fmt"
	"github.com/cutecarti/db-ops/internal/models"
)

type DBService struct {
	repo RecordRepository
}

func NewDBService(repo RecordRepository) *DBService {
	return &DBService{
		repo: repo,
	}
}

func (s *DBService) CreateRecord(ctx context.Context, record models.Record) (int, error) {

	id, err := s.repo.CreateRecord(ctx, record)
	if err != nil {
		return 0, fmt.Errorf("failed to make record: %w", err)
	}
	return id, nil
}

func (s *DBService) GetRecord(ctx context.Context, id int) (models.Record, error) {
	record, err := s.repo.GetRecord(ctx, id)
	if err != nil {
		return models.Record{}, err
	}

	return record, nil
}

func (s *DBService) DeleteRecord(ctx context.Context, id int) error {
	err := s.repo.DeleteRecord(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *DBService) UpdateRecord(ctx context.Context, id int, name string) error {
	err := s.repo.UpdateRecord(ctx, id, name)
	if err != nil {
		return err
	}
	return nil
}
