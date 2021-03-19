package ec2

import "context"

//todo more instance types will be added in future
//context is for future tracing
func GetAwsEC2InstanceTypes(ctx context.Context) []string {
	return []string{
		"GENERAL_PURPOSE",
		"COMPUTE_OPTIMIZED",
		"MEMORY_OPTIMIZED",
		"ACCELERATED_COMPUTING",
	}
}
