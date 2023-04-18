package services

import (
	"fmt"
	"go-rust-drop/internal/bot/models"
	"go-rust-drop/internal/bot/repositories"
	"sync"
	"time"
)

type TradeService struct {
	tradeRepository repositories.TradeRepository
}

type Burner struct {
	queue      chan models.Trade
	processing bool
	sync.Mutex
}

var (
	burner      *Burner
	burnerOnce  sync.Once
	burnerDelay = 5 * time.Second
)

func getBurner() *Burner {
	burnerOnce.Do(func() {
		burner = &Burner{
			queue:      make(chan models.Trade, 100),
			processing: false,
		}
	})
	return burner
}

func (s *TradeService) CreateOffer(createTradeDTO models.CreateTradeDTO) (models.Trade, error) {
	trade := models.Trade{
		UUID:       createTradeDTO.UUID,
		TradeToken: createTradeDTO.TradeToken,
		SteamID:    createTradeDTO.SteamID,
		Items:      createTradeDTO.Items,
		To:         "client",
	}

	burnerAddInQueue(trade)

	result, err := s.tradeRepository.CreateOffer(trade)
	if err != nil {
		return models.Trade{}, err
	}

	return result, nil
}

func burnerAddInQueue(trade models.Trade) {
	burner := getBurner()
	burner.queue <- trade

	burner.Lock()
	if !burner.processing {
		burner.processing = true
		go burnerSendAllToStorage()
	}
	burner.Unlock()
}

func burnerSendAllToStorage() {
	burner := getBurner()

	for {
		select {
		case trade := <-burner.queue:
			// Добавить логику для отправки trade в хранилище
			fmt.Printf("Sending trade %s to storage\n", trade.UUID)
		case <-time.After(burnerDelay):
			burner.Lock()
			if len(burner.queue) == 0 {
				burner.processing = false
				burner.Unlock()
				return
			}
			burner.Unlock()
		}
	}
}
