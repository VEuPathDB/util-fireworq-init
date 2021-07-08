package main

import (
	"log"

	"util-multi-blast-queue-init/internal/setup"
)

func main() {
	cliConf := setup.ParseCLI()

	log.Println("Loading queue config from " + cliConf.ConfigFile)

	queueConf  := setup.LoadQueueConfig(cliConf.ConfigFile)
	liveQueues := setup.LoadLiveQueues(cliConf.QueueURL)
	liveRoutes := setup.LoadLiveRoutes(cliConf.QueueURL)

	for k := range queueConf {
		// If the queue doesn't exist, submit it
		if _, ok := liveQueues[k]; !ok {
			setup.SubmitQueue(cliConf.QueueURL, queueConf[k])
			setup.AwaitQueue(cliConf.QueueURL, queueConf[k])
		}

		// Ensure all the configured routes exist
		for _, cat := range queueConf[k].Categories {

			// If the route doesn't exist, submit it
			if _, ok := liveRoutes[cat]; !ok {
				setup.SubmitRoute(cliConf.QueueURL, cat, queueConf[k])
			}
		}
	}

	log.Println("Done")
}
