package main

import (
	"github.com/landonp1203/goUtils/common"
	"jobListingsWorker/utils"
)

func main() {
	utils.GetJobsJob() // do this just so that every time the container starts, we try to retrieve data

	common.CronStart([] uint64{6}, utils.GetJobsJob)
}
