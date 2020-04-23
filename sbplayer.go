// Copyright 2020 Vladislav Smirnov

package main

// SBPlayer is used as a general representation of a player
type SBPlayer struct {
	UID                string `json:"uid"`
	Power              int    `json:"power"`
	Hp                 int    `json:"hp"`
	Location           int    `json:"location"`
	HealCostMultiplier int    `json:"-"`
}

// Token was removed from Player due to being obsolete
