package dynatrace

import (
	"context"
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

	response := backend.NewQueryDataResponse()

	//for _, q := range req.Queries {
	//	res := td.query(ctx, q)
	//	response.Responses[q.RefID] = res
	//}

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
