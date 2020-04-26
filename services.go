package main

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

// SBServiceInterface is used for different services which are called in MatchLoop
type SBServiceInterface interface {
	Init()
	Run(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, messages []runtime.MatchData)
}

/* ======================== */
/* Space Battle API Service */
/* ======================== */

// SBUserMessageHandlerService is used to handle user messages
type SBUserMessageHandlerService struct{}

// Run is the main method of SBServiceInterface
func (s *SBUserMessageHandlerService) Run(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, messages []runtime.MatchData) {

}

// Init is the initializator method of SBServiceInterface
func (s *SBUserMessageHandlerService) Init() {

}

/* =========================== */
/* Space Battle Payday Service */
/* =========================== */

// SBPaydayService is used to handle user messages
type SBPaydayService struct {
	nextPaydayTime int64
}

// Run is the main method of SBServiceInterface
func (s *SBPaydayService) Run(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, messages []runtime.MatchData) {
	if s.nextPaydayTime < tick {
		s.nextPaydayTime += 10000
		mState, _ := state.(*MatchState)
		for _, player := range mState.Room.Players {
			player.Power += 5
			// TODO: add amount of power that players earned, not 5
		}
	}
}

// Init is the initializator method of SBServiceInterface
func (s *SBPaydayService) Init() {
	s.nextPaydayTime = 0
}
