package postgres

import (
	"fmt"
	"time"

	"github.com/xueyiyao/safekeep/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	DB *gorm.DB

	// Datasource Name
	DSN string

	Now func() time.Time
}

func NewDB(dsn string) *DB {
	db := &DB{DSN: dsn, Now: time.Now}
	return db
}

func (db *DB) Open() (err error) {
	if db.DSN == "" {
		return fmt.Errorf("DSN required")
	}

	if db.DB, err = gorm.Open(postgres.Open(db.DSN), &gorm.Config{}); err != nil {
		return err
	}

	if err := db.migrate(); err != nil {
		return fmt.Errorf("migrate: %w", err)
	}

	return nil
}

func (db *DB) migrate() (err error) {
	if err := db.DB.AutoMigrate(&domain.User{}); err != nil {
		return err
	}
	if err := db.DB.AutoMigrate(&domain.Container{}); err != nil {
		return err
	}
	if err := db.DB.AutoMigrate(&domain.Item{}); err != nil {
		return err
	}

	return nil
}

func (db *DB) Close() error {
	if db.DB != nil {
		db, err := db.DB.DB()
		if err != nil {
			return err
		}
		return db.Close()
	}

	return nil
}
