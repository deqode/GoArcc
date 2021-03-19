package awsDefault

import "context"

func GetAwsContainerService(ctx context.Context) []string {
	return []string{
		"ECS",
		"EKS",
	}
}
