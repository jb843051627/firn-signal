package report

import "github.com/jb843051627/firn-signal/internal/model"

func BlockingCodes(signals []model.Signal) []string {
	codes := make([]string, 0)
	for _, signal := range signals {
		if signal.BlocksRelease() {
			codes = append(codes, signal.Code)
		}
	}
	return codes
}

func WatchCodes(signals []model.Signal) []string {
	codes := make([]string, 0)
	for _, signal := range signals {
		if signal.Level == model.SignalWatch {
			codes = append(codes, signal.Code)
		}
	}
	return codes
}
