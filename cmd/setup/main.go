package main

import (
	"log"

	"util-multi-blast-queue-init/internal/setup"
)

func main() {
	cliConf := setup.ParseCLI()

	log.Println("Loading queue config from " + cliConf.ConfigFile)

	queueConf  := setup.LoadQueueConfig(cliConf.ConfigFile)

	log.Println("Fetching queues from " + cliConf.QueueURL)

	liveQueues := setup.LoadLiveQueues(cliConf.QueueURL)

	for k := range queueConf {
		if _, ok := liveQueues[k]; !ok {
			log.Println("Creating new queue " + k)
			setup.SubmitQueue(cliConf.QueueURL, queueConf[k])
			setup.SubmitCategory(cliConf.QueueURL, queueConf[k])
		}
	}

	log.Println("Done")
}
