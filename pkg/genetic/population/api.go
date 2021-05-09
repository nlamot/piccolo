package population

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	
	"cloud.google.com/go/firestore"
)

// GeneratePopulation generates the population for the genetic internship planner
func GeneratePopulation(w http.ResponseWriter, r *http.Request, client *firestore.Client) {
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