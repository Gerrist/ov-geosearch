package processor

import (
	"ov-geosearch/models"
	vehiclestore "ov-geosearch/store"
)

func ProcessPosition(positionUpdate models.PositionUpdate) {
	//log.Println(positionUpdate);

	//latFloat, _ = strconv.ParseFloat(positionUpdate.Lat, 32)

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
