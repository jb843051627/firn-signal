package policy

import "github.com/jb843051627/firn-signal/internal/model"

type LayerRule02 struct {
	Name             string
	MinimumDepth     float64
	MaximumGradient  float64
	WatchTemperature float64
	RequiredPoints   int
}

func LayerRuleFor02() LayerRule02 {
	return LayerRule02{Name: "firn-layer-02", MinimumDepth: 4.0, MaximumGradient: 5.5, WatchTemperature: -30.0, RequiredPoints: 4}
}

func EvaluateLayer02(profile model.ThermalProfile) bool {
	rule := LayerRuleFor02()
	if len(profile.Points) < rule.RequiredPoints {
		return false
	}
	return profile.Coverage() >= rule.MinimumDepth && MaxGradient(profile) <= rule.MaximumGradient
}
