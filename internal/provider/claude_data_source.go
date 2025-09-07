package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"gopkg.in/yaml.v3"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &claudeDataSource{}
	_ datasource.DataSourceWithConfigure = &claudeDataSource{}
)

// NewClaudeDataSource is a helper function to simplify the provider implementation.
func NewClaudeDataSource() datasource.DataSource {
	return &claudeDataSource{}
}

// claudeDataSource is the data source implementation.
type claudeDataSource struct {
	client *FileClient
}

// claudeDataSourceModel maps the schema to a Go type.
type claudeDataSourceModel struct {
	ID                   types.String               `tfsdk:"id"`
	Settings             *claudeSettingsModel       `tfsdk:"settings"`
	EnvironmentVariables *environmentVariablesModel `tfsdk:"environment_variables"`
	GlobalConfig         *globalConfigModel         `tfsdk:"global_config"`
	Subagents            []subagentModel            `tfsdk:"subagents"`
	Commands             []commandModel             `tfsdk:"commands"`
	HookFiles            []hookFileModel            `tfsdk:"hook_files"`
}

// claudeSettingsModel maps the settings nested attribute to a Go type.
type claudeSettingsModel struct {
	APIKeyHelper               types.String      `tfsdk:"api_key_helper"`
	CleanupPeriodDays          types.Int64       `tfsdk:"cleanup_period_days"`
	Env                        types.Map         `tfsdk:"env"`
	IncludeCoAuthoredBy        types.Bool        `tfsdk:"include_co_authored_by"`
	Permissions                *permissionsModel `tfsdk:"permissions"`
	Hooks                      types.Map         `tfsdk:"hooks"`
	DisableAllHooks            types.Bool        `tfsdk:"disable_all_hooks"`
	Model                      types.String      `tfsdk:"model"`
	StatusLine                 *statusLineModel  `tfsdk:"status_line"`
	OutputStyle                types.String      `tfsdk:"output_style"`
	ForceLoginMethod           types.String      `tfsdk:"force_login_method"`
	ForceLoginOrgUUID          types.String      `tfsdk:"force_login_org_uuid"`
	EnableAllProjectMcpServers types.Bool        `tfsdk:"enable_all_project_mcp_servers"`
	EnabledMcpjsonServers      []types.String    `tfsdk:"enabled_mcpjson_servers"`
	DisabledMcpjsonServers     []types.String    `tfsdk:"disabled_mcpjson_servers"`
	AwsAuthRefresh             types.String      `tfsdk:"aws_auth_refresh"`
	AwsCredentialExport        types.String      `tfsdk:"aws_credential_export"`
}

// permissionsModel maps the permissions nested attribute to a Go type.
type permissionsModel struct {
	Allow                        []types.String `tfsdk:"allow"`
	Ask                          []types.String `tfsdk:"ask"`
	Deny                         []types.String `tfsdk:"deny"`
	AdditionalDirectories        []types.String `tfsdk:"additional_directories"`
	DefaultMode                  types.String   `tfsdk:"default_mode"`
	DisableBypassPermissionsMode types.String   `tfsdk:"disable_bypass_permissions_mode"`
}

// statusLineModel maps the status_line nested attribute to a Go type.
type statusLineModel struct {
	Type    types.String `tfsdk:"type"`
	Command types.String `tfsdk:"command"`
}

// environmentVariablesModel maps the environment_variables nested attribute to a Go type.
type environmentVariablesModel struct {
	AnthropicAPIKey                      types.String `tfsdk:"anthropic_api_key"`
	AnthropicAuthToken                   types.String `tfsdk:"anthropic_auth_token"`
	AnthropicCustomHeaders               types.String `tfsdk:"anthropic_custom_headers"`
	AnthropicDefaultHaikuModel           types.String `tfsdk:"anthropic_default_haiku_model"`
	AnthropicDefaultOpusModel            types.String `tfsdk:"anthropic_default_opus_model"`
	AnthropicDefaultSonnetModel          types.String `tfsdk:"anthropic_default_sonnet_model"`
	AnthropicModel                       types.String `tfsdk:"anthropic_model"`
	AnthropicSmallFastModel              types.String `tfsdk:"anthropic_small_fast_model"`
	AnthropicSmallFastModelAWSRegion     types.String `tfsdk:"anthropic_small_fast_model_aws_region"`
	AWSBearerTokenBedrock                types.String `tfsdk:"aws_bearer_token_bedrock"`
	BashDefaultTimeoutMS                 types.Int64  `tfsdk:"bash_default_timeout_ms"`
	BashMaxTimeoutMS                     types.Int64  `tfsdk:"bash_max_timeout_ms"`
	BashMaxOutputLength                  types.Int64  `tfsdk:"bash_max_output_length"`
	ClaudeBashMaintainProjectWorkingDir  types.Bool   `tfsdk:"claude_bash_maintain_project_working_dir"`
	ClaudeCodeAPIKeyHelperTTLMS          types.Int64  `tfsdk:"claude_code_api_key_helper_ttl_ms"`
	ClaudeCodeIDESkipAutoInstall         types.Bool   `tfsdk:"claude_code_ide_skip_auto_install"`
	ClaudeCodeMaxOutputTokens            types.Int64  `tfsdk:"claude_code_max_output_tokens"`
	ClaudeCodeUseBedrock                 types.Bool   `tfsdk:"claude_code_use_bedrock"`
	ClaudeCodeUseVertex                  types.Bool   `tfsdk:"claude_code_use_vertex"`
	ClaudeCodeSkipBedrockAuth            types.Bool   `tfsdk:"claude_code_skip_bedrock_auth"`
	ClaudeCodeSkipVertexAuth             types.Bool   `tfsdk:"claude_code_skip_vertex_auth"`
	ClaudeCodeDisableNonessentialTraffic types.Bool   `tfsdk:"claude_code_disable_nonessential_traffic"`
	ClaudeCodeDisableTerminalTitle       types.Bool   `tfsdk:"claude_code_disable_terminal_title"`
	ClaudeCodeSubagentModel              types.String `tfsdk:"claude_code_subagent_model"`
	DisableAutoupdater                   types.Bool   `tfsdk:"disable_autoupdater"`
	DisableBugCommand                    types.Bool   `tfsdk:"disable_bug_command"`
	DisableCostWarnings                  types.Bool   `tfsdk:"disable_cost_warnings"`
	DisableErrorReporting                types.Bool   `tfsdk:"disable_error_reporting"`
	DisableNonEssentialModelCalls        types.Bool   `tfsdk:"disable_non_essential_model_calls"`
	DisableTelemetry                     types.Bool   `tfsdk:"disable_telemetry"`
	HTTPProxy                            types.String `tfsdk:"http_proxy"`
	HTTPSProxy                           types.String `tfsdk:"https_proxy"`
	MaxThinkingTokens                    types.Int64  `tfsdk:"max_thinking_tokens"`
	MCPTimeout                           types.Int64  `tfsdk:"mcp_timeout"`
	MCPToolTimeout                       types.Int64  `tfsdk:"mcp_tool_timeout"`
	MaxMCPOutputTokens                   types.Int64  `tfsdk:"max_mcp_output_tokens"`
	NoProxy                              types.String `tfsdk:"no_proxy"`
	UseBuiltinRipgrep                    types.Bool   `tfsdk:"use_builtin_ripgrep"`
	VertexRegionClaude35Haiku            types.String `tfsdk:"vertex_region_claude_3_5_haiku"`
	VertexRegionClaude35Sonnet           types.String `tfsdk:"vertex_region_claude_3_5_sonnet"`
	VertexRegionClaude37Sonnet           types.String `tfsdk:"vertex_region_claude_3_7_sonnet"`
	VertexRegionClaude40Opus             types.String `tfsdk:"vertex_region_claude_4_0_opus"`
	VertexRegionClaude40Sonnet           types.String `tfsdk:"vertex_region_claude_4_0_sonnet"`
	VertexRegionClaude41Opus             types.String `tfsdk:"vertex_region_claude_4_1_opus"`
}

// globalConfigModel maps the global_config nested attribute to a Go type.
type globalConfigModel struct {
	AutoUpdates           types.Bool   `tfsdk:"auto_updates"`
	PreferredNotifChannel types.String `tfsdk:"preferred_notif_channel"`
	Theme                 types.String `tfsdk:"theme"`
	Verbose               types.Bool   `tfsdk:"verbose"`
}

// subagentModel maps the subagents nested attribute to a Go type.
type subagentModel struct {
	Path        types.String `tfsdk:"path"`
	Name        types.String `tfsdk:"name"`
	Model       types.String `tfsdk:"model"`
	Description types.String `tfsdk:"description"`
	Color       types.String `tfsdk:"color"`
	Prompt      types.String `tfsdk:"prompt"`
}

// commandModel maps the commands nested attribute to a Go type.
type commandModel struct {
	Path types.String `tfsdk:"path"`
}

// hookFileModel maps the hook_files nested attribute to a Go type.
type hookFileModel struct {
	Path types.String `tfsdk:"path"`
}

// Metadata returns the data source type name.
func (d *claudeDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_claude"
}

// Schema defines the schema for the data source.
func (d *claudeDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Reads and consolidates the complete Claude Code configuration from the local environment. This includes settings from user, project, and enterprise files, environment variables, and discovered assets like subagents, commands, and hooks.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "A static identifier for the data source.",
				Computed:    true,
			},
			"settings": schema.SingleNestedAttribute{
				Description: "Merged settings from all found `settings.json` files, following Claude Code's precedence rules.",
				Computed:    true,
				Attributes: map[string]schema.Attribute{
					"api_key_helper": schema.StringAttribute{
						Description: "A custom script to be executed in `/bin/sh` to generate an auth value for model requests.",
						Computed:    true,
					},
					"cleanup_period_days": schema.Int64Attribute{
						Description: "How long to locally retain chat transcripts based on last activity date. Default: 30 days.",
						Computed:    true,
					},
					"env": schema.MapAttribute{
						Description: "Environment variables that will be applied to every session.",
						ElementType: types.StringType,
						Computed:    true,
					},
					"include_co_authored_by": schema.BoolAttribute{
						Description: "Whether to include the `co-authored-by Claude` byline in git commits and pull requests. Default: true.",
						Computed:    true,
					},
					"permissions": schema.SingleNestedAttribute{
						Description: "A block for configuring tool usage permissions.",
						Computed:    true,
						Attributes: map[string]schema.Attribute{
							"allow": schema.ListAttribute{
								Description: "A list of permission rules to automatically allow tool use without prompting.",
								ElementType: types.StringType,
								Computed:    true,
							},
							"ask": schema.ListAttribute{
								Description: "A list of permission rules that will cause Claude to ask for confirmation before using a tool.",
								ElementType: types.StringType,
								Computed:    true,
							},
							"deny": schema.ListAttribute{
								Description: "A list of permission rules to deny tool use. Also used to exclude sensitive files from being read.",
								ElementType: types.StringType,
								Computed:    true,
							},
							"additional_directories": schema.ListAttribute{
								Description: "A list of additional working directories that Claude has access to.",
								ElementType: types.StringType,
								Computed:    true,
							},
							"default_mode": schema.StringAttribute{
								Description: "The default permission mode when opening Claude Code (e.g., `acceptEdits`).",
								Computed:    true,
							},
							"disable_bypass_permissions_mode": schema.StringAttribute{
								Description: "Set to `disable` to prevent `bypassPermissions` mode from being activated.",
								Computed:    true,
							},
						},
					},
					"hooks": schema.MapAttribute{
						Description: "A map of custom commands to run before or after tool executions (e.g., `PreToolUse`).",
						ElementType: types.MapType{ElemType: types.StringType},
						Computed:    true,
					},
					"disable_all_hooks": schema.BoolAttribute{
						Description: "If true, disables all configured hooks.",
						Computed:    true,
					},
					"model": schema.StringAttribute{
						Description: "The name of the model to use for Claude Code (e.g., `claude-3-5-sonnet-20241022`).",
						Computed:    true,
					},
					"status_line": schema.SingleNestedAttribute{
						Description: "Configuration for a custom status line to display context.",
						Computed:    true,
						Attributes: map[string]schema.Attribute{
							"type":    schema.StringAttribute{Description: "The type of status line, e.g., `command`.", Computed: true},
							"command": schema.StringAttribute{Description: "The command to execute to generate the status line content.", Computed: true},
						},
					},
					"output_style": schema.StringAttribute{
						Description: "The output style to adjust the system prompt (e.g., `Explanatory`).",
						Computed:    true,
					},
					"force_login_method": schema.StringAttribute{
						Description: "Restricts login to a specific method. Use `claudeai` for Claude.ai accounts or `console` for Anthropic Console accounts.",
						Computed:    true,
					},
					"force_login_org_uuid": schema.StringAttribute{
						Description: "The UUID of an organization to automatically select during login, bypassing the organization selection step.",
						Computed:    true,
					},
					"enable_all_project_mcp_servers": schema.BoolAttribute{
						Description: "If true, automatically approves all MCP servers defined in project `.mcp.json` files.",
						Computed:    true,
					},
					"enabled_mcpjson_servers": schema.ListAttribute{
						Description: "A list of specific MCP servers from `.mcp.json` files to approve.",
						ElementType: types.StringType,
						Computed:    true,
					},
					"disabled_mcpjson_servers": schema.ListAttribute{
						Description: "A list of specific MCP servers from `.mcp.json` files to reject.",
						ElementType: types.StringType,
						Computed:    true,
					},
					"aws_auth_refresh": schema.StringAttribute{
						Description: "A custom script that modifies the `.aws` directory for auth refresh, e.g., `aws sso login`.",
						Computed:    true,
					},
					"aws_credential_export": schema.StringAttribute{
						Description: "A custom script that outputs JSON with temporary AWS credentials.",
						Computed:    true,
					},
				},
			},
			"environment_variables": schema.SingleNestedAttribute{
				Description: "A map of all Claude Code related environment variables found in the environment.",
				Computed:    true,
				Attributes: map[string]schema.Attribute{
					"anthropic_api_key":                        schema.StringAttribute{Description: "API key for Anthropic models.", Computed: true, Sensitive: true},
					"anthropic_auth_token":                     schema.StringAttribute{Description: "Bearer token for authentication.", Computed: true, Sensitive: true},
					"aws_bearer_token_bedrock":                 schema.StringAttribute{Description: "Bearer token for AWS Bedrock.", Computed: true, Sensitive: true},
					"anthropic_custom_headers":                 schema.StringAttribute{Description: "Custom headers to add to model requests.", Computed: true},
					"anthropic_default_haiku_model":            schema.StringAttribute{Description: "Default model name for the Haiku class.", Computed: true},
					"anthropic_default_opus_model":             schema.StringAttribute{Description: "Default model name for the Opus class.", Computed: true},
					"anthropic_default_sonnet_model":           schema.StringAttribute{Description: "Default model name for the Sonnet class.", Computed: true},
					"anthropic_model":                          schema.StringAttribute{Description: "The primary model to use.", Computed: true},
					"anthropic_small_fast_model":               schema.StringAttribute{Description: "DEPRECATED. Name of the Haiku-class model for background tasks.", Computed: true},
					"anthropic_small_fast_model_aws_region":    schema.StringAttribute{Description: "AWS region for the Haiku-class model when using Bedrock.", Computed: true},
					"bash_default_timeout_ms":                  schema.Int64Attribute{Description: "Default timeout in milliseconds for long-running bash commands.", Computed: true},
					"bash_max_timeout_ms":                      schema.Int64Attribute{Description: "Maximum timeout the model can set for long-running bash commands.", Computed: true},
					"bash_max_output_length":                   schema.Int64Attribute{Description: "Maximum number of characters in bash outputs before truncation.", Computed: true},
					"claude_bash_maintain_project_working_dir": schema.BoolAttribute{Description: "If true, returns to the original working directory after each Bash command.", Computed: true},
					"claude_code_api_key_helper_ttl_ms":        schema.Int64Attribute{Description: "Interval in milliseconds at which credentials should be refreshed when using `apiKeyHelper`.", Computed: true},
					"claude_code_ide_skip_auto_install":        schema.BoolAttribute{Description: "If true, skips auto-installation of IDE extensions.", Computed: true},
					"claude_code_max_output_tokens":            schema.Int64Attribute{Description: "Sets the maximum number of output tokens for most requests.", Computed: true},
					"claude_code_use_bedrock":                  schema.BoolAttribute{Description: "If true, uses AWS Bedrock for models.", Computed: true},
					"claude_code_use_vertex":                   schema.BoolAttribute{Description: "If true, uses Google Vertex AI for models.", Computed: true},
					"claude_code_skip_bedrock_auth":            schema.BoolAttribute{Description: "If true, skips AWS authentication for Bedrock.", Computed: true},
					"claude_code_skip_vertex_auth":             schema.BoolAttribute{Description: "If true, skips Google authentication for Vertex.", Computed: true},
					"claude_code_disable_nonessential_traffic": schema.BoolAttribute{Description: "If true, disables traffic for non-essential features like auto-updates and error reporting.", Computed: true},
					"claude_code_disable_terminal_title":       schema.BoolAttribute{Description: "If true, disables automatic terminal title updates.", Computed: true},
					"claude_code_subagent_model":               schema.StringAttribute{Description: "The model to use for subagents.", Computed: true},
					"disable_autoupdater":                      schema.BoolAttribute{Description: "If true, disables automatic updates.", Computed: true},
					"disable_bug_command":                      schema.BoolAttribute{Description: "If true, disables the `/bug` command.", Computed: true},
					"disable_cost_warnings":                    schema.BoolAttribute{Description: "If true, disables cost warning messages.", Computed: true},
					"disable_error_reporting":                  schema.BoolAttribute{Description: "If true, opts out of Sentry error reporting.", Computed: true},
					"disable_non_essential_model_calls":        schema.BoolAttribute{Description: "If true, disables model calls for non-critical paths like flavor text.", Computed: true},
					"disable_telemetry":                        schema.BoolAttribute{Description: "If true, opts out of Statsig telemetry.", Computed: true},
					"http_proxy":                               schema.StringAttribute{Description: "The URL for an HTTP proxy server.", Computed: true},
					"https_proxy":                              schema.StringAttribute{Description: "The URL for an HTTPS proxy server.", Computed: true},
					"max_thinking_tokens":                      schema.Int64Attribute{Description: "Forces a thinking token budget for the model.", Computed: true},
					"mcp_timeout":                              schema.Int64Attribute{Description: "Timeout in milliseconds for MCP server startup.", Computed: true},
					"mcp_tool_timeout":                         schema.Int64Attribute{Description: "Timeout in milliseconds for MCP tool execution.", Computed: true},
					"max_mcp_output_tokens":                    schema.Int64Attribute{Description: "Maximum number of tokens allowed in MCP tool responses. Default: 25000.", Computed: true},
					"no_proxy":                                 schema.StringAttribute{Description: "A comma-separated list of domains and IPs to bypass the proxy.", Computed: true},
					"use_builtin_ripgrep":                      schema.BoolAttribute{Description: "If false, uses the system-installed `rg` instead of the one bundled with Claude Code.", Computed: true},
					"vertex_region_claude_3_5_haiku":           schema.StringAttribute{Description: "Overrides the AWS region for Claude 3.5 Haiku when using Vertex AI.", Computed: true},
					"vertex_region_claude_3_5_sonnet":          schema.StringAttribute{Description: "Overrides the AWS region for Claude 3.5 Sonnet when using Vertex AI.", Computed: true},
					"vertex_region_claude_3_7_sonnet":          schema.StringAttribute{Description: "Overrides the AWS region for Claude 3.7 Sonnet when using Vertex AI.", Computed: true},
					"vertex_region_claude_4_0_opus":            schema.StringAttribute{Description: "Overrides the AWS region for Claude 4.0 Opus when using Vertex AI.", Computed: true},
					"vertex_region_claude_4_0_sonnet":          schema.StringAttribute{Description: "Overrides the AWS region for Claude 4.0 Sonnet when using Vertex AI.", Computed: true},
					"vertex_region_claude_4_1_opus":            schema.StringAttribute{Description: "Overrides the AWS region for Claude 4.1 Opus when using Vertex AI.", Computed: true},
				},
			},
			"global_config": schema.SingleNestedAttribute{
				Description: "Global configuration settings from `~/.claude/config.json`.",
				Computed:    true,
				Attributes: map[string]schema.Attribute{
					"auto_updates": schema.BoolAttribute{
						Description: "DEPRECATED. Use the DISABLE_AUTOUPDATER environment variable instead.",
						Computed:    true,
					},
					"preferred_notif_channel": schema.StringAttribute{
						Description: "The preferred channel for notifications (e.g., `iterm2`, `terminal_bell`).",
						Computed:    true,
					},
					"theme": schema.StringAttribute{
						Description: "The color theme for the UI (e.g., `dark`, `light`).",
						Computed:    true,
					},
					"verbose": schema.BoolAttribute{
						Description: "If true, shows full bash and command outputs. Default: false.",
						Computed:    true,
					},
				},
			},
			"subagents": schema.ListNestedAttribute{
				Description: "A list of discovered subagents, including their configuration and prompt content.",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"path":        schema.StringAttribute{Description: "The absolute path to the subagent's markdown file.", Computed: true},
						"name":        schema.StringAttribute{Description: "The name of the subagent.", Computed: true},
						"model":       schema.StringAttribute{Description: "The model the subagent is configured to use.", Computed: true},
						"description": schema.StringAttribute{Description: "A description of the subagent's purpose.", Computed: true},
						"color":       schema.StringAttribute{Description: "The color used for the subagent in the UI.", Computed: true},
						"prompt":      schema.StringAttribute{Description: "The full instructional prompt for the subagent.", Computed: true},
					},
				},
			},
			"commands": schema.ListNestedAttribute{
				Description: "A list of discovered custom command files.",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"path": schema.StringAttribute{Description: "The absolute path to the command file.", Computed: true},
					},
				},
			},
			"hook_files": schema.ListNestedAttribute{
				Description: "A list of discovered hook definition files.",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"path": schema.StringAttribute{Description: "The absolute path to the hook file.", Computed: true},
					},
				},
			},
		},
	}
}

// Read refreshes the Terraform state with the latest data.
func (d *claudeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state claudeDataSourceModel

	// Initialize nested objects
	state.EnvironmentVariables = &environmentVariablesModel{}
	state.Settings = &claudeSettingsModel{}

	// Read environment variables
	readEnvironmentVariables(state.EnvironmentVariables)

	// Read and merge settings files
	mergedSettings := d.getMergedSettings(ctx, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Populate state from merged settings
	d.populateSettingsFromMap(ctx, &resp.Diagnostics, state.Settings, mergedSettings)
	if resp.Diagnostics.HasError() {
		return
	}

	// Discover and parse subagents
	state.Subagents = d.readSubagents(ctx, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Discover commands and hooks
	state.Commands = d.readCommands(ctx, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	state.HookFiles = d.readHookFiles(ctx, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set ID
	state.ID = types.StringValue("claude-settings")

	// Save state to Terraform
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (d *claudeDataSource) getMergedSettings(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	paths := d.getSettingsFilePaths(diags)
	if diags.HasError() {
		return nil
	}

	merged := make(map[string]interface{})

	for _, path := range paths {
		if _, err := os.Stat(path); err == nil {
			data, err := d.readSettingsFile(path)
			if err != nil {
				diags.AddError(fmt.Sprintf("Error reading settings file: %s", path), err.Error())
				return nil
			}
			// Simple merge, last one wins
			for k, v := range data {
				merged[k] = v
			}
		}
	}
	return merged
}

func (d *claudeDataSource) getSettingsFilePaths(diagnostics *diag.Diagnostics) []string {
	var paths []string

	// User settings
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		diagnostics.AddError("Unable to get user home directory", err.Error())
		return nil
	}
	paths = append(paths, filepath.Join(userHomeDir, ".claude", "settings.json"))

	// Project settings (shared then local)
	if d.client != nil && d.client.workDir != "" {
		paths = append(paths, filepath.Join(d.client.workDir, ".claude", "settings.json"))
		paths = append(paths, filepath.Join(d.client.workDir, ".claude", "settings.local.json"))
	}

	// Enterprise settings
	var enterprisePath string
	switch runtime.GOOS {
	case "darwin":
		enterprisePath = "/Library/Application Support/ClaudeCode/managed-settings.json"
	case "linux":
		enterprisePath = "/etc/claude-code/managed-settings.json"
	case "windows":
		enterprisePath = "C:\\ProgramData\\ClaudeCode\\managed-settings.json"
	}
	if enterprisePath != "" {
		paths = append(paths, enterprisePath)
	}

	return paths
}

func (d *claudeDataSource) readSettingsFile(path string) (map[string]interface{}, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var data map[string]interface{}
	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (d *claudeDataSource) populateSettingsFromMap(ctx context.Context, diags *diag.Diagnostics, settingsModel *claudeSettingsModel, data map[string]interface{}) {
	if val, ok := data["apiKeyHelper"].(string); ok {
		settingsModel.APIKeyHelper = types.StringValue(val)
	}
	if val, ok := data["cleanupPeriodDays"].(float64); ok {
		settingsModel.CleanupPeriodDays = types.Int64Value(int64(val))
	}
	if val, ok := data["includeCoAuthoredBy"].(bool); ok {
		settingsModel.IncludeCoAuthoredBy = types.BoolValue(val)
	}
	if val, ok := data["disableAllHooks"].(bool); ok {
		settingsModel.DisableAllHooks = types.BoolValue(val)
	}
	if val, ok := data["model"].(string); ok {
		settingsModel.Model = types.StringValue(val)
	}
	if val, ok := data["outputStyle"].(string); ok {
		settingsModel.OutputStyle = types.StringValue(val)
	}
	if val, ok := data["forceLoginMethod"].(string); ok {
		settingsModel.ForceLoginMethod = types.StringValue(val)
	}
	if val, ok := data["forceLoginOrgUUID"].(string); ok {
		settingsModel.ForceLoginOrgUUID = types.StringValue(val)
	}
	if val, ok := data["enableAllProjectMcpServers"].(bool); ok {
		settingsModel.EnableAllProjectMcpServers = types.BoolValue(val)
	}
	if val, ok := data["awsAuthRefresh"].(string); ok {
		settingsModel.AwsAuthRefresh = types.StringValue(val)
	}
	if val, ok := data["awsCredentialExport"].(string); ok {
		settingsModel.AwsCredentialExport = types.StringValue(val)
	}

	if val, ok := data["env"].(map[string]interface{}); ok {
		if len(val) == 0 {
			settingsModel.Env = types.MapNull(types.StringType)
		} else {
			envMap, d := types.MapValueFrom(ctx, types.StringType, val)
			if d.HasError() {
				diags.Append(d...)
				return
			}
			settingsModel.Env = envMap
		}
	} else {
		settingsModel.Env = types.MapNull(types.StringType)
	}

	if val, ok := data["hooks"].(map[string]interface{}); ok {
		if len(val) == 0 {
			settingsModel.Hooks = types.MapNull(types.MapType{ElemType: types.StringType})
		} else {
			hooksMap, d := types.MapValueFrom(ctx, types.MapType{ElemType: types.StringType}, val)
			if d.HasError() {
				diags.Append(d...)
				return
			}
			settingsModel.Hooks = hooksMap
		}
	} else {
		settingsModel.Hooks = types.MapNull(types.MapType{ElemType: types.StringType})
	}

	if val, ok := data["enabledMcpjsonServers"].([]interface{}); ok {
		var servers []types.String
		for _, v := range val {
			if s, ok := v.(string); ok {
				servers = append(servers, types.StringValue(s))
			}
		}
		settingsModel.EnabledMcpjsonServers = servers
	}
	if val, ok := data["disabledMcpjsonServers"].([]interface{}); ok {
		var servers []types.String
		for _, v := range val {
			if s, ok := v.(string); ok {
				servers = append(servers, types.StringValue(s))
			}
		}
		settingsModel.DisabledMcpjsonServers = servers
	}

	if val, ok := data["permissions"].(map[string]interface{}); ok {
		permissions := &permissionsModel{}
		if s, k := val["defaultMode"].(string); k {
			permissions.DefaultMode = types.StringValue(s)
		}
		if s, k := val["disableBypassPermissionsMode"].(string); k {
			permissions.DisableBypassPermissionsMode = types.StringValue(s)
		}
		if list, k := val["allow"].([]interface{}); k {
			var items []types.String
			for _, v := range list {
				if s, ok := v.(string); ok {
					items = append(items, types.StringValue(s))
				}
			}
			permissions.Allow = items
		}
		if list, k := val["ask"].([]interface{}); k {
			var items []types.String
			for _, v := range list {
				if s, ok := v.(string); ok {
					items = append(items, types.StringValue(s))
				}
			}
			permissions.Ask = items
		}
		if list, k := val["deny"].([]interface{}); k {
			var items []types.String
			for _, v := range list {
				if s, ok := v.(string); ok {
					items = append(items, types.StringValue(s))
				}
			}
			permissions.Deny = items
		}
		if list, k := val["additionalDirectories"].([]interface{}); k {
			var items []types.String
			for _, v := range list {
				if s, ok := v.(string); ok {
					items = append(items, types.StringValue(s))
				}
			}
			permissions.AdditionalDirectories = items
		}
		settingsModel.Permissions = permissions
	}

	if val, ok := data["statusLine"].(map[string]interface{}); ok {
		statusLine := &statusLineModel{}
		if s, k := val["type"].(string); k {
			statusLine.Type = types.StringValue(s)
		}
		if s, k := val["command"].(string); k {
			statusLine.Command = types.StringValue(s)
		}
		settingsModel.StatusLine = statusLine
	}
}

func (d *claudeDataSource) readSubagents(ctx context.Context, diags *diag.Diagnostics) []subagentModel {
	var subagents []subagentModel
	dirs := d.getDiscoveryDirs(diags, "agents")
	if diags.HasError() {
		return nil
	}

	for _, dir := range dirs {
		walkErr := filepath.WalkDir(dir, func(path string, de fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if !de.IsDir() && strings.HasSuffix(de.Name(), ".md") {
				content, err := os.ReadFile(path)
				if err != nil {
					diags.AddWarning(fmt.Sprintf("Failed to read subagent file: %s", path), err.Error())
					return nil // Continue walking
				}

				parts := strings.SplitN(string(content), "---", 3)
				if len(parts) < 3 {
					diags.AddWarning(fmt.Sprintf("Invalid format for subagent file: %s", path), "Missing YAML frontmatter separators.")
					return nil // Continue walking
				}

				var frontmatter subagentFrontmatter
				if err := yaml.Unmarshal([]byte(parts[1]), &frontmatter); err != nil {
					diags.AddWarning(fmt.Sprintf("Failed to parse YAML frontmatter for: %s", path), err.Error())
					return nil // Continue walking
				}

				subagents = append(subagents, subagentModel{
					Path:        types.StringValue(path),
					Name:        types.StringValue(frontmatter.Name),
					Model:       types.StringValue(frontmatter.Model),
					Description: types.StringValue(frontmatter.Description),
					Color:       types.StringValue(frontmatter.Color),
					Prompt:      types.StringValue(strings.TrimSpace(parts[2])),
				})
			}
			return nil
		})
		if walkErr != nil {
			diags.AddWarning(fmt.Sprintf("Error walking agents directory %s", dir), walkErr.Error())
		}
	}

	return subagents
}

func (d *claudeDataSource) readCommands(ctx context.Context, diags *diag.Diagnostics) []commandModel {
	var commands []commandModel
	dirs := d.getDiscoveryDirs(diags, "commands")
	if diags.HasError() {
		return nil
	}

	for _, dir := range dirs {
		walkErr := filepath.WalkDir(dir, func(path string, de fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if !de.IsDir() {
				commands = append(commands, commandModel{Path: types.StringValue(path)})
			}
			return nil
		})
		if walkErr != nil {
			diags.AddWarning(fmt.Sprintf("Error walking commands directory %s", dir), walkErr.Error())
		}
	}
	return commands
}

func (d *claudeDataSource) readHookFiles(ctx context.Context, diags *diag.Diagnostics) []hookFileModel {
	var hooks []hookFileModel
	dirs := d.getDiscoveryDirs(diags, "hooks")
	if diags.HasError() {
		return nil
	}

	for _, dir := range dirs {
		walkErr := filepath.WalkDir(dir, func(path string, de fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if !de.IsDir() {
				hooks = append(hooks, hookFileModel{Path: types.StringValue(path)})
			}
			return nil
		})
		if walkErr != nil {
			diags.AddWarning(fmt.Sprintf("Error walking hooks directory %s", dir), walkErr.Error())
		}
	}
	return hooks
}

func (d *claudeDataSource) getDiscoveryDirs(diags *diag.Diagnostics, dirName string) []string {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		diags.AddError("Unable to get user home directory", err.Error())
		return nil
	}

	dirs := []string{
		filepath.Join(userHomeDir, ".claude", dirName),
	}
	if d.client != nil && d.client.workDir != "" {
		dirs = append(dirs, filepath.Join(d.client.workDir, ".claude", dirName))
	}

	var validDirs []string
	for _, dir := range dirs {
		if _, err := os.Stat(dir); !os.IsNotExist(err) {
			validDirs = append(validDirs, dir)
		}
	}
	return validDirs
}

func readEnvironmentVariables(envModel *environmentVariablesModel) {
	envModel.AnthropicAPIKey = types.StringValue(os.Getenv("ANTHROPIC_API_KEY"))
	envModel.AnthropicAuthToken = types.StringValue(os.Getenv("ANTHROPIC_AUTH_TOKEN"))
	envModel.AnthropicCustomHeaders = types.StringValue(os.Getenv("ANTHROPIC_CUSTOM_HEADERS"))
	envModel.AnthropicDefaultHaikuModel = types.StringValue(os.Getenv("ANTHROPIC_DEFAULT_HAIKU_MODEL"))
	envModel.AnthropicDefaultOpusModel = types.StringValue(os.Getenv("ANTHROPIC_DEFAULT_OPUS_MODEL"))
	envModel.AnthropicDefaultSonnetModel = types.StringValue(os.Getenv("ANTHROPIC_DEFAULT_SONNET_MODEL"))
	envModel.AnthropicModel = types.StringValue(os.Getenv("ANTHROPIC_MODEL"))
	envModel.AnthropicSmallFastModel = types.StringValue(os.Getenv("ANTHROPIC_SMALL_FAST_MODEL"))
	envModel.AnthropicSmallFastModelAWSRegion = types.StringValue(os.Getenv("ANTHROPIC_SMALL_FAST_MODEL_AWS_REGION"))
	envModel.AWSBearerTokenBedrock = types.StringValue(os.Getenv("AWS_BEARER_TOKEN_BEDROCK"))
	envModel.ClaudeCodeSubagentModel = types.StringValue(os.Getenv("CLAUDE_CODE_SUBAGENT_MODEL"))
	envModel.HTTPProxy = types.StringValue(os.Getenv("HTTP_PROXY"))
	envModel.HTTPSProxy = types.StringValue(os.Getenv("HTTPS_PROXY"))
	envModel.NoProxy = types.StringValue(os.Getenv("NO_PROXY"))
	envModel.VertexRegionClaude35Haiku = types.StringValue(os.Getenv("VERTEX_REGION_CLAUDE_3_5_HAIKU"))
	envModel.VertexRegionClaude35Sonnet = types.StringValue(os.Getenv("VERTEX_REGION_CLAUDE_3_5_SONNET"))
	envModel.VertexRegionClaude37Sonnet = types.StringValue(os.Getenv("VERTEX_REGION_CLAUDE_3_7_SONNET"))
	envModel.VertexRegionClaude40Opus = types.StringValue(os.Getenv("VERTEX_REGION_CLAUDE_4_0_OPUS"))
	envModel.VertexRegionClaude40Sonnet = types.StringValue(os.Getenv("VERTEX_REGION_CLAUDE_4_0_SONNET"))
	envModel.VertexRegionClaude41Opus = types.StringValue(os.Getenv("VERTEX_REGION_CLAUDE_4_1_OPUS"))

	if val, ok := os.LookupEnv("BASH_DEFAULT_TIMEOUT_MS"); ok {
		if intVal, err := strconv.ParseInt(val, 10, 64); err == nil {
			envModel.BashDefaultTimeoutMS = types.Int64Value(intVal)
		}
	}
	if val, ok := os.LookupEnv("BASH_MAX_TIMEOUT_MS"); ok {
		if intVal, err := strconv.ParseInt(val, 10, 64); err == nil {
			envModel.BashMaxTimeoutMS = types.Int64Value(intVal)
		}
	}
	if val, ok := os.LookupEnv("BASH_MAX_OUTPUT_LENGTH"); ok {
		if intVal, err := strconv.ParseInt(val, 10, 64); err == nil {
			envModel.BashMaxOutputLength = types.Int64Value(intVal)
		}
	}
	if val, ok := os.LookupEnv("CLAUDE_CODE_API_KEY_HELPER_TTL_MS"); ok {
		if intVal, err := strconv.ParseInt(val, 10, 64); err == nil {
			envModel.ClaudeCodeAPIKeyHelperTTLMS = types.Int64Value(intVal)
		}
	}
	if val, ok := os.LookupEnv("CLAUDE_CODE_MAX_OUTPUT_TOKENS"); ok {
		if intVal, err := strconv.ParseInt(val, 10, 64); err == nil {
			envModel.ClaudeCodeMaxOutputTokens = types.Int64Value(intVal)
		}
	}
	if val, ok := os.LookupEnv("MAX_THINKING_TOKENS"); ok {
		if intVal, err := strconv.ParseInt(val, 10, 64); err == nil {
			envModel.MaxThinkingTokens = types.Int64Value(intVal)
		}
	}
	if val, ok := os.LookupEnv("MCP_TIMEOUT"); ok {
		if intVal, err := strconv.ParseInt(val, 10, 64); err == nil {
			envModel.MCPTimeout = types.Int64Value(intVal)
		}
	}
	if val, ok := os.LookupEnv("MCP_TOOL_TIMEOUT"); ok {
		if intVal, err := strconv.ParseInt(val, 10, 64); err == nil {
			envModel.MCPToolTimeout = types.Int64Value(intVal)
		}
	}
	if val, ok := os.LookupEnv("MAX_MCP_OUTPUT_TOKENS"); ok {
		if intVal, err := strconv.ParseInt(val, 10, 64); err == nil {
			envModel.MaxMCPOutputTokens = types.Int64Value(intVal)
		}
	}

	if val := os.Getenv("CLAUDE_BASH_MAINTAIN_PROJECT_WORKING_DIR"); val == "1" || val == "true" {
		envModel.ClaudeBashMaintainProjectWorkingDir = types.BoolValue(true)
	}
	if val := os.Getenv("CLAUDE_CODE_IDE_SKIP_AUTO_INSTALL"); val == "1" || val == "true" {
		envModel.ClaudeCodeIDESkipAutoInstall = types.BoolValue(true)
	}
	if val := os.Getenv("CLAUDE_CODE_USE_BEDROCK"); val == "1" || val == "true" {
		envModel.ClaudeCodeUseBedrock = types.BoolValue(true)
	}
	if val := os.Getenv("CLAUDE_CODE_USE_VERTEX"); val == "1" || val == "true" {
		envModel.ClaudeCodeUseVertex = types.BoolValue(true)
	}
	if val := os.Getenv("CLAUDE_CODE_SKIP_BEDROCK_AUTH"); val == "1" || val == "true" {
		envModel.ClaudeCodeSkipBedrockAuth = types.BoolValue(true)
	}
	if val := os.Getenv("CLAUDE_CODE_SKIP_VERTEX_AUTH"); val == "1" || val == "true" {
		envModel.ClaudeCodeSkipVertexAuth = types.BoolValue(true)
	}
	if val := os.Getenv("CLAUDE_CODE_DISABLE_NONESSENTIAL_TRAFFIC"); val == "1" || val == "true" {
		envModel.ClaudeCodeDisableNonessentialTraffic = types.BoolValue(true)
	}
	if val := os.Getenv("CLAUDE_CODE_DISABLE_TERMINAL_TITLE"); val == "1" || val == "true" {
		envModel.ClaudeCodeDisableTerminalTitle = types.BoolValue(true)
	}
	if val := os.Getenv("DISABLE_AUTOUPDATER"); val == "1" || val == "true" {
		envModel.DisableAutoupdater = types.BoolValue(true)
	}
	if val := os.Getenv("DISABLE_BUG_COMMAND"); val == "1" || val == "true" {
		envModel.DisableBugCommand = types.BoolValue(true)
	}
	if val := os.Getenv("DISABLE_COST_WARNINGS"); val == "1" || val == "true" {
		envModel.DisableCostWarnings = types.BoolValue(true)
	}
	if val := os.Getenv("DISABLE_ERROR_REPORTING"); val == "1" || val == "true" {
		envModel.DisableErrorReporting = types.BoolValue(true)
	}
	if val := os.Getenv("DISABLE_NON_ESSENTIAL_MODEL_CALLS"); val == "1" || val == "true" {
		envModel.DisableNonEssentialModelCalls = types.BoolValue(true)
	}
	if val := os.Getenv("DISABLE_TELEMETRY"); val == "1" || val == "true" {
		envModel.DisableTelemetry = types.BoolValue(true)
	}
	if val := os.Getenv("USE_BUILTIN_RIPGREP"); val == "0" {
		envModel.UseBuiltinRipgrep = types.BoolValue(false)
	} else if val != "" {
		envModel.UseBuiltinRipgrep = types.BoolValue(true)
	}
}

// Configure adds the provider configured client to the data source.
func (d *claudeDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Add a nil check when handling ProviderData because Terraform
	// sets that data after it calls the ConfigureProvider RPC.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*FileClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *FileClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	d.client = client
}
