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
		Where("boxes.active = ?", 1).
		Find(&boxes)

	return boxes
}

func (b BoxRepository) FindByUUID(uuid string) models.Box {
	var (
		box models.Box
	)

	MysqlDB.
		Preload("BoxItem.Item").
		Table("boxes").
		Where("boxes.active = ?", 1).
		Where("boxes.uuid = ?", uuid).
		Find(&box)

	log.Printf("items: %v", box)

	return box
}
