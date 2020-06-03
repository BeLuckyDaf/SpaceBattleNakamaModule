package server

import (
	"context"
	"database/sql"
	"spacebattle/core"
	"spacebattle/services"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

/* =============== */
/* Nakama Handlers */
/* =============== */

// MakeMatch handler
func MakeMatch(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, entries []runtime.MatchmakerEntry) (string, error) {
	for _, e := range entries {
		logger.Info("Matched user '%s' named '%s'", e.GetPresence().GetUserId(), e.GetPresence().GetUsername())

		for k, v := range e.GetProperties() {
			logger.Info("Matched on '%s' value '%v'", k, v)
		}
	}

	matchID, err := nk.MatchCreate(ctx, "spacebattle", map[string]interface{}{"invited": entries})
	if err != nil {
		return "", err
	}

	return matchID, nil
}

// AfterAuthenticateEmail handler
func AfterAuthenticateEmail(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.Session, in *api.AuthenticateEmailRequest) error {
	logger.Info("User '%s' has successfully authenticated via Email.", in.Account.GetEmail())
	return nil
}

// MatchCreateSpaceBattle match creator handler
func MatchCreateSpaceBattle(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule) (runtime.Match, error) {
	matchServices := []services.SBServiceInterface{
		&services.SBPaydayService{},
		&services.SBUserMessageHandlerService{},
		&services.SBMatchBackupService{},
	}

	matchConfig := core.SBConfig{
		KMaxPlayers:          128,
		KWorldSize:           100,
		KMinimalDistance:     60.0,
		KEdgeDistance:        140.0,
		KPaytimeInterval:     5,
		KPlanetCost:          2,
		KAsteroidCost:        2,
		KStationCost:         2,
		KPlanetPayout:        1,
		KAsteroidPayout:      2,
		KMovementCost:        1,
		KHealAmount:          1,
		KInitialPlayerPower:  3,
		KInitialPlayerHealth: 3,
		KInitialHealingPrice: 10,
		KHealCostMultiplier:  2,
		KStationDamage:       1,
		KMaxHealth:           3,
	}

	return &Match{services: matchServices, config: matchConfig}, nil
}
