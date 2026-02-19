// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package restrictionpolicy

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RestrictionPolicyBindingsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RestrictionPolicyBindingsList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RestrictionPolicyBindingsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RestrictionPolicyBindingsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RestrictionPolicyBindingsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RestrictionPolicyBindingsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RestrictionPolicyBindingsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRestrictionPolicyBindingsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

