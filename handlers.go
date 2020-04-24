package main

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

// MakeMatch handler
func MakeMatch(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, entries []runtime.MatchmakerEntry) (string, error) {
	for _, e := range entries {
		logger.Info("Matched user '%s' named '%s'", e.GetPresence().GetUserId(), e.GetPresence().GetUsername())

		for k, v := range e.GetProperties() {
			logger.Info("Matched on '%s' value '%v'", k, v)
		}
	}

	matchID, err := nk.MatchCreate(ctx, "spacebattle", map[string]interface{}{"invited": entries})
	if err != nil {
		return "", err
	}

	return matchID, nil
}

// AfterAuthenticateEmail handler
func AfterAuthenticateEmail(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.Session, in *api.AuthenticateEmailRequest) error {
	logger.Info("User '%s' has successfully authenticated via Email.", in.Account.GetEmail())
	return nil
}
