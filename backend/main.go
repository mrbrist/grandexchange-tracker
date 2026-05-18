package main

import (
	"backend/internal/handlers"
	"backend/internal/utils"
	"log"
	"net/http"
)

func main() {
	envCfg := utils.SetupEnvCfg()

	cfg := &handlers.APIConfig{
		Env: envCfg,
	}

	mux := http.NewServeMux()

	// System
	mux.HandleFunc("/health", cfg.Health)

	srv := &http.Server{
		Addr:    ":" + cfg.Env.Port,
		Handler: mux,
	}

	log.Printf("Serving on: http://localhost:%s/\n", cfg.Env.Port)
	log.Fatal(srv.ListenAndServe())
}
