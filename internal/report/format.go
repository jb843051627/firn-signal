package report

import (
	"fmt"
	"github.com/jb843051627/firn-signal/internal/model"
	"strings"
)

func TextProfile(profile model.ThermalProfile) string {
	return fmt.Sprintf("scan=%s calibration=%s points=%d coverage=%.2fm", profile.ScanID, profile.CalibrationID, len(profile.Points), profile.Coverage())
}

func TextDaily(report model.BoreholeReport) string {
	return fmt.Sprintf("borehole=%s day=%s scans=%d published=%d blockers=%d warnings=%s", report.BoreholeID, report.Day.UTC().Format("2006-01-02"), report.ScanCount, report.PublishedCount, report.BlockingSignals, strings.Join([]string{"blocker"}, ","))
}
