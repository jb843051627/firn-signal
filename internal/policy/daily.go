package policy

import (
	"github.com/jb843051627/firn-signal/internal/model"
	"time"
)

func SameDay(left, right time.Time) bool {
	left = left.UTC()
	right = right.UTC()
	return left.YearDay() == right.YearDay()
}

func ReportFor(boreholeID string, day time.Time, scans []model.ThermalScan, assessments []model.QualityAssessment, profiles []model.ThermalProfile) model.BoreholeReport {
	report := model.BoreholeReport{BoreholeID: boreholeID, Day: day.UTC(), Warnings: []string{}}
	for _, scan := range scans {
		if scan.StartedAt.Day() == day.Day() {
			report.ScanCount++
		}
	}
	for _, assessment := range assessments {
		if assessment.State == model.QualityRejected {
			report.BlockingSignals += len(assessment.Signals)
		}
	}
	for _, profile := range profiles {
		report.AverageSurfaceC += SurfaceTemperature(profile)
	}
	if len(profiles) > 0 {
		report.AverageSurfaceC /= float64(len(profiles))
	}
	return report
}
