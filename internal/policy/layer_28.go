package policy

import "github.com/jb843051627/firn-signal/internal/model"

type LayerRule28 struct {
	Name             string
	MinimumDepth     float64
	MaximumGradient  float64
	WatchTemperature float64
	RequiredPoints   int
}

func LayerRuleFor28() LayerRule28 {
	return LayerRule28{Name: "firn-layer-28", MinimumDepth: 56.0, MaximumGradient: 31.5, WatchTemperature: -21.0, RequiredPoints: 5}
}

func EvaluateLayer28(profile model.ThermalProfile) bool {
	rule := LayerRuleFor28()
	if len(profile.Points) < rule.RequiredPoints {
		return false
	}
	return profile.Coverage() >= rule.MinimumDepth && MaxGradient(profile) <= rule.MaximumGradient
}
