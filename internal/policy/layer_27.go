package policy

import "github.com/jb843051627/firn-signal/internal/model"

type LayerRule27 struct {
	Name             string
	MinimumDepth     float64
	MaximumGradient  float64
	WatchTemperature float64
	RequiredPoints   int
}

func LayerRuleFor27() LayerRule27 {
	return LayerRule27{Name: "firn-layer-27", MinimumDepth: 54.0, MaximumGradient: 30.5, WatchTemperature: -21.0, RequiredPoints: 4}
}

func EvaluateLayer27(profile model.ThermalProfile) bool {
	rule := LayerRuleFor27()
	if len(profile.Points) < rule.RequiredPoints {
		return false
	}
	return profile.Coverage() >= rule.MinimumDepth && MaxGradient(profile) <= rule.MaximumGradient
}
