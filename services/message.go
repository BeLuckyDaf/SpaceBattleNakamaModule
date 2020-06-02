package services

import (
	"context"
	"database/sql"
	"spacebattle/commands"
	"spacebattle/core"
	"spacebattle/matchstate"
	"spacebattle/sjson"

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
	mState := state.(*matchstate.MatchState)
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
		case commands.CommandPlayerLeft:
			break
		// On player moving from one point to another
		case commands.CommandPlayerMove:
			// TODO: check if valid location
			// check if correct from location
			// check if have points to move
			// remove points
			payload := commands.PayloadPlayerInputMove{}
			if sjson.Unmarshal(data, &payload, logger) == false {
				break
			}
			out := commands.PayloadPlayerUpdateMove{
				UID:  uid,
				From: mState.Room.Players[uid].Location,
				To:   payload.Location,
			}
			outData := sjson.Marshal(out, logger)
			if outData != nil {
				if !mState.Room.GameWorld.Points[out.From].IsAdjacent(out.To) {
					// broadcast state instead
					logger.Error("Player %s can't move, since %d is not adjacent to %d.", message.GetUsername(), out.From, out.To)
					break
				}
				mState.Room.Players[uid].Location = out.To
				dispatcher.BroadcastMessage(commands.CommandPlayerMove, outData, presences, mState.Presences[uid], true)
				logger.Info("Player %s moved from %d to %d.", message.GetUsername(), out.From, out.To)
			}
			break
		// On player buying property
		case commands.CommandPlayerBuyProperty:
			// TODO: check if valid location,
			// Remove points, check if owned already
			payload := commands.PayloadPlayerInputBuyProperty{}
			if sjson.Unmarshal(data, &payload, logger) == false {
				break
			}
			out := commands.PayloadPlayerUpdateBuyProperty{
				Location: payload.Location,
			}
			outData := sjson.Marshal(out, logger)
			if outData != nil {
				mState.Room.GameWorld.Points[out.Location].OwnerUID = uid
				dispatcher.BroadcastMessage(commands.CommandPlayerBuyProperty, outData, presences, nil, true)
				logger.Info("Player %s bought %d.", message.GetUsername(), out.Location)
			}
			break
		case commands.CommandPlayerUpgradeProperty:
			break
		case commands.CommandPlayerAttackPlayer:
			break
		case commands.CommandPlayerAttackProperty:
			break
		case commands.CommandPlayerHeal:
			break
		case commands.CommandPlayerRespawned:
			if mState.Room.Players[uid].Hp <= 0 {
				mState.Room.Players[uid].Hp = s.config.KInitialPlayerHealth
				dispatcher.BroadcastMessage(commands.CommandPlayerRespawned, nil, presences, nil, true)
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
