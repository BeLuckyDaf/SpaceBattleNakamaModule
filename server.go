package main

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/heroiclabs/nakama/runtime"
)

type MatchState struct {
	presences map[string]runtime.Presence
	room      SBRoom
}

type Match struct{}

func (m *Match) MatchInit(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, params map[string]interface{}) (interface{}, int, string) {
	// create the world here
	// .. spawn all properties and stuff

	state := &MatchState{
		presences: make(map[string]runtime.Presence),
		room: NewRoom(16, 50)
	}
	tickRate := 10
	label := ""
	return state, tickRate, label
}

func (m *Match) MatchJoinAttempt(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presence runtime.Presence, metadata map[string]string) (interface{}, bool, string) {
	// check user for validity
	// check storage for user's games and check if they have more than X games playing, not necessary

	acceptUser := true
	return state, acceptUser, ""
}

func (m *Match) MatchJoin(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presences []runtime.Presence) interface{} {
	// add user to state
	// create a player struct with all needed info

	mState, _ := state.(*MatchState)
	for _, p := range presences {
		mState.presences[p.GetUserId()] = p
	}

	return mState
}

func (m *Match) MatchLeave(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presences []runtime.Presence) interface{} {
	// basically, do nothing with the user, they are still in the match
	// find how it is handled in Nakama
	// how to show users they are still in a match
	// possibly store this information in user storage

	mState, _ := state.(*MatchState)
	for _, p := range presences {
		delete(mState.presences, p.GetUserId())
	}

	return mState
}

func (m *Match) MatchLoop(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, messages []runtime.MatchData) interface{} {
	// make a list of services, following some interface
	// call some function like 'run' in all of them
	// passing the same parameters that this function gets
	// except probably for the active presences, since
	// we store all players and services should affect
	// everyone in the game, not only those online

	mState, _ := state.(*MatchState)
	for _, presence := range mState.presences {
		logger.Info("Presence %v named %v", presence.GetUserId(), presence.GetUsername())
	}

	for _, message := range messages {
		logger.Info("Received %v from %v", string(message.GetData()), message.GetUserId())

		dispatcher.BroadcastMessage(1, message.GetData(), []runtime.Presence{message}, nil)
	}

	return mState
}

func (m *Match) MatchTerminate(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, graceSeconds int) interface{} {
	// gracefully finalize everything important
	// maybe try to make it re-runnable?
	// .. meaning that services should save state somehow?

	message := "Server shutting down in " + strconv.Itoa(graceSeconds) + " seconds."
	dispatcher.BroadcastMessage(2, []byte(message), nil, nil)
	return state
}
