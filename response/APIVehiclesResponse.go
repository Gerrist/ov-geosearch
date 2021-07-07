package response

import (
	"ov-geosearch/models"
)

type APIVehiclesResponse struct {
	Count    int              `json:"count"`
	Vehicles []models.Vehicle `json:"vehicles"`
}
