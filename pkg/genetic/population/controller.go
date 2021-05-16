package population

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Generate generates the population for the genetic internship planner
func (p *populationController) Generate(w http.ResponseWriter, r *http.Request) {
	var d generatePopulationRequest
	if err := json.NewDecoder(r.Body).Decode(&d); err == nil {
		fmt.Fprint(w, "Hello, World!")
		return
	} else {
		log.Print(err)
		http.Error(w, "Internal error, please contact administrator.", http.StatusInternalServerError)
	}
}

type PopulationController interface {
	Generate(http.ResponseWriter, *http.Request)
}

type populationController struct {
	service PopulationService
}

func ProvidePopulationController(service PopulationService) PopulationController{
	return &populationController{
		service: service,
	}
}

type generatePopulationRequest struct {
}