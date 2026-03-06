package postgres

import (
	"context"
	"database/sql"
)

type UnitOfWork struct {
	DB *sql.DB
}

func NewUnitOfWork(db *sql.DB) *UnitOfWork {
	return &UnitOfWork{DB: db}
}

func (u *UnitOfWork) WithinTransaction(ctx context.Context, fn func(tx *sql.Tx) error) error {
	tx, err := u.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if err := fn(tx); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
