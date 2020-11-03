import { DataSourceInstanceSettings } from '@grafana/data';
import { DataSourceWithBackend } from '@grafana/runtime';
import { DynatraceOptions, DynatraceQuery } from './types';

export class DataSource extends DataSourceWithBackend<DynatraceQuery, DynatraceOptions> {
  constructor(instanceSettings: DataSourceInstanceSettings<DynatraceOptions>) {
    super(instanceSettings);
  }
}
