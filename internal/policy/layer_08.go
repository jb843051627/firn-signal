package policy

import "github.com/jb843051627/firn-signal/internal/model"

type LayerRule08 struct {
	Name             string
	MinimumDepth     float64
	MaximumGradient  float64
	WatchTemperature float64
	RequiredPoints   int
}

func LayerRuleFor08() LayerRule08 {
	return LayerRule08{Name: "firn-layer-08", MinimumDepth: 16.0, MaximumGradient: 11.5, WatchTemperature: -28.0, RequiredPoints: 5}
}

func EvaluateLayer08(profile model.ThermalProfile) bool {
	rule := LayerRuleFor08()
	if len(profile.Points) < rule.RequiredPoints {
		return false
	}
	return profile.Coverage() >= rule.MinimumDepth && MaxGradient(profile) <= rule.MaximumGradient
}
