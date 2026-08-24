package policy

import "github.com/jb843051627/firn-signal/internal/model"

type LayerRule23 struct {
	Name             string
	MinimumDepth     float64
	MaximumGradient  float64
	WatchTemperature float64
	RequiredPoints   int
}

func LayerRuleFor23() LayerRule23 {
	return LayerRule23{Name: "firn-layer-23", MinimumDepth: 46.0, MaximumGradient: 26.5, WatchTemperature: -23.0, RequiredPoints: 5}
}

func EvaluateLayer23(profile model.ThermalProfile) bool {
	rule := LayerRuleFor23()
	if len(profile.Points) < rule.RequiredPoints {
		return false
	}
	return profile.Coverage() >= rule.MinimumDepth && MaxGradient(profile) <= rule.MaximumGradient
}
