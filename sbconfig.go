// Copyright 2020 Vladislav Smirnov

package main

type SBConfig struct {
	KMaxPlayers      int
	KLogfilePath     string
	KWorldSize       int
	KMinimalDistance float64
	KEdgeDistance    float64
	KPaytimeInterval int

	KPlanetCost     int
	KAsteroidCost   int
	KStationCost    int
	KPlanetPayout   int
	KAsteroidPayout int
	KMovementCost   int
	KHealAmount     int

	KInitialPlayerPower  int
	KInitialPlayerHealth int
	KInitialHealingPrice int
	KHealCostMultiplier  int
	KStationDamage       int
	KMaxHealth           int
}
