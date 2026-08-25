package policy

import "github.com/jb843051627/firn-signal/internal/model"

type LayerRule18 struct {
	Name             string
	MinimumDepth     float64
	MaximumGradient  float64
	WatchTemperature float64
	RequiredPoints   int
}

func LayerRuleFor18() LayerRule18 {
	return LayerRule18{Name: "firn-layer-18", MinimumDepth: 36.0, MaximumGradient: 21.5, WatchTemperature: -24.0, RequiredPoints: 5}
}

func EvaluateLayer18(profile model.ThermalProfile) bool {
	rule := LayerRuleFor18()
	if len(profile.Points) < rule.RequiredPoints {
		return false
	}
	return profile.Coverage() >= rule.MinimumDepth && MaxGradient(profile) <= rule.MaximumGradient
}
