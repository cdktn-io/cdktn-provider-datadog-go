// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dashboardv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/dashboardv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/dashboard_v2 datadog_dashboard_v2}.
type DashboardV2 interface {
	cdktn.TerraformResource
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DashboardLists() *[]*float64
	SetDashboardLists(val *[]*float64)
	DashboardListsInput() *[]*float64
	DashboardListsRemoved() *[]*float64
	SetDashboardListsRemoved(val *[]*float64)
	DashboardListsRemovedInput() *[]*float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IsReadOnly() interface{}
	SetIsReadOnly(val interface{})
	IsReadOnlyInput() interface{}
	LayoutType() *string
	SetLayoutType(val *string)
	LayoutTypeInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	NotifyList() *[]*string
	SetNotifyList(val *[]*string)
	NotifyListInput() *[]*string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	ReflowType() *string
	SetReflowType(val *string)
	ReflowTypeInput() *string
	RestrictedRoles() *[]*string
	SetRestrictedRoles(val *[]*string)
	RestrictedRolesInput() *[]*string
	Tab() DashboardV2TabList
	TabInput() interface{}
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
	TemplateVariable() DashboardV2TemplateVariableList
	TemplateVariableInput() interface{}
	TemplateVariablePreset() DashboardV2TemplateVariablePresetList
	TemplateVariablePresetInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Title() *string
	SetTitle(val *string)
	TitleInput() *string
	Url() *string
	SetUrl(val *string)
	UrlInput() *string
	Widget() DashboardV2WidgetList
	WidgetInput() interface{}
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktn.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutTab(value interface{})
	PutTemplateVariable(value interface{})
	PutTemplateVariablePreset(value interface{})
	PutWidget(value interface{})
	ResetDashboardLists()
	ResetDashboardListsRemoved()
	ResetDescription()
	ResetId()
	ResetIsReadOnly()
	ResetNotifyList()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetReflowType()
	ResetRestrictedRoles()
	ResetTab()
	ResetTags()
	ResetTemplateVariable()
	ResetTemplateVariablePreset()
	ResetUrl()
	ResetWidget()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for DashboardV2
type jsiiProxy_DashboardV2 struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_DashboardV2) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2) DashboardLists() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"dashboardLists",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2) DashboardListsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"dashboardListsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2) DashboardListsRemoved() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"dashboardListsRemoved",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2) DashboardListsRemovedInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"dashboardListsRemovedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2) IsReadOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isReadOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2) IsReadOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isReadOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2) LayoutType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"layoutType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2) LayoutTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"layoutTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2) NotifyList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notifyList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2) NotifyListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notifyListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2) ReflowType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reflowType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2) ReflowTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reflowTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2) RestrictedRoles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"restrictedRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2) RestrictedRolesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"restrictedRolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2) Tab() DashboardV2TabList {
	var returns DashboardV2TabList
	_jsii_.Get(
		j,
		"tab",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2) TabInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tabInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2) TemplateVariable() DashboardV2TemplateVariableList {
	var returns DashboardV2TemplateVariableList
	_jsii_.Get(
		j,
		"templateVariable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2) TemplateVariableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"templateVariableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2) TemplateVariablePreset() DashboardV2TemplateVariablePresetList {
	var returns DashboardV2TemplateVariablePresetList
	_jsii_.Get(
		j,
		"templateVariablePreset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2) TemplateVariablePresetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"templateVariablePresetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2) Widget() DashboardV2WidgetList {
	var returns DashboardV2WidgetList
	_jsii_.Get(
		j,
		"widget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DashboardV2) WidgetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"widgetInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/dashboard_v2 datadog_dashboard_v2} Resource.
func NewDashboardV2(scope constructs.Construct, id *string, config *DashboardV2Config) DashboardV2 {
	_init_.Initialize()

	if err := validateNewDashboardV2Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DashboardV2{}

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/dashboard_v2 datadog_dashboard_v2} Resource.
func NewDashboardV2_Override(d DashboardV2, scope constructs.Construct, id *string, config *DashboardV2Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DashboardV2)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DashboardV2)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DashboardV2)SetDashboardLists(val *[]*float64) {
	if err := j.validateSetDashboardListsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dashboardLists",
		val,
	)
}

func (j *jsiiProxy_DashboardV2)SetDashboardListsRemoved(val *[]*float64) {
	if err := j.validateSetDashboardListsRemovedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dashboardListsRemoved",
		val,
	)
}

func (j *jsiiProxy_DashboardV2)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DashboardV2)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DashboardV2)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DashboardV2)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DashboardV2)SetIsReadOnly(val interface{}) {
	if err := j.validateSetIsReadOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isReadOnly",
		val,
	)
}

func (j *jsiiProxy_DashboardV2)SetLayoutType(val *string) {
	if err := j.validateSetLayoutTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"layoutType",
		val,
	)
}

func (j *jsiiProxy_DashboardV2)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DashboardV2)SetNotifyList(val *[]*string) {
	if err := j.validateSetNotifyListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notifyList",
		val,
	)
}

func (j *jsiiProxy_DashboardV2)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DashboardV2)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DashboardV2)SetReflowType(val *string) {
	if err := j.validateSetReflowTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reflowType",
		val,
	)
}

func (j *jsiiProxy_DashboardV2)SetRestrictedRoles(val *[]*string) {
	if err := j.validateSetRestrictedRolesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restrictedRoles",
		val,
	)
}

func (j *jsiiProxy_DashboardV2)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_DashboardV2)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (j *jsiiProxy_DashboardV2)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

// Generates CDKTN code for importing a DashboardV2 resource upon running "cdktn plan <stack-name>".
func DashboardV2_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDashboardV2_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func DashboardV2_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDashboardV2_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DashboardV2_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDashboardV2_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DashboardV2_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDashboardV2_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DashboardV2_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-datadog.dashboardV2.DashboardV2",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DashboardV2) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DashboardV2) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DashboardV2) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DashboardV2) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DashboardV2) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DashboardV2) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DashboardV2) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DashboardV2) PutTab(value interface{}) {
	if err := d.validatePutTabParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTab",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2) PutTemplateVariable(value interface{}) {
	if err := d.validatePutTemplateVariableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTemplateVariable",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2) PutTemplateVariablePreset(value interface{}) {
	if err := d.validatePutTemplateVariablePresetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTemplateVariablePreset",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2) PutWidget(value interface{}) {
	if err := d.validatePutWidgetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putWidget",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DashboardV2) ResetDashboardLists() {
	_jsii_.InvokeVoid(
		d,
		"resetDashboardLists",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2) ResetDashboardListsRemoved() {
	_jsii_.InvokeVoid(
		d,
		"resetDashboardListsRemoved",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2) ResetIsReadOnly() {
	_jsii_.InvokeVoid(
		d,
		"resetIsReadOnly",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2) ResetNotifyList() {
	_jsii_.InvokeVoid(
		d,
		"resetNotifyList",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2) ResetReflowType() {
	_jsii_.InvokeVoid(
		d,
		"resetReflowType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2) ResetRestrictedRoles() {
	_jsii_.InvokeVoid(
		d,
		"resetRestrictedRoles",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2) ResetTab() {
	_jsii_.InvokeVoid(
		d,
		"resetTab",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2) ResetTemplateVariable() {
	_jsii_.InvokeVoid(
		d,
		"resetTemplateVariable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2) ResetTemplateVariablePreset() {
	_jsii_.InvokeVoid(
		d,
		"resetTemplateVariablePreset",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2) ResetUrl() {
	_jsii_.InvokeVoid(
		d,
		"resetUrl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2) ResetWidget() {
	_jsii_.InvokeVoid(
		d,
		"resetWidget",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DashboardV2) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DashboardV2) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		d,
		"with",
		args,
		&returns,
	)

	return returns
}

