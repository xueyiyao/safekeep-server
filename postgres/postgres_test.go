package postgres_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/xueyiyao/safekeep/internal/path"
	"github.com/xueyiyao/safekeep/postgres"
	"gorm.io/gorm"
)

func TestDB(t *testing.T) {
	db, tx := MustOpenDB(t)
	MustCloseDB(t, db, tx, func() {})
}

func MustOpenDB(tb testing.TB) (*postgres.DB, *gorm.DB) {
	tb.Helper()

	if err := godotenv.Load(path.Root + "/.env"); err != nil {
		tb.Fatal(err)
	}

	dsn := getDSN()
	db := postgres.NewDB(dsn)
	if err := db.Open(); err != nil {
		tb.Fatal(err)
	}
	tx := db.DB.Begin()
	if tx.Error != nil {
		tb.Fatal(tx.Error)
	}

	return db, tx
}

func MustCloseDB(tb testing.TB, db *postgres.DB, tx *gorm.DB, teardown func()) {
	tb.Helper()

	if err := tx.Rollback(); err.Error != nil {
		tb.Fatal(err.Error)
	}
	teardown()

	if err := db.Close(); err != nil {
		tb.Fatal(err)
	}
}

func getDSN() string {
	host := os.Getenv("TEST_DB_HOST")
	port := os.Getenv("TEST_DB_PORT")
	user := os.Getenv("TEST_DB_USER")
	password := os.Getenv("TEST_DB_PASSWORD")
	dbname := os.Getenv("TEST_DB_NAME")

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
}

func resetSequence(db *gorm.DB, tableName string, columnName string, resetTo int) error {
	query := fmt.Sprintf("SELECT setval(pg_get_serial_sequence('%s', '%s'), %d, false)", tableName, columnName, resetTo)
	return db.Exec(query).Error
}
