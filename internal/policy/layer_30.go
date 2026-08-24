package policy

import "github.com/jb843051627/firn-signal/internal/model"

type LayerRule30 struct {
	Name             string
	MinimumDepth     float64
	MaximumGradient  float64
	WatchTemperature float64
	RequiredPoints   int
}

func LayerRuleFor30() LayerRule30 {
	return LayerRule30{Name: "firn-layer-30", MinimumDepth: 60.0, MaximumGradient: 33.5, WatchTemperature: -20.0, RequiredPoints: 2}
}

func EvaluateLayer30(profile model.ThermalProfile) bool {
	rule := LayerRuleFor30()
	if len(profile.Points) < rule.RequiredPoints {
		return false
	}
	return profile.Coverage() >= rule.MinimumDepth && MaxGradient(profile) <= rule.MaximumGradient
}
