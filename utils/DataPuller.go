package utils

import (
	"dataPullerWorker/types"
	"fmt"
	"github.com/landonp1203/goUtils/loggly"
	"goUtils/networking"
)

var stationsList, fileReadError = ReadStationsList()

// Retrieves data about a specific station.
func RetrieveStationData() {
	if fileReadError != nil {
		loggly.Error("There was an error reading the stationsList.txt file.")
	} else {
		for _, station := range stationsList {
			var jsonData = networking.Get("https://tidesandcurrents.noaa.gov/api/datagetter?date=today&station=" + station + "&product=water_temperature&units=english&time_zone=gmt&application=tidal_station_project&format=json")

			var waterTemperatureStation = types.WaterTemperatureStation{}
			DecodeJson(jsonData, &waterTemperatureStation)

			// Just print to console the actual API response
			fmt.Printf("%+v\n\n", waterTemperatureStation)

			loggly.Info(waterTemperatureStation.LogStruct())
		}
	}
}
