// Copyright 2020 Vladislav Smirnov

package main

import "math"

// LoctypePlanet is used for identification of the planets
// LoctypeAsteroid is used for identification of the asteroids
// LoctypeStation is used for identification of the stations
const (
	LoctypePlanet   = 0
	LoctypeAsteroid = 1
	LoctypeStation  = 2
)

// Vector2 is a general two-dimensional vector structure
type Vector2 struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// Distance calculates the distance between two vectors
func (v1 Vector2) Distance(v2 Vector2) float64 {
	y := v2.Y - v1.Y
	x := v2.X - v1.X
	return math.Sqrt(float64(x*x + y*y))
}

// WorldPoint is used as a general represention of a point in the world
type WorldPoint struct {
	LocType  int     `json:"loc_type"`
	OwnedBy  string  `json:"owned_by"`
	Position Vector2 `json:"position"`
	Adjacent []int   `json:"adjacent"`
}

// IsAdjacent is used to check if two locations are adjacent to each other
func (w *WorldPoint) IsAdjacent(p int) bool {
	for i := 0; i < len(w.Adjacent); i++ {
		if w.Adjacent[i] == p {
			return true
		}
	}
	return false
}
