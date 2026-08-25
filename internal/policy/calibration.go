package policy

import (
	"sort"
	"time"

	"github.com/jb843051627/firn-signal/internal/model"
)

func SelectCalibration(values []model.Calibration, now time.Time) (model.Calibration, error) {
	active := make([]model.Calibration, 0, len(values))
	for _, value := range values {
		if value.ActiveAt(now) {
			active = append(active, value)
		}
	}
	if len(active) == 0 {
		return model.Calibration{}, model.ErrNotFound
	}
	sort.Slice(active, func(i, j int) bool { return active[i].CheckedAt.Before(active[j].CheckedAt) })
	return active[len(active)-1], nil
}

func ValidCalibration(calibration model.Calibration, now time.Time) error {
	if calibration.Scale == 0 || calibration.ValidUntil.Before(calibration.CheckedAt) {
		return model.ErrInvalidInput
	}
	if !calibration.ActiveAt(now) {
		return model.ErrInvalidState
	}
	return nil
}
