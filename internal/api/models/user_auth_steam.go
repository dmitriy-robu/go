package models

import (
	"time"
)

type UserAuthSteam struct {
	UserUUID    string    `bson:"user_uuid"`
	SteamUserID string    `bson:"steam_user_id"`
	CreatedAt   time.Time `bson:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at"`
}
