package model

import "time"

type QualityAssessment struct {
	ID          string       `json:"id"`
	ScanID      string       `json:"scan_id"`
	State       QualityState `json:"state"`
	Score       float64      `json:"score"`
	CoverageM   float64      `json:"coverage_m"`
	Signals     []Signal     `json:"signals"`
	Reviewer    string       `json:"reviewer"`
	EvaluatedAt time.Time    `json:"evaluated_at"`
}

func (q QualityAssessment) HasBlocker() bool {
	for _, signal := range q.Signals {
		if signal.BlocksRelease() {
			return true
		}
	}
	return false
}
func (q QualityAssessment) Accepted() bool { return q.State == QualityAccepted && !q.HasBlocker() }
