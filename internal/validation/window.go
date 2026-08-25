package validation

import "time"

func Ordered(start, end time.Time) error {
	if start.IsZero() || end.IsZero() || !end.After(start) {
		return FieldError{Field: "time_window", Message: "must be ordered"}
	}
	return nil
}

func SameUTCDate(left, right time.Time) bool {
	l := left.UTC()
	r := right.UTC()
	return l.Year() == r.Year() && l.YearDay() == r.YearDay()
}
