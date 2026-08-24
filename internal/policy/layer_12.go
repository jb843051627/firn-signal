package policy

import "github.com/jb843051627/firn-signal/internal/model"

type LayerRule12 struct {
	Name             string
	MinimumDepth     float64
	MaximumGradient  float64
	WatchTemperature float64
	RequiredPoints   int
}

func LayerRuleFor12() LayerRule12 {
	return LayerRule12{Name: "firn-layer-12", MinimumDepth: 24.0, MaximumGradient: 15.5, WatchTemperature: -26.0, RequiredPoints: 4}
}

func EvaluateLayer12(profile model.ThermalProfile) bool {
	rule := LayerRuleFor12()
	if len(profile.Points) < rule.RequiredPoints {
		return false
	}
	return profile.Coverage() >= rule.MinimumDepth && MaxGradient(profile) <= rule.MaximumGradient
}
