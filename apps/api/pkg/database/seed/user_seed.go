package seed

import (
	"database/sql"
	"log"
)

func Run(db *sql.DB) error {

	log.Println("Running database seeds")

	err := seedUsers(db)
	if err != nil {
		return err
	}

	return nil
}

func seedUsers(db *sql.DB) error {

	query := `
	INSERT INTO users (name, email)
	VALUES ('Admin', 'admin@example.com')
	ON CONFLICT DO NOTHING
	`

	_, err := db.Exec(query)

	return err
}