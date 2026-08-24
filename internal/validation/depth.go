package validation

import "fmt"

func Depth(value float64, field string) error {
	if value < 0 || value > 5000 {
		return FieldError{Field: field, Message: "is outside borehole range"}
	}
	return nil
}

func Span(start, end float64) error {
	if err := Depth(start, "depth_start_m"); err != nil {
		return err
	}
	if err := Depth(end, "depth_end_m"); err != nil {
		return err
	}
	if end <= start {
		return fmt.Errorf("depth_end_m must be greater than depth_start_m")
	}
	return nil
}

func Covers(start, end, depth float64) bool { return depth >= start && depth <= end }
