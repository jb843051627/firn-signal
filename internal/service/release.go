package service

import (
	"context"
	"fmt"

	"github.com/jb843051627/firn-signal/internal/model"
	"github.com/jb843051627/firn-signal/internal/policy"
)

func (l *Lab) PrepareRelease(ctx context.Context, scanID string) (model.ReleaseManifest, error) {
	scan, err := l.Scan(ctx, scanID)
	if err != nil {
		return model.ReleaseManifest{}, err
	}
	borehole, err := l.Borehole(ctx, scan.BoreholeID)
	if err != nil {
		return model.ReleaseManifest{}, err
	}
	profile, err := l.Profile(ctx, scanID)
	if err != nil {
		return model.ReleaseManifest{}, err
	}
	assessment, err := l.Assessment(ctx, scanID)
	if err != nil {
		return model.ReleaseManifest{}, err
	}
	release, err := policy.PrepareRelease(scan, borehole, profile, assessment, l.clock.Now())
	if err != nil {
		return model.ReleaseManifest{}, err
	}
	if err := l.store.Save(ctx, "release", release.ID, release); err != nil {
		return model.ReleaseManifest{}, err
	}
	l.metrics.Add("releases.prepared", 1)
	return release, l.store.Event(ctx, scanID, "release-prepared", release)
}

func (l *Lab) PublishRelease(ctx context.Context, scanID string) (model.ReleaseManifest, error) {
	lock := scanLock(l, scanID)
	lock.Lock()
	defer lock.Unlock()
	scan, err := l.Scan(ctx, scanID)
	if err != nil {
		return model.ReleaseManifest{}, err
	}
	release, err := l.Release(ctx, scanID)
	if err != nil {
		return model.ReleaseManifest{}, err
	}
	if release.State != model.ReleasePrepared || scan.State != model.ScanEvaluated {
		return model.ReleaseManifest{}, fmt.Errorf("release precondition: %w", model.ErrInvalidState)
	}
	release, err = policy.PublishRelease(release, l.clock.Now())
	if err != nil {
		return model.ReleaseManifest{}, err
	}
	if err := l.store.Save(ctx, "release", release.ID, release); err != nil {
		return model.ReleaseManifest{}, err
	}
	scan.State = model.ScanReleased
	if err := l.store.Save(ctx, "scan", scan.ID, scan); err != nil {
		return model.ReleaseManifest{}, err
	}
	l.metrics.Add("releases.published", 1)
	return release, l.store.Event(ctx, scanID, "release-published", release)
}

func (l *Lab) Release(ctx context.Context, scanID string) (model.ReleaseManifest, error) {
	return load[model.ReleaseManifest](ctx, l.store, "release", model.StableID("release", scanID))
}
func (l *Lab) Manifest(ctx context.Context, scanID string) (model.ReleaseManifest, error) {
	return l.Release(ctx, scanID)
}
func (l *Lab) ListReleases(ctx context.Context) ([]model.ReleaseManifest, error) {
	return list[model.ReleaseManifest](ctx, l.store, "release")
}
