package main

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	// Register as matchmaker matched hook, this call should be in InitModule.
	if err := initializer.RegisterMatchmakerMatched(MakeMatch); err != nil {
		logger.Error("Unable to register: %v", err)
		return err
	}

	logger.Info("SpaceBattle module created.")

	return nil
}
