package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ provider.Provider = &agentsmithProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &agentsmithProvider{
			version: version,
		}
	}
}

// agentsmithProvider is the provider implementation.
type agentsmithProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// agentsmithProviderModel maps provider schema data to a Go type.
type agentsmithProviderModel struct {
	Workdir types.String `tfsdk:"workdir"`
}

// Metadata returns the provider type name.
func (p *agentsmithProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "agentsmith"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *agentsmithProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"workdir": schema.StringAttribute{
				Optional:    true,
				Description: "The workdir you want to manage",
			},
		},
	}
}

// Configure prepares a agentsmith API client for data sources and resources.
func (p *agentsmithProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

	// Retrieve provider data from configuration
	var config agentsmithProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// If practitioner provided a configuration value for any of the
	// attributes, it must be a known value.

	workdir := os.Getenv("AGENTSMITH_WORKDIR")

	if !config.Workdir.IsNull() {
		workdir = config.Workdir.ValueString()
	}

	if workdir == "" {
		userHomeDir, err := os.UserHomeDir()
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to Get User Home Directory",
				"An unexpected error occurred when getting the user home directory. "+
					"Please specify the 'workdir' provider argument or set the AGENTSMITH_WORKDIR environment variable.\n\n"+
					"Error: "+err.Error(),
			)
			return
		}
		workdir = userHomeDir
		resp.Diagnostics.AddAttributeWarning(
			path.Root("workdir"),
			"No workdir specified, this instance of the provider will operate on your home directory.",
			fmt.Sprintf("This provider instance will operate on your %s directory.", userHomeDir),
		)
	}

	// Create a new file system client using the configuration values
	client, err := NewFileClient(workdir)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Agentsmith Filesystem Client",
			"An unexpected error occurred when creating the Agentsmith Filesystem client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"Agentsmith Filesystem Client Error: "+err.Error(),
		)
		return
	}

	// Make the Agentsmith client available during DataSource and Resource type Configure methods.
	resp.DataSourceData = client
	resp.ResourceData = client
}

// DataSources defines the data sources implemented in the provider.
func (p *agentsmithProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewClaudeDataSource,
	}
}

// Resources defines the resources implemented in the provider.
func (p *agentsmithProvider) Resources(_ context.Context) []func() resource.Resource {
	return nil
}
