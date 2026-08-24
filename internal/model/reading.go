package model

import "time"

type ThermalReading struct {
	ID           string            `json:"id"`
	ScanID       string            `json:"scan_id"`
	ProbeID      string            `json:"probe_id"`
	DepthM       float64           `json:"depth_m"`
	TempC        float64           `json:"temp_c"`
	Conductivity float64           `json:"conductivity"`
	CollectedAt  time.Time         `json:"collected_at"`
	Labels       map[string]string `json:"labels"`
}

func (r ThermalReading) Valid() bool {
	return r.ID != "" && r.ScanID != "" && r.ProbeID != "" && r.DepthM >= 0 && r.TempC > -80 && r.TempC < 30
}
