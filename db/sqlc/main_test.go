package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testQueries *Queries
var testDB *pgxpool.Pool

const (
	dbSource = "postgresql://root:secret@172.26.5.169:5432/simple_bank?sslmode=disable"
)

func TestMain(m *testing.M) {
	ctx := context.Background()
	var err error
	testDB, err = pgxpool.New(ctx, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	testQueries = New(testDB)

	code := m.Run()
	testDB.Close()
	os.Exit(code)
}
