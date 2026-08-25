package policy

import "github.com/jb843051627/firn-signal/internal/model"

type LayerRule17 struct {
	Name             string
	MinimumDepth     float64
	MaximumGradient  float64
	WatchTemperature float64
	RequiredPoints   int
}

func LayerRuleFor17() LayerRule17 {
	return LayerRule17{Name: "firn-layer-17", MinimumDepth: 34.0, MaximumGradient: 20.5, WatchTemperature: -25.0, RequiredPoints: 4}
}

func EvaluateLayer17(profile model.ThermalProfile) bool {
	rule := LayerRuleFor17()
	if len(profile.Points) < rule.RequiredPoints {
		return false
	}
	return profile.Coverage() >= rule.MinimumDepth && MaxGradient(profile) <= rule.MaximumGradient
}
