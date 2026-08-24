package service

import (
	"context"
	"fmt"

	"github.com/jb843051627/firn-signal/internal/model"
	"github.com/jb843051627/firn-signal/internal/validation"
)

func (l *Lab) RegisterProbe(ctx context.Context, input model.RegisterProbeInput) (model.Probe, error) {
	if err := validation.Identifier(input.ID, "id"); err != nil {
		return model.Probe{}, err
	}
	if err := validation.Identifier(input.BoreholeID, "borehole_id"); err != nil {
		return model.Probe{}, err
	}
	if err := validation.Required(input.Serial, "serial"); err != nil {
		return model.Probe{}, err
	}
	if err := validation.Span(input.DepthStartM, input.DepthEndM); err != nil {
		return model.Probe{}, err
	}
	borehole, err := l.Borehole(ctx, input.BoreholeID)
	if err != nil {
		return model.Probe{}, err
	}
	if !borehole.Active() {
		return model.Probe{}, fmt.Errorf("borehole is not active: %w", model.ErrInvalidState)
	}
	probe := model.Probe{ID: input.ID, BoreholeID: input.BoreholeID, Serial: input.Serial, DepthStartM: input.DepthStartM, DepthEndM: input.DepthEndM, State: model.ProbeInstalled, InstalledAt: l.clock.Now(), Notes: input.Notes}
	if err := l.store.Save(ctx, "probe", probe.ID, probe); err != nil {
		return model.Probe{}, err
	}
	l.metrics.Add("probes.registered", 1)
	return probe, l.store.Event(ctx, probe.ID, "probe-registered", probe)
}

func (l *Lab) Probe(ctx context.Context, id string) (model.Probe, error) {
	return load[model.Probe](ctx, l.store, "probe", id)
}
func (l *Lab) ListProbes(ctx context.Context, boreholeID string) ([]model.Probe, error) {
	values, err := list[model.Probe](ctx, l.store, "probe")
	if err != nil {
		return nil, err
	}
	filtered := make([]model.Probe, 0, len(values))
	for _, value := range values {
		if boreholeID == "" || value.BoreholeID == boreholeID {
			filtered = append(filtered, value)
		}
	}
	return filtered, nil
}

func (l *Lab) RemoveProbe(ctx context.Context, id string) (model.Probe, error) {
	probe, err := l.Probe(ctx, id)
	if err != nil {
		return model.Probe{}, err
	}
	if probe.State == model.ProbeRemoved {
		return probe, nil
	}
	probe.State = model.ProbeRemoved
	if err := l.store.Save(ctx, "probe", id, probe); err != nil {
		return model.Probe{}, err
	}
	return probe, nil
}
