package testdb_test

import (
	"testing"

	"github.com/GIT_USER_ID/GIT_REPO_ID/go/testdb"
)

func TestCreateTest(t *testing.T) {
	dsn := "host=localhost user=postgres password=postgres port=5432 sslmode=disable"
	db, err := testdb.New(dsn)
	if err != nil {
		t.Fatal(err)
	}

	_ = db
}
