package utils

import (
	"github.com/grokify/html-strip-tags-go"
	"github.com/landonp1203/goUtils/loggly"
	"github.com/landonp1203/goUtils/networking"
	"jobListingsWorker/types"
)

func GetJobsJob() {
	jobs := getJobPostings()

	if jobs != nil {
		sendJobsToDB(jobs)
	} else {
		loggly.Warn("The jobs list was empty.")
	}
}

// Retrieves job postings.
func getJobPostings() [] *types.GithubJob {

	var jsonData, err = networking.Get("https://jobs.github.com/positions.json?search=java")

	if err != nil {
		loggly.Error(err)
		return nil
	} else {
		var jobs [] *types.GithubJob
		DecodeJson(jsonData, &jobs)

		for _, j := range jobs { // strips html tags from `HowToApply` field
			stripped := strip.StripTags(j.HowToApply)
			j.HowToApply = stripped
		}

		return jobs
	}
}

// Helper to send the jobs array to the database
func sendJobsToDB(jobs [] *types.GithubJob) {
	for _, j := range jobs {
		err := PutItem(*j)
		if err != nil {
			loggly.Error(err)
			break
		}
	}
}
