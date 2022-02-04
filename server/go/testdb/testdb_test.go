package testdb_test

import (
	"context"
	"fmt"
	"testing"

	openapi "github.com/GIT_USER_ID/GIT_REPO_ID/go"
	"github.com/GIT_USER_ID/GIT_REPO_ID/go/testdb"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestCreateTest(t *testing.T) {
	dsn := "host=localhost user=postgres password=postgres port=5432 sslmode=disable"
	db, err := testdb.New(dsn)
	if err != nil {
		t.Fatal(err)
	}
	/*
		defer func() {
			err = db.Drop()
			if err != nil {
				t.Fatal(err)
			}
		}()
	*/
	test := openapi.Test{
		Name:        "first test",
		Description: "description",
		ServiceUnderTest: openapi.TestServiceUnderTest{
			Url: "http://localhost:3030/hello-instrumented",
		},
		Assertions: []openapi.Assertion{{
			OperationName: "Equal",
			Duration:      "100",
			NumOfSPans:    2,
			Attributes: []openapi.Attribute{{
				Key:   "kubeshop",
				Value: "1",
			}},
		}},
		Repeats: 0,
	}
	ctx := context.Background()
	id, err := db.CreateTest(ctx, &test)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(id)
	gotTest, err := db.GetTest(ctx, id)
	if err != nil {
		t.Fatal(err)
	}
	spew.Dump(gotTest)
	assert.Equal(t, test, gotTest)

}
