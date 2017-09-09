package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"
	"fmt"
)

func main() {

	svc := ecs.New(session.New())
	input := &ecs.ListClustersInput{}

	result, err := svc.ListClusters(input)
	if err != nil {
		fmt.Printf( "error while listing clusters %v", err)
	}

	for cluster := range result.ClusterArns {
		fmt.Printf("found cluster with arn %v", cluster)
	}
}