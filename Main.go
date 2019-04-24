package main

import "jobListingsWorker/utils"

func main() {
	utils.GetJobsJob() // do this just so that every time the container starts, we try to retrieve data

	utils.CronStart()
}
