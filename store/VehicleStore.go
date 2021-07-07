package vehiclestore

import (
	"ov-geosearch/models"
	"sync"
)

type Vehicles = map[string]models.Vehicle

var VehicleStore Vehicles = make(map[string]models.Vehicle)

func Get(operator string, id string) models.Vehicle {
	return VehicleStore[operator+":"+id]
}

func Set(operator string, id string, vehicle models.Vehicle) {
	var mutex = &sync.Mutex{}
	mutex.Lock()
	VehicleStore[operator+":"+id] = vehicle
	mutex.Unlock()
}

func Count() int {
	return len(VehicleStore)
}

func All() Vehicles {
	return VehicleStore
}
