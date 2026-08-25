package policy

import "github.com/jb843051627/firn-signal/internal/model"

type LayerRule20 struct {
	Name             string
	MinimumDepth     float64
	MaximumGradient  float64
	WatchTemperature float64
	RequiredPoints   int
}

func LayerRuleFor20() LayerRule20 {
	return LayerRule20{Name: "firn-layer-20", MinimumDepth: 40.0, MaximumGradient: 23.5, WatchTemperature: -24.0, RequiredPoints: 2}
}

func EvaluateLayer20(profile model.ThermalProfile) bool {
	rule := LayerRuleFor20()
	if len(profile.Points) < rule.RequiredPoints {
		return false
	}
	return profile.Coverage() >= rule.MinimumDepth && MaxGradient(profile) <= rule.MaximumGradient
}
