package db

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/vusalalishov/manpass/internal/config"
)

type Migrator interface {
	Migrate() error
}

type migratorImpl struct {
	db *sql.DB
	config config.Config
}

func (m *migratorImpl) Migrate() error {
	driver, err := sqlite3.WithInstance(m.db, &sqlite3.Config{})
	if err != nil {
		return err
	}
	migration, err := migrate.NewWithDatabaseInstance(m.config.Get(config.MIGRATIONS_PATH), m.config.Get(config.DB_FILE), driver)
	if err != nil {
		return err
	}
	return migration.Up()
}

func ProvideMigrator(db *sql.DB, cfg config.Config) Migrator {
	return &migratorImpl{db: db, config: cfg}
}

func InjectMigrator() (Migrator, error) {
	db, err := InjectDb()
	if err != nil {
		return nil, err
	}
	return ProvideMigrator(db, config.InjectConfig()), nil
}