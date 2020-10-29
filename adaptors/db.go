package adaptors

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	"atbmarket.comfaceapp/models"

	pg "github.com/lib/pq"
)

type Store struct {
	db *sql.DB
}

func (s Store) CreateProfile(name string, image []byte, descriptor []float32, shop int) (profileId int, err error) {

	tx, err := s.db.Begin()
	if err != nil {
		return
	}
	defer tx.Rollback()

	prof_stmt, err := tx.Prepare("INSERT INTO profiles (descriptor, name, shop_num) VALUES ($1,$2,$3) RETURNING id")
	if err != nil {
		return
	}

	defer prof_stmt.Close()
	res := prof_stmt.QueryRow(pg.Array(descriptor), name, shop)
	if err := res.Scan(&profileId); err != nil {
		return 0, err
	}
	pic_stmt, err := tx.Prepare("INSERT INTO pictures (profile_id, data) VALUES($1,$2)")
	if err != nil {
		return
	}
	defer pic_stmt.Close()
	if _, err := pic_stmt.Exec(profileId, image); err != nil {
		return 0, err
	}
	err = tx.Commit()
	return

}

func (s Store) GetAllProfiles() (profiles []models.Profile, err error) {
	rows, err := s.db.Query("SELECT id, descriptor, name, shop_num, created_date FROM profiles")
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		profile := models.Profile{}
		err = rows.Scan(&profile.Id, pg.Array(&profile.Descriptor), &profile.Name, &profile.ShopNum, &profile.CreatedDate)
		if err != nil {
			return
		}
		profiles = append(profiles, profile)
	}
	return
}

func (s Store) GetImage(profileId int) (data []byte, err error) {
	var rawData interface{}
	stmt, err := s.db.Prepare("SELECT data FROM pictures WHERE profile_id=$1 ")
	if err != nil {
		return
	}
	defer stmt.Close()
	row := stmt.QueryRow(profileId)
	err = row.Scan(&rawData)
	if err != nil {
		return
	}
	data, ok := rawData.([]byte)
	if !ok {
		return nil, fmt.Errorf("Type assertion problem interface{} to []byte")
	}
	return

}

func createTables(db *sql.DB) {
	tables := map[string]string{
		"profiles":    "CREATE TABLE IF NOT EXISTS profiles (id SERIAL PRIMARY KEY ,descriptor double precision [] not null, name text not null,shop_num int NOT NULL, created_date time not null DEFAULT NOW())",
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

var thisStore Store
var once sync.Once

func dbInit() {
	conn_string := os.Getenv("DB_CONNECTION_STRING")
	db, err := sql.Open("postgres", conn_string)
	if err != nil {
		panic(err)
	}
	createTables(db)
	thisStore = Store{db}

}

func GetDB() *Store {
	once.Do(dbInit)
	return &thisStore
}
