package model

type BoreholeState string

const (
	BoreholeActive   BoreholeState = "active"
	BoreholeArchived BoreholeState = "archived"
)

type ProbeState string

const (
	ProbeInstalled ProbeState = "installed"
	ProbeRemoved   ProbeState = "removed"
)

type ScanState string

const (
	ScanOpen      ScanState = "open"
	ScanSealed    ScanState = "sealed"
	ScanEvaluated ScanState = "evaluated"
	ScanReleased  ScanState = "released"
	ScanAbandoned ScanState = "abandoned"
)

type QualityState string

const (
	QualityAccepted QualityState = "accepted"
	QualityRejected QualityState = "rejected"
)

type ReleaseState string

const (
	ReleasePrepared  ReleaseState = "prepared"
	ReleasePublished ReleaseState = "published"
)
