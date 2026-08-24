package service

import (
	"context"
	"time"

	"github.com/jb843051627/firn-signal/internal/model"
	"github.com/jb843051627/firn-signal/internal/policy"
)

func (l *Lab) DailyReport(ctx context.Context, boreholeID string, day time.Time) (model.BoreholeReport, error) {
	scans, err := l.ListScans(ctx, boreholeID)
	if err != nil {
		return model.BoreholeReport{}, err
	}
	assessments, err := l.ListAssessments(ctx)
	if err != nil {
		return model.BoreholeReport{}, err
	}
	profiles, err := l.ListProfiles(ctx)
	if err != nil {
		return model.BoreholeReport{}, err
	}
	releases, err := l.ListReleases(ctx)
	if err != nil {
		return model.BoreholeReport{}, err
	}
	report := policy.ReportFor(boreholeID, day, scans, assessments, profiles)
	for _, release := range releases {
		for _, scan := range scans {
			if release.ScanID == scan.ID && release.Published() && policy.SameDay(release.PublishedAt, day) {
				report.PublishedCount++
			}
		}
	}
	return report, nil
}

func (l *Lab) Snapshot(ctx context.Context, boreholeID string) (map[string]any, error) {
	borehole, err := l.Borehole(ctx, boreholeID)
	if err != nil {
		return nil, err
	}
	scans, err := l.ListScans(ctx, boreholeID)
	if err != nil {
		return nil, err
	}
	probes, err := l.ListProbes(ctx, boreholeID)
	if err != nil {
		return nil, err
	}
	return map[string]any{"borehole": borehole, "scans": scans, "probes": probes, "metrics": l.Metrics()}, nil
}
