package enum

type CaseBattleStatuses string

const (
	CaseBattleCreated   CaseBattleStatuses = "created"
	CaseBattleStarted   CaseBattleStatuses = "started"
	CaseBattleDestroyed CaseBattleStatuses = "destroyed"
)
