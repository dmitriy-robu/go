package seeders

import (
	"github.com/go-faker/faker/v4"
	"go-rust-drop/internal/api/database/mysql"
	"go-rust-drop/internal/api/models"
	"gorm.io/gorm"
)

type ItemSeeder struct {
}

func (b ItemSeeder) Seed() {
	var (
		db                *gorm.DB
		err               error
		model             *models.Item
		countOfIterations = 10
	)

	db, err = mysql.GetGormConnection()

	for i := 0; i < countOfIterations; i++ {
		if err = faker.FakeData(&model); err != nil {
			continue
		}

		db.Create(&models.Item{
			UUID:            model.UUID,
			Name:            model.Name,
			Price:           model.Price,
			Color:           model.Color,
			GameEnvironment: model.GameEnvironment,
			ImageUrl:        model.ImageUrl,
		})
	}
}
