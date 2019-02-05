package utils

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
)

// Starts the cron job.
func Start() {
	fmt.Println("Data retrieval has been started.")
	gocron.Every(10).Seconds().Do(RetrieveStationData)
	<-gocron.Start()
}
