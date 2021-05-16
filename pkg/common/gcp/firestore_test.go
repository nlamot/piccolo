package gcp_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"piccolo.com/planner/pkg/common/gcp"
)


func TestFirestoreClientShutsDownApplicationIfNotConnectedProperly(t *testing.T) {
	_, err := gcp.ProvideFirestoreClient(&gcp.Configuration{ProjectID: "INVALID"})
	assert.EqualError(t, err, "google: could not find default credentials. See https://developers.google.com/accounts/docs/application-default-credentials for more information.")
}