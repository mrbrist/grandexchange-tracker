package main

import (
	"backend/internal/database"
	"backend/internal/handlers"
	"backend/internal/utils"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

// Some cron jobs will need to be put here to keep the database updated as well as back fill the database in the event the server is offline for some time, the will pull from the osrs api for historic item price data with the goal being all data points being stored and updated every 5 mins with the real time api

func main() {
	envCfg := utils.SetupEnvCfg()

	itemMappingData, err := utils.GetItemMappingData(envCfg.APIAppName)

	if err != nil {
		fmt.Println("FAILED TO GET ITEM MAPPING DATA...")
		return
	}

	db, err := sql.Open("postgres", envCfg.DB_URL)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)

	cfg := &handlers.APIConfig{
		Env:              envCfg,
		DB:               dbQueries,
		GlobalItemsStore: &itemMappingData,
	}

	// Start cron scheduling
	err = utils.StartScheduling(envCfg.APIAppName, dbQueries)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	// System
	mux.HandleFunc("/health", cfg.Health)

	// API
	// mux.HandleFunc("/items_info", cfg.ItemsInfo)
	mux.HandleFunc("/items_list", cfg.ItemsList)
	mux.HandleFunc("/item/{id}", cfg.Item)

	// TESTING
	mux.HandleFunc("/test", cfg.Test)

	srv := &http.Server{
		Addr:    ":" + cfg.Env.Port,
		Handler: mux,
	}

	log.Printf("Serving on: http://localhost:%s/\n", cfg.Env.Port)
	log.Fatal(srv.ListenAndServe())
}
