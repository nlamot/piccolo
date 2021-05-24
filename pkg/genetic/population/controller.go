package population

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Create creates the population for the genetic internship planner
func (p *populationController) Create(w http.ResponseWriter, r *http.Request) {
	var d CreatePopulationRequest
	switch r.Method {
	case http.MethodPost:
		if err := json.NewDecoder(r.Body).Decode(&d); err == nil {
			p.service.Create(d)
			fmt.Fprint(w, "Population created.")
			return
		} else {
			log.Print(err)
			http.Error(w, "Internal error, please contact administrator.", http.StatusInternalServerError)
		}
	default:
		http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
	}

}

type PopulationController interface {
	Create(http.ResponseWriter, *http.Request)
}

type populationController struct {
	service PopulationService
}

func ProvidePopulationController(service PopulationService) PopulationController {
	return &populationController{
		service: service,
	}
}
