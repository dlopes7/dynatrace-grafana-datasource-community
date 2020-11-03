import defaults from 'lodash/defaults';

import React, { ChangeEvent, PureComponent } from 'react';
import { LegacyForms } from '@grafana/ui';
import { QueryEditorProps } from '@grafana/data';
import { DataSource } from './DataSource';
import { defaultQuery, DynatraceOptions, DynatraceQuery } from './types';

const { FormField } = LegacyForms;

type Props = QueryEditorProps<DataSource, DynatraceQuery, DynatraceOptions>;

export class QueryEditor extends PureComponent<Props> {
  onQueryTextChange = (event: ChangeEvent<HTMLInputElement>) => {
    const { onChange, query } = this.props;
    onChange({ ...query, metricSelector: event.target.value });
  };

  onConstantChange = (event: ChangeEvent<HTMLInputElement>) => {
    const { onChange, query, onRunQuery } = this.props;
    onChange({ ...query});
    // executes the query
    onRunQuery();
  };

  render() {
    const query = defaults(this.props.query, defaultQuery);
    const { metricSelector} = query;

    return (
      <div className="gf-form">
        <FormField
          labelWidth={8}
          value={metricSelector || ''}
          onChange={this.onQueryTextChange}
          label="Metric selector"
        />
      </div>
    );
  }
}
