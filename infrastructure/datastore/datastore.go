package datastore

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type DataBase interface {
	Close() error
}

type MysqlDB struct {
	DB *sqlx.DB
}

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

func (m *MysqlDB) Close() error {
	return m.DB.Close()
}
