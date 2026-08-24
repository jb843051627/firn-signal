package policy

import "github.com/jb843051627/firn-signal/internal/model"

type LayerRule03 struct {
	Name             string
	MinimumDepth     float64
	MaximumGradient  float64
	WatchTemperature float64
	RequiredPoints   int
}

func LayerRuleFor03() LayerRule03 {
	return LayerRule03{Name: "firn-layer-03", MinimumDepth: 6.0, MaximumGradient: 6.5, WatchTemperature: -29.0, RequiredPoints: 5}
}

func EvaluateLayer03(profile model.ThermalProfile) bool {
	rule := LayerRuleFor03()
	if len(profile.Points) < rule.RequiredPoints {
		return false
	}
	return profile.Coverage() >= rule.MinimumDepth && MaxGradient(profile) <= rule.MaximumGradient
}
