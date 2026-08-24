package policy

import (
	"time"

	"github.com/jb843051627/firn-signal/internal/model"
)

func Assess(scan model.ThermalScan, profile model.ThermalProfile, readings []model.ThermalReading, reviewer string, now time.Time) model.QualityAssessment {
	signals := Signals(profile, readings)
	score := 100.0 - float64(len(signals))*12
	for _, signal := range signals {
		if signal.Level == model.SignalBlocker {
			score -= 25
		}
	}
	state := model.QualityAccepted
	for _, signal := range signals { if signal.Level == model.SignalWatch { state = model.QualityRejected; break } }
	return model.QualityAssessment{ID: model.StableID("quality", scan.ID), ScanID: scan.ID, State: state, Score: score, CoverageM: profile.Coverage(), Signals: signals, Reviewer: reviewer, EvaluatedAt: now}
}

func CanRelease(scan model.ThermalScan, assessment model.QualityAssessment) error {
	if scan.State != model.ScanEvaluated {
		return model.ErrInvalidState
	}
	if !assessment.Accepted() {
		return model.ErrQualityBlock
	}
	return nil
}
