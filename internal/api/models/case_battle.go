package models

import (
	"go-rust-drop/internal/api/enum"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CaseBattles []CaseBattle

type CaseBattle struct {
	ID          primitive.ObjectID      `bson:"_id"`
	GameMode    enum.GameModes          `bson:"game_mode"`
	PlayersMode enum.PlayersModes       `bson:"players_mode"`
	Privacy     bool                    `bson:"privacy"`
	Status      enum.CaseBattleStatuses `bson:"status"`
	TotalCost   uint                    `bson:"total_cost"`
	Rounds      CaseBattleRounds        `bson:"rounds"`
	CreatedAt   primitive.DateTime      `bson:"created_at"`
	UpdatedAt   primitive.DateTime      `bson:"updated_at"`
}

type CaseBattleRounds []CaseBattleRound

type CaseBattleRound struct {
	ID           primitive.ObjectID `bson:"_id"`
	BoxUUID      string             `bson:"box_uuid"`
	CaseBattleID string             `bson:"case_battle_id"`
	CreatedAt    primitive.DateTime `bson:"created_at"`
	UpdatedAt    primitive.DateTime `bson:"updated_at"`
}
