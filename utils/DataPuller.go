package utils

import (
	"dataPullerWorker/types"
	"fmt"
)

// Retrieves data about a specific station.
func RetrieveStationData() {
	var jsonData = Get("https://tidesandcurrents.noaa.gov/api/datagetter?date=today&station=8454000&product=water_temperature&units=english&time_zone=gmt&application=tidal_station_project&format=json")

	var waterTemperatureStation = types.WaterTemperatureStation{}
	DecodeJson(jsonData, &waterTemperatureStation)

	// Just print to console the actual API response
	fmt.Printf("%+v\n", waterTemperatureStation)

	Info(waterTemperatureStation.LogStruct())
}
