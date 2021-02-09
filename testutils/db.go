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

// User

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

// AnonyURL
type urls struct {
	ID       string
	Original string
	Short    string
	Status   int64
	UserID   string
}

// InsertURLData inserts url data
func InsertURLData() {
	InsertUserData()
	urls := []urls{
		{ID: "id1", Original: "original1", Short: "short1", Status: 1, UserID: "id1"},
		{ID: "id2", Original: "original2", Short: "short2", Status: 1, UserID: "id1"},
		{ID: "id3", Original: "original3", Short: "short3", Status: 2, UserID: "id1"},
		{ID: "id4", Original: "original4", Short: "short4", Status: 2, UserID: "id1"},
		{ID: "id5", Original: "original5", Short: "short5", Status: 2, UserID: "id1"},
	}
	for _, p := range urls {
		_, err := testDB.DB.Exec("INSERT INTO urls (id, original, short, status, user_id) values (?, ?, ?, ?, ?)", p.ID, p.Original, p.Short, p.Status, p.UserID)
		if err != nil {
			panic(err)
		}
	}
}

// ClearURLData clears urls data
func ClearURLData() {
	_, err := testDB.DB.Exec("DELETE FROM urls")
	if err != nil {
		panic(err)
	}
}

// CountURLData counts urls data
func CountURLData() int {
	var count int
	row := testDB.DB.QueryRow("SELECT COUNT(*) FROM urls")
	err := row.Scan(&count)
	if err != nil {
		panic(err)
	}
	return count
}
