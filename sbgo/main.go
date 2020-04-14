// Copyright 2020 Vladislav Smirnov

package main

import "github.com/spf13/viper"

// Slogger is a global logger object
var Slogger *Logger

var s *Server
var api *API

func main() {
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

	Slogger = NewLogger(viper.GetString("LogfilePath"))
	s = NewServer()
	api = NewAPI(s)
	s.Room = NewRoom(viper.GetInt("MaxPlayers"), viper.GetInt("WorldSize"))

	go s.LaunchPaytimeTimer()
	Slogger.Log("Started server at port 34000.")

	api.Start()
}
