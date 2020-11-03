package dynatrace

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/datasource"
	"github.com/grafana/grafana-plugin-sdk-go/backend/instancemgmt"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

type dynatraceDatasource struct {
	im instancemgmt.InstanceManager
}

func NewDatasource() datasource.ServeOpts {
	// This is called every time a new datasource instance is created or updated
	im := datasource.NewInstanceManager(newDatasourceInstance)
	ds := &dynatraceDatasource{
		im: im,
	}

	return datasource.ServeOpts{
		QueryDataHandler:   ds,
		CheckHealthHandler: ds,
	}
}

func (ds *dynatraceDatasource) QueryData(ctx context.Context, req *backend.QueryDataRequest) (*backend.QueryDataResponse, error) {
	log.DefaultLogger.Info("QueryData", "request", req)

	dsInstance, err := ds.getDSInstance(req.PluginContext)
	if err != nil {
		return nil, err
	}

	response := backend.NewQueryDataResponse()

	for _, q := range req.Queries {
		res := ds.query(ctx, q, *dsInstance)
		response.Responses[q.RefID] = res
	}

	return response, nil
}

func (ds *dynatraceDatasource) CheckHealth(ctx context.Context, req *backend.CheckHealthRequest) (*backend.CheckHealthResult, error) {

	res := &backend.CheckHealthResult{}

	dsInstance, err := ds.getDSInstance(req.PluginContext)
	if err != nil {
		res.Status = backend.HealthStatusError
		res.Message = "Error getting datasource instance"
		log.DefaultLogger.Error("Error getting datasource instance", "err", err)
		return res, nil
	}

	testResult, err := dsInstance.testConnection(ctx)
	if err != nil {
		res.Status = backend.HealthStatusError
		res.Message = err.Error()
		log.DefaultLogger.Error("Error connecting to Dynatrace", "err", err)
		return res, nil
	}

	res.Status = backend.HealthStatusOk
	res.Message = testResult
	return res, nil
}

func (ds *dynatraceDatasource) getDSInstance(pluginContext backend.PluginContext) (*dynatraceDatasourceInstance, error) {
	instance, err := ds.im.Get(pluginContext)
	if err != nil {
		return nil, err
	}
	return instance.(*dynatraceDatasourceInstance), nil
}

func (ds *dynatraceDatasource) query(ctx context.Context, q backend.DataQuery, dsInstance dynatraceDatasourceInstance) backend.DataResponse {

	res := backend.DataResponse{}
	metricQuery, err := ReadQuery(q)
	if err != nil {
		res.Error = err
	} else {
		log.DefaultLogger.Info("Processing query", "query", metricQuery, "timerange", metricQuery.TimeRange)
		frames, err := dsInstance.query(ctx, metricQuery)
		if err != nil {
			res.Error = err
		} else {
			res.Frames = frames
		}
	}

	return res

}

func ReadQuery(query backend.DataQuery) (MetricQuery, error) {
	model := MetricQuery{}
	if err := json.Unmarshal(query.JSON, &model); err != nil {
		return model, fmt.Errorf("could not read query: %w", err)
	}

	model.TimeRange = query.TimeRange
	return model, nil
}
