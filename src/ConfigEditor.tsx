import React, { ChangeEvent, PureComponent } from 'react';
import { LegacyForms } from '@grafana/ui';
import { DataSourcePluginOptionsEditorProps } from '@grafana/data';
import { DynatraceOptions, DynatraceTokenData } from './types';

const { SecretFormField, FormField } = LegacyForms;

interface Props extends DataSourcePluginOptionsEditorProps<DynatraceOptions> {}

interface State {}

export class ConfigEditor extends PureComponent<Props, State> {
  onTenantURLChange = (event: ChangeEvent<HTMLInputElement>) => {
    const { onOptionsChange, options } = this.props;
    const jsonData = {
      ...options.jsonData,
      tenantURL: event.target.value,
    };
    onOptionsChange({ ...options, jsonData });
  };

  // Secure field (only sent to the backend)
  onChangeAPIToken = (event: ChangeEvent<HTMLInputElement>) => {
    const { onOptionsChange, options } = this.props;
    onOptionsChange({
      ...options,
      secureJsonData: {
        apiToken: event.target.value,
      },
    });
  };

  onResetAPIToken = () => {
    const { onOptionsChange, options } = this.props;
    onOptionsChange({
      ...options,
      secureJsonFields: {
        ...options.secureJsonFields,
        apiToken: false,
      },
      secureJsonData: {
        ...options.secureJsonData,
        apiToken: '',
      },
    });
  };

  render() {
    const { options } = this.props;
    const { jsonData, secureJsonFields } = options;
    const secureJsonData = (options.secureJsonData || {}) as DynatraceTokenData;

    return (
      <div className="gf-form-group">
        <div className="gf-form">
          <FormField
            label="Tenant URL"
            labelWidth={10}
            inputWidth={20}
            onChange={this.onTenantURLChange}
            value={jsonData.tenantURL || ''}
            placeholder="https://<id>.live.dynatrace.com"
            tooltip="For managed use https://<my_domain>/e/<my_environment_id>"
          />
        </div>

        <div className="gf-form-inline">
          <div className="gf-form">
            <SecretFormField
              isConfigured={(secureJsonFields && secureJsonFields.apiToken) as boolean}
              value={secureJsonData.apiToken || ''}
              label="API Token"
              placeholder=""
              labelWidth={10}
              inputWidth={20}
              onReset={this.onResetAPIToken}
              onChange={this.onChangeAPIToken}
            />
          </div>
        </div>
      </div>
    );
  }
}
