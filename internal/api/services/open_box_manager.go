package services

import (
	"github.com/gin-gonic/gin"
	"go-rust-drop/internal/api/enum"
	"go-rust-drop/internal/api/models"
	"go-rust-drop/internal/api/repositories"
	"go-rust-drop/internal/api/utils"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

type OpenBoxManager struct {
	boxManager          BoxManager
	openBoxRepository   repositories.OpenBoxRepository
	provablyFairManager ProvablyFairManager
	mysql               *gorm.DB
}

func NewOpenBoxManager(
	provablyFairManager ProvablyFairManager,
	boxManager BoxManager,
	openBoxRepository repositories.OpenBoxRepository,
) OpenBoxManager {
	return OpenBoxManager{
		boxManager:          boxManager,
		openBoxRepository:   openBoxRepository,
		provablyFairManager: provablyFairManager,
	}
}

func (obm OpenBoxManager) Open(c *gin.Context) (
	item models.BoxItem,
	serverSeed string,
	err *utils.Errors,
) {
	var (
		winPercent float64
		box        models.Box
	)

	box, err = obm.boxManager.FindByUUIDWithItems(c.Param("uuid"))
	if err != nil {
		return
	}

	serverSeed = obm.provablyFairManager.GenerateServerSeed()
	winPercent = obm.provablyFairManager.GetWinPercent(c.PostForm("client_seed"), serverSeed)

	item = getWinItemByWinPercentAndRarity(box.BoxItems, winPercent)

	return
}

func getWinItemByWinPercentAndRarity(
	boxItems []models.BoxItem,
	randomNumber float64,
) (item models.BoxItem) {
	var (
		rarity                enum.BoxItemRarity
		cumulativeProbability float64
		itemsWithSameRarity   []models.BoxItem
	)

	randomNumber = randomNumber / 100 // delimit to 100 to get range from 0 to 1

	var probabilities = []utils.RarityProbability{
		{enum.BoxItemRarityLow, 0.8},
		{enum.BoxItemRarityCommon, 0.5},
		{enum.BoxItemRarityUncommon, 0.3},
		{enum.BoxItemRarityRare, 0.1},
		{enum.BoxItemRarityEpic, 0.075},
		{enum.BoxItemRarityLegendary, 0.025},
	}

	cumulativeProbability = 0.0
	for _, prob := range probabilities {
		cumulativeProbability += prob.Probability
		if randomNumber < cumulativeProbability {
			rarity = prob.Rarity
			break
		}
	}

	for _, boxItem := range boxItems {
		if boxItem.Rarity == rarity {
			itemsWithSameRarity = append(itemsWithSameRarity, boxItem)
		}
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIndex := r.Intn(len(itemsWithSameRarity))

	return itemsWithSameRarity[randomIndex-1]
}
