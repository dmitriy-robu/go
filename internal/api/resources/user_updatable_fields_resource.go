package resources

import (
	"go-rust-drop/internal/api/models"
	"go-rust-drop/internal/api/utils"
)

type UserUpdatableFieldsResource struct {
	User  models.User
	Level models.Level
	util  utils.MoneyConvert
}

func (u *UserUpdatableFieldsResource) ToJSON() (map[string]interface{}, error) {
	var updatableFields map[string]interface{}

	levelResource := LevelResource{
		Level: u.Level,
	}

	levelResourceJSON, _ := levelResource.ToJSON()

	updatableFields = map[string]interface{}{
		"balance": u.util.FromCentsToVault(u.User.UserBalance.Balance),
		"level":   levelResourceJSON,
	}

	return updatableFields, nil
}
