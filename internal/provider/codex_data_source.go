package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	toml "github.com/pelletier/go-toml/v2"
	"os"
	"path/filepath"
	"sort"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &codexDataSource{}
	_ datasource.DataSourceWithConfigure = &codexDataSource{}
)

// NewCodexDataSource is a helper function to simplify the provider implementation.
func NewCodexDataSource() datasource.DataSource { return &codexDataSource{} }

// codexDataSource is the data source implementation.
type codexDataSource struct {
	client *FileClient
}

// codexDataSourceModel maps the schema to a Go type.
type codexDataSourceModel struct {
	ID            types.String           `tfsdk:"id"`
	SourceFiles   []types.String         `tfsdk:"source_files"`
	ActiveProfile types.String           `tfsdk:"active_profile"`
	Environment   *codexEnvironmentModel `tfsdk:"environment"`
	Effective     *codexEffectiveModel   `tfsdk:"effective_config"`
}

type codexEnvironmentModel struct {
	CodexHome                   types.String `tfsdk:"codex_home"`
	OpenAIBaseURL               types.String `tfsdk:"openai_base_url"`
	CodexSandboxNetworkDisabled types.Bool   `tfsdk:"codex_sandbox_network_disabled"`
}

// Metadata returns the data source type name.
func (d *codexDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_codex"
}

// Schema defines the schema for the data source.
func (d *codexDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Reads and consolidates the complete Codex CLI configuration from the local environment. This includes settings from `config.toml` files and relevant environment variables.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "A static identifier for the data source.",
				Computed:    true,
			},
			"source_files": schema.ListAttribute{
				Description: "A list of the absolute paths of the `config.toml` files that were found and merged, in order of precedence (project file takes precedence over home file).",
				ElementType: types.StringType,
				Computed:    true,
			},
			"active_profile": schema.StringAttribute{
				Description: "The name of the active Codex profile, if one is configured in `config.toml`.",
				Computed:    true,
			},
			"effective_config": schema.SingleNestedAttribute{
				Description: "The final, effective Codex configuration after merging all `config.toml` files, applying the active profile, and considering environment variable overrides.",
				Computed:    true,
				Attributes: map[string]schema.Attribute{
					"model":                              schema.StringAttribute{Description: "The model that Codex should use (e.g., `o3`, `gpt-5`).", Computed: true},
					"model_provider":                     schema.StringAttribute{Description: "The identifier of the model provider to use from the `model_providers` map. Defaults to `openai`.", Computed: true},
					"model_context_window":               schema.Int64Attribute{Description: "The context window size for the model, in tokens.", Computed: true},
					"model_max_output_tokens":            schema.Int64Attribute{Description: "The maximum number of output tokens for the model.", Computed: true},
					"approval_policy":                    schema.StringAttribute{Description: "The approval policy for executing commands (`untrusted`, `on-failure`, `on-request`, `never`).", Computed: true},
					"sandbox_mode":                       schema.StringAttribute{Description: "The OS-level sandbox policy (`read-only`, `workspace-write`, `danger-full-access`).", Computed: true},
					"file_opener":                        schema.StringAttribute{Description: "The editor/URI scheme for hyperlinking file citations (e.g., `vscode`, `cursor`).", Computed: true},
					"hide_agent_reasoning":               schema.BoolAttribute{Description: "If true, suppresses the model's internal 'thinking' events from the output.", Computed: true},
					"show_raw_agent_reasoning":           schema.BoolAttribute{Description: "If true, surfaces the model’s raw chain-of-thought, if available.", Computed: true},
					"model_reasoning_effort":             schema.StringAttribute{Description: "Reasoning effort for Responses API models (`minimal`, `low`, `medium`, `high`).", Computed: true},
					"model_reasoning_summary":            schema.StringAttribute{Description: "Reasoning summary detail for Responses API models (`auto`, `concise`, `detailed`, `none`).", Computed: true},
					"model_verbosity":                    schema.StringAttribute{Description: "Output verbosity for GPT-5 family models (`low`, `medium`, `high`).", Computed: true},
					"model_supports_reasoning_summaries": schema.BoolAttribute{Description: "If true, forces reasoning to be set on requests to the current model.", Computed: true},
					"project_doc_max_bytes":              schema.Int64Attribute{Description: "Maximum number of bytes to read from an `AGENTS.md` file. Defaults to 32 KiB.", Computed: true},
					"notify":                             schema.ListAttribute{Description: "A command and arguments to execute for notifications.", ElementType: types.StringType, Computed: true},
					"sandbox_workspace_write": schema.SingleNestedAttribute{
						Description: "Specific settings for the `workspace-write` sandbox mode.",
						Computed:    true,
						Attributes: map[string]schema.Attribute{
							"writable_roots":         schema.ListAttribute{Description: "A list of additional writable root paths beyond the defaults.", ElementType: types.StringType, Computed: true},
							"network_access":         schema.BoolAttribute{Description: "If true, allows the command being run inside the sandbox to make outbound network requests. Default: false.", Computed: true},
							"exclude_tmpdir_env_var": schema.BoolAttribute{Description: "If true, excludes the `$TMPDIR` environment variable from writable roots.", Computed: true},
							"exclude_slash_tmp":      schema.BoolAttribute{Description: "If true, excludes the `/tmp` directory from writable roots.", Computed: true},
						},
					},
					"history": schema.SingleNestedAttribute{
						Description: "Settings for command history persistence.",
						Computed:    true,
						Attributes: map[string]schema.Attribute{
							"persistence": schema.StringAttribute{Description: "History persistence mode. `save-all` (default) saves history, `none` disables it.", Computed: true},
							"max_bytes":   schema.Int64Attribute{Description: "The maximum size of the history file in bytes.", Computed: true},
						},
					},
					"shell_environment_policy": schema.SingleNestedAttribute{
						Description: "Policy for managing environment variables passed to subprocesses.",
						Computed:    true,
						Attributes: map[string]schema.Attribute{
							"inherit":                 schema.StringAttribute{Description: "The starting template for the environment: `all` (default), `core`, or `none`.", Computed: true},
							"ignore_default_excludes": schema.BoolAttribute{Description: "If false (default), automatically removes variables containing `KEY`, `SECRET`, or `TOKEN`.", Computed: true},
							"exclude":                 schema.ListAttribute{Description: "A list of case-insensitive glob patterns for environment variables to exclude.", ElementType: types.StringType, Computed: true},
							"set":                     schema.MapAttribute{Description: "A map of key/value pairs to explicitly set or override.", ElementType: types.StringType, Computed: true},
							"include_only":            schema.ListAttribute{Description: "If non-empty, acts as a whitelist of glob patterns for variables to keep.", ElementType: types.StringType, Computed: true},
						},
					},
					"model_providers": schema.ListNestedAttribute{
						Description: "A list of configured model providers.",
						Computed:    true,
						NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{
							"id":                     schema.StringAttribute{Description: "The unique identifier for the model provider.", Computed: true},
							"name":                   schema.StringAttribute{Description: "The display name of the provider.", Computed: true},
							"base_url":               schema.StringAttribute{Description: "The base URL for the provider's API.", Computed: true},
							"env_key":                schema.StringAttribute{Description: "The environment variable that holds the API key for this provider.", Computed: true},
							"env_key_is_set":         schema.BoolAttribute{Description: "Indicates whether the API key environment variable is set.", Computed: true},
							"wire_api":               schema.StringAttribute{Description: "The wire protocol to use (`chat` or `responses`). Defaults to `chat`.", Computed: true},
							"query_params":           schema.MapAttribute{Description: "A map of extra query parameters to add to requests (e.g., `api-version` for Azure).", ElementType: types.StringType, Computed: true},
							"http_headers":           schema.MapAttribute{Description: "A map of static HTTP headers to add to requests.", ElementType: types.StringType, Computed: true},
							"env_http_headers":       schema.MapAttribute{Description: "A map of HTTP headers to add to requests, with values sourced from environment variables.", ElementType: types.StringType, Computed: true},
							"request_max_retries":    schema.Int64Attribute{Description: "How many times to retry a failed HTTP request. Default: 4.", Computed: true},
							"stream_max_retries":     schema.Int64Attribute{Description: "How many times to reconnect a dropped streaming response. Default: 5.", Computed: true},
							"stream_idle_timeout_ms": schema.Int64Attribute{Description: "How long in milliseconds to wait for activity on a streaming response before timing out. Default: 300000 (5 minutes).", Computed: true},
						}},
					},
					"mcp_servers": schema.ListNestedAttribute{
						Description: "A list of configured MCP (Model-Context Protocol) servers for custom tools.",
						Computed:    true,
						NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{
							"id":      schema.StringAttribute{Description: "The unique identifier for the MCP server.", Computed: true},
							"command": schema.StringAttribute{Description: "The command to execute to start the server.", Computed: true},
							"args":    schema.ListAttribute{Description: "A list of arguments for the command.", ElementType: types.StringType, Computed: true},
							"env":     schema.MapAttribute{Description: "A map of environment variables to set for the server process.", ElementType: types.StringType, Computed: true, Sensitive: true},
						}},
					},
					"profiles": schema.ListNestedAttribute{
						Description: "A list of all configuration profiles defined in `config.toml`.",
						Computed:    true,
						NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{
							"name":            schema.StringAttribute{Description: "The name of the profile.", Computed: true},
							"model":           schema.StringAttribute{Description: "The model associated with this profile.", Computed: true},
							"model_provider":  schema.StringAttribute{Description: "The model provider associated with this profile.", Computed: true},
							"approval_policy": schema.StringAttribute{Description: "The approval policy associated with this profile.", Computed: true},
							"sandbox_mode":    schema.StringAttribute{Description: "The sandbox mode associated with this profile.", Computed: true},
						}},
					},
				},
			},
			"environment": schema.SingleNestedAttribute{
				Description: "Environment signals used by Codex that influence configuration resolution.",
				Computed:    true,
				Attributes: map[string]schema.Attribute{
					"codex_home":                     schema.StringAttribute{Description: "The resolved path to the Codex home directory (from `$CODEX_HOME` or `~/.codex`).", Computed: true},
					"openai_base_url":                schema.StringAttribute{Description: "The value of the `OPENAI_BASE_URL` environment variable, if set.", Computed: true},
					"codex_sandbox_network_disabled": schema.BoolAttribute{Description: "The value of the `CODEX_SANDBOX_NETWORK_DISABLED` environment variable, if set.", Computed: true},
				},
			},
		},
	}
}

// Read refreshes the Terraform state with the latest data.
func (d *codexDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state codexDataSourceModel

	// Initialize nested objects
	state.Environment = &codexEnvironmentModel{}
	state.Effective = &codexEffectiveModel{}

	// Resolve CODEX_HOME
	codexHome := os.Getenv("CODEX_HOME")
	if codexHome == "" {
		userHomeDir, err := os.UserHomeDir()
		if err != nil {
			resp.Diagnostics.AddError("Unable to get user home directory", err.Error())
			return
		}
		codexHome = filepath.Join(userHomeDir, ".codex")
	}
	state.Environment.CodexHome = types.StringValue(codexHome)

	// Discover config files in precedence order: home then project
	var paths []string
	homeConfig := filepath.Join(codexHome, "config.toml")
	if fileExists(homeConfig) {
		paths = append(paths, homeConfig)
	}
	if d.client != nil && d.client.workDir != "" {
		projectConfig := filepath.Join(d.client.workDir, ".codex", "config.toml")
		if fileExists(projectConfig) {
			// project has higher precedence, so list after home
			paths = append(paths, projectConfig)
		}
	}
	// Convert to Terraform values
	state.SourceFiles = toTFStringList(paths)

	// Environment derived values (non‑sensitive)
	state.Environment.OpenAIBaseURL = types.StringValue(os.Getenv("OPENAI_BASE_URL"))
	if v := os.Getenv("CODEX_SANDBOX_NETWORK_DISABLED"); v == "1" || v == "true" {
		state.Environment.CodexSandboxNetworkDisabled = types.BoolValue(true)
	} else if v != "" {
		state.Environment.CodexSandboxNetworkDisabled = types.BoolValue(false)
	}

	// Parse and merge TOML configs
	base := rawCodexConfig{}
	for _, p := range paths {
		bytes, err := os.ReadFile(p)
		if err != nil {
			resp.Diagnostics.AddError(fmt.Sprintf("Failed to read Codex config file: %s", p), err.Error())
			return
		}
		var cfg rawCodexConfig
		if err := toml.Unmarshal(bytes, &cfg); err != nil {
			resp.Diagnostics.AddError(fmt.Sprintf("Failed to parse TOML in: %s", p), err.Error())
			return
		}
		base = mergeCodexConfig(base, cfg)
	}

	// Determine active profile
	if base.Profile != "" {
		state.ActiveProfile = types.StringValue(base.Profile)
		if prof, ok := base.Profiles[base.Profile]; ok {
			base = mergeCodexConfig(base, prof)
		}
	} else {
		state.ActiveProfile = types.StringNull()
	}

	// Environment overrides
	openaiBaseURL := os.Getenv("OPENAI_BASE_URL")
	if openaiBaseURL != "" {
		// Override built-in openai provider base_url if present
		if base.ModelProviders == nil {
			base.ModelProviders = map[string]modelProviderConfig{}
		}
		mp := base.ModelProviders["openai"]
		mp.BaseURL = openaiBaseURL
		base.ModelProviders["openai"] = mp
	}

	// Compute env_key_is_set flags
	for id, mp := range base.ModelProviders {
		if mp.EnvKey != "" {
			if _, ok := os.LookupEnv(mp.EnvKey); ok {
				mp.EnvKeyIsSet = boolPtr(true)
			} else {
				mp.EnvKeyIsSet = boolPtr(false)
			}
			base.ModelProviders[id] = mp
		}
	}

	// Populate Effective model
	populateEffectiveModel(state.Effective, &base)

	// Set ID
	state.ID = types.StringValue("codex-config")

	// Save state to Terraform
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// Configure adds the provider configured client to the data source.
func (d *codexDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

// Helpers
func fileExists(path string) bool {
	if path == "" {
		return false
	}
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

func toTFStringList(items []string) []types.String {
	out := make([]types.String, 0, len(items))
	for _, s := range items {
		out = append(out, types.StringValue(s))
	}
	return out
}

// ---------- Internal config types and helpers ----------

type rawCodexConfig struct {
	Model                           string                 `toml:"model"`
	ModelProvider                   string                 `toml:"model_provider"`
	ModelContextWindow              *int64                 `toml:"model_context_window"`
	ModelMaxOutputTokens            *int64                 `toml:"model_max_output_tokens"`
	ApprovalPolicy                  string                 `toml:"approval_policy"`
	SandboxMode                     string                 `toml:"sandbox_mode"`
	SandboxWorkspaceWrite           *sandboxWorkspaceWrite `toml:"sandbox_workspace_write"`
	Notify                          []string               `toml:"notify"`
	History                         *historyConfig         `toml:"history"`
	FileOpener                      string                 `toml:"file_opener"`
	HideAgentReasoning              *bool                  `toml:"hide_agent_reasoning"`
	ShowRawAgentReasoning           *bool                  `toml:"show_raw_agent_reasoning"`
	ModelReasoningEffort            string                 `toml:"model_reasoning_effort"`
	ModelReasoningSummary           string                 `toml:"model_reasoning_summary"`
	ModelVerbosity                  string                 `toml:"model_verbosity"`
	ModelSupportsReasoningSummaries *bool                  `toml:"model_supports_reasoning_summaries"`
	ProjectDocMaxBytes              *int64                 `toml:"project_doc_max_bytes"`
	ShellEnvironmentPolicy          *shellEnvPolicy        `toml:"shell_environment_policy"`

	// Dynamic tables
	ModelProviders map[string]modelProviderConfig `toml:"model_providers"`
	MCPServers     map[string]mcpServerConfig     `toml:"mcp_servers"`

	// Profiles
	Profile  string                    `toml:"profile"`
	Profiles map[string]rawCodexConfig `toml:"profiles"`
}

type sandboxWorkspaceWrite struct {
	WritableRoots       []string `toml:"writable_roots"`
	NetworkAccess       *bool    `toml:"network_access"`
	ExcludeTmpdirEnvVar *bool    `toml:"exclude_tmpdir_env_var"`
	ExcludeSlashTmp     *bool    `toml:"exclude_slash_tmp"`
}

type historyConfig struct {
	Persistence string `toml:"persistence"`
	MaxBytes    *int64 `toml:"max_bytes"`
}

type shellEnvPolicy struct {
	Inherit               string            `toml:"inherit"`
	IgnoreDefaultExcludes *bool             `toml:"ignore_default_excludes"`
	Exclude               []string          `toml:"exclude"`
	Set                   map[string]string `toml:"set"`
	IncludeOnly           []string          `toml:"include_only"`
}

type modelProviderConfig struct {
	Name                string            `toml:"name"`
	BaseURL             string            `toml:"base_url"`
	EnvKey              string            `toml:"env_key"`
	WireAPI             string            `toml:"wire_api"`
	QueryParams         map[string]string `toml:"query_params"`
	HTTPHeaders         map[string]string `toml:"http_headers"`
	EnvHTTPHeaders      map[string]string `toml:"env_http_headers"`
	RequestMaxRetries   *int64            `toml:"request_max_retries"`
	StreamMaxRetries    *int64            `toml:"stream_max_retries"`
	StreamIdleTimeoutMS *int64            `toml:"stream_idle_timeout_ms"`

	// Computed, not from TOML
	EnvKeyIsSet *bool `toml:"-"`
}

type mcpServerConfig struct {
	Command string            `toml:"command"`
	Args    []string          `toml:"args"`
	Env     map[string]string `toml:"env"`
}

// Merge cfg2 onto cfg1, returning a new config
func mergeCodexConfig(cfg1, cfg2 rawCodexConfig) rawCodexConfig {
	out := cfg1
	// Scalars
	if cfg2.Model != "" {
		out.Model = cfg2.Model
	}
	if cfg2.ModelProvider != "" {
		out.ModelProvider = cfg2.ModelProvider
	}
	if cfg2.ModelContextWindow != nil {
		out.ModelContextWindow = cfg2.ModelContextWindow
	}
	if cfg2.ModelMaxOutputTokens != nil {
		out.ModelMaxOutputTokens = cfg2.ModelMaxOutputTokens
	}
	if cfg2.ApprovalPolicy != "" {
		out.ApprovalPolicy = cfg2.ApprovalPolicy
	}
	if cfg2.SandboxMode != "" {
		out.SandboxMode = cfg2.SandboxMode
	}
	if cfg2.FileOpener != "" {
		out.FileOpener = cfg2.FileOpener
	}
	if cfg2.ModelReasoningEffort != "" {
		out.ModelReasoningEffort = cfg2.ModelReasoningEffort
	}
	if cfg2.ModelReasoningSummary != "" {
		out.ModelReasoningSummary = cfg2.ModelReasoningSummary
	}
	if cfg2.ModelVerbosity != "" {
		out.ModelVerbosity = cfg2.ModelVerbosity
	}
	if cfg2.ModelSupportsReasoningSummaries != nil {
		out.ModelSupportsReasoningSummaries = cfg2.ModelSupportsReasoningSummaries
	}
	if cfg2.ProjectDocMaxBytes != nil {
		out.ProjectDocMaxBytes = cfg2.ProjectDocMaxBytes
	}
	if cfg2.HideAgentReasoning != nil {
		out.HideAgentReasoning = cfg2.HideAgentReasoning
	}
	if cfg2.ShowRawAgentReasoning != nil {
		out.ShowRawAgentReasoning = cfg2.ShowRawAgentReasoning
	}
	if len(cfg2.Notify) > 0 {
		out.Notify = cfg2.Notify
	}

	// Nested structs
	if cfg2.SandboxWorkspaceWrite != nil {
		if out.SandboxWorkspaceWrite == nil {
			out.SandboxWorkspaceWrite = &sandboxWorkspaceWrite{}
		}
		out.SandboxWorkspaceWrite = mergeSandbox(out.SandboxWorkspaceWrite, cfg2.SandboxWorkspaceWrite)
	}
	if cfg2.History != nil {
		if out.History == nil {
			out.History = &historyConfig{}
		}
		out.History = mergeHistory(out.History, cfg2.History)
	}
	if cfg2.ShellEnvironmentPolicy != nil {
		if out.ShellEnvironmentPolicy == nil {
			out.ShellEnvironmentPolicy = &shellEnvPolicy{}
		}
		out.ShellEnvironmentPolicy = mergeShellEnv(out.ShellEnvironmentPolicy, cfg2.ShellEnvironmentPolicy)
	}

	// Maps with merging behavior
	if cfg2.ModelProviders != nil {
		if out.ModelProviders == nil {
			out.ModelProviders = map[string]modelProviderConfig{}
		}
		for k, v := range cfg2.ModelProviders {
			if existing, ok := out.ModelProviders[k]; ok {
				out.ModelProviders[k] = mergeModelProvider(existing, v)
			} else {
				out.ModelProviders[k] = v
			}
		}
	}
	if cfg2.MCPServers != nil {
		if out.MCPServers == nil {
			out.MCPServers = map[string]mcpServerConfig{}
		}
		for k, v := range cfg2.MCPServers {
			if existing, ok := out.MCPServers[k]; ok {
				out.MCPServers[k] = mergeMCP(existing, v)
			} else {
				out.MCPServers[k] = v
			}
		}
	}

	// Profiles
	if cfg2.Profile != "" {
		out.Profile = cfg2.Profile
	}
	if cfg2.Profiles != nil {
		if out.Profiles == nil {
			out.Profiles = map[string]rawCodexConfig{}
		}
		for k, v := range cfg2.Profiles {
			out.Profiles[k] = mergeCodexConfig(out.Profiles[k], v)
		}
	}

	return out
}

func mergeSandbox(a, b *sandboxWorkspaceWrite) *sandboxWorkspaceWrite {
	out := *a
	if len(b.WritableRoots) > 0 {
		out.WritableRoots = b.WritableRoots
	}
	if b.NetworkAccess != nil {
		out.NetworkAccess = b.NetworkAccess
	}
	if b.ExcludeTmpdirEnvVar != nil {
		out.ExcludeTmpdirEnvVar = b.ExcludeTmpdirEnvVar
	}
	if b.ExcludeSlashTmp != nil {
		out.ExcludeSlashTmp = b.ExcludeSlashTmp
	}
	return &out
}

func mergeHistory(a, b *historyConfig) *historyConfig {
	out := *a
	if b.Persistence != "" {
		out.Persistence = b.Persistence
	}
	if b.MaxBytes != nil {
		out.MaxBytes = b.MaxBytes
	}
	return &out
}

func mergeShellEnv(a, b *shellEnvPolicy) *shellEnvPolicy {
	out := *a
	if b.Inherit != "" {
		out.Inherit = b.Inherit
	}
	if b.IgnoreDefaultExcludes != nil {
		out.IgnoreDefaultExcludes = b.IgnoreDefaultExcludes
	}
	if len(b.Exclude) > 0 {
		out.Exclude = b.Exclude
	}
	if len(b.Set) > 0 {
		out.Set = b.Set
	}
	if len(b.IncludeOnly) > 0 {
		out.IncludeOnly = b.IncludeOnly
	}
	return &out
}

func mergeModelProvider(a, b modelProviderConfig) modelProviderConfig {
	out := a
	if b.Name != "" {
		out.Name = b.Name
	}
	if b.BaseURL != "" {
		out.BaseURL = b.BaseURL
	}
	if b.EnvKey != "" {
		out.EnvKey = b.EnvKey
	}
	if b.WireAPI != "" {
		out.WireAPI = b.WireAPI
	}
	if len(b.QueryParams) > 0 {
		out.QueryParams = b.QueryParams
	}
	if len(b.HTTPHeaders) > 0 {
		out.HTTPHeaders = b.HTTPHeaders
	}
	if len(b.EnvHTTPHeaders) > 0 {
		out.EnvHTTPHeaders = b.EnvHTTPHeaders
	}
	if b.RequestMaxRetries != nil {
		out.RequestMaxRetries = b.RequestMaxRetries
	}
	if b.StreamMaxRetries != nil {
		out.StreamMaxRetries = b.StreamMaxRetries
	}
	if b.StreamIdleTimeoutMS != nil {
		out.StreamIdleTimeoutMS = b.StreamIdleTimeoutMS
	}
	return out
}

func mergeMCP(a, b mcpServerConfig) mcpServerConfig {
	out := a
	if b.Command != "" {
		out.Command = b.Command
	}
	if len(b.Args) > 0 {
		out.Args = b.Args
	}
	if len(b.Env) > 0 {
		out.Env = b.Env
	}
	return out
}

func boolPtr(b bool) *bool { return &b }

// ---------- Terraform model population ----------

type codexEffectiveModel struct {
	Model                           types.String   `tfsdk:"model"`
	ModelProvider                   types.String   `tfsdk:"model_provider"`
	ModelContextWindow              types.Int64    `tfsdk:"model_context_window"`
	ModelMaxOutputTokens            types.Int64    `tfsdk:"model_max_output_tokens"`
	ApprovalPolicy                  types.String   `tfsdk:"approval_policy"`
	SandboxMode                     types.String   `tfsdk:"sandbox_mode"`
	FileOpener                      types.String   `tfsdk:"file_opener"`
	HideAgentReasoning              types.Bool     `tfsdk:"hide_agent_reasoning"`
	ShowRawAgentReasoning           types.Bool     `tfsdk:"show_raw_agent_reasoning"`
	ModelReasoningEffort            types.String   `tfsdk:"model_reasoning_effort"`
	ModelReasoningSummary           types.String   `tfsdk:"model_reasoning_summary"`
	ModelVerbosity                  types.String   `tfsdk:"model_verbosity"`
	ModelSupportsReasoningSummaries types.Bool     `tfsdk:"model_supports_reasoning_summaries"`
	ProjectDocMaxBytes              types.Int64    `tfsdk:"project_doc_max_bytes"`
	Notify                          []types.String `tfsdk:"notify"`

	SandboxWorkspaceWrite  *codexSandboxWorkspaceWriteModel `tfsdk:"sandbox_workspace_write"`
	History                *codexHistoryModel               `tfsdk:"history"`
	ShellEnvironmentPolicy *codexShellEnvPolicyModel        `tfsdk:"shell_environment_policy"`

	// Dynamic tables
	ModelProviders []codexModelProviderModel `tfsdk:"model_providers"`
	MCPServers     []codexMCPServerModel     `tfsdk:"mcp_servers"`
	Profiles       []codexProfileModel       `tfsdk:"profiles"`
}

type codexSandboxWorkspaceWriteModel struct {
	WritableRoots       []types.String `tfsdk:"writable_roots"`
	NetworkAccess       types.Bool     `tfsdk:"network_access"`
	ExcludeTmpdirEnvVar types.Bool     `tfsdk:"exclude_tmpdir_env_var"`
	ExcludeSlashTmp     types.Bool     `tfsdk:"exclude_slash_tmp"`
}

type codexHistoryModel struct {
	Persistence types.String `tfsdk:"persistence"`
	MaxBytes    types.Int64  `tfsdk:"max_bytes"`
}

type codexShellEnvPolicyModel struct {
	Inherit               types.String            `tfsdk:"inherit"`
	IgnoreDefaultExcludes types.Bool              `tfsdk:"ignore_default_excludes"`
	Exclude               []types.String          `tfsdk:"exclude"`
	Set                   map[string]types.String `tfsdk:"set"`
	IncludeOnly           []types.String          `tfsdk:"include_only"`
}

func populateEffectiveModel(dst *codexEffectiveModel, src *rawCodexConfig) {
	// Scalars
	if src.Model != "" {
		dst.Model = types.StringValue(src.Model)
	}
	if src.ModelProvider != "" {
		dst.ModelProvider = types.StringValue(src.ModelProvider)
	}
	if src.ModelContextWindow != nil {
		dst.ModelContextWindow = types.Int64Value(*src.ModelContextWindow)
	}
	if src.ModelMaxOutputTokens != nil {
		dst.ModelMaxOutputTokens = types.Int64Value(*src.ModelMaxOutputTokens)
	}
	if src.ApprovalPolicy != "" {
		dst.ApprovalPolicy = types.StringValue(src.ApprovalPolicy)
	}
	if src.SandboxMode != "" {
		dst.SandboxMode = types.StringValue(src.SandboxMode)
	}
	if src.FileOpener != "" {
		dst.FileOpener = types.StringValue(src.FileOpener)
	}
	if src.HideAgentReasoning != nil {
		dst.HideAgentReasoning = types.BoolValue(*src.HideAgentReasoning)
	}
	if src.ShowRawAgentReasoning != nil {
		dst.ShowRawAgentReasoning = types.BoolValue(*src.ShowRawAgentReasoning)
	}
	if src.ModelReasoningEffort != "" {
		dst.ModelReasoningEffort = types.StringValue(src.ModelReasoningEffort)
	}
	if src.ModelReasoningSummary != "" {
		dst.ModelReasoningSummary = types.StringValue(src.ModelReasoningSummary)
	}
	if src.ModelVerbosity != "" {
		dst.ModelVerbosity = types.StringValue(src.ModelVerbosity)
	}
	if src.ModelSupportsReasoningSummaries != nil {
		dst.ModelSupportsReasoningSummaries = types.BoolValue(*src.ModelSupportsReasoningSummaries)
	}
	if src.ProjectDocMaxBytes != nil {
		dst.ProjectDocMaxBytes = types.Int64Value(*src.ProjectDocMaxBytes)
	}
	if len(src.Notify) > 0 {
		dst.Notify = toTFStringList(src.Notify)
	}

	// Sandbox workspace write
	if src.SandboxWorkspaceWrite != nil {
		dst.SandboxWorkspaceWrite = &codexSandboxWorkspaceWriteModel{}
		if len(src.SandboxWorkspaceWrite.WritableRoots) > 0 {
			dst.SandboxWorkspaceWrite.WritableRoots = toTFStringList(src.SandboxWorkspaceWrite.WritableRoots)
		}
		if src.SandboxWorkspaceWrite.NetworkAccess != nil {
			dst.SandboxWorkspaceWrite.NetworkAccess = types.BoolValue(*src.SandboxWorkspaceWrite.NetworkAccess)
		}
		if src.SandboxWorkspaceWrite.ExcludeTmpdirEnvVar != nil {
			dst.SandboxWorkspaceWrite.ExcludeTmpdirEnvVar = types.BoolValue(*src.SandboxWorkspaceWrite.ExcludeTmpdirEnvVar)
		}
		if src.SandboxWorkspaceWrite.ExcludeSlashTmp != nil {
			dst.SandboxWorkspaceWrite.ExcludeSlashTmp = types.BoolValue(*src.SandboxWorkspaceWrite.ExcludeSlashTmp)
		}
	}

	// History
	if src.History != nil {
		dst.History = &codexHistoryModel{}
		if src.History.Persistence != "" {
			dst.History.Persistence = types.StringValue(src.History.Persistence)
		}
		if src.History.MaxBytes != nil {
			dst.History.MaxBytes = types.Int64Value(*src.History.MaxBytes)
		}
	}

	// Shell env policy
	if src.ShellEnvironmentPolicy != nil {
		dst.ShellEnvironmentPolicy = &codexShellEnvPolicyModel{}
		if src.ShellEnvironmentPolicy.Inherit != "" {
			dst.ShellEnvironmentPolicy.Inherit = types.StringValue(src.ShellEnvironmentPolicy.Inherit)
		}
		if src.ShellEnvironmentPolicy.IgnoreDefaultExcludes != nil {
			dst.ShellEnvironmentPolicy.IgnoreDefaultExcludes = types.BoolValue(*src.ShellEnvironmentPolicy.IgnoreDefaultExcludes)
		}
		if len(src.ShellEnvironmentPolicy.Exclude) > 0 {
			dst.ShellEnvironmentPolicy.Exclude = toTFStringList(src.ShellEnvironmentPolicy.Exclude)
		}
		if len(src.ShellEnvironmentPolicy.Set) > 0 {
			dst.ShellEnvironmentPolicy.Set = make(map[string]types.String, len(src.ShellEnvironmentPolicy.Set))
			for k, v := range src.ShellEnvironmentPolicy.Set {
				dst.ShellEnvironmentPolicy.Set[k] = types.StringValue(v)
			}
		}
		if len(src.ShellEnvironmentPolicy.IncludeOnly) > 0 {
			dst.ShellEnvironmentPolicy.IncludeOnly = toTFStringList(src.ShellEnvironmentPolicy.IncludeOnly)
		}
	}

	// Dynamic tables
	dst.ModelProviders = flattenModelProviders(src)
	dst.MCPServers = flattenMCPServers(src)
	dst.Profiles = flattenProfiles(src)
}

type codexModelProviderModel struct {
	ID                  types.String `tfsdk:"id"`
	Name                types.String `tfsdk:"name"`
	BaseURL             types.String `tfsdk:"base_url"`
	EnvKey              types.String `tfsdk:"env_key"`
	EnvKeyIsSet         types.Bool   `tfsdk:"env_key_is_set"`
	WireAPI             types.String `tfsdk:"wire_api"`
	QueryParams         types.Map    `tfsdk:"query_params"`
	HTTPHeaders         types.Map    `tfsdk:"http_headers"`
	EnvHTTPHeaders      types.Map    `tfsdk:"env_http_headers"`
	RequestMaxRetries   types.Int64  `tfsdk:"request_max_retries"`
	StreamMaxRetries    types.Int64  `tfsdk:"stream_max_retries"`
	StreamIdleTimeoutMS types.Int64  `tfsdk:"stream_idle_timeout_ms"`
}

type codexMCPServerModel struct {
	ID      types.String   `tfsdk:"id"`
	Command types.String   `tfsdk:"command"`
	Args    []types.String `tfsdk:"args"`
	Env     types.Map      `tfsdk:"env"`
}

type codexProfileModel struct {
	Name           types.String `tfsdk:"name"`
	Model          types.String `tfsdk:"model"`
	ModelProvider  types.String `tfsdk:"model_provider"`
	ApprovalPolicy types.String `tfsdk:"approval_policy"`
	SandboxMode    types.String `tfsdk:"sandbox_mode"`
}

func flattenModelProviders(src *rawCodexConfig) []codexModelProviderModel {
	if len(src.ModelProviders) == 0 {
		return nil
	}
	ids := make([]string, 0, len(src.ModelProviders))
	for id := range src.ModelProviders {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	out := make([]codexModelProviderModel, 0, len(ids))
	for _, id := range ids {
		mp := src.ModelProviders[id]
		var q, h, eh types.Map
		// Build maps lazily using helper wrapper; since we don't have ctx here, construct via types.MapValueMust
		// Use types.MapValueFrom requires context, which we avoid. The zero value (null) is acceptable when empty.
		if len(mp.QueryParams) > 0 {
			q = mapToTypesMapString(mp.QueryParams)
		}
		if len(mp.HTTPHeaders) > 0 {
			h = mapToTypesMapString(mp.HTTPHeaders)
		}
		if len(mp.EnvHTTPHeaders) > 0 {
			eh = mapToTypesMapString(mp.EnvHTTPHeaders)
		}
		item := codexModelProviderModel{
			ID:             types.StringValue(id),
			Name:           types.StringValue(mp.Name),
			BaseURL:        types.StringValue(mp.BaseURL),
			EnvKey:         types.StringValue(mp.EnvKey),
			WireAPI:        types.StringValue(mp.WireAPI),
			QueryParams:    q,
			HTTPHeaders:    h,
			EnvHTTPHeaders: eh,
		}
		if mp.EnvKeyIsSet != nil {
			item.EnvKeyIsSet = types.BoolValue(*mp.EnvKeyIsSet)
		}
		if mp.RequestMaxRetries != nil {
			item.RequestMaxRetries = types.Int64Value(*mp.RequestMaxRetries)
		}
		if mp.StreamMaxRetries != nil {
			item.StreamMaxRetries = types.Int64Value(*mp.StreamMaxRetries)
		}
		if mp.StreamIdleTimeoutMS != nil {
			item.StreamIdleTimeoutMS = types.Int64Value(*mp.StreamIdleTimeoutMS)
		}
		out = append(out, item)
	}
	return out
}

func flattenMCPServers(src *rawCodexConfig) []codexMCPServerModel {
	if len(src.MCPServers) == 0 {
		return nil
	}
	ids := make([]string, 0, len(src.MCPServers))
	for id := range src.MCPServers {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	out := make([]codexMCPServerModel, 0, len(ids))
	for _, id := range ids {
		s := src.MCPServers[id]
		var env types.Map
		if len(s.Env) > 0 {
			env = mapToTypesMapString(s.Env)
		}
		out = append(out, codexMCPServerModel{
			ID:      types.StringValue(id),
			Command: types.StringValue(s.Command),
			Args:    toTFStringList(s.Args),
			Env:     env,
		})
	}
	return out
}

func flattenProfiles(src *rawCodexConfig) []codexProfileModel {
	if len(src.Profiles) == 0 {
		return nil
	}
	names := make([]string, 0, len(src.Profiles))
	for name := range src.Profiles {
		names = append(names, name)
	}
	sort.Strings(names)
	out := make([]codexProfileModel, 0, len(names))
	for _, name := range names {
		p := src.Profiles[name]
		item := codexProfileModel{Name: types.StringValue(name)}
		if p.Model != "" {
			item.Model = types.StringValue(p.Model)
		}
		if p.ModelProvider != "" {
			item.ModelProvider = types.StringValue(p.ModelProvider)
		}
		if p.ApprovalPolicy != "" {
			item.ApprovalPolicy = types.StringValue(p.ApprovalPolicy)
		}
		if p.SandboxMode != "" {
			item.SandboxMode = types.StringValue(p.SandboxMode)
		}
		out = append(out, item)
	}
	return out
}

// Build a types.Map[string]string without context; use Must constructor
func mapToTypesMapString(in map[string]string) types.Map {
	return types.MapValueMust(types.StringType, mapStringToAttrValue(in))
}

func mapStringToAttrValue(in map[string]string) map[string]attr.Value {
	out := make(map[string]attr.Value, len(in))
	for k, v := range in {
		out[k] = types.StringValue(v)
	}
	return out
}
