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
		positionList, diags := types.ListValueFrom(ctx, types.Float32Type, n.Position)
		if diags.HasError() {
			return nil, fmt.Errorf("failed to convert position: %v", diags)
		}

		parametersMap, err := convertMapToTypesMap(ctx, *n.Parameters)
		if err != nil {
			return nil, fmt.Errorf("failed to convert parameters: %w", err)
		}

		credentialsMap, err := convertMapToTypesMap(ctx, *n.Credentials)
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
	connections, diags := types.MapValueFrom(ctx, types.StringType, workflow.Connections)
	if diags.HasError() {
		return nil, fmt.Errorf("failed to convert connections: %v", diags)
	}
	wfModel.Connections = connections

	// Map Settings
	wfModel.Settings = settings{
		SaveExecutionProgress:    types.BoolPointerValue(workflow.Settings.SaveExecutionProgress),
		SaveManualExecutions:     types.BoolPointerValue(workflow.Settings.SaveManualExecutions),
		SaveDataErrorExecution:   types.StringValue(string(*workflow.Settings.SaveDataErrorExecution)),
		SaveDataSuccessExecution: types.StringValue(string(*workflow.Settings.SaveDataSuccessExecution)),
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

func convertInterfaceMapToStringMap(input map[string]interface{}) (map[string]string, error) {
	output := make(map[string]string)
	for k, v := range input {
		str, ok := v.(string)
		if !ok {
			return nil, fmt.Errorf("value for key '%s' is not a string", k)
		}
		output[k] = str
	}
	return output, nil
}

func convertMapToTypesMap(ctx context.Context, data map[string]interface{}) (types.Map, error) {
	stringMap, err := convertInterfaceMapToStringMap(data)
	if err != nil {
		return types.Map{}, err
	}
	mapValue, diags := types.MapValueFrom(ctx, types.StringType, stringMap)
	if diags.HasError() {
		return types.Map{}, fmt.Errorf("failed to convert map: %v", diags)
	}
	return mapValue, nil
}
