package adaptors

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	"go.uber.org/zap"

	log "github.com/torvald2/faces/app_logger"
	"github.com/torvald2/faces/config"
	"github.com/torvald2/faces/models"

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

func (s Store) GetShopProfiles(shopId int) (profiles []models.Profile, err error) {
	stmt, err := s.db.Prepare("SELECT id, descriptor, name, shop_num, created_date FROM profiles WHERE shop_num=$1")
	if err != nil {
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(shopId)
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

func (s Store) GetProfileById(profileId int) (profile models.Profile, err error) {
	stmt, err := s.db.Prepare("SELECT id, descriptor, name, shop_num, created_date FROM profiles WHERE id=$1")
	if err != nil {
		return
	}
	defer stmt.Close()

	row := stmt.QueryRow(profileId)
	if err != nil {
		return
	}
	err = row.Scan(&profile.Id, pg.Array(&profile.Descriptor), &profile.Name, &profile.ShopNum, &profile.CreatedDate)
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

func (s Store) NewJornalRecord(oper models.JornalOperation) error {
	stmt, err := s.db.Prepare("INSERT INTO workjornal (profile_id, operation_type, operation_date) VALUES ($1, $2, $3)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(oper.UserId, oper.OperationType, oper.OperationDate)
	return err
}

func (s Store) LogBadRequest(request models.BadRequest) error {
	stmt, err := s.db.Prepare("INSERT INTO badrequest (profile_id, recognized_profiles, current_face, error_type, recognized_time, shop_id, request_id) VALUES ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(request.UserId, pg.Array(request.RecognizedUsers), request.CurrentFace, request.ErrorType, request.RecognizeTime, request.Shop, request.RequestId)
	return err
}
func (s Store) GetJornalRecords(start, end time.Time) (data []models.JornalOperationDB, err error) {
	stmt, err := s.db.Prepare(`SELECT wj.operation_type, wj.operation_date, pf.name, pf.shop_num
                           FROM workjornal as wj
                           JOIN profiles as pf 
                           ON wj.profile_id = pf.id
						   WHERE wj.operation_date BETWEEN $1 AND $2`)
	if err != nil {
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(start, end)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		op := models.JornalOperationDB{}

		err = rows.Scan(&op.OperationType, &op.OperationDete, &op.UserName, &op.ShopNum)
		if err != nil {
			return
		}
		data = append(data, op)
	}
	return
}

func (s Store) GetBadRequests(dateFrom, dateTo time.Time) (data []models.BadRequest, err error) {
	stmt, err := s.db.Prepare(`SELECT  id,
							   profile_id,
							   COALESCE(recognized_profiles, array[]::bigint[])
							   recognized_profiles,
							   error_type,
							   recognized_time
							   FROM badrequest WHERE created_date BETWEEN $1 AND $2`)
	if err != nil {
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(dateFrom, dateTo)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var ru pg.Int64Array
		record := models.BadRequest{}
		err = rows.Scan(&record.Id, &record.UserId, &ru, &record.ErrorType, &record.RecognizeTime)
		if err != nil {
			return
		}
		r := []int64(ru)
		var recognizedUsers []int
		for _, v := range r {
			recognizedUsers = append(recognizedUsers, int(v))

		}
		record.RecognizedUsers = recognizedUsers
		data = append(data, record)
	}
	return
}

func (s Store) GetBadRequestImage(recordId int) (data []byte, err error) {
	var rawData interface{}
	stmt, err := s.db.Prepare("SELECT current_face FROM badrequest WHERE id=$1 ")
	if err != nil {
		return
	}
	defer stmt.Close()
	row := stmt.QueryRow(recordId)
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
	var tx *sql.Tx
	var err error
	tables := map[string]string{
		"profiles":    "CREATE TABLE IF NOT EXISTS profiles (id SERIAL PRIMARY KEY ,descriptor double precision [] not null, name text not null,shop_num int NOT NULL, created_date timestamp not null DEFAULT NOW())",
		"pictures":    "CREATE TABLE IF NOT EXISTS pictures (id SERIAL PRIMARY KEY ,profile_id INT, data bytea NOT NULL)",
		"workjornal":  "CREATE TABLE IF NOT EXISTS workjornal (id SERIAL  PRIMARY KEY, profile_id INT NOT NULL,operation_type TEXT NOT NULL, operation_date timestamp NOT NULL, created_date timestamp not null DEFAULT NOW())",
		"badRequests": "CREATE TABLE IF NOT EXISTS badrequest (id SERIAL  PRIMARY KEY, profile_id INT, recognized_profiles INT[], current_face bytea, error_type TEXT,recognized_time timestamp, shop_id INT, request_id TEXT, created_date timestamp not null DEFAULT NOW())",
	}
	for i := 1; i < 10; i++ {
		tx, err = db.Begin()
		if err == nil {
			break
		}
		log.Logger.Warn("Connection db. New attemt", zap.Error(err))
		time.Sleep(time.Duration(i) * time.Second)
	}
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
	conf := config.GetConfig()
	var db *sql.DB
	var err error
	for i := 1; i < 6; i++ {
		db, err = sql.Open("postgres", conf.DBConnectionString)
		if err == nil {
			break
		}
		log.Logger.Warn("Connection db. New attemt", zap.Error(err))

		time.Sleep(time.Duration(i) * time.Second)
	}
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
