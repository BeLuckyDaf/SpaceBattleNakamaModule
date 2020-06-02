package services

import (
	"context"
	"database/sql"
	"spacebattle/core"

	"github.com/heroiclabs/nakama-common/runtime"
)

/* ========================== */
/* Server & Gameplay Services */
/* ========================== */

// SBServiceInterface is used for different services which are called in MatchLoop
type SBServiceInterface interface {
	Init(config *core.SBConfig)
	Update(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, messages []runtime.MatchData)
}
