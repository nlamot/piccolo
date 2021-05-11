package population

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"

	"piccolo.com/planner/pkg/common/gcp"
)

type PopulationHandler interface {
	Generate(http.ResponseWriter, *http.Request)
}

type populationHandler struct {
	firestoreClient gcp.FirestoreClient
}

func ProvidePopulationHandler(firestoreClient gcp.FirestoreClient) PopulationHandler{
	return &populationHandler{
		firestoreClient: firestoreClient,
	}
}

// Generate generates the population for the genetic internship planner
func (p *populationHandler) Generate(w http.ResponseWriter, r *http.Request) {
	if p.firestoreClient == nil {
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