package response

import (
	"ov-geosearch/models"
)

type APIVehicleResponse struct {
	Found   bool           `json:"found"`
	Vehicle models.Vehicle `json:"vehicle"`
}
