package model

type SignalLevel string

const (
	SignalInfo    SignalLevel = "info"
	SignalWatch   SignalLevel = "watch"
	SignalBlocker SignalLevel = "blocker"
)

type Signal struct {
	Code     string      `json:"code"`
	Level    SignalLevel `json:"level"`
	Message  string      `json:"message"`
	Blocking bool        `json:"blocking"`
}

func (s Signal) BlocksRelease() bool { return s.Blocking || s.Level == SignalBlocker }
