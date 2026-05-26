// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package integrationdatabricksaccount

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IntegrationDatabricksAccountConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// A human-readable name for the account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/integration_databricks_account#name IntegrationDatabricksAccount#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The URL of your Databricks workspace (e.g., https://your-workspace.cloud.databricks.com).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/integration_databricks_account#workspace_url IntegrationDatabricksAccount#workspace_url}
	WorkspaceUrl *string `field:"required" json:"workspaceUrl" yaml:"workspaceUrl"`
	// auth_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/integration_databricks_account#auth_config IntegrationDatabricksAccount#auth_config}
	AuthConfig *IntegrationDatabricksAccountAuthConfig `field:"optional" json:"authConfig" yaml:"authConfig"`
	// Enable Cloud Cost Management to collect cost data from Databricks System Tables. Requires `system_tables_sql_warehouse_id`. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/integration_databricks_account#ccm_enabled IntegrationDatabricksAccount#ccm_enabled}
	CcmEnabled interface{} `field:"optional" json:"ccmEnabled" yaml:"ccmEnabled"`
	// Datadog API Key ID used for the Data Jobs Monitoring init script when managed by Datadog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/integration_databricks_account#dd_api_key_id IntegrationDatabricksAccount#dd_api_key_id}
	DdApiKeyId *string `field:"optional" json:"ddApiKeyId" yaml:"ddApiKeyId"`
	// Datadog API Key value (not ID) used for the Data Jobs Monitoring init script when managed by Datadog.
	//
	// This value is write-only; changes made outside of Terraform will not be drift-detected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/integration_databricks_account#dd_api_key_secret IntegrationDatabricksAccount#dd_api_key_secret}
	DdApiKeySecret *string `field:"optional" json:"ddApiKeySecret" yaml:"ddApiKeySecret"`
	// When enabled, Datadog installs and manages the Agent using a cluster policy and Unity Catalog Volume.
	//
	// Requires a Unity Catalog-enabled workspace with DBR 13.3 LTS+ and `uc_volume_path`. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/integration_databricks_account#djm_cluster_policy_enabled IntegrationDatabricksAccount#djm_cluster_policy_enabled}
	DjmClusterPolicyEnabled interface{} `field:"optional" json:"djmClusterPolicyEnabled" yaml:"djmClusterPolicyEnabled"`
	// Enable Data Jobs Monitoring for this workspace. Defaults to true. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/integration_databricks_account#djm_enabled IntegrationDatabricksAccount#djm_enabled}
	DjmEnabled interface{} `field:"optional" json:"djmEnabled" yaml:"djmEnabled"`
	// When enabled, Datadog installs and manages the Agent with a global init script in the workspace.
	//
	// Installation can take up to 15 minutes. Requires Workspace Admin permissions. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/integration_databricks_account#djm_global_init_script_enabled IntegrationDatabricksAccount#djm_global_init_script_enabled}
	DjmGlobalInitScriptEnabled interface{} `field:"optional" json:"djmGlobalInitScriptEnabled" yaml:"djmGlobalInitScriptEnabled"`
	// Cron schedule controlling how often Datadog crawls the Databricks warehouse for metadata.
	//
	// Defaults to hourly. Defaults to `"0 * * * *"`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/integration_databricks_account#do_crawlers_cron IntegrationDatabricksAccount#do_crawlers_cron}
	DoCrawlersCron *string `field:"optional" json:"doCrawlersCron" yaml:"doCrawlersCron"`
	// Enable Data Observability to collect data for viewing in Datadog Data Observability. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/integration_databricks_account#do_enabled IntegrationDatabricksAccount#do_enabled}
	DoEnabled interface{} `field:"optional" json:"doEnabled" yaml:"doEnabled"`
	// Name of the Databricks model serving endpoint to monitor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/integration_databricks_account#model_serving_endpoint_name IntegrationDatabricksAccount#model_serving_endpoint_name}
	ModelServingEndpointName *string `field:"optional" json:"modelServingEndpointName" yaml:"modelServingEndpointName"`
	// Retrieve health and usage metrics from Databricks model serving endpoints. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/integration_databricks_account#model_serving_metrics_enabled IntegrationDatabricksAccount#model_serving_metrics_enabled}
	ModelServingMetricsEnabled interface{} `field:"optional" json:"modelServingMetricsEnabled" yaml:"modelServingMetricsEnabled"`
	// private_action_runner_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/integration_databricks_account#private_action_runner_configuration IntegrationDatabricksAccount#private_action_runner_configuration}
	PrivateActionRunnerConfiguration *IntegrationDatabricksAccountPrivateActionRunnerConfiguration `field:"optional" json:"privateActionRunnerConfiguration" yaml:"privateActionRunnerConfiguration"`
	// Collect GPU metrics from Databricks clusters when using a Datadog-managed init script. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/integration_databricks_account#script_gpum_enabled IntegrationDatabricksAccount#script_gpum_enabled}
	ScriptGpumEnabled interface{} `field:"optional" json:"scriptGpumEnabled" yaml:"scriptGpumEnabled"`
	// Collect driver and worker logs from Databricks clusters when using a Datadog-managed init script. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/integration_databricks_account#script_logs_enabled IntegrationDatabricksAccount#script_logs_enabled}
	ScriptLogsEnabled interface{} `field:"optional" json:"scriptLogsEnabled" yaml:"scriptLogsEnabled"`
	// Serverless opt-in for Data Jobs Monitoring. Defaults to true. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/integration_databricks_account#serverless_jobs_enabled IntegrationDatabricksAccount#serverless_jobs_enabled}
	ServerlessJobsEnabled interface{} `field:"optional" json:"serverlessJobsEnabled" yaml:"serverlessJobsEnabled"`
	// SQL Warehouse ID for querying Databricks System Tables. Required for Cloud Cost Management.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/integration_databricks_account#system_tables_sql_warehouse_id IntegrationDatabricksAccount#system_tables_sql_warehouse_id}
	SystemTablesSqlWarehouseId *string `field:"optional" json:"systemTablesSqlWarehouseId" yaml:"systemTablesSqlWarehouseId"`
	// Enable table lineage tracking for Databricks tables. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/integration_databricks_account#table_lineage_enabled IntegrationDatabricksAccount#table_lineage_enabled}
	TableLineageEnabled interface{} `field:"optional" json:"tableLineageEnabled" yaml:"tableLineageEnabled"`
	// Unity Catalog volume path in `catalog.schema.volume` format where the Datadog init script will be stored (e.g. `main.default.datadog_volume`). Required when `djm_cluster_policy_enabled` is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/4.10.0/docs/resources/integration_databricks_account#uc_volume_path IntegrationDatabricksAccount#uc_volume_path}
	UcVolumePath *string `field:"optional" json:"ucVolumePath" yaml:"ucVolumePath"`
}

