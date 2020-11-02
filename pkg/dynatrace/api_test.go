package dynatrace

import (
	"context"
	"os"
	"testing"
)

func TestGetClusterVersion(t *testing.T) {

	api := APIClient{
		TenantURL: os.Getenv("DYNATRACE_TENANT_URL"),
		HttpClient: HttpClient{
			Token: os.Getenv("DYNATRACE_API_TOKEN"),
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

func TestQueryMetrics(t *testing.T) {

	api := APIClient{
		TenantURL: os.Getenv("DYNATRACE_TENANT_URL"),
		HttpClient: HttpClient{
			Token: os.Getenv("DYNATRACE_API_TOKEN"),
		},
	}

	ctx := context.Background()
	metrics, err := api.QueryMetrics(ctx)
	if err != nil {
		t.Errorf("error querying metrics %s", err.Error())
		t.FailNow()
	}
	t.Logf("Got %d pages of metrics", len(metrics))

}
