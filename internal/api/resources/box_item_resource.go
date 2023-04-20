package resources

import (
	"fmt"
	"go-rust-drop/internal/api/models"
)

type BoxItemResource struct {
	BoxItem models.BoxItem
}

func (bir BoxItemResource) ToJSON() map[string]interface{} {
	var (
		boxItemResource map[string]interface{}
		item            models.Item
	)

	item = bir.BoxItem.Item

	fmt.Println(item)

	boxItemResource = map[string]interface{}{
		"uuid":   item.UUID,
		"title":  item.Name,
		"image":  item.ImageUrl,
		"price":  item.Price,
		"color":  item.Color,
		"rarity": bir.BoxItem.Rarity,
	}

	return boxItemResource
}
