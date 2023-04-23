package repositories

import (
	"github.com/pkg/errors"
	"go-rust-drop/internal/api/models"
	"gorm.io/gorm"
)

type BoxRepository struct {
	MysqlDB *gorm.DB
}

func (b BoxRepository) FindAllWithItems() models.Boxes {
	var (
		boxes models.Boxes
	)

	b.MysqlDB.
		Preload("BoxItem.Item").
		Table("boxes").
		Where("boxes.active = ?", 1).
		Find(&boxes)

	return boxes
}

func (b BoxRepository) FindByUUIDWithItems(uuid string) (models.Box, error) {
	var (
		err error
		box models.Box
	)

	err = b.MysqlDB.
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

func (b BoxRepository) FindByUUID(uuid string) (models.Box, error) {
	var (
		err error
		box models.Box
	)

	err = b.MysqlDB.
		Table("boxes").
		Where("boxes.uuid = ?", uuid).
		Where("boxes.active = ?", 1).
		First(&box).Error

	if err != nil {
		return box, errors.Wrap(err, "Error finding box by UUID")
	}

	return box, nil
}
