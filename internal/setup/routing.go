package setup

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Foxcapades/Go-Chainrequest/simple"
)

const (
	routeListURL = "/routings"
	routeBaseURL = "/routing"
	routeTargURL = routeBaseURL + "/%s"
)

type LiveRoute struct {
	QueueName   string `json:"queue_name"`
	JobCategory string `json:"job_category"`
}

type NewRoute struct {
	QueueName string `json:"queue_name"`
}

// LoadLiveRoutes loads a mapping of all currently configured routes indexed on
// each route's category name.
func LoadLiveRoutes(url string) map[string]LiveRoute {
	log.Printf("Loading previously configured routes from Fireworq")

	// Load list of routes
	res := simple.GetRequest(prefixUrl(url) + routeListURL).Submit()
	bail(res.GetError())

	// Bail if Fireworq returned a non-success code.
	if res.MustGetResponseCode() != 200 {
		panic("unexpected response from queue server: route list lookup failed")
	}

	// Deserialize the json response into a list of live routes
	tmp := make([]LiveRoute, 0, 5)
	bail(res.UnmarshalBody(&tmp, simple.UnmarshallerFunc(json.Unmarshal)))

	// Index the live routes on the `job_category` field
	out := make(map[string]LiveRoute, len(tmp))
	for i := range tmp {
		out[tmp[i].JobCategory] = tmp[i]
	}

	return out
}


func SubmitRoute(url, category string, q QueueConfig) {
	log.Printf(`Submitting route "%s" for queue "%s".`, category, q.QueueName)

	res := simple.PutRequest(prefixUrl(url)+fmt.Sprintf(routeTargURL, category)).
		MarshalBody(CategoryPut{
			QueueName: q.QueueName,
		}, simple.MarshallerFunc(json.Marshal)).
		Submit()
	bail(res.GetError())

	if res.MustGetResponseCode() != 200 {
		panic("unexpected response from queue server: " + string(res.MustGetBody()))
	}
}
