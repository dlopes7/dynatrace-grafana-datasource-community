package main

import (
	"github.com/dlopes7/dynatrace-community-grafana-datasource/pkg/dynatrace"
	"github.com/grafana/grafana-plugin-sdk-go/backend/datasource"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"os"
)

func main() {
	err := datasource.Serve(dynatrace.NewDatasource())

	if err != nil {
		log.DefaultLogger.Error(err.Error())
		os.Exit(1)
	}
}
