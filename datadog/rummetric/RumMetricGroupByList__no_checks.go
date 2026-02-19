// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package rummetric

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RumMetricGroupByList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RumMetricGroupByList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RumMetricGroupByList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RumMetricGroupByList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RumMetricGroupByList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RumMetricGroupByList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RumMetricGroupByList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRumMetricGroupByListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

