package provider

import (
	"context"
	"fmt"

	"github.com/edenreich/n8n-cli/n8n"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Client -
type client struct {
	N8NClient *n8n.Client
}

func (c *client) getWorkflow(ctx context.Context, workflowID string) (*workflowDataSourceModel, error) {
	workflow, err := c.N8NClient.GetWorkflow(workflowID)
	if err != nil {
		return nil, fmt.Errorf("failed to get workflow: %w", err)
	}

	var wfModel workflowDataSourceModel

	// Map top-level attributes
	wfModel.ID = types.StringPointerValue(workflow.Id)
	wfModel.Name = types.StringValue(workflow.Name)
	wfModel.Active = types.BoolPointerValue(workflow.Active)
	wfModel.CreatedAt = types.StringValue(workflow.CreatedAt.String())
	wfModel.UpdatedAt = types.StringValue(workflow.UpdatedAt.String())

	// Map Nodes
	nodes := make([]node, len(workflow.Nodes))
	for i, n := range workflow.Nodes {
		nodes[i] = node{
			ID:               types.StringPointerValue(n.Id),
			Name:             types.StringPointerValue(n.Name),
			WebhookID:        types.StringPointerValue(n.WebhookId),
			Disabled:         types.BoolPointerValue(n.Disabled),
			NotesInFlow:      types.BoolPointerValue(n.NotesInFlow),
			Notes:            types.StringPointerValue(n.Notes),
			Type:             types.StringPointerValue(n.Type),
			TypeVersion:      types.Float32PointerValue(n.TypeVersion),
			ExecuteOnce:      types.BoolPointerValue(n.ExecuteOnce),
			AlwaysOutputData: types.BoolPointerValue(n.AlwaysOutputData),
			RetryOnFail:      types.BoolPointerValue(n.RetryOnFail),
			MaxTries:         types.Float32PointerValue(n.MaxTries),
			WaitBetweenTries: types.Float32PointerValue(n.WaitBetweenTries),
			ContinueOnFail:   types.BoolPointerValue(n.ContinueOnFail),
			OnError:          types.StringPointerValue(n.OnError),
			Position:         convertInt64SliceToTypesInt64Slice(n.Position),
			Parameters:       convertMapToTypesMap(ctx, n.Parameters),
			Credentials:      convertMapToTypesMap(ctx, n.Credentials),
			CreatedAt:        types.StringValue(n.CreatedAt.String()),
			UpdatedAt:        types.StringValue(n.UpdatedAt.String()),
		}
	}
	wfModel.Nodes = nodes

	// Map Connections
	connections, diags := types.MapValueFrom(ctx, types.StringType, workflow.Connections)
	if diags.HasError() {
		return nil, fmt.Errorf("failed to convert connections: %v", diags)
	}
	wfModel.Connections = connections

	// Map Settings
	wfModel.Settings = settings{
		SaveExecutionProgress:    types.BoolPointerValue(workflow.Settings.SaveExecutionProgress),
		SaveManualExecutions:     types.BoolPointerValue(workflow.Settings.SaveManualExecutions),
		SaveDataErrorExecution:   types.StringPointerValue(workflow.Settings.SaveDataErrorExecution),
		SaveDataSuccessExecution: types.StringPointerValue(workflow.Settings.SaveDataSuccessExecution),
		ExecutionTimeout:         types.Int64PointerValue(workflow.Settings.ExecutionTimeout),
		ErrorWorkflow:            types.StringPointerValue(workflow.Settings.ErrorWorkflow),
		Timezone:                 types.StringPointerValue(workflow.Settings.Timezone),
		ExecutionOrder:           types.StringPointerValue(workflow.Settings.ExecutionOrder),
	}

	// Map Tags
	tags := make([]tag, len(workflow.Tags))
	for i, t := range workflow.Tags {
		tags[i] = tag{
			ID:        types.StringValue(t.ID),
			Name:      types.StringValue(t.Name),
			CreatedAt: types.StringValue(t.CreatedAt.String()),
			UpdatedAt: types.StringValue(t.UpdatedAt.String()),
		}
	}
	wfModel.Tags = tags

	return &wfModel, nil
}
