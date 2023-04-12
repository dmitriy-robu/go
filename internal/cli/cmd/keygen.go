package cmd

import (
	"github.com/spf13/cobra"
	"go-rust-drop/internal/cli/services"
	"log"
)

func init() {
	rootCmd.AddCommand(genkeyCmd)
}

var genkeyCmd = &cobra.Command{
	Use:   "genkey",
	Short: "Generate a new session secret key and save it to .env",
	Run: func(cmd *cobra.Command, args []string) {
		key, err := services.GenerateKeyService{}.GenerateRandomKey(32)
		if err != nil {
			log.Fatalf("Failed to generate random key: %v", err)
			return
		}

		err = services.GenerateKeyService{}.AppendKeyToFile(key, ".env")
		if err != nil {
			log.Fatalf("Failed to append key to .env file: %v", err)
			return
		}

		log.Printf("New session secret key generated and saved to .env: %s", key)
	},
}
