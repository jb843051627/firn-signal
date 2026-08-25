package report

import (
	"github.com/jb843051627/firn-signal/internal/model"
	"strings"
)

func TextDiagnostic(value model.DiagnosticReport) string { return strings.Join(value.Checks, " | ") }
func HasCheck(value model.DiagnosticReport, expected string) bool {
	for _, check := range value.Checks {
		if check == expected {
			return true
		}
	}
	return false
}
