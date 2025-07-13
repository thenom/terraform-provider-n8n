package provider

import (
	"context"
	"fmt"

	"github.com/edenreich/n8n-cli/n8n"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Client -
type client struct {
	N8NClient *n8n.Client
}

func (c *client) getWorkflows(ctx context.Context) (*workflowListDataSourceModel, error) {
	workflows, err := c.N8NClient.GetWorkflows()
	if err != nil {
		return nil, fmt.Errorf("failed to get workflow: %w", err)
	}
	tflog.Debug(ctx, "received workflows", map[string]any{"workflows": workflows})

	wfList := workflowListDataSourceModel{}
	wfListData := []workflowDataSourceModel{}
	for _, wf := range *workflows.Data {
		wfModel, err := workflowToModel(ctx, &wf)
		if err != nil {
			return nil, err
		}

		wfListData = append(wfListData, *wfModel)
	}
	wfList.Data = &wfListData

	return &wfList, nil
}

func workflowToModel(ctx context.Context, wf *n8n.Workflow) (*workflowDataSourceModel, error) {
	var wfModel workflowDataSourceModel

	// Map top-level attributes
	wfModel.ID = types.StringPointerValue(wf.Id)
	wfModel.Name = types.StringValue(wf.Name)
	wfModel.Active = types.BoolPointerValue(wf.Active)
	wfModel.CreatedAt = types.StringValue(wf.CreatedAt.String())
	wfModel.UpdatedAt = types.StringValue(wf.UpdatedAt.String())

	// Map Nodes
	nodes := make([]node, len(wf.Nodes))
	for i, n := range wf.Nodes {
		positionList, diags := types.ListValueFrom(ctx, types.Float32Type, n.Position)
		if diags.HasError() {
			return nil, fmt.Errorf("failed to convert position: %v", diags)
		}

		parametersMap, err := convertInterfaceMapToStringMap(n.Parameters)
		if err != nil {
			return nil, fmt.Errorf("failed to convert parameters: %w", err)
		}

		credentialsMap, err := convertInterfaceMapToStringMap(n.Credentials)
		if err != nil {
			return nil, fmt.Errorf("failed to convert credentials: %w", err)
		}

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
			Position:         positionList,
			Parameters:       parametersMap,
			Credentials:      credentialsMap,
			CreatedAt:        types.StringValue(n.CreatedAt.String()),
			UpdatedAt:        types.StringValue(n.UpdatedAt.String()),
		}
	}
	wfModel.Nodes = nodes

	// Map Connections
	connections, diags := types.MapValueFrom(ctx, types.StringType, wf.Connections)
	if diags.HasError() {
		return nil, fmt.Errorf("failed to convert connections: %v", diags)
	}
	wfModel.Connections = connections

	// Map Settings
	wfModel.Settings = settings{
		SaveExecutionProgress:    types.BoolPointerValue(wf.Settings.SaveExecutionProgress),
		SaveManualExecutions:     types.BoolPointerValue(wf.Settings.SaveManualExecutions),
		SaveDataErrorExecution:   types.StringValue(string(*wf.Settings.SaveDataErrorExecution)),
		SaveDataSuccessExecution: types.StringValue(string(*wf.Settings.SaveDataSuccessExecution)),
		ExecutionTimeout:         types.Float32PointerValue(wf.Settings.ExecutionTimeout),
		ErrorWorkflow:            types.StringPointerValue(wf.Settings.ErrorWorkflow),
		Timezone:                 types.StringPointerValue(wf.Settings.Timezone),
		ExecutionOrder:           types.StringPointerValue(wf.Settings.ExecutionOrder),
	}

	// Map Tags
	tags := make([]tag, len(*wf.Tags))
	for i, t := range *wf.Tags {
		tags[i] = tag{
			ID:        types.StringPointerValue(t.Id),
			Name:      types.StringValue(t.Name),
			CreatedAt: types.StringValue(t.CreatedAt.String()),
			UpdatedAt: types.StringValue(t.UpdatedAt.String()),
		}
	}
	wfModel.Tags = tags

	return &wfModel, nil
}
