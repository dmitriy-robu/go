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
			Url:                env.GetEnvOrDefault("STEAM_APIS_URL", "https://api.steamapis.com/steam/inventory").(string),
			APIKey:             env.GetEnvOrDefault("STEAM_APIS_API_KEY", "k-ENKG-syMEX_NuE2i0gKnzJff4").(string),
			SteamAccountAPIKey: env.GetEnvOrDefault("STEAM_ACCOUNT_API_KEY", "3_FCxv4dfoq2gdEreuJTQdIyUIM").(string),
		},
		GameInventory: GameInventory{
			SteamID:   env.GetEnvOrDefault("GAME_INVENTORY_STEAM_ID", "76561199222590363").(string),
			AppID:     env.GetEnvOrDefault("GAME_INVENTORY_APP_ID", "252490").(string),
			ContextID: 2,
		},
	}
}
