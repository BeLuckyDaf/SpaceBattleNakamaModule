package services

import (
	"context"
	"database/sql"
	"spacebattle/backup"
	"spacebattle/core"
	"spacebattle/types"

	"github.com/heroiclabs/nakama-common/runtime"
)

/* =========================== */
/* Space Battle Backup Service */
/* =========================== */

// SBMatchBackupService is used to autosave match state for recovery
type SBMatchBackupService struct {
	nextBackupTime  int64
	backupTimeDelay int64
	name            string
}

// Update is the main method of SBServiceInterface
func (s *SBMatchBackupService) Update(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state *types.MatchState, messages []runtime.MatchData) {
	if tick == 0 {
		s.name = state.Name
	}

	if s.nextBackupTime < tick {
		s.nextBackupTime += s.backupTimeDelay

		if s.name != state.Name {
			logger.Error("SBMatchBackupService: match name is different than before %v -> %v!", s.name, state.Name)
			return
		}

		saved := backup.SaveMatchState(ctx, s.name, state, nk)
		if saved {
			logger.Info("Match saved: %v", s.name)
		} else {
			logger.Error("Could not save match: %v", s.name)
		}
	}
}

// Init is the initializator method of SBServiceInterface
func (s *SBMatchBackupService) Init(config *core.SBConfig) {
	s.nextBackupTime = 0
	s.backupTimeDelay = 1000
	s.name = "nil"
}
