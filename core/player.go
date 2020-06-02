// Copyright 2020 Vladislav Smirnov

package core

// SBPlayer is used as a general representation of a player
type SBPlayer struct {
	UID                string `json:"UID"`
	Power              int    `json:"Power"`
	Hp                 int    `json:"HP"`
	Location           int    `json:"Location"`
	HealCostMultiplier int    `json:"HealCostMultiplier"`
}

// Token was removed from Player due to being obsolete
