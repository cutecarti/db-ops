package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/cutecarti/db-ops/internal/models"
	"github.com/cutecarti/db-ops/internal/service"
)

type fakeRecordRepo struct {
	createRecordFunc func(ctx context.Context, record models.Record) (int, error)
}

func (f *fakeRecordRepo) CreateRecord(ctx context.Context, record models.Record) (int, error) {
	return f.createRecordFunc(ctx, record)
}

func (f *fakeRecordRepo) GetRecord(ctx context.Context, id int) (models.Record, error) {
	return models.Record{}, nil
}

func (f *fakeRecordRepo) DeleteRecord(ctx context.Context, id int) error {
	return nil
}

func (f *fakeRecordRepo) UpdateRecord(ctx context.Context, id int, name string) error {
	return nil
}

func TestDBService_CreateRecord(t *testing.T) {
	tests := []struct {
		name      string
		record    models.Record
		repoID    int
		repoErr   error
		wantID    int
		wantError bool
	}{
		{
			name:   "success",
			record: models.Record{Name: "test"},
			repoID: 1,
			wantID: 1,
		},
		{
			name:      "repo error",
			record:    models.Record{Name: "test"},
			repoErr:   errors.New("db error"),
			wantID:    0,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &fakeRecordRepo{
				createRecordFunc: func(ctx context.Context, record models.Record) (int, error) {
					return tt.repoID, tt.repoErr
				},
			}

			s := service.NewDBService(repo)

			gotID, err := s.CreateRecord(context.Background(), tt.record)

			if tt.wantError && err == nil {
				t.Fatal("expected error, got nil")
			}

			if !tt.wantError && err != nil {
				t.Fatalf("expected nil error, got %v", err)
			}

			if gotID != tt.wantID {
				t.Fatalf("expected id %d, got %d", tt.wantID, gotID)
			}
		})
	}
}
