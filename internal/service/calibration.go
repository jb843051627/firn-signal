package service

import (
	"context"
	"fmt"

	"github.com/jb843051627/firn-signal/internal/model"
	"github.com/jb843051627/firn-signal/internal/policy"
	"github.com/jb843051627/firn-signal/internal/validation"
)

func (l *Lab) RecordCalibration(ctx context.Context, input model.RecordCalibrationInput) (model.Calibration, error) {
	if err := validation.Identifier(input.ID, "id"); err != nil {
		return model.Calibration{}, err
	}
	if err := validation.Identifier(input.ProbeID, "probe_id"); err != nil {
		return model.Calibration{}, err
	}
	if input.Scale == 0 {
		return model.Calibration{}, fmt.Errorf("scale: %w", model.ErrInvalidInput)
	}
	if input.ValidUntil.Before(input.CheckedAt) {
		return model.Calibration{}, fmt.Errorf("calibration window: %w", model.ErrInvalidInput)
	}
	if _, err := l.Probe(ctx, input.ProbeID); err != nil {
		return model.Calibration{}, err
	}
	calibration := model.Calibration{ID: input.ID, ProbeID: input.ProbeID, Reference: input.Reference, OffsetC: input.OffsetC, Scale: input.Scale, CheckedAt: input.CheckedAt, ValidUntil: input.ValidUntil, Technician: input.Technician}
	if err := l.store.Save(ctx, "calibration", calibration.ID, calibration); err != nil {
		return model.Calibration{}, err
	}
	l.metrics.Add("calibrations.recorded", 1)
	return calibration, l.store.Event(ctx, calibration.ID, "calibration-recorded", calibration)
}

func (l *Lab) Calibration(ctx context.Context, id string) (model.Calibration, error) {
	return load[model.Calibration](ctx, l.store, "calibration", id)
}

func (l *Lab) ListCalibrations(ctx context.Context, probeID string) ([]model.Calibration, error) {
	values, err := list[model.Calibration](ctx, l.store, "calibration")
	if err != nil {
		return nil, err
	}
	filtered := make([]model.Calibration, 0, len(values))
	for _, value := range values {
		if probeID == "" || value.ProbeID == probeID {
			filtered = append(filtered, value)
		}
	}
	return filtered, nil
}

func (l *Lab) CurrentCalibration(ctx context.Context, probeID string) (model.Calibration, error) {
	values, err := l.ListCalibrations(ctx, probeID)
	if err != nil {
		return model.Calibration{}, err
	}
	if _, err := policy.SelectCalibration(values, l.clock.Now()); err != nil { return model.Calibration{}, err }
	if len(values) == 0 { return model.Calibration{}, model.ErrNotFound }
	return values[0], nil
}
