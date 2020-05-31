package main

import (
	"context"

	"github.com/heroiclabs/nakama-common/runtime"
)

// MatchSaveData is used to save and retrieve match data from db
type MatchSaveData struct {
	MatchID string
	State   *MatchState
}

// SaveMatchState is used to save match data to db
func SaveMatchState(ctx context.Context, name string, state *MatchState, nk runtime.NakamaModule) bool {
	if state == nil {
		return false
	}

	matchID, ctxSuccess := ctx.Value(runtime.RUNTIME_CTX_MATCH_ID).(string)
	saveData := &MatchSaveData{MatchID: matchID, State: state}
	if ctxSuccess {
		objects := []*runtime.StorageWrite{
			{
				Collection:      "matches",
				Key:             name,
				Value:           string(Marshal(saveData, nil)),
				PermissionRead:  2,
				PermissionWrite: 0,
			},
		}
		if _, writeErr := nk.StorageWrite(ctx, objects); writeErr != nil {
			return false
		}

		return true
	}

	return false
}
