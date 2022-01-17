package store

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"testing"
)

func TestStore(t *testing.T, databaseUrl string) (*sql.DB, func(...string)) {
	t.Helper()
	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db, func(s ...string) {
		if len(s) > 0 {
			db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(s, ", ")))
		}
		db.Close()
	}
}
