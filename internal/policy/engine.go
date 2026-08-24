package policy

import "github.com/jb843051627/firn-signal/internal/model"

type Engine struct{ profileNames []string }

func New() Engine {
	return Engine{profileNames: []string{"surface", "deep", "gradient", "void", "drift"}}
}
func (e Engine) ProfileNames() []string { return append([]string(nil), e.profileNames...) }
func (e Engine) ProfileCount() int      { return len(e.profileNames) }
func (e Engine) IsKnownProfile(name string) bool {
	for _, value := range e.profileNames {
		if value == name {
			return true
		}
	}
	return false
}
func (e Engine) Summary(profile model.ThermalProfile) string {
	return profile.CalibrationID + ":" + string(rune(len(profile.Points)))
}
