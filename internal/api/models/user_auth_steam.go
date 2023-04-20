package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserAuthSteam struct {
	ID          primitive.ObjectID `bson:"_id"`
	UserUUID    string             `bson:"user_uuid"`
	SteamUserID string             `bson:"steam_user_id"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}
