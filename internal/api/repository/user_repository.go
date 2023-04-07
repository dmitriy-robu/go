package repository

import (
	"database/sql"
	"github.com/doug-martin/goqu/v9"
	"go-rust-drop/internal/api/models"
)

type UserRepository struct {
}

func (u UserRepository) FindUserByIDWithBalance(db *sql.DB, userID uint64) (models.UserWithBalance, error) {
	var err error

	ds := goqu.From("users").Select(
		"users.id as user_id",
		"users.uuid",
		"users.name",
		"users.avatar_url",
		"users.email",
		"users.email_verified_at",
		"users.password",
		"users.steam_trade_url",
		"users.experience",
		"users.active",
		"users.is_bot",
		"users.remember_token",
		"users.deleted_at",
		"users.created_at",
		"users.updated_at",
		"user_balances.id as balance_id",
		"user_balances.balance",
	).InnerJoin(
		goqu.T("user_balances"),
		goqu.On(goqu.Ex{"users.id": goqu.I("user_balances.user_id")}),
	).Where(
		goqu.Ex{"users.id": userID},
	)

	query, _, err := ds.ToSQL()
	if err != nil {
		return models.UserWithBalance{}, err
	}

	row := db.QueryRow(query)

	var user models.UserWithBalance
	err = row.Scan(
		&user.User.ID,
		&user.UserBalance.Balance,
	)
	if err != nil {
		return models.UserWithBalance{}, err
	}

	return user, nil
}

func (u UserRepository) FindUserByID(db *sql.DB, userID uint64) (models.User, error) {
	var err error

	ds := goqu.From("users").Select(
		"id",
		"uuid",
		"name",
		"avatar_url",
		"email",
		"email_verified_at",
		"password",
		"steam_trade_url",
		"experience",
		"active",
		"is_bot",
		"remember_token",
		"deleted_at",
		"created_at",
		"updated_at",
	).Where(
		goqu.Ex{"id": userID},
	)

	query, _, err := ds.ToSQL()
	if err != nil {
		return models.User{}, err
	}

	row := db.QueryRow(query, userID)

	var user models.User
	err = row.Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.AvatarURL,
		&user.Email,
		&user.EmailVerifiedAt,
		&user.Password,
		&user.SteamTradeURL,
		&user.Experience,
		&user.Active,
		&user.IsBot,
		&user.RememberToken,
		&user.DeletedAt,
		&user.CreatedAt,
		&user.UpdatedAt)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (u UserRepository) GetUserBalance(db *sql.DB, userID uint64) (models.UserBalance, error) {
	var err error

	ds := goqu.From("user_balances").Select(
		"id",
		"user_id",
		"balance",
	).Where(
		goqu.Ex{"user_id": userID},
	)

	query, _, err := ds.ToSQL()
	if err != nil {
		return models.UserBalance{}, err
	}

	row := db.QueryRow(query, userID)

	var userBalance models.UserBalance
	err = row.Scan(&userBalance.ID, &userBalance.UserID, &userBalance.Balance)
	if err != nil {
		return models.UserBalance{}, err
	}

	return userBalance, nil
}
