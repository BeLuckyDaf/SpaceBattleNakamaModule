package services

import (
	"context"
	"database/sql"
	"spacebattle/core"
	"spacebattle/types"

	"github.com/heroiclabs/nakama-common/runtime"
)

/* =========================== */
/* Space Battle Payday Service */
/* =========================== */

// SBPaydayService is used to handle user messages
type SBPaydayService struct {
	nextPaydayTime int64
}

// Init is the initializator method of SBServiceInterface
func (s *SBPaydayService) Init(config *core.SBConfig) {
	s.nextPaydayTime = 0
}

// Update is the main method of SBServiceInterface
func (s *SBPaydayService) Update(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state *types.MatchState, messages []runtime.MatchData) {
	if s.nextPaydayTime < tick {
		s.nextPaydayTime += 10000
		for _, player := range state.Room.Players {
			player.Power += 5
			// TODO: add amount of power that players earned, not 5
		}
	}
}
