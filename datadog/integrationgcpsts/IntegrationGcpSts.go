// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package integrationgcpsts

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v13/integrationgcpsts/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/integration_gcp_sts datadog_integration_gcp_sts}.
type IntegrationGcpSts interface {
	cdktn.TerraformResource
	AccountTags() *[]*string
	SetAccountTags(val *[]*string)
	AccountTagsInput() *[]*string
	Automute() interface{}
	SetAutomute(val interface{})
	AutomuteInput() interface{}
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	ClientEmail() *string
	SetClientEmail(val *string)
	ClientEmailInput() *string
	CloudRunRevisionFilters() *[]*string
	SetCloudRunRevisionFilters(val *[]*string)
	CloudRunRevisionFiltersInput() *[]*string
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
	DelegateAccountEmail() *string
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
	HostFilters() *[]*string
	SetHostFilters(val *[]*string)
	HostFiltersInput() *[]*string
	Id() *string
	IsCspmEnabled() interface{}
	SetIsCspmEnabled(val interface{})
	IsCspmEnabledInput() interface{}
	IsGlobalLocationEnabled() interface{}
	SetIsGlobalLocationEnabled(val interface{})
	IsGlobalLocationEnabledInput() interface{}
	IsPerProjectQuotaEnabled() interface{}
	SetIsPerProjectQuotaEnabled(val interface{})
	IsPerProjectQuotaEnabledInput() interface{}
	IsResourceChangeCollectionEnabled() interface{}
	SetIsResourceChangeCollectionEnabled(val interface{})
	IsResourceChangeCollectionEnabledInput() interface{}
	IsSecurityCommandCenterEnabled() interface{}
	SetIsSecurityCommandCenterEnabled(val interface{})
	IsSecurityCommandCenterEnabledInput() interface{}
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	MetricNamespaceConfigs() IntegrationGcpStsMetricNamespaceConfigsList
	MetricNamespaceConfigsInput() interface{}
	MonitoredResourceConfigs() IntegrationGcpStsMonitoredResourceConfigsList
	MonitoredResourceConfigsInput() interface{}
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
	RegionFilterConfigs() *[]*string
	SetRegionFilterConfigs(val *[]*string)
	RegionFilterConfigsInput() *[]*string
	ResourceCollectionEnabled() interface{}
	SetResourceCollectionEnabled(val interface{})
	ResourceCollectionEnabledInput() interface{}
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
	PutMetricNamespaceConfigs(value interface{})
	PutMonitoredResourceConfigs(value interface{})
	ResetAccountTags()
	ResetAutomute()
	ResetCloudRunRevisionFilters()
	ResetHostFilters()
	ResetIsCspmEnabled()
	ResetIsGlobalLocationEnabled()
	ResetIsPerProjectQuotaEnabled()
	ResetIsResourceChangeCollectionEnabled()
	ResetIsSecurityCommandCenterEnabled()
	ResetMetricNamespaceConfigs()
	ResetMonitoredResourceConfigs()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegionFilterConfigs()
	ResetResourceCollectionEnabled()
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

// The jsii proxy struct for IntegrationGcpSts
type jsiiProxy_IntegrationGcpSts struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_IntegrationGcpSts) AccountTags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accountTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcpSts) AccountTagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accountTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcpSts) Automute() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcpSts) AutomuteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automuteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcpSts) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcpSts) ClientEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcpSts) ClientEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcpSts) CloudRunRevisionFilters() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cloudRunRevisionFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcpSts) CloudRunRevisionFiltersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cloudRunRevisionFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcpSts) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcpSts) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcpSts) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcpSts) DelegateAccountEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delegateAccountEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcpSts) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcpSts) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcpSts) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcpSts) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcpSts) HostFilters() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hostFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcpSts) HostFiltersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hostFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcpSts) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcpSts) IsCspmEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isCspmEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcpSts) IsCspmEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isCspmEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcpSts) IsGlobalLocationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isGlobalLocationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcpSts) IsGlobalLocationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isGlobalLocationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcpSts) IsPerProjectQuotaEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isPerProjectQuotaEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcpSts) IsPerProjectQuotaEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isPerProjectQuotaEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcpSts) IsResourceChangeCollectionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isResourceChangeCollectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcpSts) IsResourceChangeCollectionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isResourceChangeCollectionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcpSts) IsSecurityCommandCenterEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSecurityCommandCenterEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcpSts) IsSecurityCommandCenterEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSecurityCommandCenterEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcpSts) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcpSts) MetricNamespaceConfigs() IntegrationGcpStsMetricNamespaceConfigsList {
	var returns IntegrationGcpStsMetricNamespaceConfigsList
	_jsii_.Get(
		j,
		"metricNamespaceConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcpSts) MetricNamespaceConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricNamespaceConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcpSts) MonitoredResourceConfigs() IntegrationGcpStsMonitoredResourceConfigsList {
	var returns IntegrationGcpStsMonitoredResourceConfigsList
	_jsii_.Get(
		j,
		"monitoredResourceConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcpSts) MonitoredResourceConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"monitoredResourceConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcpSts) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcpSts) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcpSts) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcpSts) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcpSts) RegionFilterConfigs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regionFilterConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcpSts) RegionFilterConfigsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regionFilterConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcpSts) ResourceCollectionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceCollectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcpSts) ResourceCollectionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceCollectionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcpSts) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcpSts) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationGcpSts) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/integration_gcp_sts datadog_integration_gcp_sts} Resource.
func NewIntegrationGcpSts(scope constructs.Construct, id *string, config *IntegrationGcpStsConfig) IntegrationGcpSts {
	_init_.Initialize()

	if err := validateNewIntegrationGcpStsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_IntegrationGcpSts{}

	_jsii_.Create(
		"@cdktn/provider-datadog.integrationGcpSts.IntegrationGcpSts",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/datadog/datadog/3.91.0/docs/resources/integration_gcp_sts datadog_integration_gcp_sts} Resource.
func NewIntegrationGcpSts_Override(i IntegrationGcpSts, scope constructs.Construct, id *string, config *IntegrationGcpStsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.integrationGcpSts.IntegrationGcpSts",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IntegrationGcpSts)SetAccountTags(val *[]*string) {
	if err := j.validateSetAccountTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountTags",
		val,
	)
}

func (j *jsiiProxy_IntegrationGcpSts)SetAutomute(val interface{}) {
	if err := j.validateSetAutomuteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automute",
		val,
	)
}

func (j *jsiiProxy_IntegrationGcpSts)SetClientEmail(val *string) {
	if err := j.validateSetClientEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientEmail",
		val,
	)
}

func (j *jsiiProxy_IntegrationGcpSts)SetCloudRunRevisionFilters(val *[]*string) {
	if err := j.validateSetCloudRunRevisionFiltersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudRunRevisionFilters",
		val,
	)
}

func (j *jsiiProxy_IntegrationGcpSts)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_IntegrationGcpSts)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IntegrationGcpSts)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IntegrationGcpSts)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_IntegrationGcpSts)SetHostFilters(val *[]*string) {
	if err := j.validateSetHostFiltersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostFilters",
		val,
	)
}

func (j *jsiiProxy_IntegrationGcpSts)SetIsCspmEnabled(val interface{}) {
	if err := j.validateSetIsCspmEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isCspmEnabled",
		val,
	)
}

func (j *jsiiProxy_IntegrationGcpSts)SetIsGlobalLocationEnabled(val interface{}) {
	if err := j.validateSetIsGlobalLocationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isGlobalLocationEnabled",
		val,
	)
}

func (j *jsiiProxy_IntegrationGcpSts)SetIsPerProjectQuotaEnabled(val interface{}) {
	if err := j.validateSetIsPerProjectQuotaEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isPerProjectQuotaEnabled",
		val,
	)
}

func (j *jsiiProxy_IntegrationGcpSts)SetIsResourceChangeCollectionEnabled(val interface{}) {
	if err := j.validateSetIsResourceChangeCollectionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isResourceChangeCollectionEnabled",
		val,
	)
}

func (j *jsiiProxy_IntegrationGcpSts)SetIsSecurityCommandCenterEnabled(val interface{}) {
	if err := j.validateSetIsSecurityCommandCenterEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isSecurityCommandCenterEnabled",
		val,
	)
}

func (j *jsiiProxy_IntegrationGcpSts)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IntegrationGcpSts)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_IntegrationGcpSts)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_IntegrationGcpSts)SetRegionFilterConfigs(val *[]*string) {
	if err := j.validateSetRegionFilterConfigsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regionFilterConfigs",
		val,
	)
}

func (j *jsiiProxy_IntegrationGcpSts)SetResourceCollectionEnabled(val interface{}) {
	if err := j.validateSetResourceCollectionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceCollectionEnabled",
		val,
	)
}

// Generates CDKTN code for importing a IntegrationGcpSts resource upon running "cdktn plan <stack-name>".
func IntegrationGcpSts_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateIntegrationGcpSts_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-datadog.integrationGcpSts.IntegrationGcpSts",
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
func IntegrationGcpSts_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIntegrationGcpSts_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-datadog.integrationGcpSts.IntegrationGcpSts",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IntegrationGcpSts_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIntegrationGcpSts_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-datadog.integrationGcpSts.IntegrationGcpSts",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IntegrationGcpSts_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIntegrationGcpSts_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-datadog.integrationGcpSts.IntegrationGcpSts",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IntegrationGcpSts_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-datadog.integrationGcpSts.IntegrationGcpSts",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_IntegrationGcpSts) AddMoveTarget(moveTarget *string) {
	if err := i.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (i *jsiiProxy_IntegrationGcpSts) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_IntegrationGcpSts) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IntegrationGcpSts) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IntegrationGcpSts) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IntegrationGcpSts) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IntegrationGcpSts) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IntegrationGcpSts) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IntegrationGcpSts) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IntegrationGcpSts) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IntegrationGcpSts) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IntegrationGcpSts) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationGcpSts) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := i.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (i *jsiiProxy_IntegrationGcpSts) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IntegrationGcpSts) MoveFromId(id *string) {
	if err := i.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveFromId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IntegrationGcpSts) MoveTo(moveTarget *string, index interface{}) {
	if err := i.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (i *jsiiProxy_IntegrationGcpSts) MoveToId(id *string) {
	if err := i.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveToId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IntegrationGcpSts) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_IntegrationGcpSts) PutMetricNamespaceConfigs(value interface{}) {
	if err := i.validatePutMetricNamespaceConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putMetricNamespaceConfigs",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationGcpSts) PutMonitoredResourceConfigs(value interface{}) {
	if err := i.validatePutMonitoredResourceConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putMonitoredResourceConfigs",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationGcpSts) ResetAccountTags() {
	_jsii_.InvokeVoid(
		i,
		"resetAccountTags",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationGcpSts) ResetAutomute() {
	_jsii_.InvokeVoid(
		i,
		"resetAutomute",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationGcpSts) ResetCloudRunRevisionFilters() {
	_jsii_.InvokeVoid(
		i,
		"resetCloudRunRevisionFilters",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationGcpSts) ResetHostFilters() {
	_jsii_.InvokeVoid(
		i,
		"resetHostFilters",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationGcpSts) ResetIsCspmEnabled() {
	_jsii_.InvokeVoid(
		i,
		"resetIsCspmEnabled",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationGcpSts) ResetIsGlobalLocationEnabled() {
	_jsii_.InvokeVoid(
		i,
		"resetIsGlobalLocationEnabled",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationGcpSts) ResetIsPerProjectQuotaEnabled() {
	_jsii_.InvokeVoid(
		i,
		"resetIsPerProjectQuotaEnabled",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationGcpSts) ResetIsResourceChangeCollectionEnabled() {
	_jsii_.InvokeVoid(
		i,
		"resetIsResourceChangeCollectionEnabled",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationGcpSts) ResetIsSecurityCommandCenterEnabled() {
	_jsii_.InvokeVoid(
		i,
		"resetIsSecurityCommandCenterEnabled",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationGcpSts) ResetMetricNamespaceConfigs() {
	_jsii_.InvokeVoid(
		i,
		"resetMetricNamespaceConfigs",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationGcpSts) ResetMonitoredResourceConfigs() {
	_jsii_.InvokeVoid(
		i,
		"resetMonitoredResourceConfigs",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationGcpSts) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationGcpSts) ResetRegionFilterConfigs() {
	_jsii_.InvokeVoid(
		i,
		"resetRegionFilterConfigs",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationGcpSts) ResetResourceCollectionEnabled() {
	_jsii_.InvokeVoid(
		i,
		"resetResourceCollectionEnabled",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationGcpSts) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationGcpSts) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationGcpSts) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationGcpSts) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationGcpSts) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationGcpSts) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationGcpSts) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

