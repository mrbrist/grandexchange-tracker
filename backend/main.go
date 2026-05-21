package main

import (
	"backend/internal/handlers"
	"backend/internal/utils"
	"fmt"
	"log"
	"net/http"
)

// Some cron jobs will need to be put here to keep the database updated as well as back fill the database in the event the server is offline for some time, the will pull from the osrs api for historic item price data with the goal being all data points being stored and updated every 5 mins with the real time api

func main() {
	envCfg := utils.SetupEnvCfg()

	itemMappingData, err := utils.GetItemMappingData(envCfg.APIAppName)

	if err != nil {
		fmt.Println("FAILED TO GET ITEM MAPPING DATA...")
		return
	}

	cfg := &handlers.APIConfig{
		Env:              envCfg,
		GlobalItemsStore: &itemMappingData,
	}

	mux := http.NewServeMux()

	// System
	mux.HandleFunc("/health", cfg.Health)

	// API
	// mux.HandleFunc("/items_info", cfg.ItemsInfo)
	mux.HandleFunc("/items_list", cfg.ItemsList)
	mux.HandleFunc("/item/{id}", cfg.Item)

	srv := &http.Server{
		Addr:    ":" + cfg.Env.Port,
		Handler: mux,
	}

	log.Printf("Serving on: http://localhost:%s/\n", cfg.Env.Port)
	log.Fatal(srv.ListenAndServe())
}
