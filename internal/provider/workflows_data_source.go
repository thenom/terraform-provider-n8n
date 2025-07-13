package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &workflowsDataSource{}
	_ datasource.DataSourceWithConfigure = &workflowsDataSource{}
)

// NewWorkflowsDataSource is a helper function to simplify the provider implementation.
func NewWorkflowsDataSource() datasource.DataSource {
	return &workflowsDataSource{}
}

// workflowDataSource is the data source implementation.
type workflowsDataSource struct {
	client *client
}

// Configure adds the provider configured client to the data source.
func (d *workflowsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
func (d *workflowsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_workflows"
}

// Schema defines the schema for the data source.
func (d *workflowsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	// Get workflow schema
	wfSchema := datasource.SchemaResponse{}
	wf := workflowDataSource{}
	wf.Schema(context.Background(), datasource.SchemaRequest{}, &wfSchema)

	// Setup list of workflow schemas
	resp.Schema = schema.Schema{
		Description: "List of workflows.",
		Attributes: map[string]schema.Attribute{
			"data": schema.ListNestedAttribute{
				Description: "Workflows",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: wfSchema.Schema.Attributes,
				},
			},
		},
	}
}

// Read refreshes the Terraform state with the latest data.
func (d *workflowsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	state, err := d.client.getWorkflows(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read workflows",
			err.Error(),
		)
		return
	}

	// Set state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
