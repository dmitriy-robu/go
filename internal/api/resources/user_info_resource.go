package resources

import (
	"encoding/json"
	"go-rust-drop/internal/api/models"
)

type UserResources struct {
	UserBalance models.UserBalance
	User        models.User
}

func (u UserResources) UserInfo() ([]byte, error) {
	var resource struct {
		ID            *uint64 `json:"id"`
		UUID          string  `json:"uuid"`
		Name          *string `json:"name"`
		AvatarURL     *string `json:"avatar_url"`
		Balance       *int    `json:"balance,omitempty"`
		SteamTradeURL *string `json:"steam_trade_url,omitempty"`
		Experience    *int    `json:"experience,omitempty"`
	}

	resource.ID = u.User.ID
	resource.UUID = u.User.UUID
	resource.Name = u.User.Name
	resource.AvatarURL = u.User.AvatarURL
	resource.Balance = u.UserBalance.Balance
	resource.SteamTradeURL = u.User.SteamTradeURL

	return json.Marshal(&resource)
}
