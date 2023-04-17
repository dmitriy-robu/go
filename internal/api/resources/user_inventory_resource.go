package resources

import (
	"go-rust-drop/internal/api/models"
	"go-rust-drop/internal/api/utils"
)

type UserInventoryResources struct {
	AssetData []models.AssetData
	util      utils.MoneyConvert
}

func (ur *UserInventoryResources) ToJSON() ([]map[string]interface{}, error) {
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

	return assetMaps, nil
}
