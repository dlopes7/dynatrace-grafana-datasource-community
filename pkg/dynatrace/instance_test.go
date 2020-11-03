package dynatrace

import (
	"context"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestQuery(t *testing.T) {

	c := APIClient{
		TenantURL: os.Getenv("DYNATRACE_TENANT_URL"),
		HttpClient: HttpClient{
			Client: http.Client{},
			Token:  os.Getenv("DYNATRACE_API_TOKEN"),
		},
	}

	ds := dynatraceDatasourceInstance{
		c,
	}
	ctx := context.Background()
	query := MetricQuery{
		MetricSelector: "builtin:host.cpu.idle",
		Resolution:     "",
		EntitySelector: "",
		TimeRange: backend.TimeRange{
			From: time.Now().Add(-30 * time.Minute),
			To:   time.Now(),
		},
	}
	frames, err := ds.query(ctx, query)
	if err != nil {
		t.Errorf("Error querying: %s", err.Error())
		t.FailNow()
	}
	t.Logf("Received %d frames", len(frames))

}
