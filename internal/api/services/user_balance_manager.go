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
	userBalanceRepository *repositories.UserBalanceRepository
}

func NewUserBalanceManager(user *models.User, ubr *repositories.UserBalanceRepository) *UserBalanceManager {
	return &UserBalanceManager{
		user:                  user,
		userBalanceRepository: ubr,
	}
}

func (ubm UserBalanceManager) AddBalance(amount uint) utils.Errors {
	var (
		err error
	)

	ubm.userBalanceRepository.MysqlDB = MysqlDB

	ubm.user.UserBalance.Balance += amount

	if err = ubm.userBalanceRepository.UpdateUserBalance(ubm.user.UserBalance); err != nil {
		return utils.Errors{
			Code:    http.StatusInternalServerError,
			Message: "Error updating user balance",
			Err:     err,
		}
	}

	return utils.Errors{}
}

func (ubm UserBalanceManager) SubtractBalance(amount uint) utils.Errors {
	var (
		err error
	)

	ubm.userBalanceRepository.MysqlDB = MysqlDB

	if ubm.user.UserBalance.Balance < amount {
		return utils.Errors{
			Code:    http.StatusBadRequest,
			Message: "User balance is not enough",
			Err:     errors.New("User balance is not enough"),
		}
	}

	ubm.user.UserBalance.Balance -= amount

	if err = ubm.userBalanceRepository.UpdateUserBalance(ubm.user.UserBalance); err != nil {
		return utils.Errors{
			Code:    http.StatusInternalServerError,
			Message: "Error updating user balance",
			Err:     err,
		}
	}

	return utils.Errors{}
}
