// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package integrationaws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v13/integrationaws/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/integration_aws datadog_integration_aws}.
type IntegrationAws interface {
	cdktn.TerraformResource
	AccessKeyId() *string
	SetAccessKeyId(val *string)
	AccessKeyIdInput() *string
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	AccountSpecificNamespaceRules() *map[string]interface{}
	SetAccountSpecificNamespaceRules(val *map[string]interface{})
	AccountSpecificNamespaceRulesInput() *map[string]interface{}
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
	CspmResourceCollectionEnabled() *string
	SetCspmResourceCollectionEnabled(val *string)
	CspmResourceCollectionEnabledInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	ExcludedRegions() *[]*string
	SetExcludedRegions(val *[]*string)
	ExcludedRegionsInput() *[]*string
	ExtendedResourceCollectionEnabled() *string
	SetExtendedResourceCollectionEnabled(val *string)
	ExtendedResourceCollectionEnabledInput() *string
	ExternalId() *string
	FilterTags() *[]*string
	SetFilterTags(val *[]*string)
	FilterTagsInput() *[]*string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HostTags() *[]*string
	SetHostTags(val *[]*string)
	HostTagsInput() *[]*string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	MetricsCollectionEnabled() *string
	SetMetricsCollectionEnabled(val *string)
	MetricsCollectionEnabledInput() *string
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
	ResourceCollectionEnabled() *string
	SetResourceCollectionEnabled(val *string)
	ResourceCollectionEnabledInput() *string
	RoleName() *string
	SetRoleName(val *string)
	RoleNameInput() *string
	SecretAccessKey() *string
	SetSecretAccessKey(val *string)
	SecretAccessKeyInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	ResetAccessKeyId()
	ResetAccountId()
	ResetAccountSpecificNamespaceRules()
	ResetCspmResourceCollectionEnabled()
	ResetExcludedRegions()
	ResetExtendedResourceCollectionEnabled()
	ResetFilterTags()
	ResetHostTags()
	ResetId()
	ResetMetricsCollectionEnabled()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetResourceCollectionEnabled()
	ResetRoleName()
	ResetSecretAccessKey()
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

// The jsii proxy struct for IntegrationAws
type jsiiProxy_IntegrationAws struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_IntegrationAws) AccessKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAws) AccessKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAws) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAws) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAws) AccountSpecificNamespaceRules() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"accountSpecificNamespaceRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAws) AccountSpecificNamespaceRulesInput() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"accountSpecificNamespaceRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAws) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAws) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAws) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAws) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAws) CspmResourceCollectionEnabled() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cspmResourceCollectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAws) CspmResourceCollectionEnabledInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cspmResourceCollectionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAws) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAws) ExcludedRegions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedRegions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAws) ExcludedRegionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedRegionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAws) ExtendedResourceCollectionEnabled() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extendedResourceCollectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAws) ExtendedResourceCollectionEnabledInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extendedResourceCollectionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAws) ExternalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAws) FilterTags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"filterTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAws) FilterTagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"filterTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAws) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAws) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAws) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAws) HostTags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hostTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAws) HostTagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hostTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAws) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAws) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAws) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAws) MetricsCollectionEnabled() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsCollectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAws) MetricsCollectionEnabledInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsCollectionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAws) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAws) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAws) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAws) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAws) ResourceCollectionEnabled() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceCollectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAws) ResourceCollectionEnabledInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceCollectionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAws) RoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAws) RoleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAws) SecretAccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAws) SecretAccessKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretAccessKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAws) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAws) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationAws) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/integration_aws datadog_integration_aws} Resource.
func NewIntegrationAws(scope constructs.Construct, id *string, config *IntegrationAwsConfig) IntegrationAws {
	_init_.Initialize()

	if err := validateNewIntegrationAwsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_IntegrationAws{}

	_jsii_.Create(
		"@cdktn/provider-datadog.integrationAws.IntegrationAws",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/integration_aws datadog_integration_aws} Resource.
func NewIntegrationAws_Override(i IntegrationAws, scope constructs.Construct, id *string, config *IntegrationAwsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.integrationAws.IntegrationAws",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IntegrationAws)SetAccessKeyId(val *string) {
	if err := j.validateSetAccessKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessKeyId",
		val,
	)
}

func (j *jsiiProxy_IntegrationAws)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_IntegrationAws)SetAccountSpecificNamespaceRules(val *map[string]interface{}) {
	if err := j.validateSetAccountSpecificNamespaceRulesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountSpecificNamespaceRules",
		val,
	)
}

func (j *jsiiProxy_IntegrationAws)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_IntegrationAws)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IntegrationAws)SetCspmResourceCollectionEnabled(val *string) {
	if err := j.validateSetCspmResourceCollectionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cspmResourceCollectionEnabled",
		val,
	)
}

func (j *jsiiProxy_IntegrationAws)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IntegrationAws)SetExcludedRegions(val *[]*string) {
	if err := j.validateSetExcludedRegionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedRegions",
		val,
	)
}

func (j *jsiiProxy_IntegrationAws)SetExtendedResourceCollectionEnabled(val *string) {
	if err := j.validateSetExtendedResourceCollectionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extendedResourceCollectionEnabled",
		val,
	)
}

func (j *jsiiProxy_IntegrationAws)SetFilterTags(val *[]*string) {
	if err := j.validateSetFilterTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filterTags",
		val,
	)
}

func (j *jsiiProxy_IntegrationAws)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_IntegrationAws)SetHostTags(val *[]*string) {
	if err := j.validateSetHostTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostTags",
		val,
	)
}

func (j *jsiiProxy_IntegrationAws)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_IntegrationAws)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IntegrationAws)SetMetricsCollectionEnabled(val *string) {
	if err := j.validateSetMetricsCollectionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricsCollectionEnabled",
		val,
	)
}

func (j *jsiiProxy_IntegrationAws)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_IntegrationAws)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_IntegrationAws)SetResourceCollectionEnabled(val *string) {
	if err := j.validateSetResourceCollectionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceCollectionEnabled",
		val,
	)
}

func (j *jsiiProxy_IntegrationAws)SetRoleName(val *string) {
	if err := j.validateSetRoleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleName",
		val,
	)
}

func (j *jsiiProxy_IntegrationAws)SetSecretAccessKey(val *string) {
	if err := j.validateSetSecretAccessKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretAccessKey",
		val,
	)
}

// Generates CDKTN code for importing a IntegrationAws resource upon running "cdktn plan <stack-name>".
func IntegrationAws_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateIntegrationAws_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-datadog.integrationAws.IntegrationAws",
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
func IntegrationAws_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIntegrationAws_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-datadog.integrationAws.IntegrationAws",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IntegrationAws_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIntegrationAws_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-datadog.integrationAws.IntegrationAws",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IntegrationAws_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIntegrationAws_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-datadog.integrationAws.IntegrationAws",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IntegrationAws_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-datadog.integrationAws.IntegrationAws",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_IntegrationAws) AddMoveTarget(moveTarget *string) {
	if err := i.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (i *jsiiProxy_IntegrationAws) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_IntegrationAws) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAws) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAws) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAws) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAws) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAws) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAws) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAws) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAws) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAws) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAws) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := i.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (i *jsiiProxy_IntegrationAws) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAws) MoveFromId(id *string) {
	if err := i.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveFromId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IntegrationAws) MoveTo(moveTarget *string, index interface{}) {
	if err := i.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (i *jsiiProxy_IntegrationAws) MoveToId(id *string) {
	if err := i.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveToId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IntegrationAws) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_IntegrationAws) ResetAccessKeyId() {
	_jsii_.InvokeVoid(
		i,
		"resetAccessKeyId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationAws) ResetAccountId() {
	_jsii_.InvokeVoid(
		i,
		"resetAccountId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationAws) ResetAccountSpecificNamespaceRules() {
	_jsii_.InvokeVoid(
		i,
		"resetAccountSpecificNamespaceRules",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationAws) ResetCspmResourceCollectionEnabled() {
	_jsii_.InvokeVoid(
		i,
		"resetCspmResourceCollectionEnabled",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationAws) ResetExcludedRegions() {
	_jsii_.InvokeVoid(
		i,
		"resetExcludedRegions",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationAws) ResetExtendedResourceCollectionEnabled() {
	_jsii_.InvokeVoid(
		i,
		"resetExtendedResourceCollectionEnabled",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationAws) ResetFilterTags() {
	_jsii_.InvokeVoid(
		i,
		"resetFilterTags",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationAws) ResetHostTags() {
	_jsii_.InvokeVoid(
		i,
		"resetHostTags",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationAws) ResetId() {
	_jsii_.InvokeVoid(
		i,
		"resetId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationAws) ResetMetricsCollectionEnabled() {
	_jsii_.InvokeVoid(
		i,
		"resetMetricsCollectionEnabled",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationAws) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationAws) ResetResourceCollectionEnabled() {
	_jsii_.InvokeVoid(
		i,
		"resetResourceCollectionEnabled",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationAws) ResetRoleName() {
	_jsii_.InvokeVoid(
		i,
		"resetRoleName",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationAws) ResetSecretAccessKey() {
	_jsii_.InvokeVoid(
		i,
		"resetSecretAccessKey",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationAws) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAws) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAws) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAws) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAws) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAws) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationAws) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

