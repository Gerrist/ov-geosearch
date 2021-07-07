package util

import (
	"encoding/json"
	"github.com/StefanSchroeder/Golang-Ellipsoid/ellipsoid"
	"math"
	"net/http"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func Distance(lat1 float64, lon1 float64, lat2 float64, lon2 float64) (distance float64) {
	geo1 := ellipsoid.Init("WGS84", ellipsoid.Degrees, ellipsoid.Meter, ellipsoid.LongitudeIsSymmetric, ellipsoid.BearingIsSymmetric)
	distance, _ = geo1.To(lat1, lon1, lat2, lon2)

	return distance
}

func GetSearchRange(speed float64) float64 {
	return math.Max(speed*30, 150)
}