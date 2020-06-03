package main

import (
	"context"
	"database/sql"
	"spacebattle/server"

	"github.com/heroiclabs/nakama-common/runtime"
)

// InitModule initilizes and registers things
func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	// Register different hooks and rpcs
	LogRegisterError(initializer.RegisterMatchmakerMatched(server.MakeMatch), logger)
	LogRegisterError(initializer.RegisterAfterAuthenticateEmail(server.AfterAuthenticateEmail), logger)
	LogRegisterError(initializer.RegisterRpc("create_match_rpc", server.CreateMatchRPC), logger)
	LogRegisterError(initializer.RegisterRpc("get_my_active_matches", server.GetMyActiveMatchesRPC), logger)
	LogRegisterError(initializer.RegisterMatch("spacebattle", server.MatchCreateSpaceBattle), logger)

	logger.Info("SpaceBattle module created.")

	return nil
}

// LogRegisterError prints an error if any
func LogRegisterError(err error, logger runtime.Logger) {
	if err != nil {
		logger.Error("Unable to register: %v", err)
	}
}
