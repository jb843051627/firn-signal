package model

import "time"

type ReleaseManifest struct {
	ID            string       `json:"id"`
	ScanID        string       `json:"scan_id"`
	BoreholeID    string       `json:"borehole_id"`
	CalibrationID string       `json:"calibration_id"`
	State         ReleaseState `json:"state"`
	Summary       string       `json:"summary"`
	PreparedAt    time.Time    `json:"prepared_at"`
	PublishedAt   time.Time    `json:"published_at"`
}

func (r ReleaseManifest) Published() bool { return r.State == ReleasePublished && !r.PublishedAt.IsZero() }
