package repositories

import (
	"context"
	"github.com/pkg/errors"
	"go-rust-drop/internal/api/database/mongodb"
	"go-rust-drop/internal/api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type CaseBattleRoundRepository struct {
}

func NewCaseBattleRoundRepository() CaseBattleRoundRepository {
	return CaseBattleRoundRepository{}
}

func (cbrr CaseBattleRoundRepository) CreateCaseBattleRound(ctx context.Context, caseBattleRound models.CaseBattleRound) error {
	var (
		err        error
		collection *mongo.Collection
	)

	collection, err = mongodb.GetCollectionByName("case_battle_rounds")
	if err != nil {
		return errors.Wrap(err, "Error getting MongoDB collection")
	}

	_, err = collection.InsertOne(ctx, caseBattleRound)
	if err != nil {
		return errors.Wrap(err, "Error inserting UserAuthSteam into MongoDB")
	}

	return nil
}

func (cbrr CaseBattleRoundRepository) FindCaseBattleRoundByID(id primitive.ObjectID) (models.CaseBattleRound, error) {
	var (
		err        error
		collection *mongo.Collection
		result     models.CaseBattleRound
	)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection, err = mongodb.GetCollectionByName("case_battle_rounds")
	if err != nil {
		return models.CaseBattleRound{}, errors.Wrap(err, "Error getting MongoDB collection")
	}

	err = collection.FindOne(ctx, models.CaseBattleRound{ID: id}).Decode(&result)
	if err != nil {
		return models.CaseBattleRound{}, errors.Wrap(err, "Error finding user auth steam")
	}

	return result, nil
}

func (cbrr CaseBattleRoundRepository) FindAllCaseBattleRound() (models.CaseBattleRounds, error) {
	var (
		err              error
		collection       *mongo.Collection
		caseBattleRounds models.CaseBattleRounds
		cursor           *mongo.Cursor
	)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection, err = mongodb.GetCollectionByName("case_battle_rounds")
	if err != nil {
		return caseBattleRounds, errors.Wrap(err, "Error getting MongoDB collection")
	}

	cursor, err = collection.Find(ctx, models.CaseBattleRound{})
	if err != nil {
		return caseBattleRounds, errors.Wrap(err, "Error finding all user auth steam")
	}

	for cursor.Next(ctx) {
		var userAuthSteam models.CaseBattleRound
		err = cursor.Decode(&userAuthSteam)
		if err != nil {
			return caseBattleRounds, errors.Wrap(err, "Error decoding user auth steam")
		}

		caseBattleRounds = append(caseBattleRounds, userAuthSteam)
	}

	return caseBattleRounds, nil
}
