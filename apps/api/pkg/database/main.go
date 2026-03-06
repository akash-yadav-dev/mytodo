package database

import (
	"mytodo/apps/api/pkg/database/postgres"
)

type Database interface {
	Connect() error
	Close() error
}

type PostgresDB struct {
	Connection postgres.Connection
}

func NewDb(connection postgres.Connection) *PostgresDB {
	return &PostgresDB{
		Connection: connection,
	}
}
func (p *PostgresDB) Connect() error {
	// Implement connection logic here
	return nil
}
