package server

import (
	"context"
	"database/sql"
	"encoding/json"
	"spacebattle/core"
	"spacebattle/services"
	"spacebattle/types"
	"strconv"

	"github.com/heroiclabs/nakama-common/runtime"
)

/* =========== */
/* Match Logic */
/* =========== */

// Match represents the match object
type Match struct {
	services []services.SBServiceInterface
	config   core.SBConfig
}

// MatchInit is called whenever the match is created
func (m *Match) MatchInit(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, params map[string]interface{}) (interface{}, int, string) {
	// create the world here
	// .. spawn all properties and stuff

	tickRate := 5
	label := ""

	// Initialize service variables
	for _, service := range m.services {
		service.Init(&m.config)
	}

	state := &types.MatchState{
		Presences: make(map[string]runtime.Presence),
		Room:      core.NewRoom(&m.config, m.config.KMaxPlayers, m.config.KWorldSize),
		Status:    types.GameStatus{Status: types.KGameStatusNotStarted, WinnerID: ""},
	}

	if name, ok := params["name"].(string); ok {
		state.Name = name
	} else {
		return nil, tickRate, label
	}

	return state, tickRate, label
}

// MatchJoinAttempt is called whenever the player wants to join the game
func (m *Match) MatchJoinAttempt(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presence runtime.Presence, metadata map[string]string) (interface{}, bool, string) {
	// check user for validity
	// check storage for user's games and check if they have more than X games playing, not necessary
	// check if the server is not full

	acceptUser := true
	return state, acceptUser, ""
}

// MatchJoin is called whenever the player is allowed to server and is joining
func (m *Match) MatchJoin(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presences []runtime.Presence) interface{} {
	// add user to state
	// create a player struct with all needed info

	mState, _ := state.(*types.MatchState)
	for _, p := range presences {
		// First, tell all existing players that a new player is coming
		mState.Presences[p.GetUserId()] = p
		// add the player if first time
		if mState.Room.Players[p.GetUserId()] == nil {
			mState.Room.AddPlayer(p.GetUserId(), &m.config)
		}
		player := mState.Room.Players[p.GetUserId()]
		data, _ := json.Marshal(player)
		dispatcher.BroadcastMessage(types.CommandPlayerJoined, data, nil, nil, true)
	}

	data, err := json.Marshal(mState)
	if err != nil {
		logger.Error("Could not json.Marshal the state.")
	}
	dispatcher.BroadcastMessage(types.CommandStateSnapshot, data, presences, nil, true)

	return mState
}

// MatchLeave is called whenever the player left the game or disconnected from server
func (m *Match) MatchLeave(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presences []runtime.Presence) interface{} {
	// basically, do nothing with the user, they are still in the match
	// find how it is handled in Nakama
	// how to show users they are still in a match
	// possibly store this information in user storage

	mState, _ := state.(*types.MatchState)
	for _, p := range presences {
		logger.Info("Player %v left.", p.GetUserId())
		delete(mState.Presences, p.GetUserId())
		data, err := json.Marshal(types.PayloadPlayerUpdateLeft{UID: p.GetUserId()})
		if err != nil {
			logger.Error("Could not json.Marshal PayloadPlayerUpdateLeft.")
		}
		dispatcher.BroadcastMessage(types.CommandPlayerLeft, data, nil, nil, true)
	}

	return mState
}

// MatchLoop is called on every tick of the game
func (m *Match) MatchLoop(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, messages []runtime.MatchData) interface{} {
	// make a list of services, following some interface
	// call some function like 'Update' in all of them
	// passing the same parameters that this function gets
	// except probably for the active presences, since
	// we store all players and services should affect
	// everyone in the game, not only those online

	// active presences must receive a message about
	// a particular action taking place in the world

	mState, ok := state.(*types.MatchState)
	if ok {
		// log info every 128 ticks | ~24 sec
		if tick&0b1111111 == 0 {
			matchID, _ := ctx.Value(runtime.RUNTIME_CTX_MATCH_ID).(string)
			logger.Info("MatchID: %v, Players (online): %v, Players (total): %v", matchID, len(mState.Presences), len(mState.Room.Players))
		}
		// run all services
		for _, service := range m.services {
			service.Update(ctx, logger, db, nk, dispatcher, tick, mState, messages)
		}
	} else {
		logger.Error("Could not cast state, services not called!")
		return state // we got this state, so just spit it back
	}

	return mState
}

// MatchTerminate is called whenever the game has ended or the server is shutting down
func (m *Match) MatchTerminate(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, graceSeconds int) interface{} {
	// gracefully finalize everything important
	// maybe try to make it re-runnable?
	// .. meaning that services should save state somehow?

	message := "Server shutting down in " + strconv.Itoa(graceSeconds) + " seconds."
	dispatcher.BroadcastMessage(2, []byte(message), nil, nil, false)
	return state
}
