package backup

import (
	"context"
	"spacebattle/sjson"
	"spacebattle/types"

	"github.com/heroiclabs/nakama-common/runtime"
)

// MatchSaveData is used to save and retrieve match data from db
type MatchSaveData struct {
	MatchID string
	State   *types.MatchState
}

// SaveMatchState is used to save match data to db
func SaveMatchState(ctx context.Context, name string, state *types.MatchState, nk runtime.NakamaModule) bool {
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
				Value:           string(sjson.Marshal(saveData, nil)),
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

// LoadMatchState is used to load match data from db
func LoadMatchState(name string, state *types.MatchState) bool {
	return false
}
