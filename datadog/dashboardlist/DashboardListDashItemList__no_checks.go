// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dashboardlist

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DashboardListDashItemList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DashboardListDashItemList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DashboardListDashItemList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DashboardListDashItemList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DashboardListDashItemList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DashboardListDashItemList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DashboardListDashItemList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDashboardListDashItemListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

