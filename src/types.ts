import { DataQuery, DataSourceJsonData } from '@grafana/data';

export interface DynatraceQuery extends DataQuery {
  metricSelector: string;
}

export const defaultQuery: Partial<DynatraceQuery> = {
  metricSelector: ""
};

/**
 * These are options configured for each DataSource instance
 */
export interface DynatraceOptions extends DataSourceJsonData {
  tenantURL?: string;
}

/**
 * Value that is used in the backend, but never sent over HTTP to the frontend
 */
export interface DynatraceTokenData {
  apiToken?: string;
}
