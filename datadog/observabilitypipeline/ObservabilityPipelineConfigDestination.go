// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline


type ObservabilityPipelineConfigDestination struct {
	// The unique identifier for this destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#id ObservabilityPipeline#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// A list of component IDs whose output is used as the `input` for this component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#inputs ObservabilityPipeline#inputs}
	Inputs *[]*string `field:"required" json:"inputs" yaml:"inputs"`
	// amazon_opensearch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#amazon_opensearch ObservabilityPipeline#amazon_opensearch}
	AmazonOpensearch interface{} `field:"optional" json:"amazonOpensearch" yaml:"amazonOpensearch"`
	// amazon_s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#amazon_s3 ObservabilityPipeline#amazon_s3}
	AmazonS3 interface{} `field:"optional" json:"amazonS3" yaml:"amazonS3"`
	// amazon_security_lake block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#amazon_security_lake ObservabilityPipeline#amazon_security_lake}
	AmazonSecurityLake interface{} `field:"optional" json:"amazonSecurityLake" yaml:"amazonSecurityLake"`
	// azure_storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#azure_storage ObservabilityPipeline#azure_storage}
	AzureStorage interface{} `field:"optional" json:"azureStorage" yaml:"azureStorage"`
	// cloud_prem block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#cloud_prem ObservabilityPipeline#cloud_prem}
	CloudPrem interface{} `field:"optional" json:"cloudPrem" yaml:"cloudPrem"`
	// crowdstrike_next_gen_siem block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#crowdstrike_next_gen_siem ObservabilityPipeline#crowdstrike_next_gen_siem}
	CrowdstrikeNextGenSiem interface{} `field:"optional" json:"crowdstrikeNextGenSiem" yaml:"crowdstrikeNextGenSiem"`
	// datadog_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#datadog_logs ObservabilityPipeline#datadog_logs}
	DatadogLogs interface{} `field:"optional" json:"datadogLogs" yaml:"datadogLogs"`
	// datadog_metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#datadog_metrics ObservabilityPipeline#datadog_metrics}
	DatadogMetrics interface{} `field:"optional" json:"datadogMetrics" yaml:"datadogMetrics"`
	// elasticsearch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#elasticsearch ObservabilityPipeline#elasticsearch}
	Elasticsearch interface{} `field:"optional" json:"elasticsearch" yaml:"elasticsearch"`
	// google_cloud_storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#google_cloud_storage ObservabilityPipeline#google_cloud_storage}
	GoogleCloudStorage interface{} `field:"optional" json:"googleCloudStorage" yaml:"googleCloudStorage"`
	// google_pubsub block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#google_pubsub ObservabilityPipeline#google_pubsub}
	GooglePubsub interface{} `field:"optional" json:"googlePubsub" yaml:"googlePubsub"`
	// google_secops block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#google_secops ObservabilityPipeline#google_secops}
	GoogleSecops interface{} `field:"optional" json:"googleSecops" yaml:"googleSecops"`
	// http_client block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#http_client ObservabilityPipeline#http_client}
	HttpClient interface{} `field:"optional" json:"httpClient" yaml:"httpClient"`
	// kafka block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#kafka ObservabilityPipeline#kafka}
	Kafka interface{} `field:"optional" json:"kafka" yaml:"kafka"`
	// microsoft_sentinel block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#microsoft_sentinel ObservabilityPipeline#microsoft_sentinel}
	MicrosoftSentinel interface{} `field:"optional" json:"microsoftSentinel" yaml:"microsoftSentinel"`
	// new_relic block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#new_relic ObservabilityPipeline#new_relic}
	NewRelic interface{} `field:"optional" json:"newRelic" yaml:"newRelic"`
	// opensearch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#opensearch ObservabilityPipeline#opensearch}
	Opensearch interface{} `field:"optional" json:"opensearch" yaml:"opensearch"`
	// rsyslog block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#rsyslog ObservabilityPipeline#rsyslog}
	Rsyslog interface{} `field:"optional" json:"rsyslog" yaml:"rsyslog"`
	// sentinel_one block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#sentinel_one ObservabilityPipeline#sentinel_one}
	SentinelOne interface{} `field:"optional" json:"sentinelOne" yaml:"sentinelOne"`
	// socket block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#socket ObservabilityPipeline#socket}
	Socket interface{} `field:"optional" json:"socket" yaml:"socket"`
	// splunk_hec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#splunk_hec ObservabilityPipeline#splunk_hec}
	SplunkHec interface{} `field:"optional" json:"splunkHec" yaml:"splunkHec"`
	// sumo_logic block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#sumo_logic ObservabilityPipeline#sumo_logic}
	SumoLogic interface{} `field:"optional" json:"sumoLogic" yaml:"sumoLogic"`
	// syslog_ng block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/datadog/datadog/3.89.0/docs/resources/observability_pipeline#syslog_ng ObservabilityPipeline#syslog_ng}
	SyslogNg interface{} `field:"optional" json:"syslogNg" yaml:"syslogNg"`
}

