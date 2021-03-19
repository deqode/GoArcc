package ecsDefaults

import "context"

//todo more launch types will be added in future
//context is for future tracing
func GetAwsECSLaunchTypes(ctx context.Context) []string {
	return []string{
		"EC2",
		"FRAGRATE",
	}
}
