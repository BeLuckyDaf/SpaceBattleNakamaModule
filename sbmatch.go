// Copyright 2020 Vladislav Smirnov

package main

import (
	"time"
)

// SBMatch is used as a general representation of a server
type SBMatch struct {
	Room           SBRoom `json:"room"`
	PaytimeEnabled bool   `json:"paytime_enabled"`
	timer          *time.Timer
	timerRunning   bool
}

// NewSBMatch creates a server
func NewSBMatch() *SBMatch {
	return new(SBMatch)
}
