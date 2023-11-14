package kubernetes

import (
	"log"
)

var resources = map[string]func() string{
	"cluster.resource.nodes": GetNodes,
	"cluster.resource.pods":  GetPods,
}

func GetResource(id string) string {
	f := resources[id]

	log.Printf(" [%s]: sending resource", id)

	return f()
}
