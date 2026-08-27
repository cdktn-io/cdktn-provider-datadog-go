// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package powerpackv2

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (p *jsiiProxy_PowerpackV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PowerpackV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PowerpackV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PowerpackV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PowerpackV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PowerpackV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PowerpackV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PowerpackV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PowerpackV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PowerpackV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PowerpackV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) validatePutAccountParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*PowerpackV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersAccount:
		value := value.(*[]*PowerpackV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersAccount)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*PowerpackV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersAccount:
		value_ := value.([]*PowerpackV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersAccount)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*PowerpackV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersAccount; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (p *jsiiProxy_PowerpackV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) validatePutSegmentParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*PowerpackV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersSegment:
		value := value.(*[]*PowerpackV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersSegment)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*PowerpackV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersSegment:
		value_ := value.([]*PowerpackV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersSegment)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*PowerpackV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersSegment; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (p *jsiiProxy_PowerpackV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) validatePutUserParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*PowerpackV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersUser:
		value := value.(*[]*PowerpackV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersUser)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*PowerpackV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersUser:
		value_ := value.([]*PowerpackV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersUser)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*PowerpackV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersUser; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (p *jsiiProxy_PowerpackV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_PowerpackV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_PowerpackV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_PowerpackV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) validateSetFilterConditionParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_PowerpackV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) validateSetInternalValueParameters(val *PowerpackV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFilters) error {
	if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_PowerpackV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_PowerpackV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewPowerpackV2WidgetQueryTableDefinitionRequestQueryUserJourneyQuerySearchFiltersAudienceFiltersOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

