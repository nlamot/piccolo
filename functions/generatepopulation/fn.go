package generatepopulation

import (
	"context"
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/tkanos/gonfig"
)

type Configuration struct {
    Project string `env:"GCP_PROJECT"`
}

var configuration Configuration
var client *firestore.Client

// Initialize connection to firestore
func init() {
	err := gonfig.GetConf("config.json", &configuration)
	if err != nil {
		log.Fatalf("Configuration Load: %v", err)
	}

	client, err = firestore.NewClient(context.Background(), configuration.Project)
	if err != nil {
		log.Printf("firestore.NewClient: %v", err)
		return
	}
}

// GeneratePopulation generates the population for the genetic internship planner
func GeneratePopulation(w http.ResponseWriter, r *http.Request) {
	if (client == nil) {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		log.Printf("firestore.NewClient: Firestore wasn't initialized correctly.")
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
