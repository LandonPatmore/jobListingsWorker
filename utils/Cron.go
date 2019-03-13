package utils

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
)

// Starts the cron job.
func CronStart() {
	fmt.Println("Data retrieval has been started.")
	gocron.Every(6).Hours().Do(GetJobsJob)
	<-gocron.Start()
}
