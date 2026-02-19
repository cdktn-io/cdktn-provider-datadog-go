// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dashboard

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DashboardTemplateVariablePresetList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DashboardTemplateVariablePresetList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DashboardTemplateVariablePresetList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DashboardTemplateVariablePresetList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DashboardTemplateVariablePresetList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DashboardTemplateVariablePresetList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DashboardTemplateVariablePresetList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDashboardTemplateVariablePresetListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

