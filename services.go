package main

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
)

// SBServiceInterface is used for different services which are called in MatchLoop
type SBServiceInterface interface {
	Init(m *Match)
	Update(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, messages []runtime.MatchData)
}

/* ======================== */
/* Space Battle API Service */
/* ======================== */

// SBUserMessageHandlerService is used to handle user messages
type SBUserMessageHandlerService struct {
	match *Match
}

// Update is the main method of SBServiceInterface
func (s *SBUserMessageHandlerService) Update(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, messages []runtime.MatchData) {
	mState := state.(*MatchState)
	presences := []runtime.Presence{}
	for _, p := range mState.Presences {
		presences = append(presences, p)
	}

	for _, message := range messages {
		uid := message.GetUserId()
		op := message.GetOpCode()
		data := message.GetData()
		switch op {
		// On player leaving the match
		case CommandPlayerLeft:
			break
		// On player moving from one point to another
		case CommandPlayerMove:
			// TODO: check if valid location
			// check if correct from location
			// check if have points to move
			// remove points
			payload := PayloadPlayerInputMove{}
			if Unmarshal(data, &payload, logger) == false {
				break
			}
			out := PayloadPlayerUpdateMove{
				UID:  uid,
				From: mState.Room.Players[uid].Location,
				To:   payload.Location,
			}
			outData := Marshal(out, logger)
			if outData != nil {
				if !mState.Room.GameWorld.Points[out.From].IsAdjacent(out.To) {
					// broadcast state instead
					logger.Error("Player %s can't move, since %d is not adjacent to %d.", message.GetUsername(), out.From, out.To)
					break
				}
				mState.Room.Players[uid].Location = out.To
				dispatcher.BroadcastMessage(CommandPlayerMove, outData, presences, mState.Presences[uid], true)
				logger.Info("Player %s moved from %d to %d.", message.GetUsername(), out.From, out.To)
			}
			break
		// On player buying property
		case CommandPlayerBuyProperty:
			// TODO: check if valid location,
			// Remove points, check if owned already
			payload := PayloadPlayerInputBuyProperty{}
			if Unmarshal(data, &payload, logger) == false {
				break
			}
			out := PayloadPlayerUpdateBuyProperty{
				Location: payload.Location,
			}
			outData := Marshal(out, logger)
			if outData != nil {
				mState.Room.GameWorld.Points[out.Location].OwnerUID = uid
				dispatcher.BroadcastMessage(CommandPlayerBuyProperty, outData, presences, nil, true)
				logger.Info("Player %s bought %d.", message.GetUsername(), out.Location)
			}
			break
		case CommandPlayerUpgradeProperty:
			break
		case CommandPlayerAttackPlayer:
			break
		case CommandPlayerAttackProperty:
			break
		case CommandPlayerHeal:
			break
		case CommandPlayerRespawned:
			if mState.Room.Players[uid].Hp <= 0 {
				mState.Room.Players[uid].Hp = s.match.config.KInitialPlayerHealth
				dispatcher.BroadcastMessage(CommandPlayerRespawned, nil, presences, nil, true)
				// add spawning on random non-owned location
				// maybe restrict spawning if all locations are
				// owned by other players to eliminate players
			}
			break
		}
	}
}

// Init is the initializator method of SBServiceInterface
func (s *SBUserMessageHandlerService) Init(m *Match) {
	s.match = m
}

/* =========================== */
/* Space Battle Payday Service */
/* =========================== */

// SBPaydayService is used to handle user messages
type SBPaydayService struct {
	nextPaydayTime int64
	match          *Match
}

// Update is the main method of SBServiceInterface
func (s *SBPaydayService) Update(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, messages []runtime.MatchData) {
	if s.nextPaydayTime < tick {
		s.nextPaydayTime += 10000
		mState, _ := state.(*MatchState)
		for _, player := range mState.Room.Players {
			player.Power += 5
			// TODO: add amount of power that players earned, not 5
		}
	}
}

// Init is the initializator method of SBServiceInterface
func (s *SBPaydayService) Init(m *Match) {
	s.nextPaydayTime = 0
	s.match = m
}
