package services

import (
	"context"
	"database/sql"
	"spacebattle/core"
	"spacebattle/types"

	"github.com/heroiclabs/nakama-common/runtime"
)

/* ========================== */
/* Server & Gameplay Services */
/* ========================== */

/**
 * Services include bots and other active in-game entities as well, not only system events.
 *
 * They can be used to change the state of the game due to different events or simply to
 * serve as a notifier and periodically send in-game messages. Basically, on every tick a service
 * has an ability to change game state, broadcast messages to online players, write to the
 * persistent storage, access the match context, see players messages that arrived since the last tick.
 */

// SBServiceInterface is used for different services which are called in MatchLoop
type SBServiceInterface interface {
	Init(config *core.SBConfig)
	Update(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state *types.MatchState, messages []runtime.MatchData)
}
