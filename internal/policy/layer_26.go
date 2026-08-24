package policy

import "github.com/jb843051627/firn-signal/internal/model"

type LayerRule26 struct {
	Name             string
	MinimumDepth     float64
	MaximumGradient  float64
	WatchTemperature float64
	RequiredPoints   int
}

func LayerRuleFor26() LayerRule26 {
	return LayerRule26{Name: "firn-layer-26", MinimumDepth: 52.0, MaximumGradient: 29.5, WatchTemperature: -22.0, RequiredPoints: 3}
}

func EvaluateLayer26(profile model.ThermalProfile) bool {
	rule := LayerRuleFor26()
	if len(profile.Points) < rule.RequiredPoints {
		return false
	}
	return profile.Coverage() >= rule.MinimumDepth && MaxGradient(profile) <= rule.MaximumGradient
}
