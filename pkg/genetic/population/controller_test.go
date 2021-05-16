package population_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"piccolo.com/planner/pkg/genetic/population"
)

var request *http.Request
var response *httptest.ResponseRecorder
var service = new(PopulationServiceMock)
var controller = population.ProvidePopulationController(service)

func TestGeneratePopulation(t *testing.T) {
	givenAValidRequestToGeneratePopulations()

	whenTheControllerIsCalled()

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "Hello, World!", response.Body.String())
}

func TestGeneratePopulationReturnsInternalErrorWhenGenerateFails(t *testing.T) {
	givenAnInvalidRequestToGeneratePopulations()

	whenTheControllerIsCalled()

	assert.Equal(t, http.StatusInternalServerError, response.Code)
	assert.Equal(t, "Internal error, please contact administrator.\n", response.Body.String())

}

func givenAValidRequestToGeneratePopulations() {
	request = httptest.NewRequest("POST", "/", strings.NewReader(`{}`))
	request.Header.Add("Content-Type", "application/json")
}

func givenAnInvalidRequestToGeneratePopulations(){
	request = httptest.NewRequest("POST", "/", strings.NewReader(`INVALID`))
	request.Header.Add("Content-Type", "application/json")
}

func whenTheControllerIsCalled() {
	response = httptest.NewRecorder()
	controller.Generate(response, request)
}
