package types

import (
	"spacebattle/core"

	"github.com/heroiclabs/nakama-common/runtime"
)

// MatchState represents the state object
type MatchState struct {
	Presences map[string]runtime.Presence
	Room      core.SBRoom
	Name      string
	Status    GameStatus
}

// GameStatus represents the current game status
type GameStatus struct {
	Status   int
	WinnerID string
}

// GameStatus values for
const (
	KGameStatusNotStarted = 0
	KGameStatusRunning    = 1
	KGameStatusPaused     = 2
	KGameStatusFinished   = 3
)
