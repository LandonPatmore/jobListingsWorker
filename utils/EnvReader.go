package utils

import (
	"log"
	"os"
	"strings"
)

func ReadStationsEnv() [] string {
	stationsList, found := os.LookupEnv("STATIONS_LIST")

	if found {
		return strings.Split(stationsList, ",")
	}
	
	log.Fatal(`The "STATIONS_LIST" env variable was not set.  Please set it.`)
	return nil
}
