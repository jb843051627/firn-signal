package model

import "time"

type ProfilePoint struct {
	DepthM     float64 `json:"depth_m"`
	RawTempC   float64 `json:"raw_temp_c"`
	CorrectedC float64 `json:"corrected_c"`
	Gradient   float64 `json:"gradient"`
	ReadingID  string  `json:"reading_id"`
}

type ThermalProfile struct {
	ID            string         `json:"id"`
	ScanID        string         `json:"scan_id"`
	BoreholeID    string         `json:"borehole_id"`
	CalibrationID string         `json:"calibration_id"`
	Points        []ProfilePoint `json:"points"`
	CreatedAt     time.Time      `json:"created_at"`
}

func (p ThermalProfile) Coverage() float64 {
	if len(p.Points) == 0 {
		return 0
	}
	return p.Points[len(p.Points)-1].DepthM - p.Points[0].DepthM
}
