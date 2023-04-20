package resources

import (
	"go-rust-drop/internal/api/models"
	"go-rust-drop/internal/api/utils"
)

type BoxesResource struct {
	Boxes        models.Boxes
	moneyConvert utils.MoneyConvert
}

func (b BoxesResource) ToJSON() []map[string]interface{} {
	var boxesResource []map[string]interface{}

	for _, box := range b.Boxes {
		boxResource := map[string]interface{}{
			"uuid":      box.UUID,
			"title":     box.Title,
			"image":     box.Image,
			"alt_image": box.AltImage,
			"price":     box.Price,
		}

		boxesResource = append(boxesResource, boxResource)
	}

	return boxesResource
}
