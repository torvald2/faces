package adaptors

import (
	"database/sql"
)

type Store struct {
	db *sql.DB
}

func createTables(db *sql.DB) {
	tables := map[string]string{
		"profiles":    "CREATE TABLE IF NOT EXISTS profiles (id SERIAL PRIMARY KEY ,descriptor double precision[] not null, name text not null,image_id int NOT NULL,shop_num int NOT NULL, created_date time not null DEFAULT NOW())",
		"pictures":    "CREATE TABLE IF NOT EXISTS pictures (id SERIAL PRIMARY KEY ,profile_id INT, data bytea NOT NULL)",
		"workjornal":  "CREATE TABLE IF NOT EXISTS workjornal (id SERIAL  PRIMARY KEY, profile_id INT NOT NULL, operation_date TIME NOT NULL, created_date time not null DEFAULT NOW())",
		"badRequests": "CREATE TABLE IF NOT EXISTS badrequest (id SERIAL  PRIMARY KEY, profile_id INT, recognized_profiles INT[], current_face bytea, error_type INT,recognized_time time, created_date time not null DEFAULT NOW() )",
	}
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()
	for _, v := range tables {
		_, err := tx.Exec(v)
		if err != nil {
			panic(err)
		}
	}
	if err = tx.Commit(); err != nil {
		panic(err)
	}
}
