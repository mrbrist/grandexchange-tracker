package handlers

import (
	"backend/internal/utils"
	"net/http"
	"strconv"
)

// Get the data for a specific item, this includes all the price history and current price, o that the data can be displayed on a graph in the frontend

// This just pulls the data from the database
func (cfg *APIConfig) Item(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		utils.RespondWithError(w, 500, "Cannot convert id to int", err)
		return
	}

	// FOR TESTING
	// err = cfg.DB.AddItemHistory(r.Context(), database.AddItemHistoryParams{
	// 	ItemID:         int32(id),
	// 	PriceTimestamp: time.Now(),
	// 	AvgHighPrice:   0,
	// 	AvgLowPrice:    0,
	// 	LowVolume:      0,
	// 	HighVolume:     0,
	// })

	history, err := cfg.DB.GetItemHistory(r.Context(), int32(id))
	if err != nil {
		utils.RespondWithError(w, 500, "db error", err)
		return
	}

	utils.RespondWithJSON(w, 200, history)
}
