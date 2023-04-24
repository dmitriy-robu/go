package services

import (
	"context"
	"github.com/pkg/errors"
	"go-rust-drop/internal/api/database/mongodb"
	"go-rust-drop/internal/api/enum"
	"go-rust-drop/internal/api/models"
	"go-rust-drop/internal/api/repositories"
	"go-rust-drop/internal/api/requests"
	"go-rust-drop/internal/api/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type CaseBattleManager struct {
	caseBattleRepository      repositories.CaseBattleRepository
	caseBattleRoundRepository repositories.CaseBattleRoundRepository
	boxRepository             repositories.BoxRepository
}

func NewCaseBattleManager(
	caseBattleRepo repositories.CaseBattleRepository,
	caseBattleRoundRepo repositories.CaseBattleRoundRepository,
	boxRepo repositories.BoxRepository,
) CaseBattleManager {
	return CaseBattleManager{
		caseBattleRepository:      caseBattleRepo,
		caseBattleRoundRepository: caseBattleRoundRepo,
		boxRepository:             boxRepo,
	}
}

func (cbm *CaseBattleManager) Create(caseBattleRequest requests.CaseBattleStoreRequest, user models.User) (string, *utils.Errors) {
	var (
		err             error
		caseBattle      models.CaseBattle
		caseBattleRound models.CaseBattleRound
		tx              *gorm.DB
		caseBattleID    primitive.ObjectID
		totalCost       uint
		getBox          models.Box
		errorHandler    *utils.Errors
		now             time.Time
	)

	tx = MysqlDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err = cbm.withTransaction(func(ctx mongo.SessionContext) error {
		caseBattleID = primitive.NewObjectID()

		now = utils.GetTimeNow()

		totalCost = 0
		for _, box := range caseBattleRequest.Boxes {
			getBox, err = cbm.boxRepository.FindByUUID(box.UUID)
			if err != nil {
				return errors.New("Error finding box")
			}
			totalCost += getBox.Price * box.Quantity

			caseBattleRound = models.CaseBattleRound{
				ID:           primitive.NewObjectID(),
				CaseBattleID: caseBattleID.String(),
				BoxUUID:      box.UUID,
				CreatedAt:    primitive.NewDateTimeFromTime(now),
				UpdatedAt:    primitive.NewDateTimeFromTime(now),
			}

			if err = cbm.caseBattleRoundRepository.CreateCaseBattleRound(ctx, caseBattleRound); err != nil {
				return errors.New("Error creating case battle round")
			}
		}

		caseBattle = models.CaseBattle{
			ID:          caseBattleID,
			GameMode:    caseBattleRequest.GameMode,
			PlayersMode: caseBattleRequest.PlayersMode,
			Privacy:     caseBattleRequest.Privacy,
			Status:      enum.CaseBattleCreated,
			TotalCost:   totalCost,
			CreatedAt:   primitive.NewDateTimeFromTime(now),
			UpdatedAt:   primitive.NewDateTimeFromTime(now),
		}

		if err = cbm.caseBattleRepository.CreateCaseBattle(ctx, caseBattle); err != nil {
			return errors.New("Error creating case battle")
		}

		return nil
	})

	if err != nil {
		tx.Rollback()
		return "", utils.NewErrors(http.StatusInternalServerError, "Error creating case battle", err)
	}

	userBalanceManager := NewUserBalanceManager(&user, repositories.UserBalanceRepository{MysqlDB: tx})

	if errorHandler = userBalanceManager.SubtractBalance(totalCost); errorHandler != nil {
		tx.Rollback()
		return "", errorHandler
	}

	if err = tx.Commit().Error; err != nil {
		return "", utils.NewErrors(http.StatusInternalServerError, "Error committing transaction", err)
	}

	return caseBattleID.String(), nil

}

func (cbm *CaseBattleManager) withTransaction(fn func(ctx mongo.SessionContext) error) error {
	mongoClient, err := mongodb.GetMongoDBConnection()
	if err != nil {
		return errors.New("Error getting MongoDB connection")
	}

	session, err := mongoClient.Client.StartSession()
	if err != nil {
		return errors.New("Error starting MongoDB session")
	}
	defer session.EndSession(context.Background())

	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		err = fn(sessCtx)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}

	_, err = session.WithTransaction(context.Background(), callback)
	return err
}
