package cmd

import (
	"github.com/spf13/cobra"
	"go-rust-drop/internal/api/database/seeders"
)

func init() {
	rootCmd.AddCommand(dbSeedCmd)
}

var dbSeedCmd = &cobra.Command{
	Use:   "db:seed",
	Short: "Seed the database with some initial data",
	Run: func(cmd *cobra.Command, args []string) {
		//seeders.BoxSeeder{}.Seed()
		//seeders.ItemSeeder{}.Seed()
		seeders.BoxItemSeeder{}.Seed()
	},
}
