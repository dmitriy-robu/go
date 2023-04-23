package services

import (
	"go-rust-drop/internal/api/models"
	"go-rust-drop/internal/api/repositories"
	"go-rust-drop/internal/api/utils"
	"net/http"
)

type BoxManager struct {
	BoxRepository repositories.BoxRepository
}

func (b BoxManager) FindAllWithItems() models.Boxes {
	var (
		boxes models.Boxes
	)

	b.BoxRepository.MysqlDB = MysqlDB

	boxes = b.BoxRepository.FindAllWithItems()

	return boxes
}

func (b BoxManager) FindByUUIDWithItems(uuid string) (models.Box, utils.Errors) {
	var (
		err error
		box models.Box
	)

	b.BoxRepository.MysqlDB = MysqlDB

	box, err = b.BoxRepository.FindByUUIDWithItems(uuid)
	if err != nil {
		return box, utils.Errors{
			Code:    http.StatusNotFound,
			Message: "Box not found",
			Err:     err,
		}
	}

	return box, utils.Errors{}
}

func (b BoxManager) Open(uuid string) (models.Box, utils.Errors) {
	var (
		err error
		box models.Box
	)

	b.BoxRepository.MysqlDB = MysqlDB

	box, err = b.BoxRepository.FindByUUIDWithItems(uuid)
	if err != nil {
		return box, utils.Errors{
			Code:    http.StatusNotFound,
			Message: "Box not found",
			Err:     err,
		}
	}

	return box, utils.Errors{}
}
