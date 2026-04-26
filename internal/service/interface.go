package service

import (
	"context"

	"github.com/cutecarti/db-ops/internal/models"
)

type RecordRepository interface {
	CreateRecord(ctx context.Context, record models.Record) (int, error)
	GetRecord(ctx context.Context, id int) (models.Record, error)
	DeleteRecord(ctx context.Context, id int) error
	UpdateRecord(ctx context.Context, id int, name string) error
}
