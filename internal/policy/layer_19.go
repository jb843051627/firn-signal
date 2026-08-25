package policy

import "github.com/jb843051627/firn-signal/internal/model"

type LayerRule19 struct {
	Name             string
	MinimumDepth     float64
	MaximumGradient  float64
	WatchTemperature float64
	RequiredPoints   int
}

func LayerRuleFor19() LayerRule19 {
	return LayerRule19{Name: "firn-layer-19", MinimumDepth: 38.0, MaximumGradient: 22.5, WatchTemperature: -24.0, RequiredPoints: 6}
}

func EvaluateLayer19(profile model.ThermalProfile) bool {
	rule := LayerRuleFor19()
	if len(profile.Points) < rule.RequiredPoints {
		return false
	}
	return profile.Coverage() >= rule.MinimumDepth && MaxGradient(profile) <= rule.MaximumGradient
}
