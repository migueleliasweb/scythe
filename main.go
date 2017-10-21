package main

import (
	"log"

	"google.golang.org/grpc"
	"k8s.io/helm/pkg/helm"
	"k8s.io/helm/pkg/proto/hapi/services"
)

const (
	//kubectl --namespace=kube-system port-forward tiller-deploy-1936853538-rmm9n 8080:44134
	address = "localhost:8080"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	} else {
		log.Println("Connected!")
	}
	defer conn.Close()

	c := services.NewReleaseServiceClient(conn)
	request := services.ListReleasesRequest{}
	result, err := c.ListReleases(helm.NewContext(), &request)

	if err == nil {
		response, err2 := result.Recv()
		log.Println("Got ", len(response.GetReleases()), "releases !")

		if err2 == nil {
			for _, r := range response.GetReleases() {
				log.Println(r.GetInfo())
			}
		}
	} else {
		log.Println("Got error: ", err)
	}
}
