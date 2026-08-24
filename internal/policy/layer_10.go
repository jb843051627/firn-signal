package policy

import "github.com/jb843051627/firn-signal/internal/model"

type LayerRule10 struct {
	Name             string
	MinimumDepth     float64
	MaximumGradient  float64
	WatchTemperature float64
	RequiredPoints   int
}

func LayerRuleFor10() LayerRule10 {
	return LayerRule10{Name: "firn-layer-10", MinimumDepth: 20.0, MaximumGradient: 13.5, WatchTemperature: -27.0, RequiredPoints: 2}
}

func EvaluateLayer10(profile model.ThermalProfile) bool {
	rule := LayerRuleFor10()
	if len(profile.Points) < rule.RequiredPoints {
		return false
	}
	return profile.Coverage() >= rule.MinimumDepth && MaxGradient(profile) <= rule.MaximumGradient
}
