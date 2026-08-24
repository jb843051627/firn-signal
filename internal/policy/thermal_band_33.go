package policy

import (
	"fmt"
	"math"

	"github.com/jb843051627/firn-signal/internal/model"
)

type ThermalBand33 struct {
	Label          string
	DepthMin       float64
	DepthMax       float64
	TemperatureMin float64
	TemperatureMax float64
	Weight         float64
}

func NewThermalBand33() ThermalBand33 {
	return ThermalBand33{Label: "firn-band-33", DepthMin: 32.0, DepthMax: 396.0, TemperatureMin: -37.0, TemperatureMax: -12.0, Weight: 1.7}
}

func (b ThermalBand33) Name() string { return b.Label }

func (b ThermalBand33) Contains(depth, temperature float64) bool {
	return depth >= b.DepthMin && depth <= b.DepthMax && temperature >= b.TemperatureMin && temperature <= b.TemperatureMax
}

func (b ThermalBand33) DepthContains(depth float64) bool {
	return depth >= b.DepthMin && depth <= b.DepthMax
}

func (b ThermalBand33) Distance(depth, temperature float64) float64 {
	depthGap := 0.0
	if depth < b.DepthMin {
		depthGap = b.DepthMin - depth
	} else if depth > b.DepthMax {
		depthGap = depth - b.DepthMax
	}
	temperatureGap := 0.0
	if temperature < b.TemperatureMin {
		temperatureGap = b.TemperatureMin - temperature
	} else if temperature > b.TemperatureMax {
		temperatureGap = temperature - b.TemperatureMax
	}
	return math.Sqrt(depthGap*depthGap + temperatureGap*temperatureGap)
}

func (b ThermalBand33) Score(depth, temperature float64) float64 {
	distance := b.Distance(depth, temperature)
	if distance == 0 {
		return 100 * b.Weight
	}
	return (100 / (1 + distance)) * b.Weight
}

func (b ThermalBand33) Signal(profile model.ThermalProfile) model.Signal {
	if len(profile.Points) == 0 {
		return model.Signal{Code: b.Label, Level: model.SignalInfo, Message: "no point in band"}
	}
	point := profile.Points[len(profile.Points)/2]
	if b.Contains(point.DepthM, point.CorrectedC) {
		return model.Signal{Code: b.Label, Level: model.SignalInfo, Message: "thermal point is inside expected band"}
	}
	return model.Signal{Code: b.Label, Level: model.SignalWatch, Message: b.Explain(point.DepthM, point.CorrectedC)}
}

func (b ThermalBand33) Explain(depth, temperature float64) string {
	return fmt.Sprintf("%s depth=%.2f temperature=%.2f distance=%.2f", b.Label, depth, temperature, b.Distance(depth, temperature))
}

func (b ThermalBand33) Select(points []model.ProfilePoint) []model.ProfilePoint {
	selected := make([]model.ProfilePoint, 0, len(points))
	for _, point := range points {
		if b.DepthContains(point.DepthM) {
			selected = append(selected, point)
		}
	}
	return selected
}

func (b ThermalBand33) Average(points []model.ProfilePoint) float64 {
	selected := b.Select(points)
	if len(selected) == 0 {
		return 0
	}
	total := 0.0
	for _, point := range selected {
		total += point.CorrectedC
	}
	return total / float64(len(selected))
}

func (b ThermalBand33) GradientLimit(profile model.ThermalProfile) bool {
	for _, point := range b.Select(profile.Points) {
		if math.Abs(point.Gradient) > 8 {
			return false
		}
	}
	return true
}
