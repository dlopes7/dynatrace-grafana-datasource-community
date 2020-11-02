package dynatrace

import (
	"context"
	"fmt"
)

type APIClient struct {
	TenantURL  string
	HttpClient HttpClient
}

func (api *APIClient) GetClusterVersion(ctx context.Context) (*ClusterVersion, error) {
	u := fmt.Sprintf("%s/api/v1/config/clusterversion", api.TenantURL)

	var c *ClusterVersion
	err := api.HttpClient.makeRequest(ctx, "GET", u, &c, nil)
	if err != nil {
		return nil, err
	}
	return c, nil

}

func (api *APIClient) QueryMetrics(ctx context.Context) ([]MetricSeriesCollection, error) {
	u := fmt.Sprintf("%s/api/v2/metrics/query?metricSelector=builtin:tech.generic.processCount&from=now-6M&to=now", api.TenantURL)
	var m *MetricData
	var metrics []MetricSeriesCollection

	err := api.HttpClient.makeRequest(ctx, "GET", u, &m, nil)
	if err != nil {
		return nil, err
	}

	// Add the current metrics to the final list
	metrics = append(metrics, m.Result...)

	for m.NextPageKey != "" {
		// Get all pages of metrics
		u := fmt.Sprintf("%s/api/v2/metrics/query?nextPageKey=%s", api.TenantURL, m.NextPageKey)
		m = &MetricData{}
		err := api.HttpClient.makeRequest(ctx, "GET", u, &m, nil)
		if err != nil {
			return nil, err
		}
		metrics = append(metrics, m.Result...)
	}

	return metrics, nil

}
