package db

import (
	"database/sql"
	"github.com/BigListRyRy/harbourlivingapi/util"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

var testQueries *Queries

func TestMain(m *testing.M) {

	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to database, ", err)
	}

	testQueries = New(conn)
	os.Exit(m.Run())
}
