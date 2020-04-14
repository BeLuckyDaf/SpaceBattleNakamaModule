package main

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama/runtime"
)

func MakeMatch(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, entries []runtime.MatchmakerEntry) (string, error) {
	for _, e := range entries {
		logger.Info("Matched user '%s' named '%s'", e.GetPresence().GetUserId(), e.GetPresence().GetUsername())

		for k, v := range e.GetProperties() {
			logger.Info("Matched on '%s' value '%v'", k, v)
		}
	}

	matchId, err := nk.MatchCreate(ctx, "spacebattle", map[string]interface{}{"invited": entries})
	if err != nil {
		return "", err
	}

	return matchId, nil
}
