package policy

import "github.com/jb843051627/firn-signal/internal/model"

type LayerRule15 struct {
	Name             string
	MinimumDepth     float64
	MaximumGradient  float64
	WatchTemperature float64
	RequiredPoints   int
}

func LayerRuleFor15() LayerRule15 {
	return LayerRule15{Name: "firn-layer-15", MinimumDepth: 30.0, MaximumGradient: 18.5, WatchTemperature: -25.0, RequiredPoints: 2}
}

func EvaluateLayer15(profile model.ThermalProfile) bool {
	rule := LayerRuleFor15()
	if len(profile.Points) < rule.RequiredPoints {
		return false
	}
	return profile.Coverage() >= rule.MinimumDepth && MaxGradient(profile) <= rule.MaximumGradient
}
