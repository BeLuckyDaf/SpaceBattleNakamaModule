package services

import (
	"context"
	"database/sql"
	"encoding/json"
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
		// On player leaving the match entirely, not just disconnecting
		case types.CommandPlayerLeft:
			break
		// On player moving from one point to another
		case types.CommandPlayerMove:
			s.handlePlayerMove(logger, dispatcher, state, message)
			break
		// On player buying property
		case types.CommandPlayerBuyProperty:
			s.handlePlayerBuyProperty(logger, dispatcher, state, message)
			break
		case types.CommandPlayerUpgradeProperty:
			s.handlePlayerUpgradeProperty(logger, dispatcher, state, message)
			break
		case types.CommandPlayerAttackPlayer:
			break
		case types.CommandPlayerAttackProperty:
			s.handlePlayerAttackProperty(logger, dispatcher, state, message)
			break
		case types.CommandPlayerHeal:
			s.handlePlayerHeal(logger, dispatcher, state, message)
			break
		case types.CommandPlayerRespawned:
			s.handlePlayerRespawned(logger, dispatcher, state, message)
			break
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

	out := types.PayloadPlayerUpdateMove{
		UID:    uid,
		From:   state.Room.Players[uid].Location,
		To:     payload.Location,
		Result: types.PayloadResult{Ok: true, Message: "OK"},
	}

	// validity checks here
	if !state.Room.GameWorld.Points[out.From].IsAdjacent(out.To) {
		out.Result.Ok = false
		out.Result.Message = "POINT_NOT_ADJACENT"
	} else if state.Room.Players[uid].Power <= 0 {
		out.Result.Ok = false
		out.Result.Message = "NOT_ENOUGH_POWER"
	}

	outData, err := json.Marshal(out)
	if err != nil {
		// TODO: log error
	} else {
		if out.Result.Ok { // make sure not to make changes if can't send back
			state.Room.Players[uid].Power -= s.config.KMovementCost
			state.Room.Players[uid].Location = out.To
			logger.Info("Player %s moved from %d to %d.", message.GetUsername(), out.From, out.To)
		}
		dispatcher.BroadcastMessage(types.CommandPlayerMove, outData, nil, nil, true)
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
		UID:      uid,
		Location: payload.Location,
		Result:   types.PayloadResult{Ok: true, Message: "OK"},
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

	// or not enough power points
	if state.Room.GameWorld.Points[out.Location].OwnerUID != "" {
		out.Result.Ok = false
		out.Result.Message = "POINT_ALREADY_OWNED"
	} else if state.Room.Players[uid].Location != out.Location {
		out.Result.Ok = false
		out.Result.Message = "POINT_TOO_FAR"
	} else if state.Room.Players[uid].Power < cost {
		out.Result.Ok = false
		out.Result.Message = "NOT_ENOUGH_POWER"
	}

	outData, err := json.Marshal(out)
	if err != nil {
		// TODO: log error
	} else {
		if out.Result.Ok {
			state.Room.GameWorld.Points[out.Location].OwnerUID = uid
			state.Room.Players[uid].Power -= cost
			logger.Info("Player %s bought %d.", message.GetUsername(), out.Location)
		} // TODO: if error send only to that client
		dispatcher.BroadcastMessage(types.CommandPlayerBuyProperty, outData, nil, nil, true)
	}
}

func (s *SBUserMessageHandlerService) handlePlayerUpgradeProperty(logger runtime.Logger, dispatcher runtime.MatchDispatcher, state *types.MatchState, message runtime.MatchData) {
	uid := message.GetUserId()
	data := message.GetData()
	payload := types.PayloadPlayerInputUpgradeProperty{}
	if serialization.Deserialize(data, &payload, logger) == false {
		return
	}
	out := types.PayloadPlayerUpdateUpgradeProperty{
		UID:      uid,
		Location: payload.Location,
		Result:   types.PayloadResult{Ok: true, Message: "OK"},
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

	// or not enough power points
	if state.Room.GameWorld.Points[out.Location].OwnerUID != "" { // <- TODO: change to some upgrade property
		out.Result.Ok = false
		out.Result.Message = "POINT_ALREADY_UPGRADED"
	} else if state.Room.Players[uid].Location != out.Location {
		out.Result.Ok = false
		out.Result.Message = "POINT_TOO_FAR"
	} else if state.Room.Players[uid].Power < cost {
		out.Result.Ok = false
		out.Result.Message = "NOT_ENOUGH_POWER"
	}

	// TODO: actually implement upgrading because it's not in the core yet
	// TEMPORARY!
	out.Result.Ok = false
	out.Result.Message = "POINT_ALREADY_UPGRADED"

	outData, err := json.Marshal(out)
	if err != nil {
		// TODO: log error
	} else {
		if out.Result.Ok {
			state.Room.GameWorld.Points[out.Location].OwnerUID = uid
			state.Room.Players[uid].Power -= cost
			logger.Info("Player %s upgraded %d.", message.GetUsername(), out.Location)
		} // TODO: if error send only to that client
		dispatcher.BroadcastMessage(types.CommandPlayerUpgradeProperty, outData, nil, nil, true)
	}
}

func (s *SBUserMessageHandlerService) handlePlayerAttackProperty(logger runtime.Logger, dispatcher runtime.MatchDispatcher, state *types.MatchState, message runtime.MatchData) {
	uid := message.GetUserId()
	data := message.GetData()
	payload := types.PayloadPlayerInputAttackProperty{}
	if serialization.Deserialize(data, &payload, logger) == false {
		return
	}
	out := types.PayloadPlayerUpdateAttackProperty{
		UID:      uid,
		Location: payload.Location,
		Result:   types.PayloadResult{Ok: true, Message: "OK"},
	}

	cost := 1

	// check if there is something to attack
	if state.Room.GameWorld.Points[out.Location].OwnerUID == "" {
		out.Result.Ok = false
		out.Result.Message = "POINT_NOT_OWNED"
	} else if state.Room.Players[uid].Location != out.Location {
		out.Result.Ok = false
		out.Result.Message = "POINT_TOO_FAR"
	} else if state.Room.Players[uid].Power < cost {
		out.Result.Ok = false
		out.Result.Message = "NOT_ENOUGH_POWER"
	}

	outData, err := json.Marshal(out)
	if err != nil {
		// TODO: log error
	} else {
		if out.Result.Ok {
			state.Room.GameWorld.Points[out.Location].OwnerUID = ""
			state.Room.Players[uid].Power -= cost
		} // TODO: if error send only to that client
		dispatcher.BroadcastMessage(types.CommandPlayerAttackProperty, outData, nil, nil, true)
	}
}

func (s *SBUserMessageHandlerService) handlePlayerRespawned(logger runtime.Logger, dispatcher runtime.MatchDispatcher, state *types.MatchState, message runtime.MatchData) {
	uid := message.GetUserId()

	out := types.PayloadPlayerUpdateRespawned{
		UID:      uid,
		Location: state.Room.Players[uid].Location,
		Result:   types.PayloadResult{Ok: true, Message: "OK"},
	}

	if state.Room.Players[uid].Hp > 0 {
		out.Result.Ok = false
		out.Result.Message = "NOT_DEAD"
	}

	outData, err := json.Marshal(out)
	if err != nil {
		// TODO: log error
	} else {
		if out.Result.Ok {
			state.Room.Players[uid].Hp = s.config.KInitialPlayerHealth
			// add spawning on random non-owned location
			// maybe restrict spawning if all locations are
			// owned by other players to eliminate players
		}
		dispatcher.BroadcastMessage(types.CommandPlayerRespawned, outData, nil, nil, true)
	}
}

func (s *SBUserMessageHandlerService) handlePlayerHeal(logger runtime.Logger, dispatcher runtime.MatchDispatcher, state *types.MatchState, message runtime.MatchData) {
	uid := message.GetUserId()

	out := types.PayloadPlayerUpdateHeal{
		UID:      uid,
		Location: state.Room.Players[uid].Location,
		NewHp:    s.config.KMaxHealth,
		Result:   types.PayloadResult{Ok: true, Message: "OK"},
	}

	if state.Room.Players[uid].Hp <= 0 {
		out.Result.Ok = false
		out.Result.Message = "PLAYER_DEAD"
	} else if state.Room.Players[uid].Hp >= s.config.KMaxHealth {
		out.Result.Ok = false
		out.Result.Message = "PLAYER_HEALTH_MAX"
	}

	outData, err := json.Marshal(out)
	if err != nil {
		// TODO: log error
	} else {
		if out.Result.Ok {
			state.Room.Players[uid].Hp = s.config.KMaxHealth
			// add spawning on random non-owned location
			// maybe restrict spawning if all locations are
			// owned by other players to eliminate players
		}
		dispatcher.BroadcastMessage(types.CommandPlayerRespawned, outData, nil, nil, true)
	}
}
