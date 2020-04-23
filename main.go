// Copyright 2020 Vladislav Smirnov

package main

import "github.com/spf13/viper"

var s *Server
var api *API

func main() {
	// TODO: move all this into a separate structure, some fields of which
	// could be passed by the player or matchmaker to create
	// matches with different parameters

	// and remove main() because it's invalid here

	viper.SetDefault("MaxPlayers", 128)
	viper.SetDefault("LogfilePath", "logs.txt")
	viper.SetDefault("WorldSize", 100)
	viper.SetDefault("MinimalDistance", 60.0)
	viper.SetDefault("EdgeDistance", 140.0)
	viper.SetDefault("PaytimeInterval", 5)

	viper.SetDefault("PlanetCost", 2)
	viper.SetDefault("AsteroidCost", 2)
	viper.SetDefault("StationCost", 2)
	viper.SetDefault("PlanetPayout", 1)
	viper.SetDefault("AsteroidPayout", 2)
	viper.SetDefault("MovementCost", 1)
	viper.SetDefault("HealAmount", 1)

	viper.SetDefault("InitialPlayerPower", 3)
	viper.SetDefault("InitialPlayerHealth", 3)
	viper.SetDefault("InitialHealingPrice", 10)
	viper.SetDefault("HealCostMultiplier", 2)
	viper.SetDefault("StationDamage", 1)
	viper.SetDefault("MaxHealth", 3)

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	viper.ReadInConfig()
}
