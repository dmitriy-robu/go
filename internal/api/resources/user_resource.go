package resources

import (
	"go-rust-drop/internal/api/models"
	"go-rust-drop/internal/api/utils"
)

type UserResources struct {
	User *models.User
	util utils.MoneyConvert
}

//исрпавить level

func (ur *UserResources) ToJSON() (map[string]interface{}, error) {
	return map[string]interface{}{
		"user": map[string]interface{}{
			"id":         ur.User.ID,
			"uuid":       ur.User.UUID,
			"name":       ur.User.Name,
			"avatar_url": ur.User.AvatarURL,
			"balance":    ur.util.FromCentsToVault(ur.User.UserBalance.Balance),
			"trade_url":  ur.User.SteamTradeURL,
			//"level":      services.LevelManager{}.GetLevelForByExperience(ur.User.Experience),
		},
	}, nil
}
