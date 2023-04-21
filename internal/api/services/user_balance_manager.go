package services

import (
	"github.com/pkg/errors"
	"go-rust-drop/internal/api/models"
	"go-rust-drop/internal/api/repositories"
)

type UserBalanceManager struct {
	userRepository        repositories.UserRepository
	userBalanceRepository repositories.UserBalanceRepository
}

func (ubm UserBalanceManager) AddBalance(user models.User, amount int) error {
	var (
		err error
	)

	user.UserBalance.Balance += amount

	if err = ubm.userBalanceRepository.UpdateUserBalance(user.UserBalance); err != nil {
		return errors.Wrap(err, "Error updating user balance")
	}

	return nil
}

func (ubm UserBalanceManager) SubtractBalance(user models.User, amount int) error {
	var (
		err error
	)

	if user.UserBalance.Balance < amount {
		return errors.New("Insufficient user balance")
	}

	user.UserBalance.Balance -= amount

	if err = ubm.userBalanceRepository.UpdateUserBalance(user.UserBalance); err != nil {
		return errors.Wrap(err, "Error updating user balance")
	}

	return nil
}
