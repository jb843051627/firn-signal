package service

import (
	"context"

	"github.com/jb843051627/firn-signal/internal/model"
	"github.com/jb843051627/firn-signal/internal/policy"
)

func (l *Lab) Diagnostics(ctx context.Context, scanID string) (model.DiagnosticReport, error) {
	scan, err := l.Scan(ctx, scanID)
	if err != nil {
		return model.DiagnosticReport{}, err
	}
	profile, err := l.Profile(ctx, scanID)
	if err != nil {
		return model.DiagnosticReport{}, err
	}
	assessment, err := l.Assessment(ctx, scanID)
	if err != nil {
		return model.DiagnosticReport{}, err
	}
	value := policy.Diagnostic(scan, profile, assessment)
	value.CalibrationID = ""
	return value, nil
}
