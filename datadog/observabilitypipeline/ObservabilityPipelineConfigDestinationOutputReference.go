// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v13/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v13/observabilitypipeline/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ObservabilityPipelineConfigDestinationOutputReference interface {
	cdktn.ComplexObject
	AmazonOpensearch() ObservabilityPipelineConfigDestinationAmazonOpensearchList
	AmazonOpensearchInput() interface{}
	AmazonS3() ObservabilityPipelineConfigDestinationAmazonS3List
	AmazonS3Input() interface{}
	AmazonSecurityLake() ObservabilityPipelineConfigDestinationAmazonSecurityLakeList
	AmazonSecurityLakeInput() interface{}
	AzureStorage() ObservabilityPipelineConfigDestinationAzureStorageList
	AzureStorageInput() interface{}
	CloudPrem() ObservabilityPipelineConfigDestinationCloudPremList
	CloudPremInput() interface{}
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CrowdstrikeNextGenSiem() ObservabilityPipelineConfigDestinationCrowdstrikeNextGenSiemList
	CrowdstrikeNextGenSiemInput() interface{}
	DatadogLogs() ObservabilityPipelineConfigDestinationDatadogLogsList
	DatadogLogsInput() interface{}
	DatadogMetrics() ObservabilityPipelineConfigDestinationDatadogMetricsList
	DatadogMetricsInput() interface{}
	Elasticsearch() ObservabilityPipelineConfigDestinationElasticsearchList
	ElasticsearchInput() interface{}
	// Experimental.
	Fqn() *string
	GoogleCloudStorage() ObservabilityPipelineConfigDestinationGoogleCloudStorageList
	GoogleCloudStorageInput() interface{}
	GooglePubsub() ObservabilityPipelineConfigDestinationGooglePubsubList
	GooglePubsubInput() interface{}
	GoogleSecops() ObservabilityPipelineConfigDestinationGoogleSecopsList
	GoogleSecopsInput() interface{}
	HttpClient() ObservabilityPipelineConfigDestinationHttpClientList
	HttpClientInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	Inputs() *[]*string
	SetInputs(val *[]*string)
	InputsInput() *[]*string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Kafka() ObservabilityPipelineConfigDestinationKafkaList
	KafkaInput() interface{}
	MicrosoftSentinel() ObservabilityPipelineConfigDestinationMicrosoftSentinelList
	MicrosoftSentinelInput() interface{}
	NewRelic() ObservabilityPipelineConfigDestinationNewRelicList
	NewRelicInput() interface{}
	Opensearch() ObservabilityPipelineConfigDestinationOpensearchList
	OpensearchInput() interface{}
	Rsyslog() ObservabilityPipelineConfigDestinationRsyslogList
	RsyslogInput() interface{}
	SentinelOne() ObservabilityPipelineConfigDestinationSentinelOneList
	SentinelOneInput() interface{}
	Socket() ObservabilityPipelineConfigDestinationSocketList
	SocketInput() interface{}
	SplunkHec() ObservabilityPipelineConfigDestinationSplunkHecList
	SplunkHecInput() interface{}
	SumoLogic() ObservabilityPipelineConfigDestinationSumoLogicList
	SumoLogicInput() interface{}
	SyslogNg() ObservabilityPipelineConfigDestinationSyslogNgList
	SyslogNgInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	PutAmazonOpensearch(value interface{})
	PutAmazonS3(value interface{})
	PutAmazonSecurityLake(value interface{})
	PutAzureStorage(value interface{})
	PutCloudPrem(value interface{})
	PutCrowdstrikeNextGenSiem(value interface{})
	PutDatadogLogs(value interface{})
	PutDatadogMetrics(value interface{})
	PutElasticsearch(value interface{})
	PutGoogleCloudStorage(value interface{})
	PutGooglePubsub(value interface{})
	PutGoogleSecops(value interface{})
	PutHttpClient(value interface{})
	PutKafka(value interface{})
	PutMicrosoftSentinel(value interface{})
	PutNewRelic(value interface{})
	PutOpensearch(value interface{})
	PutRsyslog(value interface{})
	PutSentinelOne(value interface{})
	PutSocket(value interface{})
	PutSplunkHec(value interface{})
	PutSumoLogic(value interface{})
	PutSyslogNg(value interface{})
	ResetAmazonOpensearch()
	ResetAmazonS3()
	ResetAmazonSecurityLake()
	ResetAzureStorage()
	ResetCloudPrem()
	ResetCrowdstrikeNextGenSiem()
	ResetDatadogLogs()
	ResetDatadogMetrics()
	ResetElasticsearch()
	ResetGoogleCloudStorage()
	ResetGooglePubsub()
	ResetGoogleSecops()
	ResetHttpClient()
	ResetKafka()
	ResetMicrosoftSentinel()
	ResetNewRelic()
	ResetOpensearch()
	ResetRsyslog()
	ResetSentinelOne()
	ResetSocket()
	ResetSplunkHec()
	ResetSumoLogic()
	ResetSyslogNg()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ObservabilityPipelineConfigDestinationOutputReference
type jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) AmazonOpensearch() ObservabilityPipelineConfigDestinationAmazonOpensearchList {
	var returns ObservabilityPipelineConfigDestinationAmazonOpensearchList
	_jsii_.Get(
		j,
		"amazonOpensearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) AmazonOpensearchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"amazonOpensearchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) AmazonS3() ObservabilityPipelineConfigDestinationAmazonS3List {
	var returns ObservabilityPipelineConfigDestinationAmazonS3List
	_jsii_.Get(
		j,
		"amazonS3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) AmazonS3Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"amazonS3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) AmazonSecurityLake() ObservabilityPipelineConfigDestinationAmazonSecurityLakeList {
	var returns ObservabilityPipelineConfigDestinationAmazonSecurityLakeList
	_jsii_.Get(
		j,
		"amazonSecurityLake",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) AmazonSecurityLakeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"amazonSecurityLakeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) AzureStorage() ObservabilityPipelineConfigDestinationAzureStorageList {
	var returns ObservabilityPipelineConfigDestinationAzureStorageList
	_jsii_.Get(
		j,
		"azureStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) AzureStorageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azureStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) CloudPrem() ObservabilityPipelineConfigDestinationCloudPremList {
	var returns ObservabilityPipelineConfigDestinationCloudPremList
	_jsii_.Get(
		j,
		"cloudPrem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) CloudPremInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudPremInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) CrowdstrikeNextGenSiem() ObservabilityPipelineConfigDestinationCrowdstrikeNextGenSiemList {
	var returns ObservabilityPipelineConfigDestinationCrowdstrikeNextGenSiemList
	_jsii_.Get(
		j,
		"crowdstrikeNextGenSiem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) CrowdstrikeNextGenSiemInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crowdstrikeNextGenSiemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) DatadogLogs() ObservabilityPipelineConfigDestinationDatadogLogsList {
	var returns ObservabilityPipelineConfigDestinationDatadogLogsList
	_jsii_.Get(
		j,
		"datadogLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) DatadogLogsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"datadogLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) DatadogMetrics() ObservabilityPipelineConfigDestinationDatadogMetricsList {
	var returns ObservabilityPipelineConfigDestinationDatadogMetricsList
	_jsii_.Get(
		j,
		"datadogMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) DatadogMetricsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"datadogMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) Elasticsearch() ObservabilityPipelineConfigDestinationElasticsearchList {
	var returns ObservabilityPipelineConfigDestinationElasticsearchList
	_jsii_.Get(
		j,
		"elasticsearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) ElasticsearchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elasticsearchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) GoogleCloudStorage() ObservabilityPipelineConfigDestinationGoogleCloudStorageList {
	var returns ObservabilityPipelineConfigDestinationGoogleCloudStorageList
	_jsii_.Get(
		j,
		"googleCloudStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) GoogleCloudStorageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"googleCloudStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) GooglePubsub() ObservabilityPipelineConfigDestinationGooglePubsubList {
	var returns ObservabilityPipelineConfigDestinationGooglePubsubList
	_jsii_.Get(
		j,
		"googlePubsub",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) GooglePubsubInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"googlePubsubInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) GoogleSecops() ObservabilityPipelineConfigDestinationGoogleSecopsList {
	var returns ObservabilityPipelineConfigDestinationGoogleSecopsList
	_jsii_.Get(
		j,
		"googleSecops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) GoogleSecopsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"googleSecopsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) HttpClient() ObservabilityPipelineConfigDestinationHttpClientList {
	var returns ObservabilityPipelineConfigDestinationHttpClientList
	_jsii_.Get(
		j,
		"httpClient",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) HttpClientInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpClientInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) Inputs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) InputsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inputsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) Kafka() ObservabilityPipelineConfigDestinationKafkaList {
	var returns ObservabilityPipelineConfigDestinationKafkaList
	_jsii_.Get(
		j,
		"kafka",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) KafkaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kafkaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) MicrosoftSentinel() ObservabilityPipelineConfigDestinationMicrosoftSentinelList {
	var returns ObservabilityPipelineConfigDestinationMicrosoftSentinelList
	_jsii_.Get(
		j,
		"microsoftSentinel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) MicrosoftSentinelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"microsoftSentinelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) NewRelic() ObservabilityPipelineConfigDestinationNewRelicList {
	var returns ObservabilityPipelineConfigDestinationNewRelicList
	_jsii_.Get(
		j,
		"newRelic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) NewRelicInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"newRelicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) Opensearch() ObservabilityPipelineConfigDestinationOpensearchList {
	var returns ObservabilityPipelineConfigDestinationOpensearchList
	_jsii_.Get(
		j,
		"opensearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) OpensearchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"opensearchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) Rsyslog() ObservabilityPipelineConfigDestinationRsyslogList {
	var returns ObservabilityPipelineConfigDestinationRsyslogList
	_jsii_.Get(
		j,
		"rsyslog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) RsyslogInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rsyslogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) SentinelOne() ObservabilityPipelineConfigDestinationSentinelOneList {
	var returns ObservabilityPipelineConfigDestinationSentinelOneList
	_jsii_.Get(
		j,
		"sentinelOne",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) SentinelOneInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sentinelOneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) Socket() ObservabilityPipelineConfigDestinationSocketList {
	var returns ObservabilityPipelineConfigDestinationSocketList
	_jsii_.Get(
		j,
		"socket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) SocketInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"socketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) SplunkHec() ObservabilityPipelineConfigDestinationSplunkHecList {
	var returns ObservabilityPipelineConfigDestinationSplunkHecList
	_jsii_.Get(
		j,
		"splunkHec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) SplunkHecInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"splunkHecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) SumoLogic() ObservabilityPipelineConfigDestinationSumoLogicList {
	var returns ObservabilityPipelineConfigDestinationSumoLogicList
	_jsii_.Get(
		j,
		"sumoLogic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) SumoLogicInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sumoLogicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) SyslogNg() ObservabilityPipelineConfigDestinationSyslogNgList {
	var returns ObservabilityPipelineConfigDestinationSyslogNgList
	_jsii_.Get(
		j,
		"syslogNg",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) SyslogNgInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"syslogNgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewObservabilityPipelineConfigDestinationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ObservabilityPipelineConfigDestinationOutputReference {
	_init_.Initialize()

	if err := validateNewObservabilityPipelineConfigDestinationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.observabilityPipeline.ObservabilityPipelineConfigDestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewObservabilityPipelineConfigDestinationOutputReference_Override(o ObservabilityPipelineConfigDestinationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.observabilityPipeline.ObservabilityPipelineConfigDestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference)SetInputs(val *[]*string) {
	if err := j.validateSetInputsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputs",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) PutAmazonOpensearch(value interface{}) {
	if err := o.validatePutAmazonOpensearchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAmazonOpensearch",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) PutAmazonS3(value interface{}) {
	if err := o.validatePutAmazonS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAmazonS3",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) PutAmazonSecurityLake(value interface{}) {
	if err := o.validatePutAmazonSecurityLakeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAmazonSecurityLake",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) PutAzureStorage(value interface{}) {
	if err := o.validatePutAzureStorageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAzureStorage",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) PutCloudPrem(value interface{}) {
	if err := o.validatePutCloudPremParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putCloudPrem",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) PutCrowdstrikeNextGenSiem(value interface{}) {
	if err := o.validatePutCrowdstrikeNextGenSiemParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putCrowdstrikeNextGenSiem",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) PutDatadogLogs(value interface{}) {
	if err := o.validatePutDatadogLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putDatadogLogs",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) PutDatadogMetrics(value interface{}) {
	if err := o.validatePutDatadogMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putDatadogMetrics",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) PutElasticsearch(value interface{}) {
	if err := o.validatePutElasticsearchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putElasticsearch",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) PutGoogleCloudStorage(value interface{}) {
	if err := o.validatePutGoogleCloudStorageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putGoogleCloudStorage",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) PutGooglePubsub(value interface{}) {
	if err := o.validatePutGooglePubsubParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putGooglePubsub",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) PutGoogleSecops(value interface{}) {
	if err := o.validatePutGoogleSecopsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putGoogleSecops",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) PutHttpClient(value interface{}) {
	if err := o.validatePutHttpClientParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putHttpClient",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) PutKafka(value interface{}) {
	if err := o.validatePutKafkaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putKafka",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) PutMicrosoftSentinel(value interface{}) {
	if err := o.validatePutMicrosoftSentinelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putMicrosoftSentinel",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) PutNewRelic(value interface{}) {
	if err := o.validatePutNewRelicParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putNewRelic",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) PutOpensearch(value interface{}) {
	if err := o.validatePutOpensearchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putOpensearch",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) PutRsyslog(value interface{}) {
	if err := o.validatePutRsyslogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putRsyslog",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) PutSentinelOne(value interface{}) {
	if err := o.validatePutSentinelOneParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSentinelOne",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) PutSocket(value interface{}) {
	if err := o.validatePutSocketParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSocket",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) PutSplunkHec(value interface{}) {
	if err := o.validatePutSplunkHecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSplunkHec",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) PutSumoLogic(value interface{}) {
	if err := o.validatePutSumoLogicParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSumoLogic",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) PutSyslogNg(value interface{}) {
	if err := o.validatePutSyslogNgParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSyslogNg",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) ResetAmazonOpensearch() {
	_jsii_.InvokeVoid(
		o,
		"resetAmazonOpensearch",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) ResetAmazonS3() {
	_jsii_.InvokeVoid(
		o,
		"resetAmazonS3",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) ResetAmazonSecurityLake() {
	_jsii_.InvokeVoid(
		o,
		"resetAmazonSecurityLake",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) ResetAzureStorage() {
	_jsii_.InvokeVoid(
		o,
		"resetAzureStorage",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) ResetCloudPrem() {
	_jsii_.InvokeVoid(
		o,
		"resetCloudPrem",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) ResetCrowdstrikeNextGenSiem() {
	_jsii_.InvokeVoid(
		o,
		"resetCrowdstrikeNextGenSiem",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) ResetDatadogLogs() {
	_jsii_.InvokeVoid(
		o,
		"resetDatadogLogs",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) ResetDatadogMetrics() {
	_jsii_.InvokeVoid(
		o,
		"resetDatadogMetrics",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) ResetElasticsearch() {
	_jsii_.InvokeVoid(
		o,
		"resetElasticsearch",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) ResetGoogleCloudStorage() {
	_jsii_.InvokeVoid(
		o,
		"resetGoogleCloudStorage",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) ResetGooglePubsub() {
	_jsii_.InvokeVoid(
		o,
		"resetGooglePubsub",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) ResetGoogleSecops() {
	_jsii_.InvokeVoid(
		o,
		"resetGoogleSecops",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) ResetHttpClient() {
	_jsii_.InvokeVoid(
		o,
		"resetHttpClient",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) ResetKafka() {
	_jsii_.InvokeVoid(
		o,
		"resetKafka",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) ResetMicrosoftSentinel() {
	_jsii_.InvokeVoid(
		o,
		"resetMicrosoftSentinel",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) ResetNewRelic() {
	_jsii_.InvokeVoid(
		o,
		"resetNewRelic",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) ResetOpensearch() {
	_jsii_.InvokeVoid(
		o,
		"resetOpensearch",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) ResetRsyslog() {
	_jsii_.InvokeVoid(
		o,
		"resetRsyslog",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) ResetSentinelOne() {
	_jsii_.InvokeVoid(
		o,
		"resetSentinelOne",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) ResetSocket() {
	_jsii_.InvokeVoid(
		o,
		"resetSocket",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) ResetSplunkHec() {
	_jsii_.InvokeVoid(
		o,
		"resetSplunkHec",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) ResetSumoLogic() {
	_jsii_.InvokeVoid(
		o,
		"resetSumoLogic",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) ResetSyslogNg() {
	_jsii_.InvokeVoid(
		o,
		"resetSyslogNg",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := o.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigDestinationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

