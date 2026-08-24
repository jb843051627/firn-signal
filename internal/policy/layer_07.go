package policy

import "github.com/jb843051627/firn-signal/internal/model"

type LayerRule07 struct {
	Name             string
	MinimumDepth     float64
	MaximumGradient  float64
	WatchTemperature float64
	RequiredPoints   int
}

func LayerRuleFor07() LayerRule07 {
	return LayerRule07{Name: "firn-layer-07", MinimumDepth: 14.0, MaximumGradient: 10.5, WatchTemperature: -28.0, RequiredPoints: 4}
}

func EvaluateLayer07(profile model.ThermalProfile) bool {
	rule := LayerRuleFor07()
	if len(profile.Points) < rule.RequiredPoints {
		return false
	}
	return profile.Coverage() >= rule.MinimumDepth && MaxGradient(profile) <= rule.MaximumGradient
}
