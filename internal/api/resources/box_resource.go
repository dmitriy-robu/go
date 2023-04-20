package resources

import (
	"go-rust-drop/internal/api/models"
)

type BoxResource struct {
	Box models.Box
}

func (br BoxResource) ToJSON() map[string]interface{} {
	var (
		boxResource map[string]interface{}
		items       []map[string]interface{}
	)

	for _, boxItem := range br.Box.BoxItems {
		boxItemResource := BoxItemResource{
			BoxItem: boxItem,
		}

		items = append(items, boxItemResource.ToJSON())
	}

	boxResource = map[string]interface{}{
		"uuid":      br.Box.UUID,
		"title":     br.Box.Title,
		"image":     br.Box.Image,
		"alt_image": br.Box.AltImage,
		"price":     br.Box.Price,
		"items":     items,
	}

	return boxResource
}
