package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/vusalalishov/manpass/internal/config"
	"sync"
)

var (
	once sync.Once
	db *sql.DB
)

func ProvideDb(cfg config.Config) (*sql.DB, error) {
	var err error
	var dbFile = cfg.Get(config.DbFile)
	once.Do(func() {
		db, err = sql.Open("sqlite3", fmt.Sprintf("./%s?mode=rw", dbFile))
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InjectDb()	(*sql.DB, error) {
	cfg := config.InjectConfig()
	return ProvideDb(cfg)
}