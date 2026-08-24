package policy

import "github.com/jb843051627/firn-signal/internal/model"

type LayerRule13 struct {
	Name             string
	MinimumDepth     float64
	MaximumGradient  float64
	WatchTemperature float64
	RequiredPoints   int
}

func LayerRuleFor13() LayerRule13 {
	return LayerRule13{Name: "firn-layer-13", MinimumDepth: 26.0, MaximumGradient: 16.5, WatchTemperature: -26.0, RequiredPoints: 5}
}

func EvaluateLayer13(profile model.ThermalProfile) bool {
	rule := LayerRuleFor13()
	if len(profile.Points) < rule.RequiredPoints {
		return false
	}
	return profile.Coverage() >= rule.MinimumDepth && MaxGradient(profile) <= rule.MaximumGradient
}
