package main

import (
	"cron/internal/database"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	envCfg := SetupEnvCfg()

	db, err := sql.Open("postgres", envCfg.DB_URL)
	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)

	scheduler, err := StartScheduling(envCfg.APIAppName, dbQueries)
	if err != nil {
		log.Fatal(err)
	}

	defer scheduler.Shutdown()

	select {} // keep process alive
}
