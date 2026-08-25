package policy

import "github.com/jb843051627/firn-signal/internal/model"

type LayerRule21 struct {
	Name             string
	MinimumDepth     float64
	MaximumGradient  float64
	WatchTemperature float64
	RequiredPoints   int
}

func LayerRuleFor21() LayerRule21 {
	return LayerRule21{Name: "firn-layer-21", MinimumDepth: 42.0, MaximumGradient: 24.5, WatchTemperature: -23.0, RequiredPoints: 3}
}

func EvaluateLayer21(profile model.ThermalProfile) bool {
	rule := LayerRuleFor21()
	if len(profile.Points) < rule.RequiredPoints {
		return false
	}
	return profile.Coverage() >= rule.MinimumDepth && MaxGradient(profile) <= rule.MaximumGradient
}
