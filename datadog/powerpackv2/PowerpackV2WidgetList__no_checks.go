// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package powerpackv2

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PowerpackV2WidgetList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_PowerpackV2WidgetList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PowerpackV2WidgetList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PowerpackV2WidgetList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PowerpackV2WidgetList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PowerpackV2WidgetList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PowerpackV2WidgetList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPowerpackV2WidgetListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

