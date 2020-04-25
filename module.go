package main

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

// InitModule initilizes and registers things
func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	// Register different hooks and rpcs
	LogRegisterError(initializer.RegisterMatchmakerMatched(MakeMatch), logger)
	LogRegisterError(initializer.RegisterAfterAuthenticateEmail(AfterAuthenticateEmail), logger)
	LogRegisterError(initializer.RegisterRpc("create_match_rpc", CreateMatchRPC), logger)
	LogRegisterError(initializer.RegisterMatch("spacebattle", MatchCreateSpaceBattle), logger)

	logger.Info("SpaceBattle module created.")

	return nil
}

// LogRegisterError prints an error if any
func LogRegisterError(err error, logger runtime.Logger) {
	if err != nil {
		logger.Error("Unable to register: %v", err)
	}
}
