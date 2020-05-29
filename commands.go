package main

/* ============================= */
/* Commands & Payload Structures */
/* ============================= */

// Commands represent opcodes for the server
const (
	CommandStateSnapshot         = 7 // server only
	CommandPlayerJoined          = 8 // server only
	CommandPlayerLeft            = 9
	CommandPlayerMove            = 10
	CommandPlayerBuyProperty     = 11
	CommandPlayerUpgradeProperty = 12
	CommandPlayerAttackPlayer    = 13
	CommandPlayerAttackProperty  = 14
	CommandPlayerHeal            = 15
	CommandPlayerKilled          = 16 // server only
	CommandPlayerRespawned       = 17
	CommandGamePause             = 18 // server only
	CommandGameUnpause           = 19 // server only
	CommandGameEnd               = 20 // server only
	CommandGameServerMessage     = 21 // server only
)

// PayloadPlayerInputMove represents new user location
type PayloadPlayerInputMove struct {
	Location int
}

// PayloadPlayerUpdateMove represents new user location
type PayloadPlayerUpdateMove struct {
	UID  string
	From int
	To   int
}

// PayloadPlayerInputBuyProperty represents new user location
type PayloadPlayerInputBuyProperty struct {
	Location int
}

// PayloadPlayerUpdateBuyProperty represents new user location
type PayloadPlayerUpdateBuyProperty struct {
	Location int
}
