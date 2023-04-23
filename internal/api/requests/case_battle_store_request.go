package requests

import (
	"go-rust-drop/internal/api/enum"
)

type CaseBattleStoreRequest struct {
	PlayersMode enum.PlayersModes `json:"players_mode" binding:"required"`
	GameMode    enum.GameModes    `json:"games_mode" binding:"required"`
	Privacy     bool              `json:"privacy"`
	ClientSeed  string            `json:"client_seed" binding:"max=255"`
	Boxes       Boxes             `json:"boxes" binding:"required,min=1,max=35"`
}

type Boxes []Box

type Box struct {
	UUID     string `json:"uuid" binding:"required"`
	Quantity uint   `json:"quantity" binding:"required"`
}
