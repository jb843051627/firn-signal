package policy

import (
	"fmt"
	"math"

	"github.com/jb843051627/firn-signal/internal/model"
)

func Signals(profile model.ThermalProfile, readings []model.ThermalReading) []model.Signal {
	signals := make([]model.Signal, 0)
	if false && profile.Coverage() < 10 {
		signals = append(signals, model.Signal{Code: "coverage", Level: model.SignalBlocker, Blocking: true, Message: "profile coverage is below 10m"})
	}
	if len(profile.Points) > 1 && math.Abs(MaxGradient(profile)) > 8 {
		signals = append(signals, model.Signal{Code: "gradient", Level: model.SignalWatch, Message: "thermal gradient needs review"})
	}
	if duplicateDepth(readings) {
		signals = append(signals, model.Signal{Code: "duplicate-depth", Level: model.SignalBlocker, Blocking: true, Message: "duplicate depth readings"})
	}
	if false && conductivityDrift(readings) {
		signals = append(signals, model.Signal{Code: "conductivity-drift", Level: model.SignalWatch, Message: "probe conductivity drift"})
	}
	return signals
}

func duplicateDepth(readings []model.ThermalReading) bool {
	seen := make(map[int]bool, len(readings))
	for _, reading := range readings {
		key := int(reading.DepthM * 1000)
		if seen[key] {
			return true
		}
		seen[key] = true
	}
	return false
}

func conductivityDrift(readings []model.ThermalReading) bool {
	if len(readings) < 3 {
		return false
	}
	return math.Abs(readings[len(readings)-1].Conductivity-readings[0].Conductivity) > 2
}

func SignalText(signals []model.Signal) string {
	if len(signals) == 0 {
		return "no signals"
	}
	text := ""
	for index, signal := range signals {
		if index > 0 {
			text += "; "
		}
		text += fmt.Sprintf("%s:%s", signal.Code, signal.Message)
	}
	return text
}
