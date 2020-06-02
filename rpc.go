package main

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/heroiclabs/nakama-common/runtime"
)

/* ================== */
/* Nakama RPC Methods */
/* ================== */

// CreateMatchRPC is an rpc method that enables players to create matches without the matchmaker
func CreateMatchRPC(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	logger.Info("Payload: %s", payload)

	params := make(map[string]interface{})
	if err := json.Unmarshal([]byte(payload), &params); err != nil {
		return "", err
	}

	// TODO: check 'name' parameter for match name collisions

	modulename := "spacebattle" // Name with which match handler was registered in InitModule, see example above.
	matchID, err := nk.MatchCreate(ctx, modulename, params)
	if err != nil {
		return "", err
	}
	return matchID, nil
}

// GetMyActiveMatchesRPC is an rpc method that enables players to create matches without the matchmaker
func GetMyActiveMatchesRPC(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	logger.Info("Payload: %s", payload)

	minSize := 0
	maxSize := 16
	matches, _ := nk.MatchList(ctx, 10, true, "", &minSize, &maxSize, "")

	// TODO: make a good search for player's active matches
	// maybe store them in persistent storage and clean up
	// whenever Nakama has just started

	data, _ := json.Marshal(matches)

	return string(data), nil
}
