package services

import (
	"go-rust-drop/internal/api/models"
	"go-rust-drop/internal/api/repositories"
)

type BoxManager struct {
	boxRepository repositories.BoxRepository
}

func (b BoxManager) FindAll() models.Boxes {
	var (
		boxes models.Boxes
	)

	boxes = b.boxRepository.FindAll()

	return boxes
}

func (b BoxManager) FindByUUID(uuid string) (models.Box, error) {
	var (
		err error
		box models.Box
	)

	box, err = b.boxRepository.FindByUUID(uuid)
	if err != nil {
		return box, err
	}

	return box, nil
}

func (b BoxManager) Open(uuid string) (models.Box, error) {
	var (
		err error
		box models.Box
	)

	box, err = b.boxRepository.FindByUUID(uuid)
	if err != nil {
		return box, err
	}

	return box, nil
}
