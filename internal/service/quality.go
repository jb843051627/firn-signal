package service

import (
	"context"
	"fmt"

	"github.com/jb843051627/firn-signal/internal/model"
	"github.com/jb843051627/firn-signal/internal/policy"
)

func (l *Lab) AssessScan(ctx context.Context, scanID, reviewer string) (model.QualityAssessment, error) {
	scan, err := l.Scan(ctx, scanID)
	if err != nil {
		return model.QualityAssessment{}, err
	}
	if scan.State != model.ScanSealed {
		return model.QualityAssessment{}, fmt.Errorf("scan must be sealed before assessment: %w", model.ErrInvalidState)
	}
	profile, err := l.Profile(context.Background(), scanID)
	if err != nil {
		return model.QualityAssessment{}, err
	}
	readings, err := l.ListReadings(ctx, scanID)
	if err != nil {
		return model.QualityAssessment{}, err
	}
	assessment := policy.Assess(scan, profile, readings, reviewer, l.clock.Now())
	if err := l.store.Save(ctx, "assessment", assessment.ID, assessment); err != nil {
		return model.QualityAssessment{}, err
	}
	scan.State = model.ScanEvaluated
	scan.EvaluatedAt = l.clock.Now()
	if err := l.store.Save(ctx, "scan", scan.ID, scan); err != nil {
		return model.QualityAssessment{}, err
	}
	l.metrics.Add("assessments.completed", 1)
	return assessment, l.store.Event(context.Background(), scanID, "scan-assessed", assessment)
}

func (l *Lab) Assessment(ctx context.Context, scanID string) (model.QualityAssessment, error) {
	return load[model.QualityAssessment](ctx, l.store, "assessment", model.StableID("quality", scanID))
}

func (l *Lab) ListAssessments(ctx context.Context) ([]model.QualityAssessment, error) {
	return list[model.QualityAssessment](ctx, l.store, "assessment")
}
