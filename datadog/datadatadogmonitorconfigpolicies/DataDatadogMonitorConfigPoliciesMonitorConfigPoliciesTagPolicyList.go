// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datadatadogmonitorconfigpolicies

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v13/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v13/datadatadogmonitorconfigpolicies/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataDatadogMonitorConfigPoliciesMonitorConfigPoliciesTagPolicyList interface {
	cdktn.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktn.IInterpolatingParent
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) DataDatadogMonitorConfigPoliciesMonitorConfigPoliciesTagPolicyOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataDatadogMonitorConfigPoliciesMonitorConfigPoliciesTagPolicyList
type jsiiProxy_DataDatadogMonitorConfigPoliciesMonitorConfigPoliciesTagPolicyList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_DataDatadogMonitorConfigPoliciesMonitorConfigPoliciesTagPolicyList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitorConfigPoliciesMonitorConfigPoliciesTagPolicyList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitorConfigPoliciesMonitorConfigPoliciesTagPolicyList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitorConfigPoliciesMonitorConfigPoliciesTagPolicyList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataDatadogMonitorConfigPoliciesMonitorConfigPoliciesTagPolicyList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDataDatadogMonitorConfigPoliciesMonitorConfigPoliciesTagPolicyList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DataDatadogMonitorConfigPoliciesMonitorConfigPoliciesTagPolicyList {
	_init_.Initialize()

	if err := validateNewDataDatadogMonitorConfigPoliciesMonitorConfigPoliciesTagPolicyListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataDatadogMonitorConfigPoliciesMonitorConfigPoliciesTagPolicyList{}

	_jsii_.Create(
		"@cdktn/provider-datadog.dataDatadogMonitorConfigPolicies.DataDatadogMonitorConfigPoliciesMonitorConfigPoliciesTagPolicyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDataDatadogMonitorConfigPoliciesMonitorConfigPoliciesTagPolicyList_Override(d DataDatadogMonitorConfigPoliciesMonitorConfigPoliciesTagPolicyList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.dataDatadogMonitorConfigPolicies.DataDatadogMonitorConfigPoliciesMonitorConfigPoliciesTagPolicyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataDatadogMonitorConfigPoliciesMonitorConfigPoliciesTagPolicyList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataDatadogMonitorConfigPoliciesMonitorConfigPoliciesTagPolicyList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataDatadogMonitorConfigPoliciesMonitorConfigPoliciesTagPolicyList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DataDatadogMonitorConfigPoliciesMonitorConfigPoliciesTagPolicyList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := d.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		d,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogMonitorConfigPoliciesMonitorConfigPoliciesTagPolicyList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogMonitorConfigPoliciesMonitorConfigPoliciesTagPolicyList) Get(index *float64) DataDatadogMonitorConfigPoliciesMonitorConfigPoliciesTagPolicyOutputReference {
	if err := d.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns DataDatadogMonitorConfigPoliciesMonitorConfigPoliciesTagPolicyOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataDatadogMonitorConfigPoliciesMonitorConfigPoliciesTagPolicyList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataDatadogMonitorConfigPoliciesMonitorConfigPoliciesTagPolicyList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

