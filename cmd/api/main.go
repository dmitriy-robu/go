package main

import (
	"crypto/rand"
	"encoding/base64"
	"flag"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/steam"
	"go-rust-drop/internal/api/database/migrations"
	"go-rust-drop/internal/api/database/mongodb"
	"go-rust-drop/internal/api/routes"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	var err error

	if err = godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
		return
	}

	genKeyFlag := flag.Bool("genkey", false, "Generate a new session secret key and save it to .env")
	flag.Parse()

	if *genKeyFlag {
		key, err := generateRandomKey(32)
		if err != nil {
			log.Fatalf("Failed to generate random key: %v", err)
			return
		}

		if err = appendKeyToFile(key, ".env"); err != nil {
			log.Fatalf("Failed to append key to .env file: %v", err)
			return
		}

		log.Printf("New session secret key generated and saved to .env: %s", key)
	}

	r := gin.Default()

	goth.UseProviders(
		steam.New(os.Getenv("STEAM_KEY"), os.Getenv("STEAM_CALLBACK_URL")),
	)

	mongoStore, err := mongodb.InitMongoSessionStore()
	if err != nil {
		log.Fatalf("Failed to initialize MongoDB session store: %v", err)
	}

	r.Use(sessions.Sessions("mysession", mongoStore))

	routes.RouteHandle(r)

	go migrations.Migrations{}.MigrateAll()

	if err = r.Run(":" + os.Getenv("GO_PORT")); err != nil {
		log.Fatalf("Failed to run server: %v", err)
		return
	}
}

func generateRandomKey(length int) (string, error) {
	key := make([]byte, length)

	_, err := io.ReadFull(rand.Reader, key)
	if err != nil {
		return "", err
	}

	keyString := base64.StdEncoding.EncodeToString(key)

	return keyString, nil
}

func appendKeyToFile(key string, envFile string) error {
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
