package resources

import (
	"encoding/json"
	"go-rust-drop/internal/api/models"
	"go-rust-drop/internal/api/services"
	"go-rust-drop/internal/api/utils"
)

type UserResources struct {
	User        *models.User
	UserBalance models.UserBalance
	Inventory   models.Inventory
	AssetData   []models.AssetData
	util        utils.MoneyConvert
}

/*func (ur UserResources) UserInfo() ([]byte, error) {
	var resource = struct {
		ID            *int    `json:"id"`
		UUID          string  `json:"uuid"`
		Name          *string `json:"name"`
		AvatarURL     *string `json:"avatar_url"`
		Balance       *int    `json:"balance"`
		SteamTradeURL *string `json:"steam_trade_url"`
		Experience    *int    `json:"experience"`
	}{
		ID:            ur.User.ID,
		UUID:          ur.User.UUID,
		Name:          ur.User.Name,
		AvatarURL:     ur.User.AvatarURL,
		Balance:       ur.UserBalance.Balance,
		SteamTradeURL: ur.User.SteamTradeURL,
		Experience:    ur.User.Experience,
	}

	return json.Marshal(&resource)
}*/

func (ur *UserResources) UserInfo() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"user": map[string]interface{}{
			"id":         ur.User.ID,
			"uuid":       ur.User.UUID,
			"name":       ur.User.Name,
			"avatar_url": ur.User.AvatarURL,
			"balance":    ur.util.FromCentsToVault(*ur.User.UserBalance.Balance),
			"trade_url":  ur.User.SteamTradeURL,
			"level":      services.LevelService{}.GetLevelForByExperience(ur.User.Experience),
		},
	})
}

func (ur *UserResources) UserInventory() ([]byte, error) {
	assetMaps := make([]map[string]interface{}, len(ur.AssetData))

	for i, assetData := range ur.AssetData {
		assetMaps[i] = map[string]interface{}{
			"asset_id":  assetData.AssetID,
			"name":      assetData.Name,
			"amount":    1,
			"price":     ur.util.FromCentsToVault(assetData.Price),
			"color":     "#" + assetData.BackgroundColor,
			"image_url": assetData.IconURL,
			"is_stack":  false,
		}
	}

	jsonData, err := json.Marshal(assetMaps)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}
