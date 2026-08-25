package model

import "time"

type Probe struct {
	ID          string     `json:"id"`
	BoreholeID  string     `json:"borehole_id"`
	Serial      string     `json:"serial"`
	DepthStartM float64    `json:"depth_start_m"`
	DepthEndM   float64    `json:"depth_end_m"`
	State       ProbeState `json:"state"`
	InstalledAt time.Time  `json:"installed_at"`
	Notes       string     `json:"notes"`
}

func (p Probe) Covers(depth float64) bool {
	return depth >= p.DepthStartM && depth <= p.DepthEndM
}
