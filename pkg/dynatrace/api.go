package dynatrace

import (
	"context"
	"fmt"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"net/url"
)

type APIClient struct {
	TenantURL  string
	HttpClient HttpClient
}

func (api *APIClient) GetClusterVersion(ctx context.Context) (*ClusterVersion, error) {
	u := fmt.Sprintf("%s/api/v1/config/clusterversion", api.TenantURL)

	var c *ClusterVersion
	err := api.HttpClient.makeRequest(ctx, "GET", u, nil, &c, nil)
	if err != nil {
		return nil, err
	}
	return c, nil

}

func (api *APIClient) QueryMetrics(ctx context.Context, metricSelector string, resolution string, from int64, to int64, entitySelector string) ([]MetricSeriesCollection, error) {
	u := fmt.Sprintf("%s/api/v2/metrics/query", api.TenantURL)
	q := url.Values{}
	q.Add("metricSelector", metricSelector)
	q.Add("from", fmt.Sprint(from))
	q.Add("to", fmt.Sprint(to))
	if resolution != "" {
		q.Add("resolution", resolution)
	}
	if entitySelector != "" {
		q.Add("entitySelector", entitySelector)
	}

	log.DefaultLogger.Info("QueryMetrics", "url", u, "QueryString", q)

	var m *MetricData
	var metrics []MetricSeriesCollection

	err := api.HttpClient.makeRequest(ctx, "GET", u, q, &m, nil)
	if err != nil {
		return nil, err
	}

	// Add the current metrics to the final list
	metrics = append(metrics, m.Result...)

	for m.NextPageKey != "" {
		// Get all pages of metrics
		u := fmt.Sprintf("%s/api/v2/metrics/query?nextPageKey=%s", api.TenantURL, m.NextPageKey)
		m = &MetricData{}
		err := api.HttpClient.makeRequest(ctx, "GET", u, nil, &m, nil)
		if err != nil {
			return nil, err
		}
		metrics = append(metrics, m.Result...)
	}

	return metrics, nil

}
