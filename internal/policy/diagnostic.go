package policy

import (
	"fmt"

	"github.com/jb843051627/firn-signal/internal/model"
)

func Diagnostic(scan model.ThermalScan, profile model.ThermalProfile, assessment model.QualityAssessment) model.DiagnosticReport {
	checks := []string{fmt.Sprintf("scan-state=%s", scan.State), fmt.Sprintf("coverage=%.2f", profile.Coverage()), fmt.Sprintf("score=%.2f", assessment.Score)}
	if assessment.HasBlocker() {
		checks = append(checks, "blocking-signal-present")
	}
	if profile.CalibrationID == "" {
		checks = append(checks, "calibration-lineage-missing")
	}
	return model.DiagnosticReport{ScanID: scan.ID, CalibrationID: profile.CalibrationID, Checks: checks}
}
