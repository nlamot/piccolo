package gcp_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"piccolo.com/planner/pkg/common/gcp"
)


func TestFirestoreClientShutsDownApplicationIfNotConnectedProperly(t *testing.T) {
	os.Setenv("GCP_PROJECT", "INVALID")
	_, err := gcp.ProvideFirestoreClient(gcp.GetConfiguration())
	assert.EqualError(t, err, "google: could not find default credentials. See https://developers.google.com/accounts/docs/application-default-credentials for more information.")
}