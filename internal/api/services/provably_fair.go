package services

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/pkg/errors"
	"go-rust-drop/internal/api/models"
	"strconv"
)

type ProvablyFairService struct {
}

func (pfs ProvablyFairService) GetProvablyFair(provablyFair models.ProvablyFair) (models.ProvablyFair, error) {
	var (
		err        error
		seed       string
		message    string
		hash       []byte
		partOfHash string
		decimal    uint64
	)

	seed, err = generateRandomServerSeed(64)
	if err != nil {
		return models.ProvablyFair{}, errors.Wrap(err, "Error generating random server seed")
	}

	provablyFair.Nonce++
	provablyFair.MinChance = 0.00
	provablyFair.ServerSeed = seed

	message = fmt.Sprintf("%s-%d", provablyFair.ClientSeed, provablyFair.Nonce)
	hash = createHmac(provablyFair.ServerSeed, message)

	partOfHash = hex.EncodeToString(hash)[:5]
	decimal, err = strconv.ParseUint(partOfHash, 16, 64)
	if err != nil {
		return models.ProvablyFair{}, errors.Wrap(err, "Error parsing int")
	}

	const maxHexValue float64 = 1048575
	result := float64(decimal) / maxHexValue
	result = provablyFair.MinChance + (provablyFair.MaxChance-provablyFair.MinChance)*result

	return models.ProvablyFair{
		ClientSeed:   provablyFair.ClientSeed,
		ServerSeed:   provablyFair.ServerSeed,
		Nonce:        provablyFair.Nonce,
		RandomNumber: result,
		MinChance:    provablyFair.MinChance,
		MaxChance:    provablyFair.MaxChance,
	}, nil
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
