// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package integrationawsaccount

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/integrationawsaccount/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IntegrationAwsAccountMetricsConfigMetricNameFiltersOutputReference interface {
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
	ExcludeOnly() *[]*string
	SetExcludeOnly(val *[]*string)
	ExcludeOnlyInput() *[]*string
	// Experimental.
	Fqn() *string
	IncludeOnly() *[]*string
	SetIncludeOnly(val *[]*string)
	IncludeOnlyInput() *[]*string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
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
	ResetExcludeOnly()
	ResetIncludeOnly()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IntegrationAwsAccountMetricsConfigMetricNameFiltersOutputReference
type jsiiProxy_IntegrationAwsAccountMetricsConfigMetricNameFiltersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IntegrationAwsAccountMetricsConfigMetricNameFiltersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccountMetricsConfigMetricNameFiltersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccountMetricsConfigMetricNameFiltersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccountMetricsConfigMetricNameFiltersOutputReference) ExcludeOnly() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccountMetricsConfigMetricNameFiltersOutputReference) ExcludeOnlyInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccountMetricsConfigMetricNameFiltersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccountMetricsConfigMetricNameFiltersOutputReference) IncludeOnly() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccountMetricsConfigMetricNameFiltersOutputReference) IncludeOnlyInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccountMetricsConfigMetricNameFiltersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccountMetricsConfigMetricNameFiltersOutputReference) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccountMetricsConfigMetricNameFiltersOutputReference) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccountMetricsConfigMetricNameFiltersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAwsAccountMetricsConfigMetricNameFiltersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIntegrationAwsAccountMetricsConfigMetricNameFiltersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) IntegrationAwsAccountMetricsConfigMetricNameFiltersOutputReference {
	_init_.Initialize()

	if err := validateNewIntegrationAwsAccountMetricsConfigMetricNameFiltersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_IntegrationAwsAccountMetricsConfigMetricNameFiltersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.integrationAwsAccount.IntegrationAwsAccountMetricsConfigMetricNameFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewIntegrationAwsAccountMetricsConfigMetricNameFiltersOutputReference_Override(i IntegrationAwsAccountMetricsConfigMetricNameFiltersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.integrationAwsAccount.IntegrationAwsAccountMetricsConfigMetricNameFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		i,
	)
}

func (j *jsiiProxy_IntegrationAwsAccountMetricsConfigMetricNameFiltersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IntegrationAwsAccountMetricsConfigMetricNameFiltersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IntegrationAwsAccountMetricsConfigMetricNameFiltersOutputReference)SetExcludeOnly(val *[]*string) {
	if err := j.validateSetExcludeOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeOnly",
		val,
	)
}

func (j *jsiiProxy_IntegrationAwsAccountMetricsConfigMetricNameFiltersOutputReference)SetIncludeOnly(val *[]*string) {
	if err := j.validateSetIncludeOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeOnly",
		val,
	)
}

func (j *jsiiProxy_IntegrationAwsAccountMetricsConfigMetricNameFiltersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IntegrationAwsAccountMetricsConfigMetricNameFiltersOutputReference)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_IntegrationAwsAccountMetricsConfigMetricNameFiltersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IntegrationAwsAccountMetricsConfigMetricNameFiltersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IntegrationAwsAccountMetricsConfigMetricNameFiltersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccountMetricsConfigMetricNameFiltersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccountMetricsConfigMetricNameFiltersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccountMetricsConfigMetricNameFiltersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccountMetricsConfigMetricNameFiltersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccountMetricsConfigMetricNameFiltersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccountMetricsConfigMetricNameFiltersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccountMetricsConfigMetricNameFiltersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccountMetricsConfigMetricNameFiltersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccountMetricsConfigMetricNameFiltersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccountMetricsConfigMetricNameFiltersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccountMetricsConfigMetricNameFiltersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccountMetricsConfigMetricNameFiltersOutputReference) ResetExcludeOnly() {
	_jsii_.InvokeVoid(
		i,
		"resetExcludeOnly",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationAwsAccountMetricsConfigMetricNameFiltersOutputReference) ResetIncludeOnly() {
	_jsii_.InvokeVoid(
		i,
		"resetIncludeOnly",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationAwsAccountMetricsConfigMetricNameFiltersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := i.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAwsAccountMetricsConfigMetricNameFiltersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

