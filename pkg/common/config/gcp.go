package config

import "os"

type GCPConfiguration struct {
	ProjectID string
}

func GetGCPConfiguration() *GCPConfiguration {
	return &GCPConfiguration{
		ProjectID: os.Getenv("GCP_PROJECT"),
	}
}