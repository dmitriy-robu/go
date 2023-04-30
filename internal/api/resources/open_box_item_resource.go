package resources

import (
	"go-rust-drop/internal/api/models"
	"go-rust-drop/internal/api/utils"
	"strings"
)

type OpenBoxItemResource struct {
	Item         models.BoxItem
	ServerSeed   string
	moneyConvert utils.MoneyConvert
}

func (obir OpenBoxItemResource) ToJSON() (map[string]interface{}, *utils.Errors) {
	return map[string]interface{}{
		"item": map[string]interface{}{
			"uuid":             obir.Item.Item.UUID,
			"title":            obir.Item.Item.Name,
			"price":            obir.moneyConvert.FromCentsToVault(uint(obir.Item.Item.Price)),
			"game_environment": obir.Item.Item.GameEnvironment,
			"color":            obir.Item.Item.Color,
			"image_url":        obir.Item.Item.ImageUrl,
			"rarity":           strings.ToTitle(string(obir.Item.Rarity)),
		},
		"server_seed": obir.ServerSeed,
	}, nil
}
