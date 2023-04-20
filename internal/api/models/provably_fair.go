package models

import (
	"go-rust-drop/internal/api/enum"
)

type ProvablyFair struct {
	Game         enum.Game `bson:"game"`
	ServerSeed   string    `bson:"server_seed"`
	ClientSeed   string    `bson:"client_seed"`
	Nonce        int       `bson:"nonce"`
	MinChance    float64   `bson:"min_chance"`
	MaxChance    float64   `bson:"max_chance"`
	RandomNumber float64   `bson:"random_number"`
}
