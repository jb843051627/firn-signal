package model

import "time"

type BoreholeReport struct {
	BoreholeID      string    `json:"borehole_id"`
	Day             time.Time `json:"day"`
	ScanCount       int       `json:"scan_count"`
	PublishedCount  int       `json:"published_count"`
	AverageSurfaceC float64   `json:"average_surface_c"`
	BlockingSignals int       `json:"blocking_signals"`
	Warnings        []string  `json:"warnings"`
}

type DiagnosticReport struct {
	ScanID        string   `json:"scan_id"`
	CalibrationID string   `json:"calibration_id"`
	Checks        []string `json:"checks"`
}
