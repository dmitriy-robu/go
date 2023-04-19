package resources

import (
	"go-rust-drop/internal/api/models"
	"go-rust-drop/internal/api/utils"
)

type BoxesResource struct {
	Boxes        models.Boxes
	moneyConvert utils.MoneyConvert
}

func (rr *BoxesResource) ToJSON() ([]map[string]interface{}, error) {
	var (
		boxes []map[string]interface{}
		asset map[string]interface{}
	)

	for _, box := range rr.Boxes {
		asset = map[string]interface{}{
			"uuid":          box.UUID,
			"title":         box.Title,
			"price":         rr.moneyConvert.FromCentsToVault(box.Price),
			"image_url":     box.Image,
			"alt_image_url": box.AltImage,
		}
		boxes = append(boxes, asset)
	}

	return boxes, nil
}
