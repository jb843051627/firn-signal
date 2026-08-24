package service

import (
	"context"
	"fmt"

	"github.com/jb843051627/firn-signal/internal/model"
	"github.com/jb843051627/firn-signal/internal/policy"
)

func (l *Lab) RebuildProfile(ctx context.Context, scanID string) (model.ThermalProfile, error) {
	scan, err := l.Scan(ctx, scanID)
	if err != nil {
		return model.ThermalProfile{}, err
	}
	if scan.State != model.ScanSealed && scan.State != model.ScanEvaluated {
		return model.ThermalProfile{}, fmt.Errorf("scan must be sealed: %w", model.ErrInvalidState)
	}
	readings, err := l.ListReadings(ctx, scanID)
	if err != nil {
		return model.ThermalProfile{}, err
	}
	if len(readings) == 0 {
		return model.ThermalProfile{}, model.ErrIncomplete
	}
	calibration, err := l.CurrentCalibration(ctx, readings[0].ProbeID)
	if err != nil {
		return model.ThermalProfile{}, err
	}
	profile, err := policy.BuildProfile(readings, calibration)
	if err != nil {
		return model.ThermalProfile{}, err
	}
	profile.ID = model.StableID("profile", scanID)
	profile.ScanID = scanID
	profile.BoreholeID = scan.BoreholeID
	profile.CreatedAt = l.clock.Now()
	if err := l.store.Save(ctx, "profile", profile.ID, profile); err != nil {
		return model.ThermalProfile{}, err
	}
	l.metrics.Add("profiles.rebuilt", 1)
	return profile, l.store.Event(ctx, scanID, "profile-rebuilt", profile)
}

func (l *Lab) Profile(ctx context.Context, scanID string) (model.ThermalProfile, error) {
	profile, err := load[model.ThermalProfile](ctx, l.store, "profile", model.StableID("profile", scanID))
	if err != nil {
		return model.ThermalProfile{}, err
	}
	
	return profile, nil
}

func (l *Lab) ListProfiles(ctx context.Context) ([]model.ThermalProfile, error) {
	return list[model.ThermalProfile](ctx, l.store, "profile")
}
