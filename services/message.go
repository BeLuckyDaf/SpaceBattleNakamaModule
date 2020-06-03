package services

import (
	"context"
	"database/sql"
	"spacebattle/core"
	"spacebattle/serialization"
	"spacebattle/types"

	"github.com/heroiclabs/nakama-common/runtime"
)

/* ======================== */
/* Space Battle API Service */
/* ======================== */

// SBUserMessageHandlerService is used to handle user messages
type SBUserMessageHandlerService struct {
	config *core.SBConfig
}

// Update is the main method of SBServiceInterface
func (s *SBUserMessageHandlerService) Update(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, messages []runtime.MatchData) {
	mState := state.(*types.MatchState)
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
		case types.CommandPlayerLeft:
			break
		// On player moving from one point to another
		case types.CommandPlayerMove:
			// TODO: check if valid location
			// check if correct from location
			// check if have points to move
			// remove points
			payload := types.PayloadPlayerInputMove{}
			if serialization.Deserialize(data, &payload, logger) == false {
				break
			}
			out := types.PayloadPlayerUpdateMove{
				UID:  uid,
				From: mState.Room.Players[uid].Location,
				To:   payload.Location,
			}
			outData := serialization.Serialize(out, logger)
			if outData != nil {
				if !mState.Room.GameWorld.Points[out.From].IsAdjacent(out.To) {
					// broadcast state instead
					logger.Error("Player %s can't move, since %d is not adjacent to %d.", message.GetUsername(), out.From, out.To)
					break
				}
				mState.Room.Players[uid].Location = out.To
				dispatcher.BroadcastMessage(types.CommandPlayerMove, outData, presences, mState.Presences[uid], true)
				logger.Info("Player %s moved from %d to %d.", message.GetUsername(), out.From, out.To)
			}
			break
		// On player buying property
		case types.CommandPlayerBuyProperty:
			// TODO: check if valid location,
			// Remove points, check if owned already
			payload := types.PayloadPlayerInputBuyProperty{}
			if serialization.Deserialize(data, &payload, logger) == false {
				break
			}
			out := types.PayloadPlayerUpdateBuyProperty{
				Location: payload.Location,
			}
			outData := serialization.Serialize(out, logger)
			if outData != nil {
				mState.Room.GameWorld.Points[out.Location].OwnerUID = uid
				dispatcher.BroadcastMessage(types.CommandPlayerBuyProperty, outData, presences, nil, true)
				logger.Info("Player %s bought %d.", message.GetUsername(), out.Location)
			}
			break
		case types.CommandPlayerUpgradeProperty:
			break
		case types.CommandPlayerAttackPlayer:
			break
		case types.CommandPlayerAttackProperty:
			break
		case types.CommandPlayerHeal:
			break
		case types.CommandPlayerRespawned:
			if mState.Room.Players[uid].Hp <= 0 {
				mState.Room.Players[uid].Hp = s.config.KInitialPlayerHealth
				dispatcher.BroadcastMessage(types.CommandPlayerRespawned, nil, presences, nil, true)
				// add spawning on random non-owned location
				// maybe restrict spawning if all locations are
				// owned by other players to eliminate players
			}
			break
		}
	}
}

// Init is the initializator method of SBServiceInterface
func (s *SBUserMessageHandlerService) Init(config *core.SBConfig) {
	s.config = config
}
