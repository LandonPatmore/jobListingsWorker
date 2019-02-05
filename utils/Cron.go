package utils

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
)

// Starts the cron job.
func CronStart() {
	fmt.Println("Data retrieval has been started.")
	gocron.Every(10).Minutes().Do(RetrieveStationData)
	<-gocron.Start()
}
