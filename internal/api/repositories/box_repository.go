package repositories

import (
	"go-rust-drop/internal/api/models"
	"log"
)

type BoxRepository struct {
}

func (b BoxRepository) FindAll() models.Boxes {
	var boxes models.Boxes

	MysqlDB.
		Table("boxes").
		Where("boxes.active", 1).
		Find(&boxes)

	return boxes
}

func (b BoxRepository) FindByUUID(uuid string) models.Box {
	var (
		box   models.Box
		items []models.BoxItemShow
	)

	MysqlDB.
		Where("boxes.active", 1).
		Where("boxes.uuid", uuid).
		First(&box)

	MysqlDB.
		Preload("Item").
		Table("box_items").
		Where("box_items.box_id", box.ID).
		Find(&items)

	log.Printf("items: %v", items)

	return box
}
