package repositories

import (
	"context"
	"github.com/pkg/errors"
	"go-rust-drop/internal/bot/database/mongodb"
	"go-rust-drop/internal/bot/models"
	"time"
)

type TradeRepository struct {
}

func (tr TradeRepository) CreateOffer(trade models.Trade) (models.Trade, error) {
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection, err := mongodb.GetCollectionByName("trades")
	if err != nil {
		return models.Trade{}, errors.Wrap(err, "Error getting MongoDB collection")
	}

	_, err = collection.InsertOne(ctx, trade)
	if err != nil {
		return models.Trade{}, errors.Wrap(err, "Error inserting UserAuthSteam into MongoDB")
	}

	return trade, nil
}
