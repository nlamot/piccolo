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
var expectedGeneratePopulationRequest population.GeneratePopulationRequest
var response *httptest.ResponseRecorder
var serviceMock = new(PopulationServiceMock)
var controller = population.ProvidePopulationController(serviceMock)

func TestGeneratePopulationPassesInformationToService(t *testing.T) {
	givenAValidRequestToGeneratePopulations()
	givenServiceGeneratesThePopulationSuccesfully()

	whenTheControllerIsCalled()

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "Population generated.", response.Body.String())
	serviceMock.AssertCalled(t, "Generate", expectedGeneratePopulationRequest)
}

func TestGeneratePopulationReturnsInternalErrorWhenGenerateFails(t *testing.T) {
	givenAnInvalidRequestToGeneratePopulations()

	whenTheControllerIsCalled()

	assert.Equal(t, http.StatusInternalServerError, response.Code)
	assert.Equal(t, "Internal error, please contact administrator.\n", response.Body.String())
}

func TestGeneratePopulationThrowsErrorIfMethodNotPost(t *testing.T) {
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
	givenAnInvalidRequestToGeneratePopulations()

	whenTheControllerIsCalled()

	assert.Equal(t, http.StatusInternalServerError, response.Code)
	assert.Equal(t, "Internal error, please contact administrator.\n", response.Body.String())
}

func givenAValidRequestToGeneratePopulations() {
	request = httptest.NewRequest("POST", "/", strings.NewReader(`{ "size": 200 }`))
	request.Header.Add("Content-Type", "application/json")
}

func givenServiceGeneratesThePopulationSuccesfully() {
	expectedGeneratePopulationRequest = population.GeneratePopulationRequest{
		Size: 200,
	}
	serviceMock.On("Generate", expectedGeneratePopulationRequest).Return()
}

func givenAnInvalidRequestToGeneratePopulations() {
	request = httptest.NewRequest("POST", "/", strings.NewReader(`INVALID`))
	request.Header.Add("Content-Type", "application/json")
}

func givenARequestWithMethod(method string) {
	request = httptest.NewRequest(method, "/", strings.NewReader(""))
}

func whenTheControllerIsCalled() {
	response = httptest.NewRecorder()
	controller.Generate(response, request)
}
