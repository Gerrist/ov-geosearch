package vehiclestore

import (
	"ov-geosearch/models"
)

type Vehicles = map[string]models.Vehicle

var VehicleStore Vehicles = make(map[string]models.Vehicle)

func Get(operator string, id string) models.Vehicle {
	return VehicleStore[operator+":"+id]
}

func Set(operator string, id string, vehicle models.Vehicle) {
	VehicleStore[operator+":"+id] = vehicle
}

func Count() int {
	return len(VehicleStore)
}

func All() Vehicles {
	return VehicleStore
}
