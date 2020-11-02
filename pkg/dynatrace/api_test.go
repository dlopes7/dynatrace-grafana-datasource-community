package dynatrace

import (
	"context"
	"os"
	"testing"
)

func TestGetClusterVersion(t *testing.T) {

	api := APIClient{
		tenantURL: os.Getenv("DYNATRACE_TENANT_URL"),
		httpClient: httpClient{
			token: os.Getenv("DYNATRACE_API_TOKEN"),
		},
	}

	ctx := context.Background()
	clusterVersion, err := api.GetClusterVersion(ctx)
	if err != nil {
		t.Errorf("error getting cluster version %s", err.Error())
		t.FailNow()
	}
	t.Logf("Got cluster version: %s", clusterVersion.Version)

}
