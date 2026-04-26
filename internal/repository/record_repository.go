package repository

import (
	"context"
	"fmt"

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
	err := r.pool.QueryRow(ctx, "INSERT INTO records (name) VALUES ($1) RETURNING id", record.Name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *RecordRepo) GetRecord(ctx context.Context, id int) (models.Record, error) {
	var record models.Record
	err := r.pool.QueryRow(ctx, "SELECT id, name FROM records WHERE id = $1", id).Scan(&record.ID, &record.Name)
	if err != nil {
		if err == pgx.ErrNoRows {
			return models.Record{}, fmt.Errorf("Not found")
		}
		return models.Record{}, err
	}
	return record, nil
}

func (r *RecordRepo) DeleteRecord(ctx context.Context, id int) error {
	cmdTag, err := r.pool.Exec(ctx, "DELETE FROM records WHERE id = $1", id)
	if err != nil {

		return err
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("Not found")
	}
	return nil
}

func (r *RecordRepo) UpdateRecord(ctx context.Context, id int, name string) error {
	cmdTag, err := r.pool.Exec(ctx, "UPDATE records SET name = $1 WHERE id = $2", name, id)
	if err != nil {

		return err
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("Not found")
	}
	return nil
}
