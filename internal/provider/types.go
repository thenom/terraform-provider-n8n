package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// workflowList map the response from GetWorkflows
type workflowListDataSourceModel struct {
	Data *[]workflowDataSourceModel `tfsdk:"data"`
}

// workflowDataSourceModel maps the data source schema data and the API response.
type workflowDataSourceModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Active      types.Bool   `tfsdk:"active"`
	Nodes       []node       `tfsdk:"nodes"`
	Connections types.Map    `tfsdk:"connections"`
	Settings    settings     `tfsdk:"settings"`
	StaticData  types.String `tfsdk:"static_data"`
	Tags        []tag        `tfsdk:"tags"`
	CreatedAt   types.String `tfsdk:"created_at"`
	UpdatedAt   types.String `tfsdk:"updated_at"`
}

// node represents a node in a workflow.
type node struct {
	ID               types.String  `tfsdk:"id"`
	Name             types.String  `tfsdk:"name"`
	WebhookID        types.String  `tfsdk:"webhook_id"`
	Disabled         types.Bool    `tfsdk:"disabled"`
	NotesInFlow      types.Bool    `tfsdk:"notes_in_flow"`
	Notes            types.String  `tfsdk:"notes"`
	Type             types.String  `tfsdk:"type"`
	TypeVersion      types.Float32 `tfsdk:"type_version"`
	ExecuteOnce      types.Bool    `tfsdk:"execute_once"`
	AlwaysOutputData types.Bool    `tfsdk:"always_output_data"`
	RetryOnFail      types.Bool    `tfsdk:"retry_on_fail"`
	MaxTries         types.Float32 `tfsdk:"max_tries"`
	WaitBetweenTries types.Float32 `tfsdk:"wait_between_tries"`
	ContinueOnFail   types.Bool    `tfsdk:"continue_on_fail"`
	OnError          types.String  `tfsdk:"on_error"`
	Position         types.List    `tfsdk:"position"`
	Parameters       types.Map     `tfsdk:"parameters"`
	Credentials      types.Map     `tfsdk:"credentials"`
	CreatedAt        types.String  `tfsdk:"created_at"`
	UpdatedAt        types.String  `tfsdk:"updated_at"`
}

// settings represents the settings of a workflow.
type settings struct {
	SaveExecutionProgress    types.Bool    `tfsdk:"save_execution_progress"`
	SaveManualExecutions     types.Bool    `tfsdk:"save_manual_executions"`
	SaveDataErrorExecution   types.String  `tfsdk:"save_data_error_execution"`
	SaveDataSuccessExecution types.String  `tfsdk:"save_data_success_execution"`
	ExecutionTimeout         types.Float32 `tfsdk:"execution_timeout"`
	ErrorWorkflow            types.String  `tfsdk:"error_workflow"`
	Timezone                 types.String  `tfsdk:"timezone"`
	ExecutionOrder           types.String  `tfsdk:"execution_order"`
}

// tag represents a tag in n8n.
type tag struct {
	ID        types.String `tfsdk:"id"`
	Name      types.String `tfsdk:"name"`
	CreatedAt types.String `tfsdk:"created_at"`
	UpdatedAt types.String `tfsdk:"updated_at"`
}
