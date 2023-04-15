package config

import (
	"go-rust-drop/internal/api/utils"
)

type SteamAPIs struct {
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
			APIKey:             env.GetEnvOrDefault("STEAM_API_KEY", "25EC47576042A4ED4E9EFC32308864C3"),
			SteamAccountAPIKey: env.GetEnvOrDefault("STEAM_ACCOUNT_API_KEY", "3_FCxv4dfoq2gdEreuJTQdIyUIM"),
		},
		GameInventory: GameInventory{
			SteamID:   env.GetEnvOrDefault("GAME_INVENTORY_STEAM_ID", "76561199222590363"),
			AppID:     env.GetEnvOrDefault("GAME_INVENTORY_APP_ID", "252490"),
			ContextID: 2,
		},
	}
}
