package services

import (
	"github.com/pkg/errors"
	"go-rust-drop/internal/api/models"
	"go-rust-drop/internal/api/repositories"
	"go-rust-drop/internal/api/request"
)

type ReferralService struct {
	referralRepository repositories.ReferralRepository
}

func (rs ReferralService) StoreReferralCode(user *models.User, store *request.StoreUserReferralCode) (*models.User, error) {
	if user.ReferralCode != nil {
		return &models.User{}, errors.New("Referral code already exists")
	}

	user, err := rs.referralRepository.StoreReferralCodeToUser(user, store)
	if err != nil {
		return &models.User{}, errors.Wrap(err, "Error storing referral code to user")
	}

	return user, nil
}
