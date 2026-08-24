package policy

import "github.com/jb843051627/firn-signal/internal/model"

type LayerRule29 struct {
	Name             string
	MinimumDepth     float64
	MaximumGradient  float64
	WatchTemperature float64
	RequiredPoints   int
}

func LayerRuleFor29() LayerRule29 {
	return LayerRule29{Name: "firn-layer-29", MinimumDepth: 58.0, MaximumGradient: 32.5, WatchTemperature: -21.0, RequiredPoints: 6}
}

func EvaluateLayer29(profile model.ThermalProfile) bool {
	rule := LayerRuleFor29()
	if len(profile.Points) < rule.RequiredPoints {
		return false
	}
	return profile.Coverage() >= rule.MinimumDepth && MaxGradient(profile) <= rule.MaximumGradient
}
