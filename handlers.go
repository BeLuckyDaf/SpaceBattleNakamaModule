package main

import (
	"context"
	"database/sql"
	"encoding/json"

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

// MatchCreateSpaceBattle match creator handler
func MatchCreateSpaceBattle(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule) (runtime.Match, error) {
	return &Match{services: nil}, nil
}

// CreateMatchRPC is an rpc method that enables players to create matches without the matchmaker
func CreateMatchRPC(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	logger.Info("Payload: %s", payload)

	params := make(map[string]interface{})
	if err := json.Unmarshal([]byte(payload), &params); err != nil {
		return "", err
	}

	modulename := "spacebattle" // Name with which match handler was registered in InitModule, see example above.
	matchID, err := nk.MatchCreate(ctx, modulename, params)
	if err != nil {
		return "", err
	}
	return matchID, nil
}
