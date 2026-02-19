// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package oncallschedule

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OnCallScheduleLayerList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (o *jsiiProxy_OnCallScheduleLayerList) validateGetParameters(index *float64) error {
	return nil
}

func (o *jsiiProxy_OnCallScheduleLayerList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_OnCallScheduleLayerList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_OnCallScheduleLayerList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_OnCallScheduleLayerList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_OnCallScheduleLayerList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewOnCallScheduleLayerListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

