package testdb

import (
	"context"
	"fmt"

	openapi "github.com/GIT_USER_ID/GIT_REPO_ID/go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type TestDB struct {
	db *gorm.DB
}

func New(dsn string) (*TestDB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("gorm open: %w", err)
	}

	// Migrate the schema
	err = db.AutoMigrate(
		&openapi.Attribute{},
		&openapi.Assertion{},
		&openapi.TestServiceUnderTest{},
		&openapi.Test{},
	)
	if err != nil {
		return nil, fmt.Errorf("gorm auto migrate: %w", err)
	}

	return &TestDB{
		db: db,
	}, nil
}

func (td *TestDB) CreateTest(ctx context.Context, test *openapi.Test) (int64, error) {
	/*	t := &Test{
		Test: test,
	}*/
	tx := td.db.Create(test)
	if tx.Error != nil {
		return 0, tx.Error
	}
	tx = td.db.Save(test)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return test.Id, nil
}

func (td *TestDB) GetTest(ctx context.Context, id int64) (*openapi.Test, error) {
	var test openapi.Test
	tx := td.db.First(&test, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &test, nil
}

func (td *TestDB) Drop() error {
	return td.db.Migrator().DropTable(
		&openapi.Attribute{},
		&openapi.TestServiceUnderTest{},
		&openapi.Assertion{},
		&openapi.Result{},
		&openapi.Test{},
	)
}
