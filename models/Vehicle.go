package models

type Vehicle struct {
	Id              string `json:"id"`
	Operator        string `json:"operator"`
	RealtimeTripId  string `json:"realtime_trip_id"`
	Date            string `json:"date"`
	UpdateTimestamp int64  `json:"update_timestamp"`
	Lat             float64  `json:"lat"`
	Lon             float64  `json:"lon"`
}
