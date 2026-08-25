package validation

import "math"

func Temperature(value float64) error {
	if math.IsNaN(value) || math.IsInf(value, 0) || value < -80 || value > 30 {
		return FieldError{Field: "temp_c", Message: "is outside firn range"}
	}
	return nil
}

func Conductivity(value float64) error {
	if math.IsNaN(value) || value < 0 || value > 20 {
		return FieldError{Field: "conductivity", Message: "is outside probe range"}
	}
	return nil
}
