// Copyright 2020 Vladislav Smirnov

package main

import (
	"math/rand"
	"strings"

	"github.com/spf13/viper"
)

// SBRoom is used as a general representation of a room is the world
type SBRoom struct {
	GameWorld  *SBWorld           `json:"game_world"`
	Players    map[string]*Player `json:"players"`
	MaxPlayers int                `json:"max_players"`
}

// NewRoom creates a new room in the world
func NewRoom(maxPlayers, worldSize int) SBRoom {
	return SBRoom{
		GameWorld:  GenerateWorld(worldSize),
		Players:    make(map[string]*Player),
		MaxPlayers: maxPlayers,
	}
}

// DeletePlayer removes the client from the room
func (r *SBRoom) DeletePlayer(uid string) {
	delete(r.Players, uid)
	for _, p := range r.GameWorld.Points {
		if strings.Compare(p.OwnedBy, uid) == 0 {
			p.OwnedBy = ""
		}
	}
}

// AddPlayer adds the client to the room
func (r *SBRoom) AddPlayer(uid string, token string) bool {
	if len(r.Players) < r.MaxPlayers {
		r.Players[uid] = &Player{
			UID:                uid,
			Power:              viper.GetInt("InitialPlayerPower"),
			Location:           rand.Intn(r.GameWorld.Size),
			Hp:                 viper.GetInt("InitialPlayerHealth"),
			HealCostMultiplier: 1,
		}
		return true
	}
	return false
}
