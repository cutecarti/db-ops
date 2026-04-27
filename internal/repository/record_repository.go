package repository

import (
	"context"

	"github.com/cutecarti/db-ops/internal/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RecordRepo struct {
	pool *pgxpool.Pool
}

func NewRecordRepo(pool *pgxpool.Pool) *RecordRepo {
	return &RecordRepo{
		pool: pool,
	}
}

func (r *RecordRepo) CreateRecord(ctx context.Context, record models.Record) (int, error) {
	var id int
	err := r.pool.QueryRow(ctx, CREATE_QUERRY, record.Name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *RecordRepo) GetRecord(ctx context.Context, id int) (models.Record, error) {
	var record models.Record
	err := r.pool.QueryRow(ctx, GET_QUERRY, id).Scan(&record.ID, &record.Name)
	if err != nil {
		if err == pgx.ErrNoRows {
			return models.Record{}, ErrRecordNotFound
		}
		return models.Record{}, err
	}
	return record, nil
}

func (r *RecordRepo) DeleteRecord(ctx context.Context, id int) error {
	cmdTag, err := r.pool.Exec(ctx, DELETE_QUERRY, id)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() == 0 {
		return ErrRecordNotFound
	}
	return nil
}

func (r *RecordRepo) UpdateRecord(ctx context.Context, id int, name string) error {
	cmdTag, err := r.pool.Exec(ctx, UPDATE_QUERRY, name, id)
	if err != nil {

		return err
	}
	if cmdTag.RowsAffected() == 0 {
		return ErrRecordNotFound
	}
	return nil
}
