// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package actionconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-datadog-go/datadog/v15/actionconnection/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/action_connection datadog_action_connection}.
type ActionConnection interface {
	cdktn.TerraformResource
	Anthropic() ActionConnectionAnthropicOutputReference
	AnthropicInput() interface{}
	Asana() ActionConnectionAsanaOutputReference
	AsanaInput() interface{}
	Aws() ActionConnectionAwsOutputReference
	AwsInput() interface{}
	Azure() ActionConnectionAzureOutputReference
	AzureInput() interface{}
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	CircleCi() ActionConnectionCircleCiOutputReference
	CircleCiInput() interface{}
	Clickup() ActionConnectionClickupOutputReference
	ClickupInput() interface{}
	Cloudflare() ActionConnectionCloudflareOutputReference
	CloudflareInput() interface{}
	ConfigCat() ActionConnectionConfigCatOutputReference
	ConfigCatInput() interface{}
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
	Datadog() ActionConnectionDatadogOutputReference
	DatadogInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Fastly() ActionConnectionFastlyOutputReference
	FastlyInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	Freshservice() ActionConnectionFreshserviceOutputReference
	FreshserviceInput() interface{}
	// Experimental.
	FriendlyUniqueId() *string
	Gcp() ActionConnectionGcpOutputReference
	GcpInput() interface{}
	Gemini() ActionConnectionGeminiOutputReference
	GeminiInput() interface{}
	Gitlab() ActionConnectionGitlabOutputReference
	GitlabInput() interface{}
	GreyNoise() ActionConnectionGreyNoiseOutputReference
	GreyNoiseInput() interface{}
	Http() ActionConnectionHttpOutputReference
	HttpInput() interface{}
	Id() *string
	LaunchDarkly() ActionConnectionLaunchDarklyOutputReference
	LaunchDarklyInput() interface{}
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Notion() ActionConnectionNotionOutputReference
	NotionInput() interface{}
	Okta() ActionConnectionOktaOutputReference
	OktaInput() interface{}
	Openai() ActionConnectionOpenaiOutputReference
	OpenaiInput() interface{}
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
	ServiceNow() ActionConnectionServiceNowOutputReference
	ServiceNowInput() interface{}
	Split() ActionConnectionSplitOutputReference
	SplitInput() interface{}
	Statsig() ActionConnectionStatsigOutputReference
	StatsigInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	VirusTotal() ActionConnectionVirusTotalOutputReference
	VirusTotalInput() interface{}
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
	PutAnthropic(value *ActionConnectionAnthropic)
	PutAsana(value *ActionConnectionAsana)
	PutAws(value *ActionConnectionAws)
	PutAzure(value *ActionConnectionAzure)
	PutCircleCi(value *ActionConnectionCircleCi)
	PutClickup(value *ActionConnectionClickup)
	PutCloudflare(value *ActionConnectionCloudflare)
	PutConfigCat(value *ActionConnectionConfigCat)
	PutDatadog(value *ActionConnectionDatadog)
	PutFastly(value *ActionConnectionFastly)
	PutFreshservice(value *ActionConnectionFreshservice)
	PutGcp(value *ActionConnectionGcp)
	PutGemini(value *ActionConnectionGemini)
	PutGitlab(value *ActionConnectionGitlab)
	PutGreyNoise(value *ActionConnectionGreyNoise)
	PutHttp(value *ActionConnectionHttp)
	PutLaunchDarkly(value *ActionConnectionLaunchDarkly)
	PutNotion(value *ActionConnectionNotion)
	PutOkta(value *ActionConnectionOkta)
	PutOpenai(value *ActionConnectionOpenai)
	PutServiceNow(value *ActionConnectionServiceNow)
	PutSplit(value *ActionConnectionSplit)
	PutStatsig(value *ActionConnectionStatsig)
	PutVirusTotal(value *ActionConnectionVirusTotal)
	ResetAnthropic()
	ResetAsana()
	ResetAws()
	ResetAzure()
	ResetCircleCi()
	ResetClickup()
	ResetCloudflare()
	ResetConfigCat()
	ResetDatadog()
	ResetFastly()
	ResetFreshservice()
	ResetGcp()
	ResetGemini()
	ResetGitlab()
	ResetGreyNoise()
	ResetHttp()
	ResetLaunchDarkly()
	ResetNotion()
	ResetOkta()
	ResetOpenai()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetServiceNow()
	ResetSplit()
	ResetStatsig()
	ResetVirusTotal()
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

// The jsii proxy struct for ActionConnection
type jsiiProxy_ActionConnection struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_ActionConnection) Anthropic() ActionConnectionAnthropicOutputReference {
	var returns ActionConnectionAnthropicOutputReference
	_jsii_.Get(
		j,
		"anthropic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) AnthropicInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"anthropicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) Asana() ActionConnectionAsanaOutputReference {
	var returns ActionConnectionAsanaOutputReference
	_jsii_.Get(
		j,
		"asana",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) AsanaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"asanaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) Aws() ActionConnectionAwsOutputReference {
	var returns ActionConnectionAwsOutputReference
	_jsii_.Get(
		j,
		"aws",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) AwsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) Azure() ActionConnectionAzureOutputReference {
	var returns ActionConnectionAzureOutputReference
	_jsii_.Get(
		j,
		"azure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) AzureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) CircleCi() ActionConnectionCircleCiOutputReference {
	var returns ActionConnectionCircleCiOutputReference
	_jsii_.Get(
		j,
		"circleCi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) CircleCiInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"circleCiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) Clickup() ActionConnectionClickupOutputReference {
	var returns ActionConnectionClickupOutputReference
	_jsii_.Get(
		j,
		"clickup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) ClickupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clickupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) Cloudflare() ActionConnectionCloudflareOutputReference {
	var returns ActionConnectionCloudflareOutputReference
	_jsii_.Get(
		j,
		"cloudflare",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) CloudflareInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudflareInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) ConfigCat() ActionConnectionConfigCatOutputReference {
	var returns ActionConnectionConfigCatOutputReference
	_jsii_.Get(
		j,
		"configCat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) ConfigCatInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configCatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) Datadog() ActionConnectionDatadogOutputReference {
	var returns ActionConnectionDatadogOutputReference
	_jsii_.Get(
		j,
		"datadog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) DatadogInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"datadogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) Fastly() ActionConnectionFastlyOutputReference {
	var returns ActionConnectionFastlyOutputReference
	_jsii_.Get(
		j,
		"fastly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) FastlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fastlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) Freshservice() ActionConnectionFreshserviceOutputReference {
	var returns ActionConnectionFreshserviceOutputReference
	_jsii_.Get(
		j,
		"freshservice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) FreshserviceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"freshserviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) Gcp() ActionConnectionGcpOutputReference {
	var returns ActionConnectionGcpOutputReference
	_jsii_.Get(
		j,
		"gcp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) GcpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gcpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) Gemini() ActionConnectionGeminiOutputReference {
	var returns ActionConnectionGeminiOutputReference
	_jsii_.Get(
		j,
		"gemini",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) GeminiInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"geminiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) Gitlab() ActionConnectionGitlabOutputReference {
	var returns ActionConnectionGitlabOutputReference
	_jsii_.Get(
		j,
		"gitlab",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) GitlabInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gitlabInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) GreyNoise() ActionConnectionGreyNoiseOutputReference {
	var returns ActionConnectionGreyNoiseOutputReference
	_jsii_.Get(
		j,
		"greyNoise",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) GreyNoiseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"greyNoiseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) Http() ActionConnectionHttpOutputReference {
	var returns ActionConnectionHttpOutputReference
	_jsii_.Get(
		j,
		"http",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) HttpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) LaunchDarkly() ActionConnectionLaunchDarklyOutputReference {
	var returns ActionConnectionLaunchDarklyOutputReference
	_jsii_.Get(
		j,
		"launchDarkly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) LaunchDarklyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"launchDarklyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) Notion() ActionConnectionNotionOutputReference {
	var returns ActionConnectionNotionOutputReference
	_jsii_.Get(
		j,
		"notion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) NotionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) Okta() ActionConnectionOktaOutputReference {
	var returns ActionConnectionOktaOutputReference
	_jsii_.Get(
		j,
		"okta",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) OktaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oktaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) Openai() ActionConnectionOpenaiOutputReference {
	var returns ActionConnectionOpenaiOutputReference
	_jsii_.Get(
		j,
		"openai",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) OpenaiInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"openaiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) ServiceNow() ActionConnectionServiceNowOutputReference {
	var returns ActionConnectionServiceNowOutputReference
	_jsii_.Get(
		j,
		"serviceNow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) ServiceNowInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceNowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) Split() ActionConnectionSplitOutputReference {
	var returns ActionConnectionSplitOutputReference
	_jsii_.Get(
		j,
		"split",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) SplitInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"splitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) Statsig() ActionConnectionStatsigOutputReference {
	var returns ActionConnectionStatsigOutputReference
	_jsii_.Get(
		j,
		"statsig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) StatsigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"statsigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) VirusTotal() ActionConnectionVirusTotalOutputReference {
	var returns ActionConnectionVirusTotalOutputReference
	_jsii_.Get(
		j,
		"virusTotal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionConnection) VirusTotalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"virusTotalInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/action_connection datadog_action_connection} Resource.
func NewActionConnection(scope constructs.Construct, id *string, config *ActionConnectionConfig) ActionConnection {
	_init_.Initialize()

	if err := validateNewActionConnectionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ActionConnection{}

	_jsii_.Create(
		"@cdktn/provider-datadog.actionConnection.ActionConnection",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/datadog/datadog/4.16.0/docs/resources/action_connection datadog_action_connection} Resource.
func NewActionConnection_Override(a ActionConnection, scope constructs.Construct, id *string, config *ActionConnectionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-datadog.actionConnection.ActionConnection",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ActionConnection)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ActionConnection)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ActionConnection)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ActionConnection)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ActionConnection)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ActionConnection)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ActionConnection)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ActionConnection)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTN code for importing a ActionConnection resource upon running "cdktn plan <stack-name>".
func ActionConnection_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateActionConnection_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-datadog.actionConnection.ActionConnection",
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
func ActionConnection_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateActionConnection_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-datadog.actionConnection.ActionConnection",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ActionConnection_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateActionConnection_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-datadog.actionConnection.ActionConnection",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ActionConnection_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateActionConnection_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-datadog.actionConnection.ActionConnection",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ActionConnection_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-datadog.actionConnection.ActionConnection",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_ActionConnection) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_ActionConnection) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_ActionConnection) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActionConnection) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActionConnection) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActionConnection) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActionConnection) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActionConnection) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActionConnection) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActionConnection) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActionConnection) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActionConnection) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActionConnection) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_ActionConnection) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActionConnection) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ActionConnection) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_ActionConnection) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ActionConnection) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ActionConnection) PutAnthropic(value *ActionConnectionAnthropic) {
	if err := a.validatePutAnthropicParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAnthropic",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ActionConnection) PutAsana(value *ActionConnectionAsana) {
	if err := a.validatePutAsanaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAsana",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ActionConnection) PutAws(value *ActionConnectionAws) {
	if err := a.validatePutAwsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAws",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ActionConnection) PutAzure(value *ActionConnectionAzure) {
	if err := a.validatePutAzureParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAzure",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ActionConnection) PutCircleCi(value *ActionConnectionCircleCi) {
	if err := a.validatePutCircleCiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCircleCi",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ActionConnection) PutClickup(value *ActionConnectionClickup) {
	if err := a.validatePutClickupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putClickup",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ActionConnection) PutCloudflare(value *ActionConnectionCloudflare) {
	if err := a.validatePutCloudflareParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCloudflare",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ActionConnection) PutConfigCat(value *ActionConnectionConfigCat) {
	if err := a.validatePutConfigCatParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putConfigCat",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ActionConnection) PutDatadog(value *ActionConnectionDatadog) {
	if err := a.validatePutDatadogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDatadog",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ActionConnection) PutFastly(value *ActionConnectionFastly) {
	if err := a.validatePutFastlyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFastly",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ActionConnection) PutFreshservice(value *ActionConnectionFreshservice) {
	if err := a.validatePutFreshserviceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFreshservice",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ActionConnection) PutGcp(value *ActionConnectionGcp) {
	if err := a.validatePutGcpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGcp",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ActionConnection) PutGemini(value *ActionConnectionGemini) {
	if err := a.validatePutGeminiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGemini",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ActionConnection) PutGitlab(value *ActionConnectionGitlab) {
	if err := a.validatePutGitlabParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGitlab",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ActionConnection) PutGreyNoise(value *ActionConnectionGreyNoise) {
	if err := a.validatePutGreyNoiseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGreyNoise",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ActionConnection) PutHttp(value *ActionConnectionHttp) {
	if err := a.validatePutHttpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHttp",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ActionConnection) PutLaunchDarkly(value *ActionConnectionLaunchDarkly) {
	if err := a.validatePutLaunchDarklyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLaunchDarkly",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ActionConnection) PutNotion(value *ActionConnectionNotion) {
	if err := a.validatePutNotionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNotion",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ActionConnection) PutOkta(value *ActionConnectionOkta) {
	if err := a.validatePutOktaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOkta",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ActionConnection) PutOpenai(value *ActionConnectionOpenai) {
	if err := a.validatePutOpenaiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOpenai",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ActionConnection) PutServiceNow(value *ActionConnectionServiceNow) {
	if err := a.validatePutServiceNowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putServiceNow",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ActionConnection) PutSplit(value *ActionConnectionSplit) {
	if err := a.validatePutSplitParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSplit",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ActionConnection) PutStatsig(value *ActionConnectionStatsig) {
	if err := a.validatePutStatsigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStatsig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ActionConnection) PutVirusTotal(value *ActionConnectionVirusTotal) {
	if err := a.validatePutVirusTotalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVirusTotal",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ActionConnection) ResetAnthropic() {
	_jsii_.InvokeVoid(
		a,
		"resetAnthropic",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActionConnection) ResetAsana() {
	_jsii_.InvokeVoid(
		a,
		"resetAsana",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActionConnection) ResetAws() {
	_jsii_.InvokeVoid(
		a,
		"resetAws",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActionConnection) ResetAzure() {
	_jsii_.InvokeVoid(
		a,
		"resetAzure",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActionConnection) ResetCircleCi() {
	_jsii_.InvokeVoid(
		a,
		"resetCircleCi",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActionConnection) ResetClickup() {
	_jsii_.InvokeVoid(
		a,
		"resetClickup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActionConnection) ResetCloudflare() {
	_jsii_.InvokeVoid(
		a,
		"resetCloudflare",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActionConnection) ResetConfigCat() {
	_jsii_.InvokeVoid(
		a,
		"resetConfigCat",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActionConnection) ResetDatadog() {
	_jsii_.InvokeVoid(
		a,
		"resetDatadog",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActionConnection) ResetFastly() {
	_jsii_.InvokeVoid(
		a,
		"resetFastly",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActionConnection) ResetFreshservice() {
	_jsii_.InvokeVoid(
		a,
		"resetFreshservice",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActionConnection) ResetGcp() {
	_jsii_.InvokeVoid(
		a,
		"resetGcp",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActionConnection) ResetGemini() {
	_jsii_.InvokeVoid(
		a,
		"resetGemini",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActionConnection) ResetGitlab() {
	_jsii_.InvokeVoid(
		a,
		"resetGitlab",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActionConnection) ResetGreyNoise() {
	_jsii_.InvokeVoid(
		a,
		"resetGreyNoise",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActionConnection) ResetHttp() {
	_jsii_.InvokeVoid(
		a,
		"resetHttp",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActionConnection) ResetLaunchDarkly() {
	_jsii_.InvokeVoid(
		a,
		"resetLaunchDarkly",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActionConnection) ResetNotion() {
	_jsii_.InvokeVoid(
		a,
		"resetNotion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActionConnection) ResetOkta() {
	_jsii_.InvokeVoid(
		a,
		"resetOkta",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActionConnection) ResetOpenai() {
	_jsii_.InvokeVoid(
		a,
		"resetOpenai",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActionConnection) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActionConnection) ResetServiceNow() {
	_jsii_.InvokeVoid(
		a,
		"resetServiceNow",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActionConnection) ResetSplit() {
	_jsii_.InvokeVoid(
		a,
		"resetSplit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActionConnection) ResetStatsig() {
	_jsii_.InvokeVoid(
		a,
		"resetStatsig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActionConnection) ResetVirusTotal() {
	_jsii_.InvokeVoid(
		a,
		"resetVirusTotal",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActionConnection) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActionConnection) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActionConnection) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActionConnection) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActionConnection) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActionConnection) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActionConnection) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		a,
		"with",
		args,
		&returns,
	)

	return returns
}

