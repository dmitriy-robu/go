package repositories

import (
	"github.com/pkg/errors"
	"go-rust-drop/internal/api/models"
)

type BoxRepository struct {
}

func (b BoxRepository) FindAll() models.Boxes {
	var (
		boxes models.Boxes
	)

	MysqlDB.
		Preload("BoxItem.Item").
		Table("boxes").
		Where("boxes.active = ?", 1).
		Find(&boxes)

	return boxes
}

func (b BoxRepository) FindByUUID(uuid string) (models.Box, error) {
	var (
		err error
		box models.Box
	)

	err = MysqlDB.
		Preload("BoxItems.Item").
		Table("boxes").
		Where("boxes.uuid = ?", uuid).
		Where("boxes.active = ?", 1).
		First(&box).Error

	if err != nil {
		return box, errors.Wrap(err, "Error finding box by UUID")
	}

	return box, nil
}
