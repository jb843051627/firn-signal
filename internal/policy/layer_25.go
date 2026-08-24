package policy

import "github.com/jb843051627/firn-signal/internal/model"

type LayerRule25 struct {
	Name             string
	MinimumDepth     float64
	MaximumGradient  float64
	WatchTemperature float64
	RequiredPoints   int
}

func LayerRuleFor25() LayerRule25 {
	return LayerRule25{Name: "firn-layer-25", MinimumDepth: 50.0, MaximumGradient: 28.5, WatchTemperature: -22.0, RequiredPoints: 2}
}

func EvaluateLayer25(profile model.ThermalProfile) bool {
	rule := LayerRuleFor25()
	if len(profile.Points) < rule.RequiredPoints {
		return false
	}
	return profile.Coverage() >= rule.MinimumDepth && MaxGradient(profile) <= rule.MaximumGradient
}
