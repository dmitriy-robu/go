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

func NewBoxManager(
	BoxRepository repositories.BoxRepository,
) BoxManager {
	return BoxManager{
		BoxRepository: BoxRepository,
	}
}

func (b BoxManager) FindAllWithItems() models.Boxes {
	var (
		boxes models.Boxes
	)

	return boxes
}

func (b BoxManager) FindByUUIDWithItems(uuid string) (models.Box, *utils.Errors) {
	var (
		err error
		box models.Box
	)

	box, err = b.BoxRepository.FindByUUIDWithItems(uuid)
	if err != nil {
		return box, utils.NewErrors(http.StatusNotFound, "Box not found", err)
	}

	return box, nil
}

func (b BoxManager) Open(uuid string) (models.Box, *utils.Errors) {
	var (
		err error
		box models.Box
	)

	box, err = b.BoxRepository.FindByUUIDWithItems(uuid)
	if err != nil {
		return box, utils.NewErrors(http.StatusNotFound, "Box not found", err)
	}

	return box, nil
}
