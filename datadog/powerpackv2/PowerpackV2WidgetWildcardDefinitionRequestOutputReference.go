// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package powerpackv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v16/powerpackv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PowerpackV2WidgetWildcardDefinitionRequestOutputReference interface {
	cdktn.ComplexObject
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
	// Experimental.
	Fqn() *string
	HistogramRequest() PowerpackV2WidgetWildcardDefinitionRequestHistogramRequestOutputReference
	HistogramRequestInput() *PowerpackV2WidgetWildcardDefinitionRequestHistogramRequest
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ListstreamRequest() PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestOutputReference
	ListstreamRequestInput() *PowerpackV2WidgetWildcardDefinitionRequestListstreamRequest
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimeseriesRequest() PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestOutputReference
	TimeseriesRequestInput() *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequest
	TreemapRequest() PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestOutputReference
	TreemapRequestInput() *PowerpackV2WidgetWildcardDefinitionRequestTreemapRequest
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
	PutHistogramRequest(value *PowerpackV2WidgetWildcardDefinitionRequestHistogramRequest)
	PutListstreamRequest(value *PowerpackV2WidgetWildcardDefinitionRequestListstreamRequest)
	PutTimeseriesRequest(value *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequest)
	PutTreemapRequest(value *PowerpackV2WidgetWildcardDefinitionRequestTreemapRequest)
	ResetHistogramRequest()
	ResetListstreamRequest()
	ResetTimeseriesRequest()
	ResetTreemapRequest()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PowerpackV2WidgetWildcardDefinitionRequestOutputReference
type jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestOutputReference) HistogramRequest() PowerpackV2WidgetWildcardDefinitionRequestHistogramRequestOutputReference {
	var returns PowerpackV2WidgetWildcardDefinitionRequestHistogramRequestOutputReference
	_jsii_.Get(
		j,
		"histogramRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestOutputReference) HistogramRequestInput() *PowerpackV2WidgetWildcardDefinitionRequestHistogramRequest {
	var returns *PowerpackV2WidgetWildcardDefinitionRequestHistogramRequest
	_jsii_.Get(
		j,
		"histogramRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestOutputReference) ListstreamRequest() PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestOutputReference {
	var returns PowerpackV2WidgetWildcardDefinitionRequestListstreamRequestOutputReference
	_jsii_.Get(
		j,
		"liststreamRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestOutputReference) ListstreamRequestInput() *PowerpackV2WidgetWildcardDefinitionRequestListstreamRequest {
	var returns *PowerpackV2WidgetWildcardDefinitionRequestListstreamRequest
	_jsii_.Get(
		j,
		"liststreamRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestOutputReference) TimeseriesRequest() PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestOutputReference {
	var returns PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequestOutputReference
	_jsii_.Get(
		j,
		"timeseriesRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestOutputReference) TimeseriesRequestInput() *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequest {
	var returns *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequest
	_jsii_.Get(
		j,
		"timeseriesRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestOutputReference) TreemapRequest() PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestOutputReference {
	var returns PowerpackV2WidgetWildcardDefinitionRequestTreemapRequestOutputReference
	_jsii_.Get(
		j,
		"treemapRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestOutputReference) TreemapRequestInput() *PowerpackV2WidgetWildcardDefinitionRequestTreemapRequest {
	var returns *PowerpackV2WidgetWildcardDefinitionRequestTreemapRequest
	_jsii_.Get(
		j,
		"treemapRequestInput",
		&returns,
	)
	return returns
}


func NewPowerpackV2WidgetWildcardDefinitionRequestOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PowerpackV2WidgetWildcardDefinitionRequestOutputReference {
	_init_.Initialize()

	if err := validateNewPowerpackV2WidgetWildcardDefinitionRequestOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetWildcardDefinitionRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPowerpackV2WidgetWildcardDefinitionRequestOutputReference_Override(p PowerpackV2WidgetWildcardDefinitionRequestOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.powerpackV2.PowerpackV2WidgetWildcardDefinitionRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestOutputReference) PutHistogramRequest(value *PowerpackV2WidgetWildcardDefinitionRequestHistogramRequest) {
	if err := p.validatePutHistogramRequestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putHistogramRequest",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestOutputReference) PutListstreamRequest(value *PowerpackV2WidgetWildcardDefinitionRequestListstreamRequest) {
	if err := p.validatePutListstreamRequestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putListstreamRequest",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestOutputReference) PutTimeseriesRequest(value *PowerpackV2WidgetWildcardDefinitionRequestTimeseriesRequest) {
	if err := p.validatePutTimeseriesRequestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putTimeseriesRequest",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestOutputReference) PutTreemapRequest(value *PowerpackV2WidgetWildcardDefinitionRequestTreemapRequest) {
	if err := p.validatePutTreemapRequestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putTreemapRequest",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestOutputReference) ResetHistogramRequest() {
	_jsii_.InvokeVoid(
		p,
		"resetHistogramRequest",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestOutputReference) ResetListstreamRequest() {
	_jsii_.InvokeVoid(
		p,
		"resetListstreamRequest",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestOutputReference) ResetTimeseriesRequest() {
	_jsii_.InvokeVoid(
		p,
		"resetTimeseriesRequest",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestOutputReference) ResetTreemapRequest() {
	_jsii_.InvokeVoid(
		p,
		"resetTreemapRequest",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := p.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PowerpackV2WidgetWildcardDefinitionRequestOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

