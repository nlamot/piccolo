package population

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var request *http.Request
var response *httptest.ResponseRecorder

func TestGeneratePopulation(t *testing.T) {
	givenAValidFirestoreClient()
	givenAValidRequestToGeneratePopulations()

	whenTheFunctionIsTriggered()

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "Hello, World!", response.Body.String())
}

func TestGeneratePopulationReturnsInternalServerErrorOnMissingFirestoreConnection(t *testing.T) {
	givenThereIsNoFirestoreConnection()
	givenAValidRequestToGeneratePopulations()

	whenTheFunctionIsTriggered()

	assert.Equal(t, http.StatusInternalServerError, response.Code)
}

func givenThereIsNoFirestoreConnection() {
	client = nil
}

func givenAValidRequestToGeneratePopulations() {
	request = httptest.NewRequest("GET", "/", strings.NewReader(`{"name": ""}`))
	request.Header.Add("Content-Type", "application/json")
}

func givenAValidFirestoreClient() {
	client = nil
}

func whenTheFunctionIsTriggered() {
	response = httptest.NewRecorder()
	GeneratePopulation(response, request)
}
