package dynatrace

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/instancemgmt"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"net/http"
)

type dynatraceDatasourceInstance struct {
	APIClient APIClient
}

func newDatasourceInstance(settings backend.DataSourceInstanceSettings) (instancemgmt.Instance, error) {

	log.DefaultLogger.Info("Got settings from Grafana", "settings", settings)

	dtSettings := &Settings{}

	err := json.Unmarshal(settings.JSONData, &dtSettings)
	if err != nil {
		return nil, err
	}

	c := APIClient{
		tenantURL: dtSettings.TenantURL,
		httpClient: httpClient{
			client: http.Client{},
			token:  settings.DecryptedSecureJSONData["apiToken"],
		},
	}
	return &dynatraceDatasourceInstance{
		APIClient: c,
	}, nil
}

func (i *dynatraceDatasourceInstance) testConnection(ctx context.Context) (string, error) {
	clusterVersion, err := i.APIClient.GetClusterVersion(ctx)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Connection OK! Dynatrace Version: %s", clusterVersion.Version), nil
}
