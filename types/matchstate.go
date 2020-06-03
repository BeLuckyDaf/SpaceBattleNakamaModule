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
}
