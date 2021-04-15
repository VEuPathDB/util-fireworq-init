package setup

import (
	"encoding/json"

	"github.com/Foxcapades/Go-Chainrequest/simple"
)

type QueueWrapper struct {
	Queues []Queue `yaml:"queues"`
}

type Queue struct {
	Name            string `json:"name" yaml:"name"`
	PollingInterval uint   `json:"polling_interval" yaml:"pollingInterval"`
	MaxWorkers      uint   `json:"max_workers" yaml:"maxWorkers"`
}

func LoadLiveQueues(url string) map[string]Queue {
	res := simple.GetRequest(prefixUrl(url) + "/queues").Submit()
	bail(res.GetError())

	if res.MustGetResponseCode() != 200 {
		panic("unexpected response from queue server")
	}

	tmp := make([]Queue, 0, 5)
	bail(res.UnmarshalBody(&tmp, simple.UnmarshallerFunc(json.Unmarshal)))

	out := make(map[string]Queue, len(tmp))
	for i := range tmp {
		out[tmp[i].Name] = tmp[i]
	}

	return out
}

func SubmitQueue(url string, q Queue) {
	res := simple.PutRequest(prefixUrl(url) + "/queue/" + q.Name).
		MarshalBody(q, simple.MarshallerFunc(json.Marshal)).
		Submit()
	bail(res.GetError())

	if res.MustGetResponseCode() != 200 {
		panic("unexpected response from queue server: " + string(res.MustGetBody()))
	}
}