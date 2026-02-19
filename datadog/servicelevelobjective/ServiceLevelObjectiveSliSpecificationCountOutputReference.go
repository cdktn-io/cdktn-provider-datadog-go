// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package servicelevelobjective

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v13/jsii"

	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v13/servicelevelobjective/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ServiceLevelObjectiveSliSpecificationCountOutputReference interface {
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
	GoodEventsFormula() *string
	SetGoodEventsFormula(val *string)
	GoodEventsFormulaInput() *string
	InternalValue() *ServiceLevelObjectiveSliSpecificationCount
	SetInternalValue(val *ServiceLevelObjectiveSliSpecificationCount)
	Queries() ServiceLevelObjectiveSliSpecificationCountQueriesList
	QueriesInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TotalEventsFormula() *string
	SetTotalEventsFormula(val *string)
	TotalEventsFormulaInput() *string
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
	PutQueries(value interface{})
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServiceLevelObjectiveSliSpecificationCountOutputReference
type jsiiProxy_ServiceLevelObjectiveSliSpecificationCountOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ServiceLevelObjectiveSliSpecificationCountOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjectiveSliSpecificationCountOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjectiveSliSpecificationCountOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjectiveSliSpecificationCountOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjectiveSliSpecificationCountOutputReference) GoodEventsFormula() *string {
	var returns *string
	_jsii_.Get(
		j,
		"goodEventsFormula",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjectiveSliSpecificationCountOutputReference) GoodEventsFormulaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"goodEventsFormulaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjectiveSliSpecificationCountOutputReference) InternalValue() *ServiceLevelObjectiveSliSpecificationCount {
	var returns *ServiceLevelObjectiveSliSpecificationCount
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjectiveSliSpecificationCountOutputReference) Queries() ServiceLevelObjectiveSliSpecificationCountQueriesList {
	var returns ServiceLevelObjectiveSliSpecificationCountQueriesList
	_jsii_.Get(
		j,
		"queries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjectiveSliSpecificationCountOutputReference) QueriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjectiveSliSpecificationCountOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjectiveSliSpecificationCountOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjectiveSliSpecificationCountOutputReference) TotalEventsFormula() *string {
	var returns *string
	_jsii_.Get(
		j,
		"totalEventsFormula",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelObjectiveSliSpecificationCountOutputReference) TotalEventsFormulaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"totalEventsFormulaInput",
		&returns,
	)
	return returns
}


func NewServiceLevelObjectiveSliSpecificationCountOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ServiceLevelObjectiveSliSpecificationCountOutputReference {
	_init_.Initialize()

	if err := validateNewServiceLevelObjectiveSliSpecificationCountOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ServiceLevelObjectiveSliSpecificationCountOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-datadog.serviceLevelObjective.ServiceLevelObjectiveSliSpecificationCountOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewServiceLevelObjectiveSliSpecificationCountOutputReference_Override(s ServiceLevelObjectiveSliSpecificationCountOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.serviceLevelObjective.ServiceLevelObjectiveSliSpecificationCountOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_ServiceLevelObjectiveSliSpecificationCountOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServiceLevelObjectiveSliSpecificationCountOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServiceLevelObjectiveSliSpecificationCountOutputReference)SetGoodEventsFormula(val *string) {
	if err := j.validateSetGoodEventsFormulaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"goodEventsFormula",
		val,
	)
}

func (j *jsiiProxy_ServiceLevelObjectiveSliSpecificationCountOutputReference)SetInternalValue(val *ServiceLevelObjectiveSliSpecificationCount) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceLevelObjectiveSliSpecificationCountOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceLevelObjectiveSliSpecificationCountOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ServiceLevelObjectiveSliSpecificationCountOutputReference)SetTotalEventsFormula(val *string) {
	if err := j.validateSetTotalEventsFormulaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"totalEventsFormula",
		val,
	)
}

func (s *jsiiProxy_ServiceLevelObjectiveSliSpecificationCountOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceLevelObjectiveSliSpecificationCountOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceLevelObjectiveSliSpecificationCountOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceLevelObjectiveSliSpecificationCountOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceLevelObjectiveSliSpecificationCountOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceLevelObjectiveSliSpecificationCountOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceLevelObjectiveSliSpecificationCountOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceLevelObjectiveSliSpecificationCountOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceLevelObjectiveSliSpecificationCountOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceLevelObjectiveSliSpecificationCountOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceLevelObjectiveSliSpecificationCountOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceLevelObjectiveSliSpecificationCountOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceLevelObjectiveSliSpecificationCountOutputReference) PutQueries(value interface{}) {
	if err := s.validatePutQueriesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putQueries",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceLevelObjectiveSliSpecificationCountOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := s.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceLevelObjectiveSliSpecificationCountOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

