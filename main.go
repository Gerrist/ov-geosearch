package main

import (
	"fmt"
	"os"
	Router "ov-geosearch/api"
	qh "ov-geosearch/queuehandler"
)

//var VehicleStore Vehicles = make(map[string]Vehicle)

func main() {
	fmt.Println("OVGeoSearch")
	fmt.Println("Connecting to database")
	fmt.Println("Starting queue handler")
	go qh.QueueHandler()
	fmt.Println("Serving REST API...")

	port := os.Getenv("PORT")
	if(port == ""){
		port = "8080"
	}

	Router.Serve(":" + port)
}
