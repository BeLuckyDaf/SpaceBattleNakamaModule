// Copyright 2020 Vladislav Smirnov

package main

// Player is used as a general representation of a player
type Player struct {
	Username           string `json:"username"`
	Token              string `json:"-"` // token is not sent on /players
	Power              int    `json:"power"`
	Hp                 int    `json:"hp"`
	Location           int    `json:"location"`
	HealCostMultiplier int    `json:"-"`
}
