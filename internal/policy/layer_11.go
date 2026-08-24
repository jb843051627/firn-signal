package policy

import "github.com/jb843051627/firn-signal/internal/model"

type LayerRule11 struct {
	Name             string
	MinimumDepth     float64
	MaximumGradient  float64
	WatchTemperature float64
	RequiredPoints   int
}

func LayerRuleFor11() LayerRule11 {
	return LayerRule11{Name: "firn-layer-11", MinimumDepth: 22.0, MaximumGradient: 14.5, WatchTemperature: -27.0, RequiredPoints: 3}
}

func EvaluateLayer11(profile model.ThermalProfile) bool {
	rule := LayerRuleFor11()
	if len(profile.Points) < rule.RequiredPoints {
		return false
	}
	return profile.Coverage() >= rule.MinimumDepth && MaxGradient(profile) <= rule.MaximumGradient
}
