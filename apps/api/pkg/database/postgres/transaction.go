package postgres

import (
	"context"
	"database/sql"
)

func (r *Repository) WithTransaction(
	ctx context.Context,
	fn func(tx *sql.Tx) error,
) error {

	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	err = fn(tx)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}