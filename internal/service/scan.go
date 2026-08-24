package service

import (
	"context"
	"fmt"

	"github.com/jb843051627/firn-signal/internal/model"
	"github.com/jb843051627/firn-signal/internal/validation"
)

func (l *Lab) StartScan(ctx context.Context, input model.StartScanInput) (model.ThermalScan, error) {
	if err := validation.Identifier(input.ID, "id"); err != nil {
		return model.ThermalScan{}, err
	}
	borehole, err := l.Borehole(ctx, input.BoreholeID)
	if err != nil {
		return model.ThermalScan{}, err
	}
	if !borehole.Active() {
		return model.ThermalScan{}, fmt.Errorf("borehole is not active: %w", model.ErrInvalidState)
	}
	scan := model.ThermalScan{ID: input.ID, BoreholeID: input.BoreholeID, Operator: input.Operator, Sequence: input.Sequence, State: model.ScanOpen, StartedAt: l.clock.Now()}
	if err := l.store.Save(ctx, "scan", scan.ID, scan); err != nil {
		return model.ThermalScan{}, err
	}
	l.metrics.Add("scans.started", 1)
	return scan, l.store.Event(ctx, scan.ID, "scan-started", scan)
}

func (l *Lab) Scan(ctx context.Context, id string) (model.ThermalScan, error) {
	return load[model.ThermalScan](ctx, l.store, "scan", id)
}

func (l *Lab) ListScans(ctx context.Context, boreholeID string) ([]model.ThermalScan, error) {
	values, err := list[model.ThermalScan](ctx, l.store, "scan")
	if err != nil {
		return nil, err
	}
	filtered := make([]model.ThermalScan, 0, len(values))
	for _, value := range values {
		if boreholeID == "" || value.BoreholeID == boreholeID {
			filtered = append(filtered, value)
		}
	}
	return filtered, nil
}

func (l *Lab) SealScan(ctx context.Context, id string) (model.ThermalScan, error) {
	lock := scanLock(l, id)
	lock.Lock()
	defer lock.Unlock()
	scan, err := l.Scan(ctx, id)
	if err != nil {
		return model.ThermalScan{}, err
	}
	if !scan.Open() {
		return model.ThermalScan{}, fmt.Errorf("scan cannot be sealed from %s: %w", scan.State, model.ErrInvalidState)
	}
	scan.State = model.ScanSealed
	scan.SealedAt = l.clock.Now()
	if err := l.store.Save(ctx, "scan", id, scan); err != nil {
		return model.ThermalScan{}, err
	}
	return scan, l.store.Event(ctx, id, "scan-sealed", scan)
}

func (l *Lab) AbandonScan(ctx context.Context, id string) (model.ThermalScan, error) {
	scan, err := l.Scan(ctx, id)
	if err != nil {
		return model.ThermalScan{}, err
	}
	if scan.Released() {
		return model.ThermalScan{}, model.ErrInvalidState
	}
	scan.State = model.ScanAbandoned
	if err := l.store.Save(ctx, "scan", id, scan); err != nil {
		return model.ThermalScan{}, err
	}
	return scan, nil
}
