package types

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

type PayloadResult struct {
	Ok      bool
	Message string
}

// PayloadPlayerInputMove represents new user location
type PayloadPlayerInputMove struct {
	Location int
}

// PayloadPlayerUpdateMove represents new user location
type PayloadPlayerUpdateMove struct {
	UID    string
	From   int
	To     int
	Result PayloadResult
}

// PayloadPlayerUpdateLeft represents the fact that a player has left
type PayloadPlayerUpdateLeft struct {
	UID string
}

// PayloadPlayerInputBuyProperty represents new user property
type PayloadPlayerInputBuyProperty struct {
	Location int
}

// PayloadPlayerUpdateBuyProperty represents new user property
type PayloadPlayerUpdateBuyProperty struct {
	UID      string
	Location int
	Result   PayloadResult
}

// PayloadPlayerInputUpgradeProperty represents new user location
type PayloadPlayerInputUpgradeProperty struct {
	Location int
}

// PayloadPlayerUpdateUpgradeProperty represents new user location
type PayloadPlayerUpdateUpgradeProperty struct {
	UID      string
	Location int
	Result   PayloadResult
}

// PayloadPlayerInputAttackProperty represents new user location
type PayloadPlayerInputAttackProperty struct {
	Location int
}

// PayloadPlayerUpdateAttackProperty represents user attacking property
type PayloadPlayerUpdateAttackProperty struct {
	UID      string
	Location int
	Result   PayloadResult
}

// PayloadPlayerUpdateHeal represents user healing
type PayloadPlayerUpdateHeal struct {
	UID      string
	Location int
	NewHp    int
	Result   PayloadResult
}

// PayloadPlayerUpdateRespawned represents newly spawned user
type PayloadPlayerUpdateRespawned struct {
	UID      string
	Location int
	Result   PayloadResult
}
