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
var expectedCreatePopulationRequest population.CreatePopulationRequest
var response *httptest.ResponseRecorder
var serviceMock = new(PopulationServiceMock)
var controller = population.ProvidePopulationController(serviceMock)

func TestCreatePopulationPassesInformationToService(t *testing.T) {
	givenAValidRequestToCreatePopulations()
	givenServiceCreatesThePopulationSuccesfully()

	whenTheControllerIsCalled()

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "Population created.", response.Body.String())
	serviceMock.AssertCalled(t, "Create", expectedCreatePopulationRequest)
}

func TestCreatePopulationReturnsInternalErrorWhenCreateFails(t *testing.T) {
	givenAnInvalidRequestToCreatePopulations()

	whenTheControllerIsCalled()

	assert.Equal(t, http.StatusInternalServerError, response.Code)
	assert.Equal(t, "Internal error, please contact administrator.\n", response.Body.String())
}

func TestCreatePopulationThrowsErrorIfMethodNotPost(t *testing.T) {
	var invalidMethods = []string{
		http.MethodConnect,
		http.MethodDelete,
		http.MethodHead,
		http.MethodOptions,
		http.MethodPatch,
		http.MethodGet,
		http.MethodPut,
		http.MethodTrace,
	}
	for _, method := range invalidMethods {
		givenARequestWithMethod(method)
		whenTheControllerIsCalled()
		assert.Equal(t, http.StatusMethodNotAllowed, response.Code)
		assert.Equal(t, "Method not allowed.\n", response.Body.String())
	}
	givenAnInvalidRequestToCreatePopulations()

	whenTheControllerIsCalled()

	assert.Equal(t, http.StatusInternalServerError, response.Code)
	assert.Equal(t, "Internal error, please contact administrator.\n", response.Body.String())
}

func givenAValidRequestToCreatePopulations() {
	request = httptest.NewRequest("POST", "/", strings.NewReader(`{ "size": 200 }`))
	request.Header.Add("Content-Type", "application/json")
}

func givenServiceCreatesThePopulationSuccesfully() {
	expectedCreatePopulationRequest = population.CreatePopulationRequest{
		Size: 200,
	}
	serviceMock.On("Create", expectedCreatePopulationRequest).Return()
}

func givenAnInvalidRequestToCreatePopulations() {
	request = httptest.NewRequest("POST", "/", strings.NewReader(`INVALID`))
	request.Header.Add("Content-Type", "application/json")
}

func givenARequestWithMethod(method string) {
	request = httptest.NewRequest(method, "/", strings.NewReader(""))
}

func whenTheControllerIsCalled() {
	response = httptest.NewRecorder()
	controller.Create(response, request)
}
