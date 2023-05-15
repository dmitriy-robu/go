package config

import (
	"go-rust-drop/internal/api/utils"
)

type SteamAPIs struct {
	Url                string
	APIKey             string
	SteamAccountAPIKey string
}

type GameInventory struct {
	SteamID   string
	AppID     string
	ContextID int
}

type SteamSettings struct {
	SteamAPIs     SteamAPIs
	GameInventory GameInventory
}

func SetSteamSettings() SteamSettings {
	env := utils.Environment{}
	return SteamSettings{
		SteamAPIs: SteamAPIs{
			Url:                env.GetEnvOrDefault("STEAM_APIS_URL", "").(string),
			APIKey:             env.GetEnvOrDefault("STEAM_APIS_API_KEY", "").(string),
			SteamAccountAPIKey: env.GetEnvOrDefault("STEAM_ACCOUNT_API_KEY", "").(string),
		},
		GameInventory: GameInventory{
			SteamID:   env.GetEnvOrDefault("GAME_INVENTORY_STEAM_ID", "").(string),
			AppID:     env.GetEnvOrDefault("GAME_INVENTORY_APP_ID", "").(string),
			ContextID: 2,
		},
	}
}
