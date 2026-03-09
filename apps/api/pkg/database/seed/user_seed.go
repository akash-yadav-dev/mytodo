package seed

import (
	"database/sql"
)

func seedUsers(db *sql.DB) error {

	query := `
	INSERT INTO users (name, email)
	VALUES ('Admin', 'admin@example.com')
	ON CONFLICT DO NOTHING
	`

	_, err := db.Exec(query)

	return err
}
