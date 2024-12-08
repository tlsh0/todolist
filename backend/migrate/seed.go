package migrate

import (
	"database/sql"
	"fmt"
)

func Seed(db *sql.DB) error {

	fmt.Println("seed succesful!")
	return nil
}

// func seedUser() {

// }

// func seedTask() {
// 	user := new(models.User)

// }
