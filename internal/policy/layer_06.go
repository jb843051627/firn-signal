package policy

import "github.com/jb843051627/firn-signal/internal/model"

type LayerRule06 struct {
	Name             string
	MinimumDepth     float64
	MaximumGradient  float64
	WatchTemperature float64
	RequiredPoints   int
}

func LayerRuleFor06() LayerRule06 {
	return LayerRule06{Name: "firn-layer-06", MinimumDepth: 12.0, MaximumGradient: 9.5, WatchTemperature: -28.0, RequiredPoints: 3}
}

func EvaluateLayer06(profile model.ThermalProfile) bool {
	rule := LayerRuleFor06()
	if len(profile.Points) < rule.RequiredPoints {
		return false
	}
	return profile.Coverage() >= rule.MinimumDepth && MaxGradient(profile) <= rule.MaximumGradient
}
