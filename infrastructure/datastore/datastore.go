package datastore

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// DataBase is interface of DB
type DataBase interface {
	Close() error
}

// MysqlDB is DB using Mysql
type MysqlDB struct {
	DB *sqlx.DB
}

// NewMysqlDB create MysqlDB
func NewMysqlDB(datasource string) (*MysqlDB, error) {
	db, err := sqlx.Open("mysql", datasource)
	if err != nil {
		return nil, fmt.Errorf("failed to open MySQL: %w", err)
	}

	// コネクションプールの設定
	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(100)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping: %w", err)
	}

	return &MysqlDB{DB: db}, nil
}

// Close close db connection
func (m *MysqlDB) Close() error {
	return m.DB.Close()
}
