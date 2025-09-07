package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &geminiDataSource{}
	_ datasource.DataSourceWithConfigure = &geminiDataSource{}
)

// NewGeminiDataSource is a helper function to simplify the provider implementation.
func NewGeminiDataSource() datasource.DataSource {
	return &geminiDataSource{}
}

// geminiDataSource is the data source implementation.
type geminiDataSource struct {
	client *FileClient
}

// geminiDataSourceModel maps the schema to a Go type.
type geminiDataSourceModel struct {
	ID                   types.String                     `tfsdk:"id"`
	Settings             *geminiSettingsModel             `tfsdk:"settings"`
	EnvironmentVariables *geminiEnvironmentVariablesModel `tfsdk:"environment_variables"`
}

// geminiSettingsModel maps the settings nested attribute to a Go type.
type geminiSettingsModel struct {
	General    *generalSettingsModel   `tfsdk:"general"`
	UI         *uiSettingsModel        `tfsdk:"ui"`
	IDE        *ideSettingsModel       `tfsdk:"ide"`
	Privacy    *privacySettingsModel   `tfsdk:"privacy"`
	Model      *modelSettingsModel     `tfsdk:"model"`
	Context    *contextSettingsModel   `tfsdk:"context"`
	Tools      *toolsSettingsModel     `tfsdk:"tools"`
	MCP        *mcpSettingsModel       `tfsdk:"mcp"`
	Security   *securitySettingsModel  `tfsdk:"security"`
	Advanced   *advancedSettingsModel  `tfsdk:"advanced"`
	MCPServers types.Map               `tfsdk:"mcp_servers"`
	Telemetry  *telemetrySettingsModel `tfsdk:"telemetry"`
}

// generalSettingsModel maps the general settings to a Go type.
type generalSettingsModel struct {
	PreferredEditor      types.String `tfsdk:"preferred_editor"`
	VimMode              types.Bool   `tfsdk:"vim_mode"`
	DisableAutoUpdate    types.Bool   `tfsdk:"disable_auto_update"`
	DisableUpdateNag     types.Bool   `tfsdk:"disable_update_nag"`
	CheckpointingEnabled types.Bool   `tfsdk:"checkpointing_enabled"`
}

// uiSettingsModel maps the ui settings to a Go type.
type uiSettingsModel struct {
	Theme                              types.String `tfsdk:"theme"`
	CustomThemes                       types.Map    `tfsdk:"custom_themes"`
	HideWindowTitle                    types.Bool   `tfsdk:"hide_window_title"`
	HideTips                           types.Bool   `tfsdk:"hide_tips"`
	HideBanner                         types.Bool   `tfsdk:"hide_banner"`
	HideFooter                         types.Bool   `tfsdk:"hide_footer"`
	ShowMemoryUsage                    types.Bool   `tfsdk:"show_memory_usage"`
	ShowLineNumbers                    types.Bool   `tfsdk:"show_line_numbers"`
	ShowCitations                      types.Bool   `tfsdk:"show_citations"`
	AccessibilityDisableLoadingPhrases types.Bool   `tfsdk:"accessibility_disable_loading_phrases"`
}

// ideSettingsModel maps the ide settings to a Go type.
type ideSettingsModel struct {
	Enabled      types.Bool `tfsdk:"enabled"`
	HasSeenNudge types.Bool `tfsdk:"has_seen_nudge"`
}

// privacySettingsModel maps the privacy settings to a Go type.
type privacySettingsModel struct {
	UsageStatisticsEnabled types.Bool `tfsdk:"usage_statistics_enabled"`
}

// modelSettingsModel maps the model settings to a Go type.
type modelSettingsModel struct {
	Name                                      types.String `tfsdk:"name"`
	MaxSessionTurns                           types.Int64  `tfsdk:"max_session_turns"`
	SummarizeToolOutput                       types.Map    `tfsdk:"summarize_tool_output"`
	ChatCompressionContextPercentageThreshold types.String `tfsdk:"chat_compression_context_percentage_threshold"`
	SkipNextSpeakerCheck                      types.Bool   `tfsdk:"skip_next_speaker_check"`
}

// contextSettingsModel maps the context settings to a Go type.
type contextSettingsModel struct {
	FileName                               types.List   `tfsdk:"file_name"`
	ImportFormat                           types.String `tfsdk:"import_format"`
	DiscoveryMaxDirs                       types.Int64  `tfsdk:"discovery_max_dirs"`
	IncludeDirectories                     types.List   `tfsdk:"include_directories"`
	LoadFromIncludeDirectories             types.Bool   `tfsdk:"load_from_include_directories"`
	FileFilteringRespectGitIgnore          types.Bool   `tfsdk:"file_filtering_respect_git_ignore"`
	FileFilteringRespectGeminiIgnore       types.Bool   `tfsdk:"file_filtering_respect_gemini_ignore"`
	FileFilteringEnableRecursiveFileSearch types.Bool   `tfsdk:"file_filtering_enable_recursive_file_search"`
}

// toolsSettingsModel maps the tools settings to a Go type.
type toolsSettingsModel struct {
	Sandbox          types.String `tfsdk:"sandbox"`
	UsePty           types.Bool   `tfsdk:"use_pty"`
	Core             types.List   `tfsdk:"core"`
	Exclude          types.List   `tfsdk:"exclude"`
	Allowed          types.List   `tfsdk:"allowed"`
	DiscoveryCommand types.String `tfsdk:"discovery_command"`
	CallCommand      types.String `tfsdk:"call_command"`
}

// mcpSettingsModel maps the mcp settings to a Go type.
type mcpSettingsModel struct {
	ServerCommand types.String `tfsdk:"server_command"`
	Allowed       types.List   `tfsdk:"allowed"`
	Excluded      types.List   `tfsdk:"excluded"`
}

// securitySettingsModel maps the security settings to a Go type.
type securitySettingsModel struct {
	FolderTrustEnabled types.Bool   `tfsdk:"folder_trust_enabled"`
	AuthSelectedType   types.String `tfsdk:"auth_selected_type"`
	AuthEnforcedType   types.String `tfsdk:"auth_enforced_type"`
	AuthUseExternal    types.Bool   `tfsdk:"auth_use_external"`
}

// advancedSettingsModel maps the advanced settings to a Go type.
type advancedSettingsModel struct {
	AutoConfigureMemory types.Bool   `tfsdk:"auto_configure_memory"`
	DNSResolutionOrder  types.String `tfsdk:"dns_resolution_order"`
	ExcludedEnvVars     types.List   `tfsdk:"excluded_env_vars"`
	BugCommand          types.Map    `tfsdk:"bug_command"`
}

// telemetrySettingsModel maps the telemetry settings to a Go type.
type telemetrySettingsModel struct {
	Enabled      types.Bool   `tfsdk:"enabled"`
	Target       types.String `tfsdk:"target"`
	OTLPEndpoint types.String `tfsdk:"otlp_endpoint"`
	OTLPProtocol types.String `tfsdk:"otlp_protocol"`
	LogPrompts   types.Bool   `tfsdk:"log_prompts"`
	Outfile      types.String `tfsdk:"outfile"`
}

// geminiEnvironmentVariablesModel maps Gemini-related environment variables.
type geminiEnvironmentVariablesModel struct {
	GeminiAPIKey                 types.String `tfsdk:"gemini_api_key"`
	GeminiModel                  types.String `tfsdk:"gemini_model"`
	GoogleAPIKey                 types.String `tfsdk:"google_api_key"`
	GoogleCloudProject           types.String `tfsdk:"google_cloud_project"`
	GoogleApplicationCredentials types.String `tfsdk:"google_application_credentials"`
	OTLPGoogleCloudProject       types.String `tfsdk:"otlp_google_cloud_project"`
	GoogleCloudLocation          types.String `tfsdk:"google_cloud_location"`
	GeminiSandbox                types.String `tfsdk:"gemini_sandbox"`
	SeatbeltProfile              types.String `tfsdk:"seatbelt_profile"`
	Debug                        types.Bool   `tfsdk:"debug"`
	DebugMode                    types.Bool   `tfsdk:"debug_mode"`
	NoColor                      types.String `tfsdk:"no_color"`
	CLITitle                     types.String `tfsdk:"cli_title"`
	CodeAssistEndpoint           types.String `tfsdk:"code_assist_endpoint"`
}

// Metadata returns the data source type name.
func (d *geminiDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gemini"
}

// Schema defines the schema for the data source.
func (d *geminiDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Reads and consolidates the complete Gemini CLI configuration from the local environment. This includes settings from system, user, and project `settings.json` files, as well as relevant environment variables.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "A static identifier for the data source.",
				Computed:    true,
			},
			"settings": schema.SingleNestedAttribute{
				Description: "The effective settings after merging all found `settings.json` files, following Gemini CLI's precedence rules.",
				Computed:    true,
				Attributes: map[string]schema.Attribute{
					"general": schema.SingleNestedAttribute{
						Description: "General application settings.",
						Computed:    true,
						Attributes: map[string]schema.Attribute{
							"preferred_editor": schema.StringAttribute{
								Description: "The preferred editor to open files in.",
								Computed:    true,
							},
							"vim_mode": schema.BoolAttribute{
								Description: "If true, enables Vim keybindings in the UI.",
								Computed:    true,
							},
							"disable_auto_update": schema.BoolAttribute{
								Description: "If true, disables automatic updates.",
								Computed:    true,
							},
							"disable_update_nag": schema.BoolAttribute{
								Description: "If true, disables update notification prompts.",
								Computed:    true,
							},
							"checkpointing_enabled": schema.BoolAttribute{
								Description: "If true, enables session checkpointing for recovery.",
								Computed:    true,
							},
						},
					},
					"ui": schema.SingleNestedAttribute{
						Description: "Settings related to the user interface.",
						Computed:    true,
						Attributes: map[string]schema.Attribute{
							"theme": schema.StringAttribute{
								Description: "The color theme for the UI.",
								Computed:    true,
							},
							"custom_themes": schema.MapAttribute{
								Description: "A map of custom theme definitions.",
								ElementType: types.StringType,
								Computed:    true,
							},
							"hide_window_title": schema.BoolAttribute{
								Description: "If true, hides the window title bar.",
								Computed:    true,
							},
							"hide_tips": schema.BoolAttribute{
								Description: "If true, hides helpful tips in the UI.",
								Computed:    true,
							},
							"hide_banner": schema.BoolAttribute{
								Description: "If true, hides the application banner.",
								Computed:    true,
							},
							"hide_footer": schema.BoolAttribute{
								Description: "If true, hides the footer from the UI.",
								Computed:    true,
							},
							"show_memory_usage": schema.BoolAttribute{
								Description: "If true, displays memory usage information in the UI.",
								Computed:    true,
							},
							"show_line_numbers": schema.BoolAttribute{
								Description: "If true, shows line numbers in the chat.",
								Computed:    true,
							},
							"show_citations": schema.BoolAttribute{
								Description: "If true, shows citations for generated text in the chat.",
								Computed:    true,
							},
							"accessibility_disable_loading_phrases": schema.BoolAttribute{
								Description: "If true, disables loading phrases for accessibility.",
								Computed:    true,
							},
						},
					},
					"ide": schema.SingleNestedAttribute{
						Description: "Settings for IDE integration.",
						Computed:    true,
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "If true, enables IDE integration mode.",
								Computed:    true,
							},
							"has_seen_nudge": schema.BoolAttribute{
								Description: "Indicates whether the user has seen the IDE integration nudge.",
								Computed:    true,
							},
						},
					},
					"privacy": schema.SingleNestedAttribute{
						Description: "Settings related to user privacy.",
						Computed:    true,
						Attributes: map[string]schema.Attribute{
							"usage_statistics_enabled": schema.BoolAttribute{
								Description: "If true, enables the collection of anonymized usage statistics. Default: true.",
								Computed:    true,
							},
						},
					},
					"model": schema.SingleNestedAttribute{
						Description: "Settings for the language model.",
						Computed:    true,
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description: "The name of the Gemini model to use for conversations.",
								Computed:    true,
							},
							"max_session_turns": schema.Int64Attribute{
								Description: "Maximum number of user/model/tool turns to keep in a session. A value of -1 means unlimited. Default: -1.",
								Computed:    true,
							},
							"summarize_tool_output": schema.MapAttribute{
								Description: "Configuration for summarizing tool output, including a token budget.",
								ElementType: types.StringType,
								Computed:    true,
							},
							"chat_compression_context_percentage_threshold": schema.StringAttribute{
								Description: "The threshold (0.0-1.0) for chat history compression as a percentage of the model's token limit. Default: 0.7.",
								Computed:    true,
							},
							"skip_next_speaker_check": schema.BoolAttribute{
								Description: "If true, skips the next speaker check.",
								Computed:    true,
							},
						},
					},
					"context": schema.SingleNestedAttribute{
						Description: "Settings for context files and memory management.",
						Computed:    true,
						Attributes: map[string]schema.Attribute{
							"file_name": schema.ListAttribute{
								Description: "The name of the context file(s) to load (e.g., `GEMINI.md`).",
								ElementType: types.StringType,
								Computed:    true,
							},
							"import_format": schema.StringAttribute{
								Description: "The format to use when importing memory.",
								Computed:    true,
							},
							"discovery_max_dirs": schema.Int64Attribute{
								Description: "Maximum number of directories to search for context files. Default: 200.",
								Computed:    true,
							},
							"include_directories": schema.ListAttribute{
								Description: "A list of additional directories to include in the workspace context.",
								ElementType: types.StringType,
								Computed:    true,
							},
							"load_from_include_directories": schema.BoolAttribute{
								Description: "If true, loads context files from all included directories.",
								Computed:    true,
							},
							"file_filtering_respect_git_ignore": schema.BoolAttribute{
								Description: "If true, respects `.gitignore` files when searching for context files. Default: true.",
								Computed:    true,
							},
							"file_filtering_respect_gemini_ignore": schema.BoolAttribute{
								Description: "If true, respects `.geminiignore` files when searching for context files. Default: true.",
								Computed:    true,
							},
							"file_filtering_enable_recursive_file_search": schema.BoolAttribute{
								Description: "If true, enables recursive search for filenames when completing `@` prefixes in the prompt. Default: true.",
								Computed:    true,
							},
						},
					},
					"tools": schema.SingleNestedAttribute{
						Description: "Settings for tool configuration and discovery.",
						Computed:    true,
						Attributes: map[string]schema.Attribute{
							"sandbox": schema.StringAttribute{
								Description: "The sandbox execution environment to use (e.g., `docker`, `podman`).",
								Computed:    true,
							},
							"use_pty": schema.BoolAttribute{
								Description: "If true, uses `node-pty` for shell command execution.",
								Computed:    true,
							},
							"core": schema.ListAttribute{
								Description: "An allowlist to restrict the set of available built-in tools.",
								ElementType: types.StringType,
								Computed:    true,
							},
							"exclude": schema.ListAttribute{
								Description: "A list of tool names to exclude from discovery.",
								ElementType: types.StringType,
								Computed:    true,
							},
							"allowed": schema.ListAttribute{
								Description: "A list of tool names that will bypass the confirmation dialog (e.g., `run_shell_command(git)`).",
								ElementType: types.StringType,
								Computed:    true,
							},
							"discovery_command": schema.StringAttribute{
								Description: "A command to run for custom tool discovery.",
								Computed:    true,
							},
							"call_command": schema.StringAttribute{
								Description: "A custom shell command for calling a discovered tool.",
								Computed:    true,
							},
						},
					},
					"mcp": schema.SingleNestedAttribute{
						Description: "Settings for the Model-Context Protocol (MCP).",
						Computed:    true,
						Attributes: map[string]schema.Attribute{
							"server_command": schema.StringAttribute{
								Description: "A command to start an MCP server.",
								Computed:    true,
							},
							"allowed": schema.ListAttribute{
								Description: "An allowlist of MCP servers to connect to.",
								ElementType: types.StringType,
								Computed:    true,
							},
							"excluded": schema.ListAttribute{
								Description: "A denylist of MCP servers to exclude.",
								ElementType: types.StringType,
								Computed:    true,
							},
						},
					},
					"security": schema.SingleNestedAttribute{
						Description: "Settings related to security and authentication.",
						Computed:    true,
						Attributes: map[string]schema.Attribute{
							"folder_trust_enabled": schema.BoolAttribute{
								Description: "Indicates whether Folder Trust is enabled.",
								Computed:    true,
							},
							"auth_selected_type": schema.StringAttribute{
								Description: "The currently selected authentication type.",
								Computed:    true,
							},
							"auth_enforced_type": schema.StringAttribute{
								Description: "The required authentication type, often used in enterprise environments.",
								Computed:    true,
							},
							"auth_use_external": schema.BoolAttribute{
								Description: "If true, uses an external authentication flow.",
								Computed:    true,
							},
						},
					},
					"advanced": schema.SingleNestedAttribute{
						Description: "Advanced configuration settings.",
						Computed:    true,
						Attributes: map[string]schema.Attribute{
							"auto_configure_memory": schema.BoolAttribute{
								Description: "If true, automatically configures Node.js memory limits.",
								Computed:    true,
							},
							"dns_resolution_order": schema.StringAttribute{
								Description: "The DNS resolution order.",
								Computed:    true,
							},
							"excluded_env_vars": schema.ListAttribute{
								Description: "A list of environment variables to exclude from the project context. Default: `[\"DEBUG\",\"DEBUG_MODE\"]`.",
								ElementType: types.StringType,
								Computed:    true,
							},
							"bug_command": schema.MapAttribute{
								Description: "Configuration for the bug report command.",
								ElementType: types.StringType,
								Computed:    true,
							},
						},
					},
					"mcp_servers": schema.MapAttribute{
						Description: "A map of configurations for individual MCP servers.",
						ElementType: types.StringType,
						Computed:    true,
					},
					"telemetry": schema.SingleNestedAttribute{
						Description: "Settings for logging and metrics configuration.",
						Computed:    true,
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "If true, enables telemetry.",
								Computed:    true,
							},
							"target": schema.StringAttribute{
								Description: "The destination for collected telemetry (`local` or `gcp`).",
								Computed:    true,
							},
							"otlp_endpoint": schema.StringAttribute{
								Description: "The endpoint for the OTLP Exporter.",
								Computed:    true,
							},
							"otlp_protocol": schema.StringAttribute{
								Description: "The protocol for the OTLP Exporter (`grpc` or `http`).",
								Computed:    true,
							},
							"log_prompts": schema.BoolAttribute{
								Description: "If true, includes the content of user prompts in the logs.",
								Computed:    true,
							},
							"outfile": schema.StringAttribute{
								Description: "The file to write telemetry to when the target is `local`.",
								Computed:    true,
							},
						},
					},
				},
			},
			"environment_variables": schema.SingleNestedAttribute{
				Description: "A map of all Gemini CLI related environment variables found in the environment.",
				Computed:    true,
				Attributes: map[string]schema.Attribute{
					"gemini_api_key":                 schema.StringAttribute{Description: "The API key for the Gemini API.", Computed: true, Sensitive: true},
					"google_api_key":                 schema.StringAttribute{Description: "The Google Cloud API key, required for Vertex AI in express mode.", Computed: true, Sensitive: true},
					"gemini_model":                   schema.StringAttribute{Description: "The default Gemini model to use.", Computed: true},
					"google_cloud_project":           schema.StringAttribute{Description: "The Google Cloud Project ID, required for Code Assist or Vertex AI.", Computed: true},
					"google_application_credentials": schema.StringAttribute{Description: "The path to the Google Application Credentials JSON file.", Computed: true},
					"otlp_google_cloud_project":      schema.StringAttribute{Description: "The Google Cloud Project ID for Telemetry in Google Cloud.", Computed: true},
					"google_cloud_location":          schema.StringAttribute{Description: "The Google Cloud Project Location (e.g., `us-central1`).", Computed: true},
					"gemini_sandbox":                 schema.StringAttribute{Description: "The sandbox execution environment to use (`true`, `false`, `docker`, `podman`, or a custom command).", Computed: true},
					"seatbelt_profile":               schema.StringAttribute{Description: "The Seatbelt (`sandbox-exec`) profile to use on macOS (`permissive-open` or `strict`).", Computed: true},
					"debug":                          schema.BoolAttribute{Description: "If true, enables verbose debug logging.", Computed: true},
					"debug_mode":                     schema.BoolAttribute{Description: "If true, enables verbose debug logging.", Computed: true},
					"no_color":                       schema.StringAttribute{Description: "If set, disables all color output in the CLI.", Computed: true},
					"cli_title":                      schema.StringAttribute{Description: "A custom title for the CLI.", Computed: true},
					"code_assist_endpoint":           schema.StringAttribute{Description: "The endpoint for the code assist server.", Computed: true},
				},
			},
		},
	}
}

// Read refreshes the Terraform state with the latest data.
func (d *geminiDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state geminiDataSourceModel

	// Initialize nested objects
	state.EnvironmentVariables = &geminiEnvironmentVariablesModel{}
	state.Settings = &geminiSettingsModel{}

	// Read environment variables
	readGeminiEnvironmentVariables(state.EnvironmentVariables)

	// Read and merge settings files
	mergedSettings := d.getMergedGeminiSettings(ctx, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Populate state from merged settings
	d.populateGeminiSettingsFromMap(ctx, &resp.Diagnostics, state.Settings, mergedSettings)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set ID
	state.ID = types.StringValue("gemini-settings")

	// Save state to Terraform
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (d *geminiDataSource) getMergedGeminiSettings(ctx context.Context, diags *diag.Diagnostics) map[string]interface{} {
	paths := d.getGeminiSettingsFilePaths(diags)
	if diags.HasError() {
		return nil
	}

	merged := make(map[string]interface{})
	foundValidConfig := false

	for _, path := range paths {
		if stat, err := os.Stat(path); err == nil {
			// Check if path is a directory instead of a file
			if stat.IsDir() {
				diags.AddWarning(fmt.Sprintf("Expected file but found directory: %s", path), "Skipping directory in settings file path")
				continue
			}

			data, err := d.readGeminiSettingsFile(path)
			if err != nil {
				// Add as warning instead of error to allow partial configuration loading
				diags.AddWarning(fmt.Sprintf("Error reading Gemini settings file: %s", path), err.Error())
				continue
			}

			if len(data) > 0 {
				foundValidConfig = true
			}

			// Deep merge settings, respecting precedence
			for k, v := range data {
				merged[k] = v
			}
		} else if !os.IsNotExist(err) {
			// File exists but can't be read (permissions, etc.)
			diags.AddWarning(fmt.Sprintf("Cannot access Gemini settings file: %s", path), err.Error())
		}
	}

	if !foundValidConfig {
		diags.AddWarning("No Gemini CLI settings files found",
			fmt.Sprintf("Searched paths: %v. The data source will return empty configuration.", paths))
	}

	return merged
}

func (d *geminiDataSource) getGeminiSettingsFilePaths(diagnostics *diag.Diagnostics) []string {
	var paths []string

	// System defaults file
	var systemDefaultsPath string
	switch runtime.GOOS {
	case "darwin":
		systemDefaultsPath = "/Library/Application Support/GeminiCli/system-defaults.json"
	case "linux":
		systemDefaultsPath = "/etc/gemini-cli/system-defaults.json"
	case "windows":
		systemDefaultsPath = "C:\\ProgramData\\gemini-cli\\system-defaults.json"
	}
	if systemDefaultsPath != "" {
		if envPath := os.Getenv("GEMINI_CLI_SYSTEM_DEFAULTS_PATH"); envPath != "" {
			systemDefaultsPath = envPath
		}
		paths = append(paths, systemDefaultsPath)
	}

	// User settings
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		diagnostics.AddError("Unable to get user home directory", err.Error())
		return nil
	}
	paths = append(paths, filepath.Join(userHomeDir, ".gemini", "settings.json"))

	// Project settings
	if d.client != nil && d.client.workDir != "" {
		paths = append(paths, filepath.Join(d.client.workDir, ".gemini", "settings.json"))
	}

	// System settings file (overrides)
	var systemSettingsPath string
	switch runtime.GOOS {
	case "darwin":
		systemSettingsPath = "/Library/Application Support/GeminiCli/settings.json"
	case "linux":
		systemSettingsPath = "/etc/gemini-cli/settings.json"
	case "windows":
		systemSettingsPath = "C:\\ProgramData\\gemini-cli\\settings.json"
	}
	if systemSettingsPath != "" {
		if envPath := os.Getenv("GEMINI_CLI_SYSTEM_SETTINGS_PATH"); envPath != "" {
			systemSettingsPath = envPath
		}
		paths = append(paths, systemSettingsPath)
	}

	return paths
}

func (d *geminiDataSource) readGeminiSettingsFile(path string) (map[string]interface{}, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read settings file: %w", err)
	}

	// Check if file is empty or contains only whitespace
	if len(bytes) == 0 || len(strings.TrimSpace(string(bytes))) == 0 {
		return make(map[string]interface{}), nil
	}

	var data map[string]interface{}
	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, fmt.Errorf("failed to parse JSON in settings file %s: %w", path, err)
	}

	// Validate that the JSON contains expected structure
	if data == nil {
		return make(map[string]interface{}), nil
	}

	return data, nil
}

func (d *geminiDataSource) populateGeminiSettingsFromMap(ctx context.Context, diags *diag.Diagnostics, settingsModel *geminiSettingsModel, data map[string]interface{}) {
	// Initialize nested objects
	settingsModel.General = &generalSettingsModel{}
	settingsModel.UI = &uiSettingsModel{}
	settingsModel.IDE = &ideSettingsModel{}
	settingsModel.Privacy = &privacySettingsModel{}
	settingsModel.Model = &modelSettingsModel{}
	settingsModel.Context = &contextSettingsModel{}
	settingsModel.Tools = &toolsSettingsModel{}
	settingsModel.MCP = &mcpSettingsModel{}
	settingsModel.Security = &securitySettingsModel{}
	settingsModel.Advanced = &advancedSettingsModel{}
	settingsModel.Telemetry = &telemetrySettingsModel{}

	// Initialize list and map fields with proper types
	settingsModel.Context.FileName = types.ListNull(types.StringType)
	settingsModel.Context.IncludeDirectories = types.ListNull(types.StringType)
	settingsModel.Tools.Core = types.ListNull(types.StringType)
	settingsModel.Tools.Exclude = types.ListNull(types.StringType)
	settingsModel.Tools.Allowed = types.ListNull(types.StringType)
	settingsModel.MCP.Allowed = types.ListNull(types.StringType)
	settingsModel.MCP.Excluded = types.ListNull(types.StringType)
	settingsModel.Advanced.ExcludedEnvVars = types.ListNull(types.StringType)
	settingsModel.UI.CustomThemes = types.MapNull(types.StringType)
	settingsModel.Model.SummarizeToolOutput = types.MapNull(types.StringType)
	settingsModel.Advanced.BugCommand = types.MapNull(types.StringType)
	settingsModel.MCPServers = types.MapNull(types.StringType)

	// Populate general settings
	if generalData, ok := data["general"].(map[string]interface{}); ok {
		if val, ok := generalData["preferredEditor"].(string); ok {
			settingsModel.General.PreferredEditor = types.StringValue(val)
		}
		if val, ok := generalData["vimMode"].(bool); ok {
			settingsModel.General.VimMode = types.BoolValue(val)
		}
		if val, ok := generalData["disableAutoUpdate"].(bool); ok {
			settingsModel.General.DisableAutoUpdate = types.BoolValue(val)
		}
		if val, ok := generalData["disableUpdateNag"].(bool); ok {
			settingsModel.General.DisableUpdateNag = types.BoolValue(val)
		}
		if checkpointingData, ok := generalData["checkpointing"].(map[string]interface{}); ok {
			if val, ok := checkpointingData["enabled"].(bool); ok {
				settingsModel.General.CheckpointingEnabled = types.BoolValue(val)
			}
		}
	}

	// Populate UI settings
	if uiData, ok := data["ui"].(map[string]interface{}); ok {
		if val, ok := uiData["theme"].(string); ok {
			settingsModel.UI.Theme = types.StringValue(val)
		}
		if val, ok := uiData["hideWindowTitle"].(bool); ok {
			settingsModel.UI.HideWindowTitle = types.BoolValue(val)
		}
		if val, ok := uiData["hideTips"].(bool); ok {
			settingsModel.UI.HideTips = types.BoolValue(val)
		}
		if val, ok := uiData["hideBanner"].(bool); ok {
			settingsModel.UI.HideBanner = types.BoolValue(val)
		}
		if val, ok := uiData["hideFooter"].(bool); ok {
			settingsModel.UI.HideFooter = types.BoolValue(val)
		}
		if val, ok := uiData["showMemoryUsage"].(bool); ok {
			settingsModel.UI.ShowMemoryUsage = types.BoolValue(val)
		}
		if val, ok := uiData["showLineNumbers"].(bool); ok {
			settingsModel.UI.ShowLineNumbers = types.BoolValue(val)
		}
		if val, ok := uiData["showCitations"].(bool); ok {
			settingsModel.UI.ShowCitations = types.BoolValue(val)
		}
		if customThemes, ok := uiData["customThemes"].(map[string]interface{}); ok {
			if len(customThemes) > 0 {
				themesMap, d := types.MapValueFrom(ctx, types.StringType, customThemes)
				if d.HasError() {
					diags.Append(d...)
					return
				}
				settingsModel.UI.CustomThemes = themesMap
			} else {
				settingsModel.UI.CustomThemes = types.MapNull(types.StringType)
			}
		}
		if accessibilityData, ok := uiData["accessibility"].(map[string]interface{}); ok {
			if val, ok := accessibilityData["disableLoadingPhrases"].(bool); ok {
				settingsModel.UI.AccessibilityDisableLoadingPhrases = types.BoolValue(val)
			}
		}
	}

	// Populate IDE settings
	if ideData, ok := data["ide"].(map[string]interface{}); ok {
		if val, ok := ideData["enabled"].(bool); ok {
			settingsModel.IDE.Enabled = types.BoolValue(val)
		}
		if val, ok := ideData["hasSeenNudge"].(bool); ok {
			settingsModel.IDE.HasSeenNudge = types.BoolValue(val)
		}
	}

	// Populate Privacy settings
	if privacyData, ok := data["privacy"].(map[string]interface{}); ok {
		if val, ok := privacyData["usageStatisticsEnabled"].(bool); ok {
			settingsModel.Privacy.UsageStatisticsEnabled = types.BoolValue(val)
		}
	}

	// Populate Model settings
	if modelData, ok := data["model"].(map[string]interface{}); ok {
		if val, ok := modelData["name"].(string); ok {
			settingsModel.Model.Name = types.StringValue(val)
		}
		if val, ok := modelData["maxSessionTurns"].(float64); ok {
			settingsModel.Model.MaxSessionTurns = types.Int64Value(int64(val))
		}
		if val, ok := modelData["skipNextSpeakerCheck"].(bool); ok {
			settingsModel.Model.SkipNextSpeakerCheck = types.BoolValue(val)
		}
		if summaryData, ok := modelData["summarizeToolOutput"].(map[string]interface{}); ok {
			if len(summaryData) > 0 {
				stringMap := make(map[string]string)
				for k, v := range summaryData {
					stringMap[k] = fmt.Sprintf("%v", v)
				}
				summaryMap, d := types.MapValueFrom(ctx, types.StringType, stringMap)
				if d.HasError() {
					diags.Append(d...)
					return
				}
				settingsModel.Model.SummarizeToolOutput = summaryMap
			} else {
				settingsModel.Model.SummarizeToolOutput = types.MapNull(types.StringType)
			}
		}
		if chatCompressionData, ok := modelData["chatCompression"].(map[string]interface{}); ok {
			if val, ok := chatCompressionData["contextPercentageThreshold"].(string); ok {
				settingsModel.Model.ChatCompressionContextPercentageThreshold = types.StringValue(val)
			}
		}
	}

	// Populate Context settings
	if contextData, ok := data["context"].(map[string]interface{}); ok {
		if val, ok := contextData["importFormat"].(string); ok {
			settingsModel.Context.ImportFormat = types.StringValue(val)
		}
		if val, ok := contextData["discoveryMaxDirs"].(float64); ok {
			settingsModel.Context.DiscoveryMaxDirs = types.Int64Value(int64(val))
		}
		if val, ok := contextData["loadFromIncludeDirectories"].(bool); ok {
			settingsModel.Context.LoadFromIncludeDirectories = types.BoolValue(val)
		}
		if fileNames, ok := contextData["fileName"].([]interface{}); ok {
			var nameList []types.String
			for i, name := range fileNames {
				if s, ok := name.(string); ok {
					nameList = append(nameList, types.StringValue(s))
				} else {
					diags.AddWarning(fmt.Sprintf("Invalid fileName at index %d", i),
						fmt.Sprintf("Expected string, got %T. Skipping this entry.", name))
				}
			}
			if len(nameList) > 0 {
				fileNamesList, d := types.ListValueFrom(ctx, types.StringType, nameList)
				if d.HasError() {
					diags.Append(d...)
					return
				}
				settingsModel.Context.FileName = fileNamesList
			}
		} else if fileName, ok := contextData["fileName"].(string); ok {
			// Handle single string fileName
			nameList := []types.String{types.StringValue(fileName)}
			fileNamesList, d := types.ListValueFrom(ctx, types.StringType, nameList)
			if d.HasError() {
				diags.Append(d...)
				return
			}
			settingsModel.Context.FileName = fileNamesList
		}
		if includeDirectories, ok := contextData["includeDirectories"].([]interface{}); ok {
			var dirList []types.String
			for _, dir := range includeDirectories {
				if s, ok := dir.(string); ok {
					dirList = append(dirList, types.StringValue(s))
				}
			}
			dirsList, d := types.ListValueFrom(ctx, types.StringType, dirList)
			if d.HasError() {
				diags.Append(d...)
				return
			}
			settingsModel.Context.IncludeDirectories = dirsList
		}
		if fileFilteringData, ok := contextData["fileFiltering"].(map[string]interface{}); ok {
			if val, ok := fileFilteringData["respectGitIgnore"].(bool); ok {
				settingsModel.Context.FileFilteringRespectGitIgnore = types.BoolValue(val)
			}
			if val, ok := fileFilteringData["respectGeminiIgnore"].(bool); ok {
				settingsModel.Context.FileFilteringRespectGeminiIgnore = types.BoolValue(val)
			}
			if val, ok := fileFilteringData["enableRecursiveFileSearch"].(bool); ok {
				settingsModel.Context.FileFilteringEnableRecursiveFileSearch = types.BoolValue(val)
			}
		}
	}

	// Populate Tools settings
	if toolsData, ok := data["tools"].(map[string]interface{}); ok {
		if val, ok := toolsData["sandbox"].(string); ok {
			settingsModel.Tools.Sandbox = types.StringValue(val)
		}
		if val, ok := toolsData["usePty"].(bool); ok {
			settingsModel.Tools.UsePty = types.BoolValue(val)
		}
		if val, ok := toolsData["discoveryCommand"].(string); ok {
			settingsModel.Tools.DiscoveryCommand = types.StringValue(val)
		}
		if val, ok := toolsData["callCommand"].(string); ok {
			settingsModel.Tools.CallCommand = types.StringValue(val)
		}
		if core, ok := toolsData["core"].([]interface{}); ok {
			var coreList []types.String
			for _, tool := range core {
				if s, ok := tool.(string); ok {
					coreList = append(coreList, types.StringValue(s))
				}
			}
			coresList, d := types.ListValueFrom(ctx, types.StringType, coreList)
			if d.HasError() {
				diags.Append(d...)
				return
			}
			settingsModel.Tools.Core = coresList
		}
		if exclude, ok := toolsData["exclude"].([]interface{}); ok {
			var excludeList []types.String
			for _, tool := range exclude {
				if s, ok := tool.(string); ok {
					excludeList = append(excludeList, types.StringValue(s))
				}
			}
			excludesList, d := types.ListValueFrom(ctx, types.StringType, excludeList)
			if d.HasError() {
				diags.Append(d...)
				return
			}
			settingsModel.Tools.Exclude = excludesList
		}
		if allowed, ok := toolsData["allowed"].([]interface{}); ok {
			var allowedList []types.String
			for _, tool := range allowed {
				if s, ok := tool.(string); ok {
					allowedList = append(allowedList, types.StringValue(s))
				}
			}
			allowedsList, d := types.ListValueFrom(ctx, types.StringType, allowedList)
			if d.HasError() {
				diags.Append(d...)
				return
			}
			settingsModel.Tools.Allowed = allowedsList
		}
	}

	// Populate MCP settings
	if mcpData, ok := data["mcp"].(map[string]interface{}); ok {
		if val, ok := mcpData["serverCommand"].(string); ok {
			settingsModel.MCP.ServerCommand = types.StringValue(val)
		}
		if allowed, ok := mcpData["allowed"].([]interface{}); ok {
			var allowedList []types.String
			for _, server := range allowed {
				if s, ok := server.(string); ok {
					allowedList = append(allowedList, types.StringValue(s))
				}
			}
			allowedsList, d := types.ListValueFrom(ctx, types.StringType, allowedList)
			if d.HasError() {
				diags.Append(d...)
				return
			}
			settingsModel.MCP.Allowed = allowedsList
		}
		if excluded, ok := mcpData["excluded"].([]interface{}); ok {
			var excludedList []types.String
			for _, server := range excluded {
				if s, ok := server.(string); ok {
					excludedList = append(excludedList, types.StringValue(s))
				}
			}
			excludedsList, d := types.ListValueFrom(ctx, types.StringType, excludedList)
			if d.HasError() {
				diags.Append(d...)
				return
			}
			settingsModel.MCP.Excluded = excludedsList
		}
	}

	// Populate Security settings
	if securityData, ok := data["security"].(map[string]interface{}); ok {
		if val, ok := securityData["folderTrust"].(map[string]interface{}); ok {
			if enabled, ok := val["enabled"].(bool); ok {
				settingsModel.Security.FolderTrustEnabled = types.BoolValue(enabled)
			}
		}
		if authData, ok := securityData["auth"].(map[string]interface{}); ok {
			if val, ok := authData["selectedType"].(string); ok {
				settingsModel.Security.AuthSelectedType = types.StringValue(val)
			}
			if val, ok := authData["enforcedType"].(string); ok {
				settingsModel.Security.AuthEnforcedType = types.StringValue(val)
			}
			if val, ok := authData["useExternal"].(bool); ok {
				settingsModel.Security.AuthUseExternal = types.BoolValue(val)
			}
		}
	}

	// Populate Advanced settings
	if advancedData, ok := data["advanced"].(map[string]interface{}); ok {
		if val, ok := advancedData["autoConfigureMemory"].(bool); ok {
			settingsModel.Advanced.AutoConfigureMemory = types.BoolValue(val)
		}
		if val, ok := advancedData["dnsResolutionOrder"].(string); ok {
			settingsModel.Advanced.DNSResolutionOrder = types.StringValue(val)
		}
		if excludedVars, ok := advancedData["excludedEnvVars"].([]interface{}); ok {
			var varsList []types.String
			for _, envVar := range excludedVars {
				if s, ok := envVar.(string); ok {
					varsList = append(varsList, types.StringValue(s))
				}
			}
			varsList_, d := types.ListValueFrom(ctx, types.StringType, varsList)
			if d.HasError() {
				diags.Append(d...)
				return
			}
			settingsModel.Advanced.ExcludedEnvVars = varsList_
		}
		if bugCommand, ok := advancedData["bugCommand"].(map[string]interface{}); ok {
			if len(bugCommand) > 0 {
				bugCommandMap, d := types.MapValueFrom(ctx, types.StringType, bugCommand)
				if d.HasError() {
					diags.Append(d...)
					return
				}
				settingsModel.Advanced.BugCommand = bugCommandMap
			} else {
				settingsModel.Advanced.BugCommand = types.MapNull(types.StringType)
			}
		}
	}

	// Populate Telemetry settings
	if telemetryData, ok := data["telemetry"].(map[string]interface{}); ok {
		if val, ok := telemetryData["enabled"].(bool); ok {
			settingsModel.Telemetry.Enabled = types.BoolValue(val)
		}
		if val, ok := telemetryData["target"].(string); ok {
			settingsModel.Telemetry.Target = types.StringValue(val)
		}
		if val, ok := telemetryData["otlpEndpoint"].(string); ok {
			settingsModel.Telemetry.OTLPEndpoint = types.StringValue(val)
		}
		if val, ok := telemetryData["otlpProtocol"].(string); ok {
			settingsModel.Telemetry.OTLPProtocol = types.StringValue(val)
		}
		if val, ok := telemetryData["logPrompts"].(bool); ok {
			settingsModel.Telemetry.LogPrompts = types.BoolValue(val)
		}
		if val, ok := telemetryData["outfile"].(string); ok {
			settingsModel.Telemetry.Outfile = types.StringValue(val)
		}
	}

	// Handle MCP servers
	if mcpServers, ok := data["mcpServers"].(map[string]interface{}); ok {
		if len(mcpServers) > 0 {
			serversMap, d := types.MapValueFrom(ctx, types.StringType, mcpServers)
			if d.HasError() {
				diags.Append(d...)
				return
			}
			settingsModel.MCPServers = serversMap
		} else {
			settingsModel.MCPServers = types.MapNull(types.StringType)
		}
	}
}

func readGeminiEnvironmentVariables(envModel *geminiEnvironmentVariablesModel) {
	envModel.GeminiAPIKey = types.StringValue(os.Getenv("GEMINI_API_KEY"))
	envModel.GeminiModel = types.StringValue(os.Getenv("GEMINI_MODEL"))
	envModel.GoogleAPIKey = types.StringValue(os.Getenv("GOOGLE_API_KEY"))
	envModel.GoogleCloudProject = types.StringValue(os.Getenv("GOOGLE_CLOUD_PROJECT"))
	envModel.GoogleApplicationCredentials = types.StringValue(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
	envModel.OTLPGoogleCloudProject = types.StringValue(os.Getenv("OTLP_GOOGLE_CLOUD_PROJECT"))
	envModel.GoogleCloudLocation = types.StringValue(os.Getenv("GOOGLE_CLOUD_LOCATION"))
	envModel.GeminiSandbox = types.StringValue(os.Getenv("GEMINI_SANDBOX"))
	envModel.SeatbeltProfile = types.StringValue(os.Getenv("SEATBELT_PROFILE"))
	envModel.NoColor = types.StringValue(os.Getenv("NO_COLOR"))
	envModel.CLITitle = types.StringValue(os.Getenv("CLI_TITLE"))
	envModel.CodeAssistEndpoint = types.StringValue(os.Getenv("CODE_ASSIST_ENDPOINT"))

	if val := os.Getenv("DEBUG"); val == "1" || val == "true" {
		envModel.Debug = types.BoolValue(true)
	}
	if val := os.Getenv("DEBUG_MODE"); val == "1" || val == "true" {
		envModel.DebugMode = types.BoolValue(true)
	}
}

// Configure adds the provider configured client to the data source.
func (d *geminiDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

// Helper function to safely convert interface{} arrays to types.String lists
func safeStringArrayToTypesList(ctx context.Context, data []interface{}, fieldName string, diags *diag.Diagnostics) (types.List, bool) {
	var stringList []types.String
	hasValidData := false

	for i, item := range data {
		if s, ok := item.(string); ok {
			stringList = append(stringList, types.StringValue(s))
			hasValidData = true
		} else {
			diags.AddWarning(fmt.Sprintf("Invalid %s at index %d", fieldName, i),
				fmt.Sprintf("Expected string, got %T. Skipping this entry.", item))
		}
	}

	if !hasValidData {
		return types.ListNull(types.StringType), false
	}

	list, d := types.ListValueFrom(ctx, types.StringType, stringList)
	if d.HasError() {
		diags.Append(d...)
		return types.ListNull(types.StringType), false
	}

	return list, true
}

// Helper function to safely convert interface{} maps to types.Map
func safeMapToTypesMap(ctx context.Context, data map[string]interface{}, fieldName string, diags *diag.Diagnostics) (types.Map, bool) {
	if len(data) == 0 {
		return types.MapNull(types.StringType), false
	}

	// Validate all values are convertible to string
	for k, v := range data {
		if _, ok := v.(string); !ok {
			diags.AddWarning(fmt.Sprintf("Invalid %s value for key %s", fieldName, k),
				fmt.Sprintf("Expected string, got %T. This may cause issues.", v))
		}
	}

	mapValue, d := types.MapValueFrom(ctx, types.StringType, data)
	if d.HasError() {
		diags.Append(d...)
		return types.MapNull(types.StringType), false
	}

	return mapValue, true
}
