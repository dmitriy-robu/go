package services

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"os"
	"strings"
)

type GenerateKeyService struct {
}

func (gks GenerateKeyService) GenerateRandomKey(length int) (string, error) {
	key := make([]byte, length)

	_, err := io.ReadFull(rand.Reader, key)
	if err != nil {
		return "", err
	}

	keyString := base64.StdEncoding.EncodeToString(key)

	return keyString, nil
}

func (gks GenerateKeyService) AppendKeyToFile(key string, envFile string) error {
	content, err := os.ReadFile(envFile)
	if err != nil {
		return err
	}

	lines := strings.Split(string(content), "\n")
	found := false

	for i, line := range lines {
		if strings.HasPrefix(line, "SESSION_SECRET=") {
			lines[i] = "SESSION_SECRET=" + key
			found = true
			break
		}
	}

	if !found {
		lines = append(lines, "SESSION_SECRET="+key)
	}

	newContent := strings.Join(lines, "\n")
	err = os.WriteFile(envFile, []byte(newContent), 0644)

	return err
}
