package config

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/usmaarn/locstique_api/internal/database"
	"log"
	"os"
)

func InitializeDatabase() *database.Queries {
	conn, err := sql.Open("postgres", os.Getenv("GOOSE_DBSTRING"))
	if err != nil {
		log.Fatal("Error connecting to database", err.Error())
	}
	return database.New(conn)
}
