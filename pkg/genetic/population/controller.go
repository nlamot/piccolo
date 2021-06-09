package population

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Generate generates the population for the genetic internship planner
func (p *populationController) Generate(w http.ResponseWriter, r *http.Request) {
	var d GeneratePopulationRequest
	switch r.Method {
	case http.MethodPost:
		if err := json.NewDecoder(r.Body).Decode(&d); err == nil {
			p.service.Generate(d)
			fmt.Fprint(w, "Population generated.")
			return
		} else {
			log.Print(err)
			http.Error(w, "Internal error, please contact administrator.", http.StatusInternalServerError)
		}
	default:
		http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
	}

}
//go:generate go run github.com/vektra/mockery/cmd/mockery -name PopulationController -output ./mock/ -outpkg mock
type PopulationController interface {
	Generate(http.ResponseWriter, *http.Request)
}

type populationController struct {
	service PopulationService
}

func ProvidePopulationController(service PopulationService) PopulationController {
	return &populationController{
		service: service,
	}
}
