package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &workflowDataSource{}
	_ datasource.DataSourceWithConfigure = &workflowDataSource{}
)

// NewWorkflowDataSource is a helper function to simplify the provider implementation.
func NewWorkflowDataSource() datasource.DataSource {
	return &workflowDataSource{}
}

// workflowDataSource is the data source implementation.
type workflowDataSource struct {
	client *client
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
	ID               types.String    `tfsdk:"id"`
	Name             types.String    `tfsdk:"name"`
	WebhookID        types.String    `tfsdk:"webhook_id"`
	Disabled         types.Bool      `tfsdk:"disabled"`
	NotesInFlow      types.Bool      `tfsdk:"notes_in_flow"`
	Notes            types.String    `tfsdk:"notes"`
	Type             types.String    `tfsdk:"type"`
	TypeVersion      types.Float32   `tfsdk:"type_version"`
	ExecuteOnce      types.Bool      `tfsdk:"execute_once"`
	AlwaysOutputData types.Bool      `tfsdk:"always_output_data"`
	RetryOnFail      types.Bool      `tfsdk:"retry_on_fail"`
	MaxTries         types.Float32   `tfsdk:"max_tries"`
	WaitBetweenTries types.Float32   `tfsdk:"wait_between_tries"`
	ContinueOnFail   types.Bool      `tfsdk:"continue_on_fail"`
	OnError          types.String    `tfsdk:"on_error"`
	Position         types.List      `tfsdk:"position"`
	Parameters       types.Map       `tfsdk:"parameters"`
	Credentials      types.Map       `tfsdk:"credentials"`
	CreatedAt        types.String    `tfsdk:"created_at"`
	UpdatedAt        types.String    `tfsdk:"updated_at"`
}

// settings represents the settings of a workflow.
type settings struct {
	SaveExecutionProgress    types.Bool   `tfsdk:"save_execution_progress"`
	SaveManualExecutions     types.Bool   `tfsdk:"save_manual_executions"`
	SaveDataErrorExecution   types.String `tfsdk:"save_data_error_execution"`
	SaveDataSuccessExecution types.String `tfsdk:"save_data_success_execution"`
	ExecutionTimeout         types.Int64  `tfsdk:"execution_timeout"`
	ErrorWorkflow            types.String `tfsdk:"error_workflow"`
	Timezone                 types.String `tfsdk:"timezone"`
	ExecutionOrder           types.String `tfsdk:"execution_order"`
}

// tag represents a tag in n8n.
type tag struct {
	ID        types.String `tfsdk:"id"`
	Name      types.String `tfsdk:"name"`
	CreatedAt types.String `tfsdk:"created_at"`
	UpdatedAt types.String `tfsdk:"updated_at"`
}

// Configure adds the provider configured client to the data source.
func (d *workflowDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}
	d.client = client
}

// Metadata returns the data source type name.
func (d *workflowDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_workflow"
}

// Schema defines the schema for the data source.
func (d *workflowDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a workflow.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Workflow ID",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the workflow.",
				Computed:    true,
			},
			"active": schema.BoolAttribute{
				Description: "Whether the workflow is active.",
				Computed:    true,
			},
			"nodes": schema.ListNestedAttribute{
				Description: "The nodes of the workflow.",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Node ID",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "Node name",
							Computed:    true,
						},
						"webhook_id": schema.StringAttribute{
							Description: "Webhook ID",
							Computed:    true,
						},
						"disabled": schema.BoolAttribute{
							Description: "Whether the node is disabled.",
							Computed:    true,
						},
						"notes_in_flow": schema.BoolAttribute{
							Description: "Whether the node has notes in the flow.",
							Computed:    true,
						},
						"notes": schema.StringAttribute{
							Description: "Node notes",
							Computed:    true,
						},
						"type": schema.StringAttribute{
							Description: "Node type",
							Computed:    true,
						},
						"type_version": schema.FloatAttribute{
							Description: "Node type version",
							Computed:    true,
						},
						"execute_once": schema.BoolAttribute{
							Description: "Whether the node executes only once.",
							Computed:    true,
						},
						"always_output_data": schema.BoolAttribute{
							Description: "Whether the node always outputs data.",
							Computed:    true,
						},
						"retry_on_fail": schema.BoolAttribute{
							Description: "Whether the node retries on fail.",
							Computed:    true,
						},
						"max_tries": schema.FloatAttribute{
							Description: "Max tries for the node.",
							Computed:    true,
						},
						"wait_between_tries": schema.FloatAttribute{
							Description: "Wait between tries for the node.",
							Computed:    true,
						},
						"continue_on_fail": schema.BoolAttribute{
							Description: "Whether the node continues on fail.",
							Computed:    true,
						},
						"on_error": schema.StringAttribute{
							Description: "On error action for the node.",
							Computed:    true,
						},
						"position": schema.ListAttribute{
							Description: "Node position",
							ElementType: types.FloatType,
							Computed:    true,
						},
						"parameters": schema.MapAttribute{
							Description: "Node parameters",
							ElementType: types.StringType,
							Computed:    true,
						},
						"credentials": schema.MapAttribute{
							Description: "Node credentials",
							ElementType: types.StringType,
							Computed:    true,
						},
						"created_at": schema.StringAttribute{
							Description: "Node creation date",
							Computed:    true,
						},
						"updated_at": schema.StringAttribute{
							Description: "Node last update date",
							Computed:    true,
						},
					},
				},
			},
			"connections": schema.MapAttribute{
				Description: "The connections of the workflow.",
				ElementType: types.StringType,
				Computed:    true,
			},
			"settings": schema.SingleNestedAttribute{
				Description: "The settings of the workflow.",
				Computed:    true,
				Attributes: map[string]schema.Attribute{
					"save_execution_progress": schema.BoolAttribute{
						Description: "Whether to save execution progress.",
						Computed:    true,
					},
					"save_manual_executions": schema.BoolAttribute{
						Description: "Whether to save manual executions.",
						Computed:    true,
					},
					"save_data_error_execution": schema.StringAttribute{
						Description: "Save data on error execution.",
						Computed:    true,
					},
					"save_data_success_execution": schema.StringAttribute{
						Description: "Save data on success execution.",
						Computed:    true,
					},
					"execution_timeout": schema.Int64Attribute{
						Description: "Execution timeout.",
						Computed:    true,
					},
					"error_workflow": schema.StringAttribute{
						Description: "Error workflow.",
						Computed:    true,
					},
					"timezone": schema.StringAttribute{
						Description: "Timezone.",
						Computed:    true,
					},
					"execution_order": schema.StringAttribute{
						Description: "Execution order.",
						Computed:    true,
					},
				},
			},
			"static_data": schema.StringAttribute{
				Description: "The static data of the workflow.",
				Computed:    true,
			},
			"tags": schema.ListNestedAttribute{
				Description: "The tags of the workflow.",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Tag ID",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "Tag name",
							Computed:    true,
						},
						"created_at": schema.StringAttribute{
							Description: "Tag creation date",
							Computed:    true,
						},
						"updated_at": schema.StringAttribute{
							Description: "Tag last update date",
							Computed:    true,
						},
					},
				},
			},
			"created_at": schema.StringAttribute{
				Description: "The creation date of the workflow.",
				Computed:    true,
			},
			"updated_at": schema.StringAttribute{
				Description: "The last update date of the workflow.",
				Computed:    true,
			},
		},
	}
}

// Read refreshes the Terraform state with the latest data.
func (d *workflowDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state workflowDataSourceModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed workflow value from n8n
	workflowResponse, err := d.client.getWorkflow(ctx, state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read n8n Workflow",
			err.Error(),
		)
		return
	}

	// Set state
	diags = resp.State.Set(ctx, &workflowResponse)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
