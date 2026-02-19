// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package observabilitypipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v13/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v13/observabilitypipeline/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ObservabilityPipelineConfigSourceOutputReference interface {
	cdktn.ComplexObject
	AmazonDataFirehose() ObservabilityPipelineConfigSourceAmazonDataFirehoseList
	AmazonDataFirehoseInput() interface{}
	AmazonS3() ObservabilityPipelineConfigSourceAmazonS3List
	AmazonS3Input() interface{}
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
	DatadogAgent() ObservabilityPipelineConfigSourceDatadogAgentList
	DatadogAgentInput() interface{}
	FluentBit() ObservabilityPipelineConfigSourceFluentBitList
	FluentBitInput() interface{}
	Fluentd() ObservabilityPipelineConfigSourceFluentdList
	FluentdInput() interface{}
	// Experimental.
	Fqn() *string
	GooglePubsub() ObservabilityPipelineConfigSourceGooglePubsubList
	GooglePubsubInput() interface{}
	HttpClient() ObservabilityPipelineConfigSourceHttpClientList
	HttpClientInput() interface{}
	HttpServer() ObservabilityPipelineConfigSourceHttpServerList
	HttpServerInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Kafka() ObservabilityPipelineConfigSourceKafkaList
	KafkaInput() interface{}
	Logstash() ObservabilityPipelineConfigSourceLogstashList
	LogstashInput() interface{}
	Opentelemetry() ObservabilityPipelineConfigSourceOpentelemetryList
	OpentelemetryInput() interface{}
	Rsyslog() ObservabilityPipelineConfigSourceRsyslogList
	RsyslogInput() interface{}
	Socket() ObservabilityPipelineConfigSourceSocketList
	SocketInput() interface{}
	SplunkHec() ObservabilityPipelineConfigSourceSplunkHecList
	SplunkHecInput() interface{}
	SplunkTcp() ObservabilityPipelineConfigSourceSplunkTcpList
	SplunkTcpInput() interface{}
	SumoLogic() ObservabilityPipelineConfigSourceSumoLogicList
	SumoLogicInput() interface{}
	SyslogNg() ObservabilityPipelineConfigSourceSyslogNgList
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
	PutAmazonDataFirehose(value interface{})
	PutAmazonS3(value interface{})
	PutDatadogAgent(value interface{})
	PutFluentBit(value interface{})
	PutFluentd(value interface{})
	PutGooglePubsub(value interface{})
	PutHttpClient(value interface{})
	PutHttpServer(value interface{})
	PutKafka(value interface{})
	PutLogstash(value interface{})
	PutOpentelemetry(value interface{})
	PutRsyslog(value interface{})
	PutSocket(value interface{})
	PutSplunkHec(value interface{})
	PutSplunkTcp(value interface{})
	PutSumoLogic(value interface{})
	PutSyslogNg(value interface{})
	ResetAmazonDataFirehose()
	ResetAmazonS3()
	ResetDatadogAgent()
	ResetFluentBit()
	ResetFluentd()
	ResetGooglePubsub()
	ResetHttpClient()
	ResetHttpServer()
	ResetKafka()
	ResetLogstash()
	ResetOpentelemetry()
	ResetRsyslog()
	ResetSocket()
	ResetSplunkHec()
	ResetSplunkTcp()
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

// The jsii proxy struct for ObservabilityPipelineConfigSourceOutputReference
type jsiiProxy_ObservabilityPipelineConfigSourceOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) AmazonDataFirehose() ObservabilityPipelineConfigSourceAmazonDataFirehoseList {
	var returns ObservabilityPipelineConfigSourceAmazonDataFirehoseList
	_jsii_.Get(
		j,
		"amazonDataFirehose",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) AmazonDataFirehoseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"amazonDataFirehoseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) AmazonS3() ObservabilityPipelineConfigSourceAmazonS3List {
	var returns ObservabilityPipelineConfigSourceAmazonS3List
	_jsii_.Get(
		j,
		"amazonS3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) AmazonS3Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"amazonS3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) DatadogAgent() ObservabilityPipelineConfigSourceDatadogAgentList {
	var returns ObservabilityPipelineConfigSourceDatadogAgentList
	_jsii_.Get(
		j,
		"datadogAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) DatadogAgentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"datadogAgentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) FluentBit() ObservabilityPipelineConfigSourceFluentBitList {
	var returns ObservabilityPipelineConfigSourceFluentBitList
	_jsii_.Get(
		j,
		"fluentBit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) FluentBitInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fluentBitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) Fluentd() ObservabilityPipelineConfigSourceFluentdList {
	var returns ObservabilityPipelineConfigSourceFluentdList
	_jsii_.Get(
		j,
		"fluentd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) FluentdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fluentdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) GooglePubsub() ObservabilityPipelineConfigSourceGooglePubsubList {
	var returns ObservabilityPipelineConfigSourceGooglePubsubList
	_jsii_.Get(
		j,
		"googlePubsub",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) GooglePubsubInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"googlePubsubInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) HttpClient() ObservabilityPipelineConfigSourceHttpClientList {
	var returns ObservabilityPipelineConfigSourceHttpClientList
	_jsii_.Get(
		j,
		"httpClient",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) HttpClientInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpClientInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) HttpServer() ObservabilityPipelineConfigSourceHttpServerList {
	var returns ObservabilityPipelineConfigSourceHttpServerList
	_jsii_.Get(
		j,
		"httpServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) HttpServerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) Kafka() ObservabilityPipelineConfigSourceKafkaList {
	var returns ObservabilityPipelineConfigSourceKafkaList
	_jsii_.Get(
		j,
		"kafka",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) KafkaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kafkaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) Logstash() ObservabilityPipelineConfigSourceLogstashList {
	var returns ObservabilityPipelineConfigSourceLogstashList
	_jsii_.Get(
		j,
		"logstash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) LogstashInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logstashInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) Opentelemetry() ObservabilityPipelineConfigSourceOpentelemetryList {
	var returns ObservabilityPipelineConfigSourceOpentelemetryList
	_jsii_.Get(
		j,
		"opentelemetry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) OpentelemetryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"opentelemetryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) Rsyslog() ObservabilityPipelineConfigSourceRsyslogList {
	var returns ObservabilityPipelineConfigSourceRsyslogList
	_jsii_.Get(
		j,
		"rsyslog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) RsyslogInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rsyslogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) Socket() ObservabilityPipelineConfigSourceSocketList {
	var returns ObservabilityPipelineConfigSourceSocketList
	_jsii_.Get(
		j,
		"socket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) SocketInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"socketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) SplunkHec() ObservabilityPipelineConfigSourceSplunkHecList {
	var returns ObservabilityPipelineConfigSourceSplunkHecList
	_jsii_.Get(
		j,
		"splunkHec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) SplunkHecInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"splunkHecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) SplunkTcp() ObservabilityPipelineConfigSourceSplunkTcpList {
	var returns ObservabilityPipelineConfigSourceSplunkTcpList
	_jsii_.Get(
		j,
		"splunkTcp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) SplunkTcpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"splunkTcpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) SumoLogic() ObservabilityPipelineConfigSourceSumoLogicList {
	var returns ObservabilityPipelineConfigSourceSumoLogicList
	_jsii_.Get(
		j,
		"sumoLogic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) SumoLogicInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sumoLogicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) SyslogNg() ObservabilityPipelineConfigSourceSyslogNgList {
	var returns ObservabilityPipelineConfigSourceSyslogNgList
	_jsii_.Get(
		j,
		"syslogNg",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) SyslogNgInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"syslogNgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewObservabilityPipelineConfigSourceOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ObservabilityPipelineConfigSourceOutputReference {
	_init_.Initialize()

	if err := validateNewObservabilityPipelineConfigSourceOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ObservabilityPipelineConfigSourceOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.observabilityPipeline.ObservabilityPipelineConfigSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewObservabilityPipelineConfigSourceOutputReference_Override(o ObservabilityPipelineConfigSourceOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.observabilityPipeline.ObservabilityPipelineConfigSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (o *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (o *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) PutAmazonDataFirehose(value interface{}) {
	if err := o.validatePutAmazonDataFirehoseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAmazonDataFirehose",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) PutAmazonS3(value interface{}) {
	if err := o.validatePutAmazonS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAmazonS3",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) PutDatadogAgent(value interface{}) {
	if err := o.validatePutDatadogAgentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putDatadogAgent",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) PutFluentBit(value interface{}) {
	if err := o.validatePutFluentBitParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putFluentBit",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) PutFluentd(value interface{}) {
	if err := o.validatePutFluentdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putFluentd",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) PutGooglePubsub(value interface{}) {
	if err := o.validatePutGooglePubsubParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putGooglePubsub",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) PutHttpClient(value interface{}) {
	if err := o.validatePutHttpClientParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putHttpClient",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) PutHttpServer(value interface{}) {
	if err := o.validatePutHttpServerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putHttpServer",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) PutKafka(value interface{}) {
	if err := o.validatePutKafkaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putKafka",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) PutLogstash(value interface{}) {
	if err := o.validatePutLogstashParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putLogstash",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) PutOpentelemetry(value interface{}) {
	if err := o.validatePutOpentelemetryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putOpentelemetry",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) PutRsyslog(value interface{}) {
	if err := o.validatePutRsyslogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putRsyslog",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) PutSocket(value interface{}) {
	if err := o.validatePutSocketParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSocket",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) PutSplunkHec(value interface{}) {
	if err := o.validatePutSplunkHecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSplunkHec",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) PutSplunkTcp(value interface{}) {
	if err := o.validatePutSplunkTcpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSplunkTcp",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) PutSumoLogic(value interface{}) {
	if err := o.validatePutSumoLogicParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSumoLogic",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) PutSyslogNg(value interface{}) {
	if err := o.validatePutSyslogNgParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSyslogNg",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) ResetAmazonDataFirehose() {
	_jsii_.InvokeVoid(
		o,
		"resetAmazonDataFirehose",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) ResetAmazonS3() {
	_jsii_.InvokeVoid(
		o,
		"resetAmazonS3",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) ResetDatadogAgent() {
	_jsii_.InvokeVoid(
		o,
		"resetDatadogAgent",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) ResetFluentBit() {
	_jsii_.InvokeVoid(
		o,
		"resetFluentBit",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) ResetFluentd() {
	_jsii_.InvokeVoid(
		o,
		"resetFluentd",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) ResetGooglePubsub() {
	_jsii_.InvokeVoid(
		o,
		"resetGooglePubsub",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) ResetHttpClient() {
	_jsii_.InvokeVoid(
		o,
		"resetHttpClient",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) ResetHttpServer() {
	_jsii_.InvokeVoid(
		o,
		"resetHttpServer",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) ResetKafka() {
	_jsii_.InvokeVoid(
		o,
		"resetKafka",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) ResetLogstash() {
	_jsii_.InvokeVoid(
		o,
		"resetLogstash",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) ResetOpentelemetry() {
	_jsii_.InvokeVoid(
		o,
		"resetOpentelemetry",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) ResetRsyslog() {
	_jsii_.InvokeVoid(
		o,
		"resetRsyslog",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) ResetSocket() {
	_jsii_.InvokeVoid(
		o,
		"resetSocket",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) ResetSplunkHec() {
	_jsii_.InvokeVoid(
		o,
		"resetSplunkHec",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) ResetSplunkTcp() {
	_jsii_.InvokeVoid(
		o,
		"resetSplunkTcp",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) ResetSumoLogic() {
	_jsii_.InvokeVoid(
		o,
		"resetSumoLogic",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) ResetSyslogNg() {
	_jsii_.InvokeVoid(
		o,
		"resetSyslogNg",
		nil, // no parameters
	)
}

func (o *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (o *jsiiProxy_ObservabilityPipelineConfigSourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

