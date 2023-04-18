package services

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"go-rust-drop/internal/api/models"
	"strconv"
)

type ProvablyFairService struct {
}

func (pfs *ProvablyFairService) GetProvablyFair(provablyFair *models.ProvablyFair) error {
	var (
		err          error
		message      string
		hash         []byte
		partOfHash   string
		decimal      uint64
		randomNumber float64
	)

	provablyFair.ServerSeed, err = generateRandomServerSeed(64)
	if err != nil {
		return err
	}

	provablyFair.Nonce++
	provablyFair.MinChance = 0.00

	if provablyFair.ClientSeed == "" {
		provablyFair.ClientSeed, err = generateRandomServerSeed(64)
		if err != nil {
			return err
		}
	}

	message = fmt.Sprintf("%s-%d", provablyFair.ClientSeed, provablyFair.Nonce)
	hash = createHmac(provablyFair.ServerSeed, message)

	partOfHash = hex.EncodeToString(hash)[:5]
	decimal, err = strconv.ParseUint(partOfHash, 16, 64)
	if err != nil {
		return err
	}

	const maxHexValue float64 = 1048575
	randomNumber = float64(decimal) / maxHexValue
	randomNumber = provablyFair.MinChance + (provablyFair.MaxChance-provablyFair.MinChance)*randomNumber

	provablyFair.RandomNumber = randomNumber

	return nil
}

func generateRandomServerSeed(length int) (string, error) {
	randomBytes := make([]byte, length/2)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(randomBytes), nil
}

func createHmac(secret, message string) []byte {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	return h.Sum(nil)
}
