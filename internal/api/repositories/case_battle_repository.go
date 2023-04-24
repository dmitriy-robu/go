package repositories

import (
	"context"
	"github.com/pkg/errors"
	"go-rust-drop/internal/api/database/mongodb"
	"go-rust-drop/internal/api/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type CaseBattleRepository struct {
}

func NewCaseBattleRepository() CaseBattleRepository {
	return CaseBattleRepository{}
}

func (cbr CaseBattleRepository) CreateCaseBattle(ctx context.Context, caseBattle models.CaseBattle) error {
	var (
		err        error
		collection *mongo.Collection
	)

	collection, err = mongodb.GetCollectionByName("case_battle")
	if err != nil {
		return errors.Wrap(err, "Error getting MongoDB collection")
	}

	_, err = collection.InsertOne(ctx, caseBattle)
	if err != nil {
		return errors.Wrap(err, "Error inserting UserAuthSteam into MongoDB")
	}

	return nil
}
