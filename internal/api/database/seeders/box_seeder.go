package seeders

import (
	"github.com/go-faker/faker/v4"
	"go-rust-drop/internal/api/database/mysql"
	"go-rust-drop/internal/api/models"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

type BoxSeeder struct {
}

func (b BoxSeeder) Seed() {
	var (
		db                *gorm.DB
		err               error
		model             *models.Box
		countOfIterations = 10
	)

	db, err = mysql.GetGormConnection()

	for i := 0; i < countOfIterations; i++ {
		if err = faker.FakeData(&model); err != nil {
		}

		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)

		if (r1.Intn(100) % 2) == 0 {
			model.Active = true
		} else {
			model.Active = false
		}

		model.ID = 0

		db.Create(&model)
	}
}
