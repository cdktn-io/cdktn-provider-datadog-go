// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package secureembeddashboard

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/secureembeddashboard/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/secure_embed_dashboard datadog_secure_embed_dashboard}.
type SecureEmbedDashboard interface {
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
	Credential() *string
	DashboardId() *string
	SetDashboardId(val *string)
	DashboardIdInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GlobalTimeLiveSpan() *string
	SetGlobalTimeLiveSpan(val *string)
	GlobalTimeLiveSpanInput() *string
	GlobalTimeSelectable() interface{}
	SetGlobalTimeSelectable(val interface{})
	GlobalTimeSelectableInput() interface{}
	Id() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
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
	SelectableTemplateVars() SecureEmbedDashboardSelectableTemplateVarsList
	SelectableTemplateVarsInput() interface{}
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Title() *string
	SetTitle(val *string)
	TitleInput() *string
	Token() *string
	Url() *string
	ViewingPreferencesHighDensity() interface{}
	SetViewingPreferencesHighDensity(val interface{})
	ViewingPreferencesHighDensityInput() interface{}
	ViewingPreferencesTheme() *string
	SetViewingPreferencesTheme(val *string)
	ViewingPreferencesThemeInput() *string
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
	PutSelectableTemplateVars(value interface{})
	ResetGlobalTimeLiveSpan()
	ResetGlobalTimeSelectable()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSelectableTemplateVars()
	ResetStatus()
	ResetViewingPreferencesHighDensity()
	ResetViewingPreferencesTheme()
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

// The jsii proxy struct for SecureEmbedDashboard
type jsiiProxy_SecureEmbedDashboard struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_SecureEmbedDashboard) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureEmbedDashboard) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureEmbedDashboard) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureEmbedDashboard) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureEmbedDashboard) Credential() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureEmbedDashboard) DashboardId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dashboardId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureEmbedDashboard) DashboardIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dashboardIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureEmbedDashboard) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureEmbedDashboard) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureEmbedDashboard) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureEmbedDashboard) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureEmbedDashboard) GlobalTimeLiveSpan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalTimeLiveSpan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureEmbedDashboard) GlobalTimeLiveSpanInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalTimeLiveSpanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureEmbedDashboard) GlobalTimeSelectable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"globalTimeSelectable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureEmbedDashboard) GlobalTimeSelectableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"globalTimeSelectableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureEmbedDashboard) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureEmbedDashboard) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureEmbedDashboard) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureEmbedDashboard) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureEmbedDashboard) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureEmbedDashboard) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureEmbedDashboard) SelectableTemplateVars() SecureEmbedDashboardSelectableTemplateVarsList {
	var returns SecureEmbedDashboardSelectableTemplateVarsList
	_jsii_.Get(
		j,
		"selectableTemplateVars",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureEmbedDashboard) SelectableTemplateVarsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"selectableTemplateVarsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureEmbedDashboard) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureEmbedDashboard) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureEmbedDashboard) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureEmbedDashboard) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureEmbedDashboard) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureEmbedDashboard) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureEmbedDashboard) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureEmbedDashboard) Token() *string {
	var returns *string
	_jsii_.Get(
		j,
		"token",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureEmbedDashboard) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureEmbedDashboard) ViewingPreferencesHighDensity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"viewingPreferencesHighDensity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureEmbedDashboard) ViewingPreferencesHighDensityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"viewingPreferencesHighDensityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureEmbedDashboard) ViewingPreferencesTheme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewingPreferencesTheme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecureEmbedDashboard) ViewingPreferencesThemeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewingPreferencesThemeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/secure_embed_dashboard datadog_secure_embed_dashboard} Resource.
func NewSecureEmbedDashboard(scope constructs.Construct, id *string, config *SecureEmbedDashboardConfig) SecureEmbedDashboard {
	_init_.Initialize()

	if err := validateNewSecureEmbedDashboardParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecureEmbedDashboard{}

	_jsii_.Create(
		"@cdktn/provider-datadog.secureEmbedDashboard.SecureEmbedDashboard",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/datadog/datadog/4.11.0/docs/resources/secure_embed_dashboard datadog_secure_embed_dashboard} Resource.
func NewSecureEmbedDashboard_Override(s SecureEmbedDashboard, scope constructs.Construct, id *string, config *SecureEmbedDashboardConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.secureEmbedDashboard.SecureEmbedDashboard",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SecureEmbedDashboard)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SecureEmbedDashboard)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SecureEmbedDashboard)SetDashboardId(val *string) {
	if err := j.validateSetDashboardIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dashboardId",
		val,
	)
}

func (j *jsiiProxy_SecureEmbedDashboard)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SecureEmbedDashboard)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SecureEmbedDashboard)SetGlobalTimeLiveSpan(val *string) {
	if err := j.validateSetGlobalTimeLiveSpanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"globalTimeLiveSpan",
		val,
	)
}

func (j *jsiiProxy_SecureEmbedDashboard)SetGlobalTimeSelectable(val interface{}) {
	if err := j.validateSetGlobalTimeSelectableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"globalTimeSelectable",
		val,
	)
}

func (j *jsiiProxy_SecureEmbedDashboard)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SecureEmbedDashboard)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SecureEmbedDashboard)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SecureEmbedDashboard)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_SecureEmbedDashboard)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (j *jsiiProxy_SecureEmbedDashboard)SetViewingPreferencesHighDensity(val interface{}) {
	if err := j.validateSetViewingPreferencesHighDensityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"viewingPreferencesHighDensity",
		val,
	)
}

func (j *jsiiProxy_SecureEmbedDashboard)SetViewingPreferencesTheme(val *string) {
	if err := j.validateSetViewingPreferencesThemeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"viewingPreferencesTheme",
		val,
	)
}

// Generates CDKTN code for importing a SecureEmbedDashboard resource upon running "cdktn plan <stack-name>".
func SecureEmbedDashboard_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateSecureEmbedDashboard_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-datadog.secureEmbedDashboard.SecureEmbedDashboard",
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
func SecureEmbedDashboard_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSecureEmbedDashboard_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-datadog.secureEmbedDashboard.SecureEmbedDashboard",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SecureEmbedDashboard_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSecureEmbedDashboard_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-datadog.secureEmbedDashboard.SecureEmbedDashboard",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SecureEmbedDashboard_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSecureEmbedDashboard_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-datadog.secureEmbedDashboard.SecureEmbedDashboard",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SecureEmbedDashboard_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-datadog.secureEmbedDashboard.SecureEmbedDashboard",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SecureEmbedDashboard) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SecureEmbedDashboard) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SecureEmbedDashboard) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecureEmbedDashboard) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecureEmbedDashboard) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecureEmbedDashboard) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecureEmbedDashboard) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecureEmbedDashboard) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecureEmbedDashboard) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecureEmbedDashboard) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecureEmbedDashboard) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecureEmbedDashboard) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecureEmbedDashboard) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SecureEmbedDashboard) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecureEmbedDashboard) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SecureEmbedDashboard) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SecureEmbedDashboard) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SecureEmbedDashboard) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SecureEmbedDashboard) PutSelectableTemplateVars(value interface{}) {
	if err := s.validatePutSelectableTemplateVarsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSelectableTemplateVars",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecureEmbedDashboard) ResetGlobalTimeLiveSpan() {
	_jsii_.InvokeVoid(
		s,
		"resetGlobalTimeLiveSpan",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecureEmbedDashboard) ResetGlobalTimeSelectable() {
	_jsii_.InvokeVoid(
		s,
		"resetGlobalTimeSelectable",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecureEmbedDashboard) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecureEmbedDashboard) ResetSelectableTemplateVars() {
	_jsii_.InvokeVoid(
		s,
		"resetSelectableTemplateVars",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecureEmbedDashboard) ResetStatus() {
	_jsii_.InvokeVoid(
		s,
		"resetStatus",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecureEmbedDashboard) ResetViewingPreferencesHighDensity() {
	_jsii_.InvokeVoid(
		s,
		"resetViewingPreferencesHighDensity",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecureEmbedDashboard) ResetViewingPreferencesTheme() {
	_jsii_.InvokeVoid(
		s,
		"resetViewingPreferencesTheme",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecureEmbedDashboard) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecureEmbedDashboard) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecureEmbedDashboard) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecureEmbedDashboard) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecureEmbedDashboard) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecureEmbedDashboard) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecureEmbedDashboard) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		s,
		"with",
		args,
		&returns,
	)

	return returns
}

