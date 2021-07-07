package models

type PositionUpdate struct {
	Vehicle string
	Operator string
	Lat float64
	Lon float64
	RealtimeTripId string
	Date string
	Timestamp int64
}
