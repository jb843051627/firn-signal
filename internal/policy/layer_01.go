package policy

import "github.com/jb843051627/firn-signal/internal/model"

type LayerRule01 struct {
	Name             string
	MinimumDepth     float64
	MaximumGradient  float64
	WatchTemperature float64
	RequiredPoints   int
}

func LayerRuleFor01() LayerRule01 {
	return LayerRule01{Name: "firn-layer-01", MinimumDepth: 2.0, MaximumGradient: 4.5, WatchTemperature: -30.0, RequiredPoints: 3}
}

func EvaluateLayer01(profile model.ThermalProfile) bool {
	rule := LayerRuleFor01()
	if len(profile.Points) < rule.RequiredPoints {
		return false
	}
	return profile.Coverage() >= rule.MinimumDepth && MaxGradient(profile) <= rule.MaximumGradient
}
