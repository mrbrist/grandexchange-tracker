package handlers

import (
	"backend/internal/utils"
	"net/http"
)

type ItemListItem struct {
	ID   int    `json:"id"`
	Icon string `json:"icon"`
	Name string `json:"name"`
}

func (cfg *APIConfig) ItemsInfo(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, 200, cfg.GlobalItemsStore)
}

func (cfg *APIConfig) ItemsList(w http.ResponseWriter, r *http.Request) {
	il := []ItemListItem{}

	for _, item := range *cfg.GlobalItemsStore {
		il = append(il, ItemListItem{
			ID:   item.ID,
			Icon: item.Icon,
			Name: item.Name,
		})
	}

	utils.RespondWithJSON(w, 200, il)
}
