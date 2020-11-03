package dynatrace

import "github.com/grafana/grafana-plugin-sdk-go/backend"

type Settings struct {
	TenantURL string `json:"TenantURL"`
}

type ClusterVersion struct {
	Version string `json:"version"`
}

type MetricData struct {
	NextPageKey string                   `json:"nextPageKey"`
	TotalCount  int64                    `json:"totalCount"`
	Warnings    []string                 `json:"warnings"`
	Result      []MetricSeriesCollection `json:"result"`
}

type MetricSeriesCollection struct {
	MetricID string         `json:"metricId"`
	Data     []MetricSeries `json:"data"`
}

type MetricSeries struct {
	Timestamps []int64    `json:"timestamps"`
	Dimensions []string   `json:"dimensions"`
	Values     []*float64 `json:"values"`
}

type MetricQuery struct {
	MetricSelector string `json:"metricSelector"`
	Resolution     string `json:"resolution"`
	EntitySelector string `json:"entitySelector"`

	TimeRange backend.TimeRange `json:"-"`
}
