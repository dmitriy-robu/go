package main

import (
	"go-rust-drop/internal/api/controllers"
	"go-rust-drop/internal/api/repositories"
	"go-rust-drop/internal/api/services"
	"go.uber.org/dig"
)

func NewContainer() *dig.Container {
	container := dig.New()

	container.Provide(repositories.NewCaseBattleRoundRepository)
	container.Provide(repositories.NewBoxRepository)
	container.Provide(repositories.NewCaseBattleRepository)
	container.Provide(repositories.NewUserBalanceRepository)
	container.Provide(repositories.NewUserRepository)
	container.Provide(repositories.NewLevelRepository)
	container.Provide(repositories.NewReferralRepository)
	container.Provide(repositories.NewSteamRepository)

	container.Provide(services.NewCaseBattleManager)
	container.Provide(services.NewUserBalanceManager)
	container.Provide(services.NewLevelManager)
	container.Provide(services.NewUserInventoryManager)
	container.Provide(services.NewReferralManager)
	container.Provide(services.NewBoxManager)
	container.Provide(services.NewSteamAuthManager)
	container.Provide(services.NewUserManager)
	container.Provide(services.NewProjectStatisticsManager)

	container.Provide(controllers.NewCaseBattleController)
	container.Provide(controllers.NewReferralController)
	container.Provide(controllers.NewUserController)
	container.Provide(controllers.NewBoxController)
	container.Provide(controllers.NewSteamAuthController)
	container.Provide(controllers.NewSteamAuthController)
	container.Provide(controllers.NewProjectStatisticController)
	container.Provide(controllers.NewControllers)

	return container
}
