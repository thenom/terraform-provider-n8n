// Package provider package is the main implementation for the N8N provider
package provider

import (
	"context"

	"github.com/edenreich/n8n-cli/n8n"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ provider.Provider = &n8nProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &n8nProvider{
			version: version,
		}
	}
}

// n8nProvider is the provider implementation.
type n8nProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
	client  *client
}

// n8nProviderModel maps provider schema data to a Go type.
type n8nProviderModel struct {
	HostURL types.String `tfsdk:"host_url"`
	APIKey  types.String `tfsdk:"api_key"`
}

// Metadata returns the provider type name.
func (p *n8nProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "n8n"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *n8nProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"host_url": schema.StringAttribute{
				MarkdownDescription: "URL of the n8n instance.",
				Required:            true,
			},
			"api_key": schema.StringAttribute{
				MarkdownDescription: "API Key for n8n instance.",
				Required:            true,
				Sensitive:           true,
			},
		},
	}
}

// Configure prepares a HashiCups API client for data sources and resources.
func (p *n8nProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config n8nProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create a new n8n client
	n8nClient := n8n.NewClient(config.HostURL.ValueString(), config.APIKey.ValueString())

	p.client = &client{
		N8NClient: n8nClient,
	}

	resp.DataSourceData = p.client
	resp.ResourceData = p.client
}

// DataSources defines the data sources implemented in the provider.
func (p *n8nProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewWorkflowsDataSource,
	}
}

// Resources defines the resources implemented in the provider.
func (p *n8nProvider) Resources(_ context.Context) []func() resource.Resource {
	return nil
}
