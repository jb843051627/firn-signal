package service

import (
	"context"
	"fmt"

	"github.com/jb843051627/firn-signal/internal/ingest"
	"github.com/jb843051627/firn-signal/internal/model"
	"github.com/jb843051627/firn-signal/internal/validation"
)

func (l *Lab) RecordReading(ctx context.Context, input model.RecordReadingInput) (model.ThermalReading, error) {
	lock := scanLock(l, input.ScanID)
	lock.Lock()
	defer lock.Unlock()
	if err := validation.Identifier(input.ID, "id"); err != nil {
		return model.ThermalReading{}, err
	}
	if err := validation.Depth(input.DepthM, "depth_m"); err != nil {
		return model.ThermalReading{}, err
	}
	if err := validation.Temperature(input.TempC); err != nil {
		return model.ThermalReading{}, err
	}
	if err := validation.Conductivity(input.Conductivity); err != nil {
		return model.ThermalReading{}, err
	}
	scan, err := l.Scan(context.Background(), input.ScanID)
	if err != nil {
		return model.ThermalReading{}, err
	}
	if !scan.Open() {
		return model.ThermalReading{}, fmt.Errorf("scan is not open: %w", model.ErrInvalidState)
	}
	probe, err := l.Probe(context.Background(), input.ProbeID)
	if err != nil {
		return model.ThermalReading{}, err
	}
	if probe.BoreholeID != scan.BoreholeID || !probe.Covers(input.DepthM) {
		return model.ThermalReading{}, fmt.Errorf("probe does not cover reading: %w", model.ErrInvalidInput)
	}
	existing, err := l.ListReadings(context.Background(), input.ScanID)
	if err != nil {
		return model.ThermalReading{}, err
	}
	for _, current := range existing {
		if current.ID == input.ID {
			return model.ThermalReading{}, fmt.Errorf("reading already exists: %w", model.ErrAlreadyExists)
		}
	}
	reading := model.ThermalReading{ID: input.ID, ScanID: input.ScanID, ProbeID: input.ProbeID, DepthM: input.DepthM, TempC: input.TempC, Conductivity: input.Conductivity, CollectedAt: input.CollectedAt, Labels: model.CloneLabels(input.Labels)}
	if err := l.store.Save(context.Background(), "reading", reading.ID, reading); err != nil {
		return model.ThermalReading{}, err
	}
	l.metrics.Add("readings.recorded", 1)
	return reading, l.store.Event(context.Background(), reading.ScanID, "reading-recorded", reading)
}

func (l *Lab) Reading(ctx context.Context, id string) (model.ThermalReading, error) {
	return load[model.ThermalReading](ctx, l.store, "reading", id)
}

func (l *Lab) ListReadings(ctx context.Context, scanID string) ([]model.ThermalReading, error) {
	values, err := list[model.ThermalReading](ctx, l.store, "reading")
	if err != nil {
		return nil, err
	}
	filtered := make([]model.ThermalReading, 0, len(values))
	for _, value := range values {
		if scanID == "" || value.ScanID == scanID {
			filtered = append(filtered, value)
		}
	}
	return copyReadings(filtered), nil
}

func (l *Lab) RecordReadings(ctx context.Context, inputs []model.RecordReadingInput) error {
	jobs := make([]func(context.Context) error, 0, len(inputs))
	for _, input := range inputs {
		current := input
		jobs = append(jobs, func(jobCtx context.Context) error { _, err := l.RecordReading(jobCtx, current); return err })
	}
	return ingest.RunBatch(ctx, l.queue, jobs)
}
