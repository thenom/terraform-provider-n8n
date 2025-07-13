package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// workflowDataSource is the data source implementation.
type workflowDataSource struct {
	client *client
}

// Schema defines the schema for the data source.
func (d *workflowDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Lists workflows.",
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
						"type_version": schema.Float32Attribute{
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
						"max_tries": schema.Float32Attribute{
							Description: "Max tries for the node.",
							Computed:    true,
						},
						"wait_between_tries": schema.Float32Attribute{
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
							ElementType: types.Float32Type,
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
