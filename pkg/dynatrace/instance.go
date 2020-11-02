package dynatrace

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/instancemgmt"
	"net/http"
	"strings"
)

type dynatraceDatasourceInstance struct {
	APIClient APIClient
}

func newDatasourceInstance(settings backend.DataSourceInstanceSettings) (instancemgmt.Instance, error) {

	dtSettings := &Settings{}

	err := json.Unmarshal(settings.JSONData, &dtSettings)
	if err != nil {
		return nil, err
	}

	c := APIClient{
		TenantURL: strings.TrimRight(dtSettings.TenantURL, "/"),
		HttpClient: HttpClient{
			Client: http.Client{},
			Token:  settings.DecryptedSecureJSONData["apiToken"],
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
