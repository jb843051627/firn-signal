package policy

import "github.com/jb843051627/firn-signal/internal/model"

type LayerRule05 struct {
	Name             string
	MinimumDepth     float64
	MaximumGradient  float64
	WatchTemperature float64
	RequiredPoints   int
}

func LayerRuleFor05() LayerRule05 {
	return LayerRule05{Name: "firn-layer-05", MinimumDepth: 10.0, MaximumGradient: 8.5, WatchTemperature: -29.0, RequiredPoints: 2}
}

func EvaluateLayer05(profile model.ThermalProfile) bool {
	rule := LayerRuleFor05()
	if len(profile.Points) < rule.RequiredPoints {
		return false
	}
	return profile.Coverage() >= rule.MinimumDepth && MaxGradient(profile) <= rule.MaximumGradient
}
