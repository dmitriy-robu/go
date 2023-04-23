package services

import (
	"go-rust-drop/internal/api/repositories"
)

type CaseBattleRoundManager struct {
	caseBattleRoundRepository repositories.CaseBattleRoundRepository
}

/*func (cbrm CaseBattleRoundManager) Create(caseBattleRound models.CaseBattleRound) utils.Errors {
	var (
		err   error
		newID primitive.ObjectID
		now   time.Time
	)

	now = utils.GetTimeNow()

	newID = primitive.NewObjectID()
	caseBattleRound.ID = newID
	caseBattleRound.CreatedAt = primitive.NewDateTimeFromTime(now)
	caseBattleRound.UpdatedAt = primitive.NewDateTimeFromTime(now)

	err = cbrm.caseBattleRoundRepository.CreateCaseBattleRound(caseBattleRound)
	if err != nil {
		return utils.Errors{
			Code:    500,
			Message: "Error creating case battle round",
			Err:     err,
		}
	}

	return utils.Errors{}
}
*/
