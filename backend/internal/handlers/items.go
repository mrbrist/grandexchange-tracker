package handlers

import (
	"backend/internal/utils"
	"net/http"
)

// Get the data from the api when the server starts then store it in memory to avoid hitting the osrs api every time a user wants to check the mapping data
func (cfg *APIConfig) Items(w http.ResponseWriter, r *http.Request) {

	utils.RespondWithJSON(w, 200, nil)
}
