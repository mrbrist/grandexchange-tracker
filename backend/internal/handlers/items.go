package handlers

import (
	"backend/internal/utils"
	"net/http"
)

func (cfg *APIConfig) Items(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, 200, cfg.GlobalItemsStore)
}
