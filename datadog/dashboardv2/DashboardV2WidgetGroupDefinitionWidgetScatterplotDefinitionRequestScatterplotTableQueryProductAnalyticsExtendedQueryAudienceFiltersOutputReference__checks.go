// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package dashboardv2

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) validatePutAccountParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProductAnalyticsExtendedQueryAudienceFiltersAccount:
		value := value.(*[]*DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProductAnalyticsExtendedQueryAudienceFiltersAccount)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProductAnalyticsExtendedQueryAudienceFiltersAccount:
		value_ := value.([]*DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProductAnalyticsExtendedQueryAudienceFiltersAccount)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProductAnalyticsExtendedQueryAudienceFiltersAccount; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) validatePutSegmentParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProductAnalyticsExtendedQueryAudienceFiltersSegment:
		value := value.(*[]*DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProductAnalyticsExtendedQueryAudienceFiltersSegment)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProductAnalyticsExtendedQueryAudienceFiltersSegment:
		value_ := value.([]*DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProductAnalyticsExtendedQueryAudienceFiltersSegment)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProductAnalyticsExtendedQueryAudienceFiltersSegment; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) validatePutUserParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProductAnalyticsExtendedQueryAudienceFiltersUser:
		value := value.(*[]*DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProductAnalyticsExtendedQueryAudienceFiltersUser)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProductAnalyticsExtendedQueryAudienceFiltersUser:
		value_ := value.([]*DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProductAnalyticsExtendedQueryAudienceFiltersUser)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProductAnalyticsExtendedQueryAudienceFiltersUser; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (d *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) validateSetFilterConditionParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) validateSetInternalValueParameters(val *DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProductAnalyticsExtendedQueryAudienceFilters) error {
	if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewDashboardV2WidgetGroupDefinitionWidgetScatterplotDefinitionRequestScatterplotTableQueryProductAnalyticsExtendedQueryAudienceFiltersOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

