package policy

import "github.com/jb843051627/firn-signal/internal/model"

type LayerRule24 struct {
	Name             string
	MinimumDepth     float64
	MaximumGradient  float64
	WatchTemperature float64
	RequiredPoints   int
}

func LayerRuleFor24() LayerRule24 {
	return LayerRule24{Name: "firn-layer-24", MinimumDepth: 48.0, MaximumGradient: 27.5, WatchTemperature: -22.0, RequiredPoints: 6}
}

func EvaluateLayer24(profile model.ThermalProfile) bool {
	rule := LayerRuleFor24()
	if len(profile.Points) < rule.RequiredPoints {
		return false
	}
	return profile.Coverage() >= rule.MinimumDepth && MaxGradient(profile) <= rule.MaximumGradient
}
