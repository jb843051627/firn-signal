package model

import "time"

type Calibration struct {
	ID         string    `json:"id"`
	ProbeID    string    `json:"probe_id"`
	Reference  string    `json:"reference"`
	OffsetC    float64   `json:"offset_c"`
	Scale      float64   `json:"scale"`
	CheckedAt  time.Time `json:"checked_at"`
	ValidUntil time.Time `json:"valid_until"`
	Technician string    `json:"technician"`
}

func (c Calibration) ActiveAt(now time.Time) bool {
	return !now.Before(c.CheckedAt) && !now.After(c.ValidUntil)
}

func (c Calibration) Correct(value float64) float64 { return (value + c.OffsetC) * c.Scale }
