package dynatrace

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/instancemgmt"
	"github.com/grafana/grafana-plugin-sdk-go/data"
	"net/http"
	"strings"
	"time"
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

func (i *dynatraceDatasourceInstance) query(ctx context.Context, q MetricQuery) (data.Frames, error) {
	metrics, err := i.APIClient.QueryMetrics(ctx, q.MetricSelector, q.Resolution, q.TimeRange.From.UnixNano()/1000000, q.TimeRange.To.UnixNano()/1000000, q.EntitySelector)
	if err != nil {
		return nil, err
	}

	frames := []*data.Frame{CollectionToFrames(metrics)}
	return frames, nil

}

// CollectionToFrames converts Dynatrace metrics to Grafana Frames
func CollectionToFrames(collections []MetricSeriesCollection) *data.Frame {

	frame := data.NewFrame("Metrics")

	// Each one of these is a "metric result" for a single time series
	for _, metricCollection := range collections {

		// Each one of these has results for a set of dimensions
		for _, metricData := range metricCollection.Data {
			timeField := data.NewFieldFromFieldType(data.FieldTypeTime, 0)
			timeField.Name = "time"
			valueField := data.NewFieldFromFieldType(data.FieldTypeNullableFloat64, 0)
			valueField.Name = metricCollection.MetricID
			valueField.Labels = make(map[string]string)

			for i, dimension := range metricData.Dimensions {
				valueField.Labels[fmt.Sprintf("Dimension %d", i)] = dimension
			}

			for i := range metricData.Timestamps {
				timeField.Append(time.Unix(0, metricData.Timestamps[i]*int64(time.Millisecond)))
				valueField.Append(metricData.Values[i])

			}
			frame.Fields = append(frame.Fields, timeField, valueField)
		}

	}

	wideFrame, err := data.LongToWide(frame, &data.FillMissing{Mode: data.FillModeNull})
	if err == nil {
		return wideFrame
	}
	return frame

}
