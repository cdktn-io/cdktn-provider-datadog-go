// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/dashboardv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DashboardV2WidgetWildcardDefinitionRequestOutputReference interface {
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
	HistogramRequest() DashboardV2WidgetWildcardDefinitionRequestHistogramRequestOutputReference
	HistogramRequestInput() *DashboardV2WidgetWildcardDefinitionRequestHistogramRequest
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ListstreamRequest() DashboardV2WidgetWildcardDefinitionRequestListstreamRequestOutputReference
	ListstreamRequestInput() *DashboardV2WidgetWildcardDefinitionRequestListstreamRequest
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimeseriesRequest() DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestOutputReference
	TimeseriesRequestInput() *DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequest
	TreemapRequest() DashboardV2WidgetWildcardDefinitionRequestTreemapRequestOutputReference
	TreemapRequestInput() *DashboardV2WidgetWildcardDefinitionRequestTreemapRequest
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
	PutHistogramRequest(value *DashboardV2WidgetWildcardDefinitionRequestHistogramRequest)
	PutListstreamRequest(value *DashboardV2WidgetWildcardDefinitionRequestListstreamRequest)
	PutTimeseriesRequest(value *DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequest)
	PutTreemapRequest(value *DashboardV2WidgetWildcardDefinitionRequestTreemapRequest)
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

// The jsii proxy struct for DashboardV2WidgetWildcardDefinitionRequestOutputReference
type jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestOutputReference) HistogramRequest() DashboardV2WidgetWildcardDefinitionRequestHistogramRequestOutputReference {
	var returns DashboardV2WidgetWildcardDefinitionRequestHistogramRequestOutputReference
	_jsii_.Get(
		j,
		"histogramRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestOutputReference) HistogramRequestInput() *DashboardV2WidgetWildcardDefinitionRequestHistogramRequest {
	var returns *DashboardV2WidgetWildcardDefinitionRequestHistogramRequest
	_jsii_.Get(
		j,
		"histogramRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestOutputReference) ListstreamRequest() DashboardV2WidgetWildcardDefinitionRequestListstreamRequestOutputReference {
	var returns DashboardV2WidgetWildcardDefinitionRequestListstreamRequestOutputReference
	_jsii_.Get(
		j,
		"liststreamRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestOutputReference) ListstreamRequestInput() *DashboardV2WidgetWildcardDefinitionRequestListstreamRequest {
	var returns *DashboardV2WidgetWildcardDefinitionRequestListstreamRequest
	_jsii_.Get(
		j,
		"liststreamRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestOutputReference) TimeseriesRequest() DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestOutputReference {
	var returns DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequestOutputReference
	_jsii_.Get(
		j,
		"timeseriesRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestOutputReference) TimeseriesRequestInput() *DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequest {
	var returns *DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequest
	_jsii_.Get(
		j,
		"timeseriesRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestOutputReference) TreemapRequest() DashboardV2WidgetWildcardDefinitionRequestTreemapRequestOutputReference {
	var returns DashboardV2WidgetWildcardDefinitionRequestTreemapRequestOutputReference
	_jsii_.Get(
		j,
		"treemapRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestOutputReference) TreemapRequestInput() *DashboardV2WidgetWildcardDefinitionRequestTreemapRequest {
	var returns *DashboardV2WidgetWildcardDefinitionRequestTreemapRequest
	_jsii_.Get(
		j,
		"treemapRequestInput",
		&returns,
	)
	return returns
}


func NewDashboardV2WidgetWildcardDefinitionRequestOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DashboardV2WidgetWildcardDefinitionRequestOutputReference {
	_init_.Initialize()

	if err := validateNewDashboardV2WidgetWildcardDefinitionRequestOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetWildcardDefinitionRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDashboardV2WidgetWildcardDefinitionRequestOutputReference_Override(d DashboardV2WidgetWildcardDefinitionRequestOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2WidgetWildcardDefinitionRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestOutputReference) PutHistogramRequest(value *DashboardV2WidgetWildcardDefinitionRequestHistogramRequest) {
	if err := d.validatePutHistogramRequestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putHistogramRequest",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestOutputReference) PutListstreamRequest(value *DashboardV2WidgetWildcardDefinitionRequestListstreamRequest) {
	if err := d.validatePutListstreamRequestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putListstreamRequest",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestOutputReference) PutTimeseriesRequest(value *DashboardV2WidgetWildcardDefinitionRequestTimeseriesRequest) {
	if err := d.validatePutTimeseriesRequestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeseriesRequest",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestOutputReference) PutTreemapRequest(value *DashboardV2WidgetWildcardDefinitionRequestTreemapRequest) {
	if err := d.validatePutTreemapRequestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTreemapRequest",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestOutputReference) ResetHistogramRequest() {
	_jsii_.InvokeVoid(
		d,
		"resetHistogramRequest",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestOutputReference) ResetListstreamRequest() {
	_jsii_.InvokeVoid(
		d,
		"resetListstreamRequest",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestOutputReference) ResetTimeseriesRequest() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeseriesRequest",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestOutputReference) ResetTreemapRequest() {
	_jsii_.InvokeVoid(
		d,
		"resetTreemapRequest",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2WidgetWildcardDefinitionRequestOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

