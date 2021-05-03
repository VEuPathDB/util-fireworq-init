package setup

import (
	"encoding/json"

	"github.com/Foxcapades/Go-Chainrequest/simple"
)

type QueueWrapper struct {
	Queues []QueueConfig `yaml:"queues"`
}

type QueueConfig struct {
	QueueName       string `json:"name" yaml:"name"`
	PollingInterval uint   `json:"polling_interval" yaml:"pollingInterval"`
	MaxWorkers      uint   `json:"max_workers" yaml:"maxWorkers"`
	Category        string `json:"job_category" yaml:"category"`
}

type QueueGet struct {
	QueueName       string `json:"name"`
	PollingInterval uint   `json:"polling_interval"`
	MaxWorkers      uint   `json:"max_workers"`
}

func LoadLiveQueues(url string) map[string]QueueGet {
	res := simple.GetRequest(prefixUrl(url) + "/queues").Submit()
	bail(res.GetError())

	if res.MustGetResponseCode() != 200 {
		panic("unexpected response from queue server")
	}

	tmp := make([]QueueGet, 0, 5)
	bail(res.UnmarshalBody(&tmp, simple.UnmarshallerFunc(json.Unmarshal)))

	out := make(map[string]QueueGet, len(tmp))
	for i := range tmp {
		out[tmp[i].QueueName] = tmp[i]
	}

	return out
}

type QueuePut struct {
	PollingInterval uint `json:"polling_interval"`
	MaxWorkers      uint `json:"max_workers"`
}

func SubmitQueue(url string, q QueueConfig) {
	res := simple.PutRequest(prefixUrl(url)+"/queue/"+q.QueueName).
		MarshalBody(QueuePut{
			PollingInterval: q.PollingInterval,
			MaxWorkers:      q.MaxWorkers,
		}, simple.MarshallerFunc(json.Marshal)).
		Submit()
	bail(res.GetError())

	if res.MustGetResponseCode() != 200 {
		panic("unexpected response from queue server: " + string(res.MustGetBody()))
	}
}

type CategoryPut struct {
	QueueName string `json:"name"`
}

func SubmitCategory(url string, q QueueConfig) {
	res := simple.PutRequest(prefixUrl(url)+"/routing/"+q.Category).
		MarshalBody(CategoryPut{
			QueueName: q.QueueName,
		}, simple.MarshallerFunc(json.Marshal)).
		Submit()
	bail(res.GetError())

	if res.MustGetResponseCode() != 200 {
		panic("unexpected response from queue server: " + string(res.MustGetBody()))
	}
}
