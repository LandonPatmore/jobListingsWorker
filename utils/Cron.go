package utils

import "github.com/jasonlvhit/gocron"

// Starts the cron job.
func Start() {
	gocron.Every(10).Seconds().Do(RetrieveStationData)
	<- gocron.Start()
}