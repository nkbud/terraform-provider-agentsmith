package provider

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var _ datasource.DataSource = &mcpStdioDataSource{}

func NewMcpStdioDataSource() datasource.DataSource {
	return &mcpStdioDataSource{}
}

type mcpStdioDataSource struct{}

type mcpStdioDataSourceModel struct {
	ID             types.String `tfsdk:"id"`
	Command        types.String `tfsdk:"command"`
	Args           types.List   `tfsdk:"args"`
	Env            types.Map    `tfsdk:"env"`
	Transport      types.String `tfsdk:"transport"`
	Type           types.String `tfsdk:"type"`
	Cwd            types.String `tfsdk:"cwd"`
	Timeout        types.Int64  `tfsdk:"timeout"`
	Description    types.String `tfsdk:"description"`
	Icon           types.String `tfsdk:"icon"`
	Authentication types.String `tfsdk:"authentication"` // Using JSON string for this complex object
	JSON           types.String `tfsdk:"json"`
}

func (d *mcpStdioDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mcp_stdio"
}

func (d *mcpStdioDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Generates a structured MCP (Model-Context Protocol) server configuration for a local server using a stdio-based transport. This data source is used to build a server definition that can be consumed by other resources.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "A unique identifier for the generated configuration, derived from the configuration content.",
				Computed:    true,
			},
			"command": schema.StringAttribute{
				Description: "The command to execute to start the MCP server.",
				Required:    true,
			},
			"args": schema.ListAttribute{
				Description: "A list of arguments to pass to the command.",
				ElementType: types.StringType,
				Optional:    true,
			},
			"env": schema.MapAttribute{
				Description: "A map of environment variables to set for the command's process.",
				ElementType: types.StringType,
				Optional:    true,
			},
			"transport": schema.StringAttribute{
				Description: "The transport mechanism. For this data source, it defaults to `stdio`.",
				Optional:    true,
			},
			"type": schema.StringAttribute{
				Description: "An alternative name for the transport field.",
				Optional:    true,
			},
			"cwd": schema.StringAttribute{
				Description: "The working directory in which to execute the command.",
				Optional:    true,
			},
			"timeout": schema.Int64Attribute{
				Description: "Maximum response time in milliseconds for the server to respond.",
				Optional:    true,
			},
			"description": schema.StringAttribute{
				Description: "A human-readable description of the server.",
				Optional:    true,
			},
			"icon": schema.StringAttribute{
				Description: "An icon path or URL for UI display.",
				Optional:    true,
			},
			"authentication": schema.StringAttribute{
				Description: "A JSON string representing a complex authentication configuration object.",
				Optional:    true,
			},
			"json": schema.StringAttribute{
				Description: "The resulting MCP server configuration, as a JSON string. This output can be used by other resources that consume MCP server definitions.",
				Computed:    true,
			},
		},
	}
}

func (d *mcpStdioDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data mcpStdioDataSourceModel
	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	serverConfig := make(map[string]interface{})
	serverConfig["command"] = data.Command.ValueString()

	if !data.Args.IsNull() && !data.Args.IsUnknown() {
		var args []string
		diags := data.Args.ElementsAs(ctx, &args, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		serverConfig["args"] = args
	}

	if !data.Env.IsNull() && !data.Env.IsUnknown() {
		var env map[string]string
		diags := data.Env.ElementsAs(ctx, &env, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		serverConfig["env"] = env
	}

	transport := "stdio"
	if !data.Transport.IsNull() && !data.Transport.IsUnknown() {
		transport = data.Transport.ValueString()
	}
	serverConfig["transport"] = transport

	if !data.Type.IsNull() && !data.Type.IsUnknown() {
		serverConfig["type"] = data.Type.ValueString()
	}
	if !data.Cwd.IsNull() && !data.Cwd.IsUnknown() {
		serverConfig["cwd"] = data.Cwd.ValueString()
	}
	if !data.Timeout.IsNull() && !data.Timeout.IsUnknown() {
		serverConfig["timeout"] = data.Timeout.ValueInt64()
	}
	if !data.Description.IsNull() && !data.Description.IsUnknown() {
		serverConfig["description"] = data.Description.ValueString()
	}
	if !data.Icon.IsNull() && !data.Icon.IsUnknown() {
		serverConfig["icon"] = data.Icon.ValueString()
	}

	if !data.Authentication.IsNull() && !data.Authentication.IsUnknown() {
		var authConfig map[string]interface{}
		err := json.Unmarshal([]byte(data.Authentication.ValueString()), &authConfig)
		if err != nil {
			resp.Diagnostics.AddError("Invalid Authentication JSON", err.Error())
			return
		}
		serverConfig["authentication"] = authConfig
	}

	jsonBytes, err := json.Marshal(serverConfig)
	if err != nil {
		resp.Diagnostics.AddError("Failed to marshal server config", err.Error())
		return
	}
	jsonString := string(jsonBytes)
	data.JSON = types.StringValue(jsonString)

	hash := sha256.Sum256([]byte(jsonString))
	data.ID = types.StringValue(hex.EncodeToString(hash[:]))

	diags = resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}
