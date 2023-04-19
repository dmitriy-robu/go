package services

import (
	"go-rust-drop/internal/api/models"
	"go-rust-drop/internal/api/repositories"
)

type BoxService struct {
	boxRepository repositories.BoxRepository
}

func (b BoxService) FindAll() models.Boxes {
	var (
		boxes models.Boxes
	)

	boxes = b.boxRepository.FindAll()

	return boxes
}

func (b BoxService) FindByUUID(uuid string) models.Box {
	var (
		box models.Box
	)

	box = b.boxRepository.FindByUUID(uuid)

	return box
}
