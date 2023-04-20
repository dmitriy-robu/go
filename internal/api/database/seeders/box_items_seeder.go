package seeders

import (
	"go-rust-drop/internal/api/database/mysql"
	"go-rust-drop/internal/api/enum"
	"go-rust-drop/internal/api/models"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

type BoxItemSeeder struct {
}

func (b BoxItemSeeder) Seed() {
	var (
		db                *gorm.DB
		err               error
		countOfIterations int = 4
	)

	db, err = mysql.GetGormConnection()

	if err != nil {
		panic(err)
	}

	for i := 0; i < countOfIterations; i++ {
		boxId := getRandBox()
		items := getRandItems()

		for _, item := range items {
			db.Create(&models.BoxItem{
				BoxID:  boxId,
				ItemID: item,
				Rarity: getRandRarity(),
			})
		}
	}
}

func getRandBox() uint {
	var (
		db  *gorm.DB
		err error
	)

	db, err = mysql.GetGormConnection()

	if err != nil {
		panic(err)
	}

	var box models.Box
	db.Table("boxes").Where("active = ?", 1).Order("RAND()").First(&box)

	return box.ID
}

func getRandItems() (Ids []uint) {
	var (
		db  *gorm.DB
		err error
	)

	db, err = mysql.GetGormConnection()

	if err != nil {
		panic(err)
	}

	var items []models.Item
	db.Table("items").Order("RAND()").Limit(4).Find(&items)

	for _, item := range items {
		Ids = append(Ids, item.ID)
	}

	return
}

func getRandRarity() enum.BoxItemRarity {
	rand.NewSource(time.Now().UnixNano())
	randIndex := rand.Intn(6)

	switch randIndex {
	case 0:
		return enum.BoxItemRarityLow
	case 1:
		return enum.BoxItemRarityCommon
	case 2:
		return enum.BoxItemRarityUncommon
	case 3:
		return enum.BoxItemRarityRare
	case 4:
		return enum.BoxItemRarityEpic
	default:
		return enum.BoxItemRarityLegendary
	}
}
