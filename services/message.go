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

// Init is the initializator method of SBServiceInterface
func (s *SBUserMessageHandlerService) Init(config *core.SBConfig) {
	s.config = config
}

// Update is the main method of SBServiceInterface
func (s *SBUserMessageHandlerService) Update(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state *types.MatchState, messages []runtime.MatchData) {
	for _, message := range messages {
		op := message.GetOpCode()
		switch op {
		// On player leaving the match
		case types.CommandPlayerLeft:
			break
		// On player moving from one point to another
		case types.CommandPlayerMove:
			s.handlePlayerMove(logger, dispatcher, state, message)
		// On player buying property
		case types.CommandPlayerBuyProperty:
			s.handlePlayerBuyProperty(logger, dispatcher, state, message)
		case types.CommandPlayerUpgradeProperty:
			break
		case types.CommandPlayerAttackPlayer:
			break
		case types.CommandPlayerAttackProperty:
			break
		case types.CommandPlayerHeal:
			break
		case types.CommandPlayerRespawned:
			s.handlePlayerRespawned(logger, dispatcher, state, message)
		}
	}
}

func (s *SBUserMessageHandlerService) handlePlayerMove(logger runtime.Logger, dispatcher runtime.MatchDispatcher, state *types.MatchState, message runtime.MatchData) {
	uid := message.GetUserId()
	data := message.GetData()
	payload := types.PayloadPlayerInputMove{}
	if serialization.Deserialize(data, &payload, logger) == false {
		return
	}

	// validity checks here
	if !state.Room.GameWorld.Points[payload.Location].IsAdjacent(state.Room.Players[uid].Location) || state.Room.Players[uid].Power <= 0 {
		payload.Location = state.Room.Players[uid].Location
	} else { // if no problem, remove power for movement
		state.Room.Players[uid].Power -= s.config.KMovementCost
	}

	out := types.PayloadPlayerUpdateMove{
		UID:  uid,
		From: state.Room.Players[uid].Location,
		To:   payload.Location,
	}
	outData := serialization.Serialize(out, logger)
	if outData != nil {
		if !state.Room.GameWorld.Points[out.From].IsAdjacent(out.To) {
			// broadcast state instead
			logger.Error("Player %s can't move, since %d is not adjacent to %d.", message.GetUsername(), out.From, out.To)
			return
		}
		state.Room.Players[uid].Location = out.To
		dispatcher.BroadcastMessage(types.CommandPlayerMove, outData, nil, state.Presences[uid], true)
		logger.Info("Player %s moved from %d to %d.", message.GetUsername(), out.From, out.To)
	}
}

func (s *SBUserMessageHandlerService) handlePlayerBuyProperty(logger runtime.Logger, dispatcher runtime.MatchDispatcher, state *types.MatchState, message runtime.MatchData) {
	uid := message.GetUserId()
	data := message.GetData()
	payload := types.PayloadPlayerInputBuyProperty{}
	if serialization.Deserialize(data, &payload, logger) == false {
		return
	}
	out := types.PayloadPlayerUpdateBuyProperty{
		Location: payload.Location,
	}

	cost := 0
	switch state.Room.GameWorld.Points[out.Location].LocType {
	case core.LoctypeAsteroid:
		cost = s.config.KAsteroidCost
	case core.LoctypePlanet:
		cost = s.config.KPlanetCost
	case core.LoctypeStation:
		cost = s.config.KStationCost
	default:
		cost = 1
	}

	// set out location to -1 if validity checks don't pass, such as if owned already or wrong location
	// or not enough power points
	if state.Room.GameWorld.Points[out.Location].OwnerUID != "" || state.Room.Players[uid].Location != out.Location || state.Room.Players[uid].Power < cost {
		out.Location = -1
	}

	outData := serialization.Serialize(out, logger)
	if outData != nil {
		if out.Location >= 0 {
			state.Room.GameWorld.Points[out.Location].OwnerUID = uid
			state.Room.Players[uid].Power -= cost
		}
		dispatcher.BroadcastMessage(types.CommandPlayerBuyProperty, outData, nil, nil, true)
		logger.Info("Player %s bought %d.", message.GetUsername(), out.Location)
	}
}

func (s *SBUserMessageHandlerService) handlePlayerRespawned(logger runtime.Logger, dispatcher runtime.MatchDispatcher, state *types.MatchState, message runtime.MatchData) {
	uid := message.GetUserId()
	if state.Room.Players[uid].Hp <= 0 {
		state.Room.Players[uid].Hp = s.config.KInitialPlayerHealth
		dispatcher.BroadcastMessage(types.CommandPlayerRespawned, nil, nil, nil, true)
		// add spawning on random non-owned location
		// maybe restrict spawning if all locations are
		// owned by other players to eliminate players
	}
}
