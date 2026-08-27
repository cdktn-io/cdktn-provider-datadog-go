// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package dashboardv2

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) validatePutAccountParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersAccount:
		value := value.(*[]*DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersAccount)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersAccount:
		value_ := value.([]*DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersAccount)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersAccount; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) validatePutSegmentParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersSegment:
		value := value.(*[]*DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersSegment)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersSegment:
		value_ := value.([]*DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersSegment)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersSegment; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) validatePutUserParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersUser:
		value := value.(*[]*DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersUser)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersUser:
		value_ := value.([]*DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersUser)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersUser; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
	switch val.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *float64:
		// ok
	case float64:
		// ok
	case *int:
		// ok
	case int:
		// ok
	case *uint:
		// ok
	case uint:
		// ok
	case *int8:
		// ok
	case int8:
		// ok
	case *int16:
		// ok
	case int16:
		// ok
	case *int32:
		// ok
	case int32:
		// ok
	case *int64:
		// ok
	case int64:
		// ok
	case *uint8:
		// ok
	case uint8:
		// ok
	case *uint16:
		// ok
	case uint16:
		// ok
	case *uint32:
		// ok
	case uint32:
		// ok
	case *uint64:
		// ok
	case uint64:
		// ok
	default:
		return fmt.Errorf("parameter val must be one of the allowed types: *string, *float64; received %#v (a %T)", val, val)
	}

	return nil
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) validateSetFilterConditionParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) validateSetInternalValueParameters(val *DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFilters) error {
	if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewDashboardV2WidgetGroupDefinitionWidgetQueryValueDefinitionRequestQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

