package planner

import (
	"context"
	"log"
	"net/http"
	"os"
	"piccolo.com/planner/pkg/genetic/population"

	"cloud.google.com/go/firestore"
)

type Configuration struct {
	Project string
}

var configuration Configuration
var client *firestore.Client

// Initialize connection to firestore
func init() {
	configuration.Project = os.Getenv("GCP_PROJECT")

	var err error
	if configuration.Project != "" {
		client, err = firestore.NewClient(context.Background(), configuration.Project)
		if err != nil {
			log.Printf("firestore.NewClient: %v", err)
			return
		}
	}
}

// GeneratePopulation generates the population for the genetic internship planner
func GeneratePopulation(w http.ResponseWriter, r *http.Request) {
	population.Generate(w, r, client)
}
