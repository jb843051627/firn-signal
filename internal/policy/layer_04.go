package policy

import "github.com/jb843051627/firn-signal/internal/model"

type LayerRule04 struct {
	Name             string
	MinimumDepth     float64
	MaximumGradient  float64
	WatchTemperature float64
	RequiredPoints   int
}

func LayerRuleFor04() LayerRule04 {
	return LayerRule04{Name: "firn-layer-04", MinimumDepth: 8.0, MaximumGradient: 7.5, WatchTemperature: -29.0, RequiredPoints: 6}
}

func EvaluateLayer04(profile model.ThermalProfile) bool {
	rule := LayerRuleFor04()
	if len(profile.Points) < rule.RequiredPoints {
		return false
	}
	return profile.Coverage() >= rule.MinimumDepth && MaxGradient(profile) <= rule.MaximumGradient
}
