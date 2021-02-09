package testutils

import (
	"fmt"
	"log"

	"github.com/Tatsuemon/anony/config"
	"github.com/Tatsuemon/anony/domain/model"
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

// InsertUserData inserts user data
func InsertUserData() {
	passes := make([]string, 5)
	for i := 1; i < 6; i++ {
		passes[i-1], _ = model.EncryptPassword(fmt.Sprintf("password%v", i))
	}
	users := []model.User{
		{ID: "id1", Name: "name1", Email: "email1", EncryptedPass: passes[0]},
		{ID: "id2", Name: "name2", Email: "email2", EncryptedPass: passes[1]},
		{ID: "id3", Name: "name3", Email: "email3", EncryptedPass: passes[2]},
		{ID: "id4", Name: "name4", Email: "email4", EncryptedPass: passes[3]},
		{ID: "id5", Name: "name5", Email: "email5", EncryptedPass: passes[4]},
	}
	for _, p := range users {
		_, err := testDB.DB.Exec("INSERT INTO users (id, name, email, password) values (?, ?, ?, ?)", p.ID, p.Name, p.Email, p.EncryptedPass)
		if err != nil {
			panic(err)
		}
	}
}

// ClearUserData clears users data
func ClearUserData() {
	_, err := testDB.DB.Exec("DELETE FROM users")
	if err != nil {
		panic(err)
	}
}

// CountUserData counts user data
func CountUserData() int {
	var count int
	row := testDB.DB.QueryRow("SELECT COUNT(*) FROM users")
	err := row.Scan(&count)
	if err != nil {
		panic(err)
	}
	return count
}
