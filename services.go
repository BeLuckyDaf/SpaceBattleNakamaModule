package main

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

// SBServiceInterface is used for different services which are called in MatchLoop
type SBServiceInterface interface {
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

/* =========================== */
/* Space Battle Payday Service */
/* =========================== */

// SBPaydayService is used to handle user messages
type SBPaydayService struct{}

// Run is the main method of SBServiceInterface
func (s *SBPaydayService) Run(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, messages []runtime.MatchData) {

}
