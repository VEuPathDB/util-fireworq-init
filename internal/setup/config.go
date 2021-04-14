package setup

import (
	"os"

	"gopkg.in/yaml.v3"
)

type CLIConfig struct {
	ConfigFile string
	QueueURL   string
}

func LoadQueueConfig(confPath string) map[string]Queue {
	file, err := os.Open(confPath)
	bail(err)
	defer file.Close()

	queues := new(QueueWrapper)
	dec := yaml.NewDecoder(file)
	bail(dec.Decode(&queues))

	out := make(map[string]Queue, len(queues.Queues))
	for i := range queues.Queues {
		if _, ok := out[queues.Queues[i].Name]; ok {
			panic("Duplicate queue names in the queue config file.")
		}
		out[queues.Queues[i].Name] = queues.Queues[i]
	}

	return out
}