package utils

import (
	"dataPullerWorker/types"
	"github.com/landonp1203/goUtils/loggly"
	"github.com/landonp1203/goUtils/networking"
	"log"
)

var stationsList = ReadStationsEnv()

// Retrieves data about a specific station.
func RetrieveStationData() {

	for _, station := range stationsList {
		var jsonData, err = networking.Get("https://tidesandcurrents.noaa.gov/api/datagetter?date=today&station=" + station + "&product=water_temperature&units=english&time_zone=gmt&application=tidal_station_project&format=json")

		if err != nil {
			log.Println(err)
		} else {
			var waterTemperatureStation = types.WaterTemperatureStation{}
			DecodeJson(jsonData, &waterTemperatureStation)

			loggly.InfoEcho(waterTemperatureStation.LogStruct())
		}
	}

}
