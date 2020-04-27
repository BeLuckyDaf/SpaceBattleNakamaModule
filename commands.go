package main

// Commands represent opcodes for the server
const (
	CommandStateSnapshot               = 7 // server only
	CommandPlayerJoined                = 8 // server only
	CommandPlayerLeft                  = 9
	CommandPlayerUpdateMove            = 10
	CommandPlayerUpdateBuyProperty     = 11
	CommandPlayerUpdateUpgradeProperty = 12
	CommandPlayerUpdateAttackPlayer    = 13
	CommandPlayerUpdateAttackProperty  = 14
	CommandPlayerUpdateHeal            = 15
	CommandPlayerUpdateKilled          = 16 // server only
	CommandPlayerRespawned             = 17
	CommandGameUpdatePause             = 18 // server only
	CommandGameUpdateUnpause           = 19 // server only
	CommandGameUpdateEnd               = 20 // server only
	CommandGameServerMessage           = 21 // server only
)

// PayloadPlayerUpdateMove represents new user location
type PayloadPlayerUpdateMove struct {
	Location int
}
