package model

import "time"

type Borehole struct {
	ID          string        `json:"id"`
	Site        string        `json:"site"`
	Label       string        `json:"label"`
	Latitude    float64       `json:"latitude"`
	Longitude   float64       `json:"longitude"`
	DepthM      float64       `json:"depth_m"`
	State       BoreholeState `json:"state"`
	InstalledAt time.Time     `json:"installed_at"`
	Notes       string        `json:"notes"`
}

func (b Borehole) Active() bool   { return b.State == BoreholeActive }
func (b Borehole) Archived() bool { return b.State == BoreholeArchived }
