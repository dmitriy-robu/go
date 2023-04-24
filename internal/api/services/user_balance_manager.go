package services

import (
	"github.com/pkg/errors"
	"go-rust-drop/internal/api/models"
	"go-rust-drop/internal/api/repositories"
	"go-rust-drop/internal/api/utils"
	"net/http"
)

type UserBalanceManager struct {
	user                  *models.User
	userBalanceRepository repositories.UserBalanceRepository
}

func NewUserBalanceManager(user *models.User, ubr repositories.UserBalanceRepository) UserBalanceManager {
	return UserBalanceManager{
		user:                  user,
		userBalanceRepository: ubr,
	}
}

func (ubm UserBalanceManager) AddBalance(amount uint) *utils.Errors {
	var (
		err error
	)

	ubm.user.UserBalance.Balance += amount

	if err = ubm.userBalanceRepository.UpdateUserBalance(ubm.user.UserBalance); err != nil {
		return utils.NewErrors(http.StatusInternalServerError, "Error updating user balance", err)
	}

	return nil
}

func (ubm UserBalanceManager) SubtractBalance(amount uint) *utils.Errors {
	var (
		err error
	)

	if ubm.user.UserBalance.Balance < amount {
		return utils.NewErrors(http.StatusBadRequest, "Not enough balance", errors.New("Not enough balance"))
	}

	ubm.user.UserBalance.Balance -= amount

	if err = ubm.userBalanceRepository.UpdateUserBalance(ubm.user.UserBalance); err != nil {
		return utils.NewErrors(http.StatusInternalServerError, "Error updating user balance", err)
	}

	return nil
}
