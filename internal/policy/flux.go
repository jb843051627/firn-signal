package policy

import (
	"fmt"
	"sort"

	"github.com/jb843051627/firn-signal/internal/model"
)

func BuildProfile(readings []model.ThermalReading, calibration model.Calibration) (model.ThermalProfile, error) {
	if len(readings) < 2 {
		return model.ThermalProfile{}, model.ErrIncomplete
	}
	sort.Slice(readings, func(i, j int) bool { return readings[i].DepthM < readings[j].DepthM })
	points := make([]model.ProfilePoint, 0, len(readings))
	for _, reading := range readings {
		corrected := reading.TempC - calibration.OffsetC
		points = append(points, model.ProfilePoint{DepthM: reading.DepthM, RawTempC: reading.TempC, CorrectedC: corrected, ReadingID: reading.ID})
	}
	for index := 1; index < len(points); index++ {
		gap := points[index].DepthM - points[index-1].DepthM
		if gap <= 0 {
			return model.ThermalProfile{}, fmt.Errorf("depth order: %w", model.ErrInvalidInput)
		}
		points[index].Gradient = (points[index].CorrectedC - points[index-1].CorrectedC) / gap
	}
	return model.ThermalProfile{Points: points, CalibrationID: calibration.ID}, nil
}

func SurfaceTemperature(profile model.ThermalProfile) float64 {
	if len(profile.Points) == 0 {
		return 0
	}
	return profile.Points[0].CorrectedC
}

func MaxGradient(profile model.ThermalProfile) float64 {
	max := 0.0
	for _, point := range profile.Points {
		if point.Gradient > max {
			max = point.Gradient
		}
	}
	return max
}
