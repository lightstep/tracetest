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

type Test struct {
	gorm.Model
	openapi.Test
}

func New(dsn string) (*TestDB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("gorm open: %w", err)
	}
	// Migrate the schema
	err = db.AutoMigrate(&Test{})
	if err != nil {
		return nil, fmt.Errorf("gorm auto migrate: %w", err)
	}
	return &TestDB{
		db: db,
	}, nil
}

func (td *TestDB) CreateTest(ctx context.Context, test openapi.Test) error {
	tx := td.db.Create(&Test{
		Test: test,
	})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
