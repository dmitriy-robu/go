package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Trade struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	UUID       string             `bson:"uuid,omitempty" validate:"required"`
	TradeToken string             `bson:"trade_token,omitempty"`
	SteamID    string             `bson:"steam_id,omitempty"`
	Items      []string           `bson:"items,omitempty"`
	To         string             `bson:"to,omitempty"`
	Status     *string            `bson:"status,omitempty"`
}

type CreateTradeDTO struct {
	TradeToken string   `json:"trade_token"`
	SteamID    string   `json:"steam_id"`
	Items      []string `json:"items"`
	UUID       string   `json:"uuid"`
}
