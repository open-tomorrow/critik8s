package amqp

import (
	"encoding/json"
	"os"
	"strings"
)

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}

func GetNodes() any {
	msg := RpcRequest("cluster.resource.nodes", "")

	var nodes map[string]interface{}
	json.Unmarshal([]byte(msg), &nodes)

	return nodes
}

func GetPods() any {
	msg := RpcRequest("cluster.resource.pods", "")

	var pods map[string]interface{}
	json.Unmarshal([]byte(msg), &pods)

	return pods
}
