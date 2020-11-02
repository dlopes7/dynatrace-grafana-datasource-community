package dynatrace

import (
	"context"
	"fmt"
)

type APIClient struct {
	tenantURL  string
	httpClient httpClient
}

func (api *APIClient) GetClusterVersion(ctx context.Context) (*ClusterVersion, error) {
	u := fmt.Sprintf("%s/api/v1/config/clusterversion", api.tenantURL)

	var c *ClusterVersion
	err := api.httpClient.makeRequest(ctx, "GET", u, &c, nil)
	if err != nil {
		return nil, err
	}
	return c, nil

}
