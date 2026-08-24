package model

import "time"

type CreateBoreholeInput struct {
	ID        string  `json:"id"`
	Site      string  `json:"site"`
	Label     string  `json:"label"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	DepthM    float64 `json:"depth_m"`
	Notes     string  `json:"notes"`
}

type RegisterProbeInput struct {
	ID          string  `json:"id"`
	BoreholeID  string  `json:"borehole_id"`
	Serial      string  `json:"serial"`
	DepthStartM float64 `json:"depth_start_m"`
	DepthEndM   float64 `json:"depth_end_m"`
	Notes       string  `json:"notes"`
}

type StartScanInput struct {
	ID         string `json:"id"`
	BoreholeID string `json:"borehole_id"`
	Operator   string `json:"operator"`
	Sequence   int    `json:"sequence"`
}

type RecordCalibrationInput struct {
	ID         string    `json:"id"`
	ProbeID    string    `json:"probe_id"`
	Reference  string    `json:"reference"`
	OffsetC    float64   `json:"offset_c"`
	Scale      float64   `json:"scale"`
	CheckedAt  time.Time `json:"checked_at"`
	ValidUntil time.Time `json:"valid_until"`
	Technician string    `json:"technician"`
}

type RecordReadingInput struct {
	ID           string            `json:"id"`
	ScanID       string            `json:"scan_id"`
	ProbeID      string            `json:"probe_id"`
	DepthM       float64           `json:"depth_m"`
	TempC        float64           `json:"temp_c"`
	Conductivity float64           `json:"conductivity"`
	CollectedAt  time.Time         `json:"collected_at"`
	Labels       map[string]string `json:"labels"`
}

type AssessInput struct {
	Reviewer string `json:"reviewer"`
}
