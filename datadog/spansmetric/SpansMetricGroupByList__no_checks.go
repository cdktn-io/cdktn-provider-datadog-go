// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package spansmetric

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SpansMetricGroupByList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SpansMetricGroupByList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SpansMetricGroupByList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SpansMetricGroupByList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SpansMetricGroupByList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SpansMetricGroupByList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SpansMetricGroupByList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSpansMetricGroupByListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

