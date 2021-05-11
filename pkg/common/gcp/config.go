package gcp

import "os"

type Configuration struct {
	ProjectID string
}

func GetConfiguration() *Configuration {
	return &Configuration{
		ProjectID: os.Getenv("GCP_PROJECT"),
	}
}