package setup

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const (
	errBadString = `Queue config entry %d missing required field "%s".`
	errBadUint   = `Queue config entry %d field "%s" is missing or set to 0.` +
		`  This value must be a positive, non-zero integer.`
)

type CLIConfig struct {
	ConfigFile string
	QueueURL   string
}

func LoadQueueConfig(confPath string) map[string]QueueConfig {
	file, err := os.Open(confPath)
	bail(err)
	defer file.Close()

	queues := new(QueueWrapper)
	dec := yaml.NewDecoder(file)
	bail(dec.Decode(&queues))

	out := make(map[string]QueueConfig, len(queues.Queues))
	for i := range queues.Queues {
		if len(queues.Queues[i].QueueName) == 0 {
			panic(fmt.Sprintf(errBadString, i+1, "name"))
		}

		if queues.Queues[i].PollingInterval == 0 {
			panic(fmt.Sprintf(errBadUint, i+1, "pollingInterval"))
		}

		if queues.Queues[i].MaxWorkers == 0 {
			panic(fmt.Sprintf(errBadUint, i+1, "maxWorkers"))
		}

		if len(queues.Queues[i].Category) == 0 {
			panic(fmt.Sprintf(errBadString, i+1, "name"))
		}

		if _, ok := out[queues.Queues[i].QueueName]; ok {
			panic("Duplicate queue names in the queue config file.")
		}
		out[queues.Queues[i].QueueName] = queues.Queues[i]
	}

	return out
}
