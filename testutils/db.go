package testutils

import (
	"log"

	"github.com/Tatsuemon/anony/config"
	"github.com/Tatsuemon/anony/infrastructure/datastore"
)

var testDB *datastore.MysqlDB

// PrepareTestDB prepare db for test
func PrepareTestDB() *datastore.MysqlDB {
	db, err := datastore.NewMysqlDB(config.DSN())
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// CloseTestDB close db connections for test
func CloseTestDB() {
	if err := testDB.Close(); err != nil {
		log.Fatal(err)
	}
}

// SetTestDB set db for test
func SetTestDB(db *datastore.MysqlDB) {
	testDB = db
	return
}

// GetTestDB get db for test
func GetTestDB() *datastore.MysqlDB {
	if testDB == nil {
		panic("mysql connection is not initialized yet")
	}
	return testDB
}

// TODO(Tatsuemon): テスト用DBに関することをかく
