package processor

import (
	"ov-geosearch/models"
	vehiclestore "ov-geosearch/store"
)

func ProcessPosition(positionUpdate models.PositionUpdate) {
	vehiclestore.Set(positionUpdate.Operator, positionUpdate.Vehicle, models.Vehicle{
		Id:              positionUpdate.Vehicle,
		Operator:        positionUpdate.Operator,
		RealtimeTripId:  positionUpdate.RealtimeTripId,
		Date:            positionUpdate.Date,
		UpdateTimestamp: positionUpdate.Timestamp,
		Lat:             positionUpdate.Lat,
		Lon:             positionUpdate.Lon,
	})
}
