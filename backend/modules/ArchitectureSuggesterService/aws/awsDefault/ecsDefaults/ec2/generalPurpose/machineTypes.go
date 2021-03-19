package generalPurpose

import "context"

//context is for future tracing
func GetAwsGeneralPurposeMachineTypes(ctx context.Context) []string {
	return []string{
		"Mac",
		"T4g",
		"T3",
		"T3a",
		"T2",
		"M6g",
		"M5",
		"M5a",
		"M5n",
		"M5zn",
		"M4",
		"A1",
	}
}

func GetAwsComputeOptimizedMachineTypes(ctx context.Context) []string {
	return []string{
		"C6g",
		"C6gn",
		"C5",
		"C5a",
		"C5n",
		"C4",
	}
}

func GetAwsMemoryOptimizedMachineTypes(ctx context.Context) []string {
	return []string{
		"R6g",
		"R5",
		"R5a",
		"R5b",
		"R5n",
		"R4",
		"X2gd",
		"X1e",
		"X1",
		"High Memory",
		"z1d",
	}
}

func GetAwsAcceleratedComputingMachineTypes(ctx context.Context) []string {
	return []string{
		"P4",
		"P3",
		"P2",
		"Inf1",
		"G4dn",
		"G4ad",
		"G3",
		"F1",
	}
}
