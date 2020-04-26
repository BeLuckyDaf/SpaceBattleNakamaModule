// Copyright 2020 Vladislav Smirnov

package main

import (
	"math/rand"
	"strings"
)

// SBRoom is used as a general representation of a room is the world
type SBRoom struct {
	GameWorld  *SBWorld             `json:"game_world"`
	Players    map[string]*SBPlayer `json:"players"`
	MaxPlayers int                  `json:"max_players"`
}

// NewRoom creates a new room in the world
func NewRoom(config *SBConfig, maxPlayers, worldSize int) SBRoom {
	return SBRoom{
		GameWorld:  GenerateWorld(config, worldSize),
		Players:    make(map[string]*SBPlayer),
		MaxPlayers: maxPlayers,
	}
}

// DeletePlayer removes the client from the room
func (r *SBRoom) DeletePlayer(uid string) {
	delete(r.Players, uid)
	for _, p := range r.GameWorld.Points {
		if strings.Compare(p.OwnerUID, uid) == 0 {
			p.OwnerUID = ""
		}
	}
}

// AddPlayer adds the client to the room
func (r *SBRoom) AddPlayer(uid string, config *SBConfig) bool {
	if len(r.Players) < r.MaxPlayers {
		r.Players[uid] = &SBPlayer{
			UID:                uid,
			Power:              config.KInitialPlayerPower,
			Location:           rand.Intn(r.GameWorld.Size),
			Hp:                 config.KInitialPlayerHealth,
			HealCostMultiplier: 1,
		}
		return true
	}
	return false
}
