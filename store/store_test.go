package store_test

import (
	"os"
	"testing"
)

var databaseURL string

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost user=usergo password=0000 dbname=dbfortests sslmode=disable"
	}
	os.Exit(m.Run()) // run возврадает код ошибки
}
