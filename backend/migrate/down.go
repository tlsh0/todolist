package migrate

import (
	"database/sql"
	"fmt"
)

func Down(db *sql.DB) error {
	query := `drop table users; drop table tasks`

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	fmt.Println("migration down!")
	return nil
}
