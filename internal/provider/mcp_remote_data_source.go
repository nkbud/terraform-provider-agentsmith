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
var _ datasource.DataSource = &mcpRemoteDataSource{}

func NewMcpRemoteDataSource() datasource.DataSource {
	return &mcpRemoteDataSource{}
}

type mcpRemoteDataSource struct{}

type mcpRemoteDataSourceModel struct {
	ID             types.String  `tfsdk:"id"`
	URL            types.String  `tfsdk:"url"`
	Transport      types.String  `tfsdk:"transport"`
	Headers        types.Map     `tfsdk:"headers"`
	Auth           types.String  `tfsdk:"auth"`
	SSEReadTimeout types.Float64 `tfsdk:"sse_read_timeout"`
	Timeout        types.Int64   `tfsdk:"timeout"`
	Description    types.String  `tfsdk:"description"`
	Icon           types.String  `tfsdk:"icon"`
	Authentication types.String  `tfsdk:"authentication"` // Using JSON string
	JSON           types.String  `tfsdk:"json"`
}

func (d *mcpRemoteDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mcp_remote"
}

func (d *mcpRemoteDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Generates a structured MCP (Model-Context Protocol) server configuration for a remote server using HTTP or SSE transports. This data source is used to build a server definition that can be consumed by other resources.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "A unique identifier for the generated configuration, derived from the configuration content.",
				Computed:    true,
			},
			"url": schema.StringAttribute{
				Description: "The URL of the remote MCP server.",
				Required:    true,
			},
			"transport": schema.StringAttribute{
				Description: "The transport mechanism. Valid values are `http`, `streamable-http`, or `sse`.",
				Optional:    true,
			},
			"headers": schema.MapAttribute{
				Description: "A map of HTTP headers to send with requests to the server.",
				ElementType: types.StringType,
				Optional:    true,
			},
			"auth": schema.StringAttribute{
				Description: "Authentication mechanism. Can be a string representing a Bearer token, or the literal `oauth` to use OAuth.",
				Optional:    true,
			},
			"sse_read_timeout": schema.Float64Attribute{
				Description: "Read timeout for Server-Sent Events (SSE) connections, in seconds.",
				Optional:    true,
			},
			"timeout": schema.Int64Attribute{
				Description: "Maximum response time for standard HTTP requests, in milliseconds.",
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

func (d *mcpRemoteDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data mcpRemoteDataSourceModel
	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	serverConfig := make(map[string]interface{})
	serverConfig["url"] = data.URL.ValueString()

	if !data.Transport.IsNull() && !data.Transport.IsUnknown() {
		serverConfig["transport"] = data.Transport.ValueString()
	}
	if !data.Headers.IsNull() && !data.Headers.IsUnknown() {
		var headers map[string]string
		diags := data.Headers.ElementsAs(ctx, &headers, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		serverConfig["headers"] = headers
	}
	if !data.Auth.IsNull() && !data.Auth.IsUnknown() {
		serverConfig["auth"] = data.Auth.ValueString()
	}
	if !data.SSEReadTimeout.IsNull() && !data.SSEReadTimeout.IsUnknown() {
		serverConfig["sse_read_timeout"] = data.SSEReadTimeout.ValueFloat64()
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
