package gcp_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"piccolo.com/planner/pkg/common/gcp"
)

func TestConfigIsLoadedFromEnvironment(t *testing.T) {
	os.Setenv("GCP_PROJECT", "PROJECT_ID")
	config := gcp.GetConfiguration()
	assert.Equal(t, "PROJECT_ID", config.ProjectID)
}
