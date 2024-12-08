package migrate

import (
	"database/sql"
	"fmt"
)

func Up(db *sql.DB) error {
	query := `
	create table if not exists users (
		id serial primary key,
		username varchar(50) not null,
		email varchar(50) not null unique,
		encrypted_password text not null
	);

	create table if not exists tasks (
		id serial primary key,
		user_id int not null,
		name varchar(50) not null,
		due_time timestamp not null,
		done boolean not null default false,
		foreign key (user_id) references users(id) on delete cascade
	);`

	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	fmt.Println("migration up!")
	return nil
}
