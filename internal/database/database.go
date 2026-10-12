package database

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

func Connect(databaseURL string) (*sql.DB, error) {
	//sql.Open("pgx", databaseURL)
	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		return nil, err
	}
	//db.Ping()
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil

}
func Migrate(db *sql.DB) error {
	//goose.SetDialect("postgres")
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}
	return goose.Up(db, "migrations")
}
