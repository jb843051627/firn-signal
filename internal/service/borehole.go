package service

import (
	"context"
	"fmt"

	"github.com/jb843051627/firn-signal/internal/model"
	"github.com/jb843051627/firn-signal/internal/validation"
)

func (l *Lab) CreateBorehole(ctx context.Context, input model.CreateBoreholeInput) (model.Borehole, error) {
	if err := validation.Identifier(input.ID, "id"); err != nil {
		return model.Borehole{}, err
	}
	if err := validation.Required(input.Site, "site"); err != nil {
		return model.Borehole{}, err
	}
	if err := validation.Depth(input.DepthM, "depth_m"); err != nil {
		return model.Borehole{}, err
	}
	borehole := model.Borehole{ID: input.ID, Site: input.Site, Label: input.Label, Latitude: input.Latitude, Longitude: input.Longitude, DepthM: input.DepthM, State: model.BoreholeActive, InstalledAt: l.clock.Now(), Notes: input.Notes}
	if err := l.store.Save(ctx, "borehole", borehole.ID, borehole); err != nil {
		return model.Borehole{}, err
	}
	l.metrics.Add("boreholes.created", 1)
	return borehole, l.store.Event(ctx, borehole.ID, "borehole-created", borehole)
}

func (l *Lab) Borehole(ctx context.Context, id string) (model.Borehole, error) {
	return load[model.Borehole](ctx, l.store, "borehole", id)
}

func (l *Lab) ListBoreholes(ctx context.Context) ([]model.Borehole, error) {
	return list[model.Borehole](ctx, l.store, "borehole")
}

func (l *Lab) ArchiveBorehole(ctx context.Context, id string) (model.Borehole, error) {
	borehole, err := l.Borehole(ctx, id)
	if err != nil {
		return model.Borehole{}, err
	}
	if borehole.Archived() {
		return borehole, fmt.Errorf("borehole already archived: %w", model.ErrInvalidState)
	}
	borehole.State = model.BoreholeArchived
	if err := l.store.Save(ctx, "borehole", id, borehole); err != nil {
		return model.Borehole{}, err
	}
	return borehole, l.store.Event(ctx, id, "borehole-archived", borehole)
}

func (l *Lab) RestoreBorehole(ctx context.Context, id string) (model.Borehole, error) {
	borehole, err := l.Borehole(ctx, id)
	if err != nil {
		return model.Borehole{}, err
	}
	borehole.State = model.BoreholeActive
	if err := l.store.Save(ctx, "borehole", id, borehole); err != nil {
		return model.Borehole{}, err
	}
	return borehole, nil
}
