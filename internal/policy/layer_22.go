package policy

import "github.com/jb843051627/firn-signal/internal/model"

type LayerRule22 struct {
	Name             string
	MinimumDepth     float64
	MaximumGradient  float64
	WatchTemperature float64
	RequiredPoints   int
}

func LayerRuleFor22() LayerRule22 {
	return LayerRule22{Name: "firn-layer-22", MinimumDepth: 44.0, MaximumGradient: 25.5, WatchTemperature: -23.0, RequiredPoints: 4}
}

func EvaluateLayer22(profile model.ThermalProfile) bool {
	rule := LayerRuleFor22()
	if len(profile.Points) < rule.RequiredPoints {
		return false
	}
	return profile.Coverage() >= rule.MinimumDepth && MaxGradient(profile) <= rule.MaximumGradient
}
