package provider

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-agentsmith/internal/codex/configio"

	toml "github.com/pelletier/go-toml/v2"
	"os"
)

var (
	_ resource.Resource                = &codexConfigResource{}
	_ resource.ResourceWithConfigure   = &codexConfigResource{}
	_ resource.ResourceWithImportState = &codexConfigResource{}
)

func NewCodexConfigResource() resource.Resource { return &codexConfigResource{} }

type codexConfigResource struct {
	client *FileClient
}

// Resource model
type codexConfigResourceModel struct {
	// Meta
	ID                      types.String `tfsdk:"id"`
	Scope                   types.String `tfsdk:"scope"`
	Path                    types.String `tfsdk:"path"`
	ResolvedPath            types.String `tfsdk:"resolved_path"`
	MergeStrategy           types.String `tfsdk:"merge_strategy"`
	CreateDirectories       types.Bool   `tfsdk:"create_directories"`
	FileMode                types.String `tfsdk:"file_mode"`
	BackupOnWrite           types.Bool   `tfsdk:"backup_on_write"`
	AllowSensitiveEnvWrites types.Bool   `tfsdk:"allow_sensitive_env_writes"`
	KeepFileOnDestroy       types.Bool   `tfsdk:"keep_file_on_destroy"`
	ValidateStrict          types.Bool   `tfsdk:"validate_strict"`

	// Config root
	Profile                         types.String   `tfsdk:"profile"`
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
	TUI                             types.Map      `tfsdk:"tui"`

	SandboxWorkspaceWrite *codexSandboxWorkspaceWriteModelR `tfsdk:"sandbox_workspace_write"`
	History               *codexHistoryModelR               `tfsdk:"history"`
	ShellEnvPolicy        *codexShellEnvPolicyModelR        `tfsdk:"shell_environment_policy"`

	ModelProviders []codexModelProviderWriteModel `tfsdk:"model_providers"`
	MCPServers     []codexMCPServerWriteModel     `tfsdk:"mcp_servers"`
	Profiles       []codexProfileWriteModel       `tfsdk:"profiles"`
}

type codexSandboxWorkspaceWriteModelR struct {
	WritableRoots       []types.String `tfsdk:"writable_roots"`
	NetworkAccess       types.Bool     `tfsdk:"network_access"`
	ExcludeTmpdirEnvVar types.Bool     `tfsdk:"exclude_tmpdir_env_var"`
	ExcludeSlashTmp     types.Bool     `tfsdk:"exclude_slash_tmp"`
}

type codexHistoryModelR struct {
	Persistence types.String `tfsdk:"persistence"`
	MaxBytes    types.Int64  `tfsdk:"max_bytes"`
}

type codexShellEnvPolicyModelR struct {
	Inherit               types.String   `tfsdk:"inherit"`
	IgnoreDefaultExcludes types.Bool     `tfsdk:"ignore_default_excludes"`
	Exclude               []types.String `tfsdk:"exclude"`
	Set                   types.Map      `tfsdk:"set"`
	IncludeOnly           []types.String `tfsdk:"include_only"`
}

type codexModelProviderWriteModel struct {
	ID                  types.String `tfsdk:"id"`
	Name                types.String `tfsdk:"name"`
	BaseURL             types.String `tfsdk:"base_url"`
	EnvKey              types.String `tfsdk:"env_key"`
	WireAPI             types.String `tfsdk:"wire_api"`
	QueryParams         types.Map    `tfsdk:"query_params"`
	HTTPHeaders         types.Map    `tfsdk:"http_headers"`
	EnvHTTPHeaders      types.Map    `tfsdk:"env_http_headers"`
	RequestMaxRetries   types.Int64  `tfsdk:"request_max_retries"`
	StreamMaxRetries    types.Int64  `tfsdk:"stream_max_retries"`
	StreamIdleTimeoutMS types.Int64  `tfsdk:"stream_idle_timeout_ms"`
}

type codexMCPServerWriteModel struct {
	ID      types.String   `tfsdk:"id"`
	Command types.String   `tfsdk:"command"`
	Args    []types.String `tfsdk:"args"`
	Env     types.Map      `tfsdk:"env"`
}

type codexProfileWriteModel struct {
	Name           types.String `tfsdk:"name"`
	Model          types.String `tfsdk:"model"`
	ModelProvider  types.String `tfsdk:"model_provider"`
	ApprovalPolicy types.String `tfsdk:"approval_policy"`
	SandboxMode    types.String `tfsdk:"sandbox_mode"`
}

func (r *codexConfigResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_codex_config"
}

func (r *codexConfigResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a Codex CLI `config.toml` file. This resource can operate at different scopes to manage home, project, or custom configurations.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "A stable identifier for the resource, derived from a SHA256 hash of the resolved file path.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"scope": schema.StringAttribute{
				Description:   "The scope of the configuration file. Must be one of `home` (~/.codex/config.toml), `project` (<workdir>/.codex/config.toml), or `custom`.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"path": schema.StringAttribute{
				Description:   "The absolute path to the `config.toml` file. Required only when `scope` is `custom`.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"resolved_path": schema.StringAttribute{
				Description:   "The fully resolved absolute path to the managed `config.toml` file.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"merge_strategy": schema.StringAttribute{
				Description: "Defines how to handle existing `config.toml` files. `preserve_unknown` (default) merges managed settings while keeping unmanaged keys. `replace_all` overwrites the file with only managed settings. `fail_on_unknown` fails if the existing file contains unmanaged keys.",
				Optional:    true,
			},
			"create_directories":         schema.BoolAttribute{Description: "If true, creates parent directories for the config file if they do not exist. Defaults to `true`.", Optional: true},
			"file_mode":                  schema.StringAttribute{Description: "The file mode to set on the `config.toml` file, in octal format (e.g., `0600`). Defaults to `0600`.", Optional: true},
			"backup_on_write":            schema.BoolAttribute{Description: "If true, creates a `.bak` file before writing changes. Defaults to `true`.", Optional: true},
			"allow_sensitive_env_writes": schema.BoolAttribute{Description: "If true, allows writing sensitive environment variables for MCP servers to the config file. Defaults to `false`.", Optional: true},
			"keep_file_on_destroy":       schema.BoolAttribute{Description: "If true, the `config.toml` file will be kept on disk when the resource is destroyed. Defaults to `true`.", Optional: true},
			"validate_strict":            schema.BoolAttribute{Description: "If true, performs strict validation of the final configuration against the Codex schema. Defaults to `true`.", Optional: true},

			// Top-level scalars
			"profile":                            schema.StringAttribute{Description: "The name of the active profile to use from the `profiles` block.", Optional: true},
			"model":                              schema.StringAttribute{Description: "The model that Codex should use (e.g., `o3`, `gpt-5`).", Optional: true},
			"model_provider":                     schema.StringAttribute{Description: "The identifier of the model provider to use. Defaults to `openai`.", Optional: true},
			"model_context_window":               schema.Int64Attribute{Description: "The context window size for the model, in tokens.", Optional: true},
			"model_max_output_tokens":            schema.Int64Attribute{Description: "The maximum number of output tokens for the model.", Optional: true},
			"approval_policy":                    schema.StringAttribute{Description: "Determines when to prompt for command execution approval. Can be `untrusted`, `on-failure`, `on-request`, or `never`.", Optional: true},
			"sandbox_mode":                       schema.StringAttribute{Description: "The OS-level sandbox policy for executing commands. Can be `read-only`, `workspace-write`, or `danger-full-access`.", Optional: true},
			"file_opener":                        schema.StringAttribute{Description: "The editor/URI scheme for hyperlinking file citations (e.g., `vscode`, `cursor`, `none`).", Optional: true},
			"hide_agent_reasoning":               schema.BoolAttribute{Description: "If true, suppresses the model's internal 'thinking' events from the output.", Optional: true},
			"show_raw_agent_reasoning":           schema.BoolAttribute{Description: "If true, surfaces the model’s raw chain-of-thought, if available.", Optional: true},
			"model_reasoning_effort":             schema.StringAttribute{Description: "Reasoning effort for Responses API models (`minimal`, `low`, `medium`, `high`).", Optional: true},
			"model_reasoning_summary":            schema.StringAttribute{Description: "Reasoning summary detail for Responses API models (`auto`, `concise`, `detailed`, `none`).", Optional: true},
			"model_verbosity":                    schema.StringAttribute{Description: "Output verbosity for GPT-5 family models (`low`, `medium`, `high`).", Optional: true},
			"model_supports_reasoning_summaries": schema.BoolAttribute{Description: "If true, forces reasoning to be set on requests to the current model.", Optional: true},
			"project_doc_max_bytes":              schema.Int64Attribute{Description: "Maximum number of bytes to read from an `AGENTS.md` file. Defaults to 32 KiB.", Optional: true},
			"notify":                             schema.ListAttribute{Description: "A command and its arguments to execute for notifications.", ElementType: types.StringType, Optional: true},
			"tui":                                schema.MapAttribute{Description: "A map of options specific to the Text User Interface (TUI).", ElementType: types.StringType, Optional: true},
		},
		Blocks: map[string]schema.Block{
			"sandbox_workspace_write": schema.SingleNestedBlock{
				Description: "Specific settings for the `workspace-write` sandbox mode.",
				Attributes: map[string]schema.Attribute{
					"writable_roots":         schema.ListAttribute{Description: "A list of additional writable root paths beyond the defaults.", ElementType: types.StringType, Optional: true},
					"network_access":         schema.BoolAttribute{Description: "If true, allows the command being run inside the sandbox to make outbound network requests. Default: false.", Optional: true},
					"exclude_tmpdir_env_var": schema.BoolAttribute{Description: "If true, excludes the `$TMPDIR` environment variable from writable roots.", Optional: true},
					"exclude_slash_tmp":      schema.BoolAttribute{Description: "If true, excludes the `/tmp` directory from writable roots.", Optional: true},
				},
			},
			"history": schema.SingleNestedBlock{
				Description: "Settings for command history persistence.",
				Attributes: map[string]schema.Attribute{
					"persistence": schema.StringAttribute{Description: "History persistence mode. `save-all` (default) saves history to `$CODEX_HOME/history.jsonl`, `none` disables it.", Optional: true},
					"max_bytes":   schema.Int64Attribute{Description: "The maximum size of the history file in bytes.", Optional: true},
				},
			},
			"shell_environment_policy": schema.SingleNestedBlock{
				Description: "Policy for managing environment variables passed to subprocesses.",
				Attributes: map[string]schema.Attribute{
					"inherit":                 schema.StringAttribute{Description: "The starting template for the environment: `all` (default), `core`, or `none`.", Optional: true},
					"ignore_default_excludes": schema.BoolAttribute{Description: "If false (default), automatically removes variables containing `KEY`, `SECRET`, or `TOKEN`.", Optional: true},
					"exclude":                 schema.ListAttribute{Description: "A list of case-insensitive glob patterns for environment variables to exclude.", ElementType: types.StringType, Optional: true},
					"set":                     schema.MapAttribute{Description: "A map of key/value pairs to explicitly set or override.", ElementType: types.StringType, Optional: true},
					"include_only":            schema.ListAttribute{Description: "If non-empty, acts as a whitelist of glob patterns for variables to keep.", ElementType: types.StringType, Optional: true},
				},
			},
			"model_providers": schema.ListNestedBlock{
				Description: "Defines a set of model providers that can be used by Codex.",
				NestedObject: schema.NestedBlockObject{Attributes: map[string]schema.Attribute{
					"id":                     schema.StringAttribute{Description: "The unique identifier for the model provider (e.g., `openai-chat-completions`).", Required: true},
					"name":                   schema.StringAttribute{Description: "The display name of the provider.", Optional: true},
					"base_url":               schema.StringAttribute{Description: "The base URL for the provider's API.", Optional: true},
					"env_key":                schema.StringAttribute{Description: "The environment variable that holds the API key for this provider.", Optional: true},
					"wire_api":               schema.StringAttribute{Description: "The wire protocol to use (`chat` or `responses`). Defaults to `chat`.", Optional: true},
					"query_params":           schema.MapAttribute{Description: "A map of extra query parameters to add to requests (e.g., `api-version` for Azure).", ElementType: types.StringType, Optional: true},
					"http_headers":           schema.MapAttribute{Description: "A map of static HTTP headers to add to requests.", ElementType: types.StringType, Optional: true},
					"env_http_headers":       schema.MapAttribute{Description: "A map of HTTP headers to add to requests, with values sourced from environment variables.", ElementType: types.StringType, Optional: true},
					"request_max_retries":    schema.Int64Attribute{Description: "How many times to retry a failed HTTP request. Default: 4.", Optional: true},
					"stream_max_retries":     schema.Int64Attribute{Description: "How many times to reconnect a dropped streaming response. Default: 5.", Optional: true},
					"stream_idle_timeout_ms": schema.Int64Attribute{Description: "How long in milliseconds to wait for activity on a streaming response before timing out. Default: 300000 (5 minutes).", Optional: true},
				}},
			},
			"mcp_servers": schema.ListNestedBlock{
				Description: "Defines a set of MCP (Model-Context Protocol) servers for custom tool discovery.",
				NestedObject: schema.NestedBlockObject{Attributes: map[string]schema.Attribute{
					"id":      schema.StringAttribute{Description: "The unique identifier for the MCP server.", Required: true},
					"command": schema.StringAttribute{Description: "The command to execute to start the server.", Required: true},
					"args":    schema.ListAttribute{Description: "A list of arguments for the command.", ElementType: types.StringType, Optional: true},
					"env":     schema.MapAttribute{Description: "A map of environment variables to set for the server process.", ElementType: types.StringType, Optional: true, Sensitive: true},
				}},
			},
			"profiles": schema.ListNestedBlock{
				Description: "Defines a set of configuration profiles.",
				NestedObject: schema.NestedBlockObject{Attributes: map[string]schema.Attribute{
					"name":            schema.StringAttribute{Description: "The name of the profile.", Required: true},
					"model":           schema.StringAttribute{Description: "The model associated with this profile.", Optional: true},
					"model_provider":  schema.StringAttribute{Description: "The model provider associated with this profile.", Optional: true},
					"approval_policy": schema.StringAttribute{Description: "The approval policy associated with this profile.", Optional: true},
					"sandbox_mode":    schema.StringAttribute{Description: "The sandbox mode associated with this profile.", Optional: true},
				}},
			},
		},
	}
}

func (r *codexConfigResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan codexConfigResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	path, err := r.resolveTargetPath(plan)
	if err != nil {
		resp.Diagnostics.AddError("Path resolution error", err.Error())
		return
	}

	merged, resolvedPath, err := r.mergeIntoExisting(ctx, &resp.Diagnostics, plan, path)
	if err != nil || resp.Diagnostics.HasError() {
		return
	}

	// Determine header comment for replace_all
	header := ""
	if strings.EqualFold(plan.MergeStrategy.ValueString(), "replace_all") {
		header = "# Managed by terraform-provider-agentsmith: agentsmith_codex_config"
	}
	if err := configio.AtomicWrite(resolvedPath, merged, strOrDefault(plan.FileMode, "0600"), boolOrDefault(plan.BackupOnWrite, true), boolOrDefault(plan.CreateDirectories, true), header); err != nil {
		resp.Diagnostics.AddError("Failed to write config.toml", err.Error())
		return
	}

	// Set computed meta
	plan.ResolvedPath = types.StringValue(resolvedPath)
	plan.ID = types.StringValue(configio.SHA256Hex(resolvedPath))

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *codexConfigResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state codexConfigResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	path, err := r.resolveTargetPath(state)
	if err != nil {
		resp.Diagnostics.AddError("Path resolution error", err.Error())
		return
	}

	// If file missing, drop from state
	if path == "" {
		resp.State.RemoveResource(ctx)
		return
	}
	if _, err := osStat(path); err != nil { // helper allows test overrides if needed
		resp.State.RemoveResource(ctx)
		return
	}

	// Keep meta; we deliberately avoid full content round-trip in MVP
	state.ResolvedPath = types.StringValue(path)
	state.ID = types.StringValue(configio.SHA256Hex(path))
	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *codexConfigResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan codexConfigResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	path, err := r.resolveTargetPath(plan)
	if err != nil {
		resp.Diagnostics.AddError("Path resolution error", err.Error())
		return
	}

	merged, resolvedPath, err := r.mergeIntoExisting(ctx, &resp.Diagnostics, plan, path)
	if err != nil || resp.Diagnostics.HasError() {
		return
	}

	header := ""
	if strings.EqualFold(plan.MergeStrategy.ValueString(), "replace_all") {
		header = "# Managed by terraform-provider-agentsmith: agentsmith_codex_config"
	}
	if err := configio.AtomicWrite(resolvedPath, merged, strOrDefault(plan.FileMode, "0600"), boolOrDefault(plan.BackupOnWrite, true), boolOrDefault(plan.CreateDirectories, true), header); err != nil {
		resp.Diagnostics.AddError("Failed to write config.toml", err.Error())
		return
	}

	plan.ResolvedPath = types.StringValue(resolvedPath)
	plan.ID = types.StringValue(configio.SHA256Hex(resolvedPath))
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *codexConfigResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state codexConfigResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	keep := boolOrDefault(state.KeepFileOnDestroy, true)
	if keep {
		// For MVP, we leave the file intact. Future: strip managed keys when preserve_unknown.
		return
	}
	path, err := r.resolveTargetPath(state)
	if err != nil {
		resp.Diagnostics.AddError("Path resolution error", err.Error())
		return
	}
	if path == "" {
		return
	}
	if err := osRemove(path); err != nil && !os.IsNotExist(err) {
		resp.Diagnostics.AddError("Failed to delete file", err.Error())
	}
}

func (r *codexConfigResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import by absolute path
	p := strings.TrimSpace(req.ID)
	if p == "" {
		resp.Diagnostics.AddError("Invalid import id", "Provide absolute path to config.toml")
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("path"), p)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("scope"), "custom")...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("resolved_path"), p)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), configio.SHA256Hex(p))...)
}

func (r *codexConfigResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client, ok := req.ProviderData.(*FileClient)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Resource Configure Type", fmt.Sprintf("Expected *FileClient, got: %T", req.ProviderData))
		return
	}
	r.client = client
}

// mergeIntoExisting performs merge according to merge_strategy and returns encoded TOML bytes.
func (r *codexConfigResource) mergeIntoExisting(ctx context.Context, diags *diag.Diagnostics, plan codexConfigResourceModel, resolvedPath string) ([]byte, string, error) {
	// Prepare desired map
	desired, derr := r.planToMap(ctx, diags, &plan)
	if derr != nil {
		return nil, resolvedPath, derr
	}

	// Read existing
	existing, err := configio.ReadTOMLMap(resolvedPath)
	if err != nil {
		diags.AddError("Failed to read existing file", err.Error())
		return nil, resolvedPath, err
	}

	strategy := strings.ToLower(plan.MergeStrategy.ValueString())
	if strategy == "" {
		strategy = "preserve_unknown"
	}

	if strings.EqualFold(strategy, "fail_on_unknown") {
		unk := configio.ValidateUnknownKeys(existing)
		if len(unk) > 0 {
			sort.Strings(unk)
			return nil, resolvedPath, fmt.Errorf("fail_on_unknown: file contains unknown keys: %s", strings.Join(unk, ", "))
		}
	}

	var merged map[string]any
	switch strategy {
	case "replace_all":
		merged = desired
	case "preserve_unknown", "fail_on_unknown":
		merged = configio.MergePreserveUnknown(existing, desired)
	default:
		merged = configio.MergePreserveUnknown(existing, desired)
	}

	// Strict validation: try marshal/unmarshal via typed struct to catch shape errors
	if boolOrDefault(plan.ValidateStrict, true) {
		var typed rawCodexConfig
		b, err := configio.MarshalTOML(merged)
		if err != nil {
			return nil, resolvedPath, err
		}
		if err := toml.Unmarshal(b, &typed); err != nil {
			return nil, resolvedPath, fmt.Errorf("strict validation failed: %w", err)
		}
		return b, resolvedPath, nil
	}

	b, err := configio.MarshalTOML(merged)
	return b, resolvedPath, err
}

func (r *codexConfigResource) resolveTargetPath(m codexConfigResourceModel) (string, error) {
	scope := m.Scope.ValueString()
	custom := m.Path.ValueString()
	workdir := ""
	if r.client != nil {
		workdir = r.client.workDir
	}
	return configio.ResolvePath(scope, workdir, custom)
}

func (r *codexConfigResource) planToMap(ctx context.Context, diags *diag.Diagnostics, m *codexConfigResourceModel) (map[string]any, error) {
	out := map[string]any{}

	setStringIfPresent := func(key string, v types.String) {
		if !v.IsNull() && !v.IsUnknown() && v.ValueString() != "" {
			out[key] = v.ValueString()
		}
	}
	setIntIfPresent := func(key string, v types.Int64) {
		if !v.IsNull() && !v.IsUnknown() {
			out[key] = v.ValueInt64()
		}
	}
	setBoolIfPresent := func(key string, v types.Bool) {
		if !v.IsNull() && !v.IsUnknown() {
			out[key] = v.ValueBool()
		}
	}

	setStringIfPresent("profile", m.Profile)
	setStringIfPresent("model", m.Model)
	setStringIfPresent("model_provider", m.ModelProvider)
	setIntIfPresent("model_context_window", m.ModelContextWindow)
	setIntIfPresent("model_max_output_tokens", m.ModelMaxOutputTokens)
	setStringIfPresent("approval_policy", m.ApprovalPolicy)
	setStringIfPresent("sandbox_mode", m.SandboxMode)
	setStringIfPresent("file_opener", m.FileOpener)
	setBoolIfPresent("hide_agent_reasoning", m.HideAgentReasoning)
	setBoolIfPresent("show_raw_agent_reasoning", m.ShowRawAgentReasoning)
	setStringIfPresent("model_reasoning_effort", m.ModelReasoningEffort)
	setStringIfPresent("model_reasoning_summary", m.ModelReasoningSummary)
	setStringIfPresent("model_verbosity", m.ModelVerbosity)
	setBoolIfPresent("model_supports_reasoning_summaries", m.ModelSupportsReasoningSummaries)
	setIntIfPresent("project_doc_max_bytes", m.ProjectDocMaxBytes)

	if m.Notify != nil && len(m.Notify) > 0 {
		var arr []string
		for _, s := range m.Notify {
			if !s.IsNull() && !s.IsUnknown() {
				arr = append(arr, s.ValueString())
			}
		}
		if len(arr) > 0 {
			out["notify"] = arr
		}
	}

	if !m.TUI.IsNull() && !m.TUI.IsUnknown() {
		var mp map[string]string
		d := m.TUI.ElementsAs(ctx, &mp, false)
		if d.HasError() {
			diags.Append(d...)
			return nil, fmt.Errorf("invalid tui map")
		}
		out["tui"] = mp
	}

	if m.SandboxWorkspaceWrite != nil {
		sw := map[string]any{}
		if m.SandboxWorkspaceWrite.WritableRoots != nil {
			var arr []string
			for _, s := range m.SandboxWorkspaceWrite.WritableRoots {
				if !s.IsNull() && !s.IsUnknown() {
					arr = append(arr, s.ValueString())
				}
			}
			if len(arr) > 0 {
				sw["writable_roots"] = arr
			}
		}
		if !m.SandboxWorkspaceWrite.NetworkAccess.IsNull() && !m.SandboxWorkspaceWrite.NetworkAccess.IsUnknown() {
			sw["network_access"] = m.SandboxWorkspaceWrite.NetworkAccess.ValueBool()
		}
		if !m.SandboxWorkspaceWrite.ExcludeTmpdirEnvVar.IsNull() && !m.SandboxWorkspaceWrite.ExcludeTmpdirEnvVar.IsUnknown() {
			sw["exclude_tmpdir_env_var"] = m.SandboxWorkspaceWrite.ExcludeTmpdirEnvVar.ValueBool()
		}
		if !m.SandboxWorkspaceWrite.ExcludeSlashTmp.IsNull() && !m.SandboxWorkspaceWrite.ExcludeSlashTmp.IsUnknown() {
			sw["exclude_slash_tmp"] = m.SandboxWorkspaceWrite.ExcludeSlashTmp.ValueBool()
		}
		if len(sw) > 0 {
			out["sandbox_workspace_write"] = sw
		}
	}

	if m.History != nil {
		h := map[string]any{}
		if !m.History.Persistence.IsNull() && !m.History.Persistence.IsUnknown() {
			h["persistence"] = m.History.Persistence.ValueString()
		}
		if !m.History.MaxBytes.IsNull() && !m.History.MaxBytes.IsUnknown() {
			h["max_bytes"] = m.History.MaxBytes.ValueInt64()
		}
		if len(h) > 0 {
			out["history"] = h
		}
	}

	if m.ShellEnvPolicy != nil {
		sp := map[string]any{}
		if !m.ShellEnvPolicy.Inherit.IsNull() && !m.ShellEnvPolicy.Inherit.IsUnknown() {
			sp["inherit"] = m.ShellEnvPolicy.Inherit.ValueString()
		}
		if !m.ShellEnvPolicy.IgnoreDefaultExcludes.IsNull() && !m.ShellEnvPolicy.IgnoreDefaultExcludes.IsUnknown() {
			sp["ignore_default_excludes"] = m.ShellEnvPolicy.IgnoreDefaultExcludes.ValueBool()
		}
		if m.ShellEnvPolicy.Exclude != nil {
			var arr []string
			for _, s := range m.ShellEnvPolicy.Exclude {
				if !s.IsNull() && !s.IsUnknown() {
					arr = append(arr, s.ValueString())
				}
			}
			if len(arr) > 0 {
				sp["exclude"] = arr
			}
		}
		if !m.ShellEnvPolicy.Set.IsNull() && !m.ShellEnvPolicy.Set.IsUnknown() {
			var sm map[string]string
			d := m.ShellEnvPolicy.Set.ElementsAs(ctx, &sm, false)
			if d.HasError() {
				diags.Append(d...)
				return nil, fmt.Errorf("invalid shell_environment_policy.set")
			}
			sp["set"] = sm
		}
		if m.ShellEnvPolicy.IncludeOnly != nil {
			var arr []string
			for _, s := range m.ShellEnvPolicy.IncludeOnly {
				if !s.IsNull() && !s.IsUnknown() {
					arr = append(arr, s.ValueString())
				}
			}
			if len(arr) > 0 {
				sp["include_only"] = arr
			}
		}
		if len(sp) > 0 {
			out["shell_environment_policy"] = sp
		}
	}

	if len(m.ModelProviders) > 0 {
		t := map[string]any{}
		for _, mp := range m.ModelProviders {
			if mp.ID.IsNull() || mp.ID.IsUnknown() || mp.ID.ValueString() == "" {
				continue
			}
			id := mp.ID.ValueString()
			row := map[string]any{}
			if !mp.Name.IsNull() && !mp.Name.IsUnknown() {
				row["name"] = mp.Name.ValueString()
			}
			if !mp.BaseURL.IsNull() && !mp.BaseURL.IsUnknown() {
				row["base_url"] = mp.BaseURL.ValueString()
			}
			if !mp.EnvKey.IsNull() && !mp.EnvKey.IsUnknown() {
				row["env_key"] = mp.EnvKey.ValueString()
			}
			if !mp.WireAPI.IsNull() && !mp.WireAPI.IsUnknown() {
				row["wire_api"] = mp.WireAPI.ValueString()
			}
			if !mp.QueryParams.IsNull() && !mp.QueryParams.IsUnknown() {
				var mm map[string]string
				d := mp.QueryParams.ElementsAs(ctx, &mm, false)
				if d.HasError() {
					diags.Append(d...)
					return nil, fmt.Errorf("invalid model_providers.query_params")
				}
				row["query_params"] = mm
			}
			if !mp.HTTPHeaders.IsNull() && !mp.HTTPHeaders.IsUnknown() {
				var mm map[string]string
				d := mp.HTTPHeaders.ElementsAs(ctx, &mm, false)
				if d.HasError() {
					diags.Append(d...)
					return nil, fmt.Errorf("invalid model_providers.http_headers")
				}
				row["http_headers"] = mm
			}
			if !mp.EnvHTTPHeaders.IsNull() && !mp.EnvHTTPHeaders.IsUnknown() {
				var mm map[string]string
				d := mp.EnvHTTPHeaders.ElementsAs(ctx, &mm, false)
				if d.HasError() {
					diags.Append(d...)
					return nil, fmt.Errorf("invalid model_providers.env_http_headers")
				}
				row["env_http_headers"] = mm
			}
			if !mp.RequestMaxRetries.IsNull() && !mp.RequestMaxRetries.IsUnknown() {
				row["request_max_retries"] = mp.RequestMaxRetries.ValueInt64()
			}
			if !mp.StreamMaxRetries.IsNull() && !mp.StreamMaxRetries.IsUnknown() {
				row["stream_max_retries"] = mp.StreamMaxRetries.ValueInt64()
			}
			if !mp.StreamIdleTimeoutMS.IsNull() && !mp.StreamIdleTimeoutMS.IsUnknown() {
				row["stream_idle_timeout_ms"] = mp.StreamIdleTimeoutMS.ValueInt64()
			}
			t[id] = row
		}
		if len(t) > 0 {
			out["model_providers"] = t
		}
	}

	if len(m.MCPServers) > 0 {
		t := map[string]any{}
		for _, s := range m.MCPServers {
			if s.ID.IsNull() || s.ID.IsUnknown() || s.ID.ValueString() == "" {
				continue
			}
			id := s.ID.ValueString()
			row := map[string]any{}
			if !s.Command.IsNull() && !s.Command.IsUnknown() {
				row["command"] = s.Command.ValueString()
			}
			if s.Args != nil {
				var arr []string
				for _, a := range s.Args {
					if !a.IsNull() && !a.IsUnknown() {
						arr = append(arr, a.ValueString())
					}
				}
				if len(arr) > 0 {
					row["args"] = arr
				}
			}
			// Only persist env when allowed
			if boolOrDefault(m.AllowSensitiveEnvWrites, false) && !s.Env.IsNull() && !s.Env.IsUnknown() {
				var mm map[string]string
				d := s.Env.ElementsAs(ctx, &mm, false)
				if d.HasError() {
					diags.Append(d...)
					return nil, fmt.Errorf("invalid mcp_servers.env")
				}
				row["env"] = mm
			}
			t[id] = row
		}
		if len(t) > 0 {
			out["mcp_servers"] = t
		}
	}

	if len(m.Profiles) > 0 {
		t := map[string]any{}
		for _, p := range m.Profiles {
			if p.Name.IsNull() || p.Name.IsUnknown() || p.Name.ValueString() == "" {
				continue
			}
			name := p.Name.ValueString()
			row := map[string]any{}
			if !p.Model.IsNull() && !p.Model.IsUnknown() {
				row["model"] = p.Model.ValueString()
			}
			if !p.ModelProvider.IsNull() && !p.ModelProvider.IsUnknown() {
				row["model_provider"] = p.ModelProvider.ValueString()
			}
			if !p.ApprovalPolicy.IsNull() && !p.ApprovalPolicy.IsUnknown() {
				row["approval_policy"] = p.ApprovalPolicy.ValueString()
			}
			if !p.SandboxMode.IsNull() && !p.SandboxMode.IsUnknown() {
				row["sandbox_mode"] = p.SandboxMode.ValueString()
			}
			t[name] = row
		}
		if len(t) > 0 {
			out["profiles"] = t
		}
	}
	return out, nil
}

// Small helpers for testability and value defaults
func strOrDefault(v types.String, def string) string {
	if !v.IsNull() && !v.IsUnknown() && v.ValueString() != "" {
		return v.ValueString()
	}
	return def
}
func boolOrDefault(v types.Bool, def bool) bool {
	if !v.IsNull() && !v.IsUnknown() {
		return v.ValueBool()
	}
	return def
}

// These wrappers make it easier to stub in tests if needed
var osRemove = func(name string) error { return os.Remove(name) }
var osStat = func(name string) (interface{}, error) { _, err := os.Stat(name); return nil, err }
