package Router

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/nleeper/goment"
	"log"
	"net/http"
	"ov-geosearch/models"
	"ov-geosearch/response"
	vehiclestore "ov-geosearch/store"
	"ov-geosearch/util"
	"strconv"
	"time"
)

type Map map[string]interface{}

var myClient = &http.Client{Timeout: 10 * time.Second}

func Serve(port string) {
	r := mux.NewRouter()

	r.HandleFunc("/api/status", getStatus).Methods("GET")
	r.HandleFunc("/api/vehicles", getVehicles).Methods("GET")
	r.HandleFunc("/api/vehicle/{operator}/{id}", getVehicle).Methods("GET")
	r.HandleFunc("/api/transferable", getTransferable).Methods("GET")

	log.Fatalln(http.ListenAndServe(port, r))
}

func getStatus(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	response := Map{
		"online": true,
	}

	json.NewEncoder(res).Encode(response)
}

func getVehicles(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	g, _ := goment.New()
	resultVehicles := make([]models.Vehicle, 0)

	for _, vehicle := range vehiclestore.All() {
		secondsSinceUpdate := g.Diff(goment.Unix(vehicle.UpdateTimestamp))

		if secondsSinceUpdate <= 30 {
			resultVehicles = append(resultVehicles, vehicle)
		}
	}

	response := response.APIVehiclesResponse{
		Count:    vehiclestore.Count(),
		Vehicles: resultVehicles,
	}

	json.NewEncoder(res).Encode(response)
}

func getVehicle(res http.ResponseWriter, req *http.Request) {
	operator := mux.Vars(req)["operator"]
	vehicleId := mux.Vars(req)["id"]

	vehicle, found := vehiclestore.All()[operator+":"+vehicleId]

	response := response.APIVehicleResponse{
		Found:   found,
		Vehicle: vehicle,
	}

	json.NewEncoder(res).Encode(response)
}

func getTransferable(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	lat := req.URL.Query().Get("lat")
	lon := req.URL.Query().Get("lon")
	speed := req.URL.Query().Get("speed")
	realtimeTripId := req.URL.Query().Get("realtime_trip_id")
	date := req.URL.Query().Get("date")

	if lat == "" || lon == "" || realtimeTripId == "" || date == "" {
		json.NewEncoder(res).Encode(Map{
			"success": false,
			"error":   "Missing required (lat, lon, realtime_trip_id, date) parameters",
		})
	} else {
		//tripUrl := "http://maglev1.travelguide.moopmobility.nl/v1/trips/" + date + "/" + realtimeTripId
		//
		//resp, err := http.Get(tripUrl)
		//if err != nil {
		//	log.Fatalln(err)
		//}
		//body, err := ioutil.ReadAll(resp.Body)
		//sb := string(body)
		//
		//errorMessage := gjson.Get(sb, "errorMessage").String()
		//
		//if errorMessage == "" { // trip found
		//	tripData := gjson.Get(sb, "tripData")
		//	stops := gjson.Get(sb, "stops")
		//
		//	log.Println(tripData)
		//} else {
		//	json.NewEncoder(res).Encode(Map{
		//		"success": false,
		//		"error":   errorMessage,
		//	})
		//}

		nearbyVehicles := make([]models.Vehicle, 0)

		latFloat, _ := strconv.ParseFloat(lat, 64)
		lonFloat, _ := strconv.ParseFloat(lon, 64)
		speedFloat, _ := strconv.ParseFloat(speed, 64)

		for _, vehicle := range vehiclestore.All() {
			if util.Distance(latFloat, lonFloat, vehicle.Lat, vehicle.Lon) < util.GetSearchRange(speedFloat) {
				nearbyVehicles = append(nearbyVehicles, vehicle)
			}
		}

		json.NewEncoder(res).Encode(Map{
			"success":      true,
			"transferable": nearbyVehicles,
		})

		//json.NewEncoder(res).Encode(Map{
		//	"success": false,
		//	"tripUrl": tripUrl,
		//})

		//json.NewDecoder(re.Body).Decode(Map{tripUrl: tripUrl})
	}

}
