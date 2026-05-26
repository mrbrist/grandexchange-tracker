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

	history, err := cfg.DB.GetItemHistory(r.Context(), int32(id))
	if err != nil {
		utils.RespondWithError(w, 500, "db error", err)
		return
	}

	utils.RespondWithJSON(w, 200, history)
}

// func (cfg *APIConfig) Test(w http.ResponseWriter, r *http.Request) {
// 	err := utils.UpdatePriceHistory1h(cfg.Env.APIAppName, cfg.DB)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }
