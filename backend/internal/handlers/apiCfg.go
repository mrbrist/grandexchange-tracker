package handlers

import (
	"backend/internal/database"
	"backend/internal/utils"
)

type APIConfig struct {
	Env              *utils.EnvCfg
	DB               *database.Queries
	GlobalItemsStore *utils.Items
}
