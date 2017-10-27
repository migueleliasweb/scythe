package main

import (
	"log"

	"k8s.io/helm/pkg/helm"
)

const (
	//kubectl --namespace=kube-system port-forward {tiller's pod} 8080:44134
	address = "localhost:8080"
)

func main() {
	c := helm.NewClient(helm.Host(address))
	response := c.ListReleases()

	for r := range response.GetReleases() {
		log.Println(r)
	}
}
