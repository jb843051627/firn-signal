package policy

import "github.com/jb843051627/firn-signal/internal/model"

type LayerRule14 struct {
	Name             string
	MinimumDepth     float64
	MaximumGradient  float64
	WatchTemperature float64
	RequiredPoints   int
}

func LayerRuleFor14() LayerRule14 {
	return LayerRule14{Name: "firn-layer-14", MinimumDepth: 28.0, MaximumGradient: 17.5, WatchTemperature: -26.0, RequiredPoints: 6}
}

func EvaluateLayer14(profile model.ThermalProfile) bool {
	rule := LayerRuleFor14()
	if len(profile.Points) < rule.RequiredPoints {
		return false
	}
	return profile.Coverage() >= rule.MinimumDepth && MaxGradient(profile) <= rule.MaximumGradient
}
