package policy

import "github.com/jb843051627/firn-signal/internal/model"

type LayerRule09 struct {
	Name             string
	MinimumDepth     float64
	MaximumGradient  float64
	WatchTemperature float64
	RequiredPoints   int
}

func LayerRuleFor09() LayerRule09 {
	return LayerRule09{Name: "firn-layer-09", MinimumDepth: 18.0, MaximumGradient: 12.5, WatchTemperature: -27.0, RequiredPoints: 6}
}

func EvaluateLayer09(profile model.ThermalProfile) bool {
	rule := LayerRuleFor09()
	if len(profile.Points) < rule.RequiredPoints {
		return false
	}
	return profile.Coverage() >= rule.MinimumDepth && MaxGradient(profile) <= rule.MaximumGradient
}
