package services

import (
	"github.com/pkg/errors"
	"go-rust-drop/internal/api/models"
	"go-rust-drop/internal/api/repositories"
)

type UserBalanceManager struct {
	user                  models.User
	userBalanceRepository repositories.UserBalanceRepository
}

func (ubm UserBalanceManager) AddBalance(amount int) error {
	var (
		err error
	)

	ubm.user.UserBalance.Balance += amount

	if err = ubm.userBalanceRepository.UpdateUserBalance(ubm.user.UserBalance); err != nil {
		return errors.Wrap(err, "Error updating user balance")
	}

	return nil
}

func (ubm UserBalanceManager) SubtractBalance(amount int) error {
	var (
		err error
	)

	if ubm.user.UserBalance.Balance < amount {
		return errors.New("Insufficient user balance")
	}

	ubm.user.UserBalance.Balance -= amount

	if err = ubm.userBalanceRepository.UpdateUserBalance(ubm.user.UserBalance); err != nil {
		return errors.Wrap(err, "Error updating user balance")
	}

	return nil
}
