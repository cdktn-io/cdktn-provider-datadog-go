// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package governancecontrol

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GovernanceControlNotificationSettingsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (g *jsiiProxy_GovernanceControlNotificationSettingsList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GovernanceControlNotificationSettingsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GovernanceControlNotificationSettingsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GovernanceControlNotificationSettingsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GovernanceControlNotificationSettingsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GovernanceControlNotificationSettingsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGovernanceControlNotificationSettingsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

