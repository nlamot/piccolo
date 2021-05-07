package generatepopulation

import (
	"context"
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"

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
	if client == nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		log.Printf("generatepopulation.GeneratePopulation: Firestore wasn't initialized correctly.")
		return
	}
	var d struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		fmt.Fprint(w, "Hello, World!")
		return
	}
	if d.Name == "" {
		fmt.Fprint(w, "Hello, World!")
		return
	}
	fmt.Fprint(w, "Hello, ", html.EscapeString(d.Name), "!")
}
