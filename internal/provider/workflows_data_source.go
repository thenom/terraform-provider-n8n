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
	_ datasource.DataSource              = &workflowsDataSource{}
	_ datasource.DataSourceWithConfigure = &workflowsDataSource{}
)

// NewWorkflowDataSource is a helper function to simplify the provider implementation.
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
	resp.Schema = schema.Schema{
		Description: "Lists workflows.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Workflow ID",
				Required:    true,
			},
}

// Read refreshes the Terraform state with the latest data.
func (d *workflowDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state workflowDataSourceModel

	workflow, err := d.client.getWorkflow()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read HashiCups Coffees",
			err.Error(),
		)
		return
	}

	// Map response body to model
	for _, coffee := range coffees {
		coffeeState := coffeesModel{
			ID:          types.Int64Value(int64(coffee.ID)),
			Name:        types.StringValue(coffee.Name),
			Teaser:      types.StringValue(coffee.Teaser),
			Description: types.StringValue(coffee.Description),
			Price:       types.Float64Value(coffee.Price),
			Image:       types.StringValue(coffee.Image),
		}

		for _, ingredient := range coffee.Ingredient {
			coffeeState.Ingredients = append(coffeeState.Ingredients, coffeesIngredientsModel{
				ID: types.Int64Value(int64(ingredient.ID)),
			})
		}

		state.Coffees = append(state.Coffees, coffeeState)
	}

	// Set state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
