package db

import (
	"database/sql"
	"path/filepath"

	"git.bode.fun/meals/pkg/migratelog"
	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
)

const driverName = "sqlite"

// TODO: Move into config.
func ResolveDBPath(path string) (string, error) {
	if path == "" {
		path = "meals.db"
	}

	var err error

	path, err = filepath.Abs(path)
	if err != nil {
		return "", err
	}

	return path, nil
}

// Pragmas and comments copied and modified from pocketbase
// Source: https://github.com/pocketbase/pocketbase/blob/f266621a0faa68edcd2def57ca6059d203ec15ad/core/db_nocgo.go#L10C1-L22C2.
func ConnectDB(dbPath string) (*sql.DB, error) {
	// TODO: Read SQLite gotchas from their website
	pragmas := ""

	// Note: the busy_timeout pragma must be first because
	// the connection needs to be set to block on busy before WAL mode
	// is set in case it hasn't been already set by another connection.
	// pragmas := "?_pragma=busy_timeout(10000)&_pragma=journal_mode(WAL)&_pragma=journal_size_limit(200000000)&_pragma=synchronous(NORMAL)&_pragma=foreign_keys(ON)&_pragma=temp_store(MEMORY)&_pragma=cache_size(-16000)"

	db, err := sql.Open(driverName, dbPath+pragmas)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func Migrate(db *sql.DB) error {
	// TODO: Make an slog compatible logger or fork goose.
	// TODO: Config logger via conf.
	goose.SetLogger(&migratelog.StructuredLogger{})

	err := goose.SetDialect(driverName)
	if err != nil {
		return err
	}

	// TODO: Move into config.
	// TODO: Embed?
	return goose.Up(db, "db/sql/migrations")
}
