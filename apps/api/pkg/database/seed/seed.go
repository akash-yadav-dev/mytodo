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
