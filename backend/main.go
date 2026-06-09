package main

import (
	"backend/internal/database"
	"backend/internal/handlers"
	"backend/internal/middleware"
	"backend/internal/utils"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

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

	mux := http.NewServeMux()

	// System
	mux.HandleFunc("/api/health", cfg.Health)

	// API
	// mux.HandleFunc("/items_info", cfg.ItemsInfo)
	mux.HandleFunc("/api/list", cfg.ItemsList)
	mux.HandleFunc("/api/item/{id}", cfg.Item)

	// TESTING
	mux.HandleFunc("/api/count", cfg.Count)
	mux.HandleFunc("/api/last", cfg.LastUpdate)

	srv := &http.Server{
		Addr:    ":" + cfg.Env.Port,
		Handler: middleware.CORSMiddleware(mux),
	}

	log.Printf("Serving on: http://localhost:%s/\n", cfg.Env.Port)
	log.Fatal(srv.ListenAndServe())
}
