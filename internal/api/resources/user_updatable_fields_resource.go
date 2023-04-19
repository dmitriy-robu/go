package resources

import (
	"go-rust-drop/internal/api/models"
	"go-rust-drop/internal/api/utils"
)

type UserUpdatableFieldsResource struct {
	User          models.User
	LevelResource LevelResource
	util          utils.MoneyConvert
}

func (u *UserUpdatableFieldsResource) ToJSON() (map[string]interface{}, error) {
	var updatableFields map[string]interface{}

	updatableFields = map[string]interface{}{
		"balance": u.util.FromCentsToVault(u.User.UserBalance.Balance),
		"level":   u.LevelResource,
	}

	return updatableFields, nil
}
