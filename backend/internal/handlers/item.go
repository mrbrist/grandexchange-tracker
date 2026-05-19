package handlers

import (
	"backend/internal/utils"
	"net/http"
)

// Get the data for a specific item, this includes all the price history and current price, o that the data can be displayed on a graph in the frontend

// This just pulls the data from the database
func (cfg *APIConfig) Item(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	utils.RespondWithJSON(w, 200, id)
}
