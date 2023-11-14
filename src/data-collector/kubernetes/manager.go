package kubernetes

import (
	"context"
	"encoding/json"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var clientset = InitKubeconfig()

func toString(v any) string {
	data, err := json.Marshal(v)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return string(data)
}

func GetNodes() string {
	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	return toString(nodes)
}

func GetPods() string {
	pods, err := clientset.CoreV1().Pods("rabbitmq").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	return toString(pods)
}
