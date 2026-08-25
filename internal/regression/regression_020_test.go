package regression

import (
	"context"
	"fmt"
	"path/filepath"
	"testing"
	"time"

	"github.com/jb843051627/firn-signal/internal/model"
	"github.com/jb843051627/firn-signal/internal/policy"
	"github.com/jb843051627/firn-signal/internal/service"
	"github.com/jb843051627/firn-signal/internal/store"
)

func lab20(t testing.TB) (*service.Lab, *store.Store) {
	t.Helper()
	repository, err := store.Open(filepath.Join(t.TempDir(), "firn.db"))
	if err != nil {
		t.Fatalf("open store: %v", err)
	}
	lab := service.NewLab(repository)
	t.Cleanup(func() { _ = lab.Close(); _ = repository.Close() })
	return lab, repository
}

func readyScan20(t testing.TB, lab *service.Lab, ctx context.Context) (model.Borehole, model.Probe, model.ThermalScan) {
	t.Helper()
	now := time.Now().UTC()
	borehole, err := lab.CreateBorehole(ctx, model.CreateBoreholeInput{ID: "borehole-20", Site: "north ice cap", Label: "ridge collar", DepthM: 120})
	if err != nil {
		t.Fatalf("borehole: %v", err)
	}
	probe, err := lab.RegisterProbe(ctx, model.RegisterProbeInput{ID: "probe-20", BoreholeID: borehole.ID, Serial: "FIRN-20", DepthStartM: 0, DepthEndM: 120})
	if err != nil {
		t.Fatalf("probe: %v", err)
	}
	scan, err := lab.StartScan(ctx, model.StartScanInput{ID: "scan-20", BoreholeID: borehole.ID, Operator: "field crew", Sequence: 1})
	if err != nil {
		t.Fatalf("scan: %v", err)
	}
	if _, err := lab.RecordCalibration(ctx, model.RecordCalibrationInput{ID: "cal-20", ProbeID: probe.ID, Reference: "ice bath", Scale: 1, CheckedAt: now.Add(-time.Hour), ValidUntil: now.Add(24 * time.Hour), Technician: "crew"}); err != nil {
		t.Fatalf("calibration: %v", err)
	}
	for index, depth := range []float64{0, 20, 40, 60} {
		if _, err := lab.RecordReading(ctx, model.RecordReadingInput{ID: fmt.Sprintf("reading-20-%d", index), ScanID: scan.ID, ProbeID: probe.ID, DepthM: depth, TempC: -20 + depth/20, Conductivity: 1, CollectedAt: now.Add(time.Duration(index) * time.Minute), Labels: map[string]string{"site": "north"}}); err != nil {
			t.Fatalf("reading %d: %v", index, err)
		}
	}
	return borehole, probe, scan
}

func TestBug20_BlockerLowersScore(t *testing.T) {
	scan := model.ThermalScan{ID: "scan"}
	profile := model.ThermalProfile{Points: []model.ProfilePoint{{DepthM: 0}, {DepthM: 20}}}
	readings := []model.ThermalReading{{DepthM: 0}, {DepthM: 0}}
	assessment := policy.Assess(scan, profile, readings, "r", time.Now())
	if assessment.Score >= 100 {
		t.Fatalf("blocker score = %v", assessment.Score)
	}
}
