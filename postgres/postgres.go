package postgres

import (
	"fmt"
	"time"

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

	return nil
}
