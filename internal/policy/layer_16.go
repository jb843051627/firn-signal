package policy

import "github.com/jb843051627/firn-signal/internal/model"

type LayerRule16 struct {
	Name             string
	MinimumDepth     float64
	MaximumGradient  float64
	WatchTemperature float64
	RequiredPoints   int
}

func LayerRuleFor16() LayerRule16 {
	return LayerRule16{Name: "firn-layer-16", MinimumDepth: 32.0, MaximumGradient: 19.5, WatchTemperature: -25.0, RequiredPoints: 3}
}

func EvaluateLayer16(profile model.ThermalProfile) bool {
	rule := LayerRuleFor16()
	if len(profile.Points) < rule.RequiredPoints {
		return false
	}
	return profile.Coverage() >= rule.MinimumDepth && MaxGradient(profile) <= rule.MaximumGradient
}
