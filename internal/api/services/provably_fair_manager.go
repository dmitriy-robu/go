package services

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"go-rust-drop/internal/api/repositories"
	"strconv"
)

type ProvablyFairManager struct {
	provableFairRepository repositories.ProvablyFairRepository
}

func NewProvablyFairManager() ProvablyFairManager {
	return ProvablyFairManager{
		provableFairRepository: repositories.NewProvablyFairRepository(),
	}
}

func (pfm ProvablyFairManager) GenerateServerSeed() string {
	serverSeed, _ := generateRandomSeed(64)

	return serverSeed
}

func (pfs *ProvablyFairManager) GetWinPercent(
	clientSeed string,
	serverSeed string,
) float64 {
	var (
		message      string
		hash         []byte
		partOfHash   string
		decimal      uint64
		randomNumber float64
		nonce        int
	)

	nonce++

	MinChance := 0.00
	MaxChance := 100.00

	message = fmt.Sprintf("%s-%d", clientSeed, nonce)
	hash = createHmac(serverSeed, message)

	partOfHash = hex.EncodeToString(hash)[:5]
	decimal, _ = strconv.ParseUint(partOfHash, 16, 64)

	const maxHexValue float64 = 1048575
	randomNumber = float64(decimal) / maxHexValue
	randomNumber = MinChance + (MaxChance-MinChance)*randomNumber

	return randomNumber
}

func generateRandomSeed(length int) (string, error) {
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
