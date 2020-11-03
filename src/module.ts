import { DataSourcePlugin } from '@grafana/data';
import { DataSource } from './DataSource';
import { ConfigEditor } from './ConfigEditor';
import { QueryEditor } from './QueryEditor';
import { DynatraceQuery, DynatraceOptions } from './types';

export const plugin = new DataSourcePlugin<DataSource, DynatraceQuery, DynatraceOptions>(DataSource)
  .setConfigEditor(ConfigEditor)
  .setQueryEditor(QueryEditor);
