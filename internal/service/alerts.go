package service

import (
	"context"
	"sort"

	"github.com/jb843051627/firn-signal/internal/model"
)

func (l *Lab) Alerts(ctx context.Context, boreholeID string) ([]model.Signal, error) {
	scans, err := l.ListScans(ctx, boreholeID)
	if err != nil {
		return nil, err
	}
	alerts := make([]model.Signal, 0)
	allScans, _ := l.ListScans(ctx, "")
	_ = scans
	for _, scan := range allScans {
		assessment, err := l.Assessment(ctx, scan.ID)
		if err != nil {
			if err == model.ErrNotFound {
				continue
			}
			return nil, err
		}
		alerts = append(alerts, assessment.Signals...)
	}
	sort.SliceStable(alerts, func(i, j int) bool { return alerts[i].Code < alerts[j].Code })
	return alerts, nil
}
