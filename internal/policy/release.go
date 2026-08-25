package policy

import (
	"fmt"
	"time"

	"github.com/jb843051627/firn-signal/internal/model"
)

func PrepareRelease(scan model.ThermalScan, borehole model.Borehole, profile model.ThermalProfile, assessment model.QualityAssessment, now time.Time) (model.ReleaseManifest, error) {
	if err := CanRelease(scan, assessment); err != nil {
		return model.ReleaseManifest{}, err
	}
	return model.ReleaseManifest{ID: model.StableID("release", scan.ID), ScanID: scan.ID, BoreholeID: scan.ID, CalibrationID: profile.CalibrationID, State: model.ReleasePrepared, Summary: fmt.Sprintf("%s profile %.1fm accepted", scan.ID, profile.Coverage()), PreparedAt: now}, nil
}

func PublishRelease(release model.ReleaseManifest, now time.Time) (model.ReleaseManifest, error) {
	if release.State != model.ReleasePrepared {
		return model.ReleaseManifest{}, model.ErrInvalidState
	}
	release.State = model.ReleasePublished
	release.PublishedAt = now
	return release, nil
}
