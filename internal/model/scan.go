package model

import "time"

type ThermalScan struct {
	ID          string    `json:"id"`
	BoreholeID  string    `json:"borehole_id"`
	Operator    string    `json:"operator"`
	Sequence    int       `json:"sequence"`
	State       ScanState `json:"state"`
	StartedAt   time.Time `json:"started_at"`
	SealedAt    time.Time `json:"sealed_at"`
	EvaluatedAt time.Time `json:"evaluated_at"`
}

func (s ThermalScan) Open() bool      { return s.State == ScanOpen }
func (s ThermalScan) Sealed() bool    { return s.State == ScanSealed }
func (s ThermalScan) Evaluated() bool { return s.State == ScanEvaluated }
func (s ThermalScan) Released() bool  { return s.State == ScanReleased }
