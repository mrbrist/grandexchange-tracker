package handlers

import (
	"backend/internal/database"
	"backend/internal/utils"
	"context"
	"net/http"
	"strconv"
)

// Get the data for a specific item, this includes all the price history and current price, o that the data can be displayed on a graph in the frontend

// This just pulls the data from the database

type itemData struct {
	Data    utils.ItemLookupData      `json:"data"`
	History []database.GePriceHistory `json:"history"`
}

func (cfg *APIConfig) Item(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.RespondWithError(w, 500, "Cannot convert id to int", err)
		return
	}

	history, err := cfg.DB.GetItemHistory(r.Context(), int32(id))
	if err != nil {
		utils.RespondWithError(w, 500, "db error", err)
		return
	}

	lookupData := utils.LookupItem(cfg.GlobalItemsStore, id)

	var data = itemData{
		Data:    lookupData,
		History: history,
	}

	utils.RespondWithJSON(w, 200, data)
}

func (cfg *APIConfig) Count(w http.ResponseWriter, r *http.Request) {
	ret, err := cfg.DB.CountAllRecords(context.Background())
	if err != nil {
		utils.RespondWithError(w, 500, "error getting datapoint count", err)
	}

	utils.RespondWithJSON(w, 200, ret)
}

func (cfg *APIConfig) LastUpdate(w http.ResponseWriter, r *http.Request) {
	ret, err := cfg.DB.GetLatestTimestamp(context.Background())
	if err != nil {
		utils.RespondWithError(w, 500, "error getting latest update timestamp", err)
	}

	utils.RespondWithJSON(w, 200, ret)
}
