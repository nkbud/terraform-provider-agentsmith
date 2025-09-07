package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = &claudeSettingsResource{}
	_ resource.ResourceWithConfigure   = &claudeSettingsResource{}
	_ resource.ResourceWithImportState = &claudeSettingsResource{}
)

func NewClaudeSettingsResource() resource.Resource {
	return &claudeSettingsResource{}
}

type claudeSettingsResource struct {
	client *FileClient
}

type claudeSettingsResourceModel struct {
	ID                         types.String                    `tfsdk:"id"`
	Scope                      types.String                    `tfsdk:"scope"`
	APIKeyHelper               types.String                    `tfsdk:"api_key_helper"`
	CleanupPeriodDays          types.Int64                     `tfsdk:"cleanup_period_days"`
	Env                        types.Map                       `tfsdk:"env"`
	IncludeCoAuthoredBy        types.Bool                      `tfsdk:"include_co_authored_by"`
	Model                      types.String                    `tfsdk:"model"`
	OutputStyle                types.String                    `tfsdk:"output_style"`
	ForceLoginMethod           types.String                    `tfsdk:"force_login_method"`
	ForceLoginOrgUUID          types.String                    `tfsdk:"force_login_org_uuid"`
	DisableAllHooks            types.Bool                      `tfsdk:"disable_all_hooks"`
	AwsAuthRefresh             types.String                    `tfsdk:"aws_auth_refresh"`
	AwsCredentialExport        types.String                    `tfsdk:"aws_credential_export"`
	EnableAllProjectMcpServers types.Bool                      `tfsdk:"enable_all_project_mcp_servers"`
	EnabledMcpjsonServers      types.List                      `tfsdk:"enabled_mcpjson_servers"`
	DisabledMcpjsonServers     types.List                      `tfsdk:"disabled_mcpjson_servers"`
	Permissions                *claudeSettingsPermissionsModel `tfsdk:"permissions"`
	StatusLine                 *claudeSettingsStatusLineModel  `tfsdk:"status_line"`
	Hooks                      types.Map                       `tfsdk:"hooks"`
}

type claudeSettingsPermissionsModel struct {
	Allow                        types.List   `tfsdk:"allow"`
	Ask                          types.List   `tfsdk:"ask"`
	Deny                         types.List   `tfsdk:"deny"`
	AdditionalDirectories        types.List   `tfsdk:"additional_directories"`
	DefaultMode                  types.String `tfsdk:"default_mode"`
	DisableBypassPermissionsMode types.String `tfsdk:"disable_bypass_permissions_mode"`
}

type claudeSettingsStatusLineModel struct {
	Type    types.String `tfsdk:"type"`
	Command types.String `tfsdk:"command"`
}

func (r *claudeSettingsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_claude_settings"
}

func (r *claudeSettingsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a Claude Code `settings.json` file. This resource can operate at different scopes to manage user-level, project-level, or local project-level settings.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "A unique identifier for this settings resource, composed of the scope.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"scope": schema.StringAttribute{
				Description: "The scope of the settings file. Must be one of `user` (~/.claude/settings.json), `project` (<workdir>/.claude/settings.json), or `local` (<workdir>/.claude/settings.local.json).",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"api_key_helper": schema.StringAttribute{
				Description: "A custom script, executed in `/bin/sh`, to generate an auth value for model requests.",
				Optional:    true,
			},
			"cleanup_period_days": schema.Int64Attribute{
				Description: "How long to locally retain chat transcripts based on last activity date. Default: 30 days.",
				Optional:    true,
			},
			"env": schema.MapAttribute{
				Description: "A map of environment variables that will be applied to every session.",
				ElementType: types.StringType,
				Optional:    true,
			},
			"include_co_authored_by": schema.BoolAttribute{
				Description: "Whether to include the `co-authored-by Claude` byline in git commits and pull requests. Default: true.",
				Optional:    true,
			},
			"model": schema.StringAttribute{
				Description: "The name of the model to use for Claude Code (e.g., `claude-3-5-sonnet-20241022`).",
				Optional:    true,
			},
			"output_style": schema.StringAttribute{
				Description: "The output style to adjust the system prompt (e.g., `Explanatory`).",
				Optional:    true,
			},
			"force_login_method": schema.StringAttribute{
				Description: "Restricts login to a specific method. Use `claudeai` for Claude.ai accounts or `console` for Anthropic Console accounts.",
				Optional:    true,
			},
			"force_login_org_uuid": schema.StringAttribute{
				Description: "The UUID of an organization to automatically select during login, bypassing the organization selection step.",
				Optional:    true,
			},
			"disable_all_hooks": schema.BoolAttribute{
				Description: "If true, disables all configured hooks.",
				Optional:    true,
			},
			"aws_auth_refresh": schema.StringAttribute{
				Description: "A custom script that modifies the `.aws` directory for auth refresh, e.g., `aws sso login`.",
				Optional:    true,
			},
			"aws_credential_export": schema.StringAttribute{
				Description: "A custom script that outputs JSON with temporary AWS credentials.",
				Optional:    true,
			},
			"enable_all_project_mcp_servers": schema.BoolAttribute{
				Description: "If true, automatically approves all MCP servers defined in project `.mcp.json` files.",
				Optional:    true,
			},
			"enabled_mcpjson_servers": schema.ListAttribute{
				Description: "A list of specific MCP servers from `.mcp.json` files to approve.",
				ElementType: types.StringType,
				Optional:    true,
			},
			"disabled_mcpjson_servers": schema.ListAttribute{
				Description: "A list of specific MCP servers from `.mcp.json` files to reject.",
				ElementType: types.StringType,
				Optional:    true,
			},
			"hooks": schema.MapAttribute{
				Description: "A map of custom commands to run before or after tool executions. The map key is the hook name (e.g., `PreToolUse`) and the value is another map containing the hook details.",
				ElementType: types.MapType{ElemType: types.StringType},
				Optional:    true,
			},
		},
		Blocks: map[string]schema.Block{
			"permissions": schema.SingleNestedBlock{
				Description: "A block for configuring tool usage permissions.",
				Attributes: map[string]schema.Attribute{
					"allow": schema.ListAttribute{
						Description: "A list of permission rules to automatically allow tool use without prompting.",
						ElementType: types.StringType,
						Optional:    true,
					},
					"ask": schema.ListAttribute{
						Description: "A list of permission rules that will cause Claude to ask for confirmation before using a tool.",
						ElementType: types.StringType,
						Optional:    true,
					},
					"deny": schema.ListAttribute{
						Description: "A list of permission rules to deny tool use. Also used to exclude sensitive files from being read.",
						ElementType: types.StringType,
						Optional:    true,
					},
					"additional_directories": schema.ListAttribute{
						Description: "A list of additional working directories that Claude has access to.",
						ElementType: types.StringType,
						Optional:    true,
					},
					"default_mode": schema.StringAttribute{
						Description: "The default permission mode when opening Claude Code (e.g., `acceptEdits`).",
						Optional:    true,
					},
					"disable_bypass_permissions_mode": schema.StringAttribute{
						Description: "Set to `disable` to prevent `bypassPermissions` mode from being activated.",
						Optional:    true,
					},
				},
			},
			"status_line": schema.SingleNestedBlock{
				Description: "Configuration for a custom status line to display context.",
				Attributes: map[string]schema.Attribute{
					"type": schema.StringAttribute{
						Description: "The type of status line. For example, `command`.",
						Optional:    true,
					},
					"command": schema.StringAttribute{
						Description: "The command to execute to generate the status line content.",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (r *claudeSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan claudeSettingsResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	filePath := r.getSettingsFilePath(plan.Scope.ValueString())
	if filePath == "" {
		resp.Diagnostics.AddError("Invalid scope", "Scope must be 'user', 'project', or 'local'")
		return
	}

	// Ensure directory exists
	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		resp.Diagnostics.AddError("Failed to create directory", err.Error())
		return
	}

	// Create settings JSON
	settingsData, err := r.modelToSettings(ctx, &resp.Diagnostics, &plan)
	if err != nil || resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError("Failed to create settings data", err.Error())
		return
	}

	jsonData, err := json.MarshalIndent(settingsData, "", "  ")
	if err != nil {
		resp.Diagnostics.AddError("Failed to marshal JSON", err.Error())
		return
	}

	if err := os.WriteFile(filePath, jsonData, 0644); err != nil {
		resp.Diagnostics.AddError("Failed to write settings file", err.Error())
		return
	}

	// Set computed attributes
	plan.ID = types.StringValue(fmt.Sprintf("claude-settings-%s", plan.Scope.ValueString()))

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *claudeSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state claudeSettingsResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	filePath := r.getSettingsFilePath(state.Scope.ValueString())
	if filePath == "" {
		resp.Diagnostics.AddError("Invalid scope", "Scope must be 'user', 'project', or 'local'")
		return
	}

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		resp.State.RemoveResource(ctx)
		return
	}

	// Read and parse settings file
	data, err := os.ReadFile(filePath)
	if err != nil {
		resp.Diagnostics.AddError("Failed to read settings file", err.Error())
		return
	}

	var settingsData map[string]interface{}
	if err := json.Unmarshal(data, &settingsData); err != nil {
		resp.Diagnostics.AddError("Failed to parse settings JSON", err.Error())
		return
	}

	// Update state with current file contents
	r.settingsToModel(ctx, &resp.Diagnostics, &state, settingsData)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *claudeSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan claudeSettingsResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	filePath := r.getSettingsFilePath(plan.Scope.ValueString())
	if filePath == "" {
		resp.Diagnostics.AddError("Invalid scope", "Scope must be 'user', 'project', or 'local'")
		return
	}

	// Create settings JSON
	settingsData, err := r.modelToSettings(ctx, &resp.Diagnostics, &plan)
	if err != nil || resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError("Failed to create settings data", err.Error())
		return
	}

	jsonData, err := json.MarshalIndent(settingsData, "", "  ")
	if err != nil {
		resp.Diagnostics.AddError("Failed to marshal JSON", err.Error())
		return
	}

	if err := os.WriteFile(filePath, jsonData, 0644); err != nil {
		resp.Diagnostics.AddError("Failed to write settings file", err.Error())
		return
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *claudeSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state claudeSettingsResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	filePath := r.getSettingsFilePath(state.Scope.ValueString())
	if filePath == "" {
		resp.Diagnostics.AddError("Invalid scope", "Scope must be 'user', 'project', or 'local'")
		return
	}

	if err := os.Remove(filePath); err != nil && !os.IsNotExist(err) {
		resp.Diagnostics.AddError("Failed to delete settings file", err.Error())
		return
	}
}

func (r *claudeSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: scope (user/project/local)
	scope := strings.TrimSpace(req.ID)
	if scope != "user" && scope != "project" && scope != "local" {
		resp.Diagnostics.AddError("Invalid import ID", "Import ID must be 'user', 'project', or 'local'")
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("scope"), scope)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), fmt.Sprintf("claude-settings-%s", scope))...)
}

func (r *claudeSettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*FileClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *FileClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = client
}

func (r *claudeSettingsResource) getSettingsFilePath(scope string) string {
	switch scope {
	case "user":
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return ""
		}
		return filepath.Join(homeDir, ".claude", "settings.json")
	case "project":
		if r.client == nil {
			return ""
		}
		return filepath.Join(r.client.workDir, ".claude", "settings.json")
	case "local":
		if r.client == nil {
			return ""
		}
		return filepath.Join(r.client.workDir, ".claude", "settings.local.json")
	default:
		return ""
	}
}

func (r *claudeSettingsResource) modelToSettings(ctx context.Context, diags *diag.Diagnostics, model *claudeSettingsResourceModel) (map[string]interface{}, error) {
	settings := make(map[string]interface{})

	if !model.APIKeyHelper.IsNull() {
		settings["apiKeyHelper"] = model.APIKeyHelper.ValueString()
	}
	if !model.CleanupPeriodDays.IsNull() {
		settings["cleanupPeriodDays"] = model.CleanupPeriodDays.ValueInt64()
	}
	if !model.IncludeCoAuthoredBy.IsNull() {
		settings["includeCoAuthoredBy"] = model.IncludeCoAuthoredBy.ValueBool()
	}
	if !model.Model.IsNull() {
		settings["model"] = model.Model.ValueString()
	}
	if !model.OutputStyle.IsNull() {
		settings["outputStyle"] = model.OutputStyle.ValueString()
	}
	if !model.ForceLoginMethod.IsNull() {
		settings["forceLoginMethod"] = model.ForceLoginMethod.ValueString()
	}
	if !model.ForceLoginOrgUUID.IsNull() {
		settings["forceLoginOrgUUID"] = model.ForceLoginOrgUUID.ValueString()
	}
	if !model.DisableAllHooks.IsNull() {
		settings["disableAllHooks"] = model.DisableAllHooks.ValueBool()
	}
	if !model.AwsAuthRefresh.IsNull() {
		settings["awsAuthRefresh"] = model.AwsAuthRefresh.ValueString()
	}
	if !model.AwsCredentialExport.IsNull() {
		settings["awsCredentialExport"] = model.AwsCredentialExport.ValueString()
	}
	if !model.EnableAllProjectMcpServers.IsNull() {
		settings["enableAllProjectMcpServers"] = model.EnableAllProjectMcpServers.ValueBool()
	}

	// Handle lists
	if !model.EnabledMcpjsonServers.IsNull() {
		var servers []string
		d := model.EnabledMcpjsonServers.ElementsAs(ctx, &servers, false)
		if d.HasError() {
			diags.Append(d...)
			return nil, fmt.Errorf("failed to convert enabled MCP servers")
		}
		settings["enabledMcpjsonServers"] = servers
	}

	if !model.DisabledMcpjsonServers.IsNull() {
		var servers []string
		d := model.DisabledMcpjsonServers.ElementsAs(ctx, &servers, false)
		if d.HasError() {
			diags.Append(d...)
			return nil, fmt.Errorf("failed to convert disabled MCP servers")
		}
		settings["disabledMcpjsonServers"] = servers
	}

	// Handle maps
	if !model.Env.IsNull() {
		var envMap map[string]string
		d := model.Env.ElementsAs(ctx, &envMap, false)
		if d.HasError() {
			diags.Append(d...)
			return nil, fmt.Errorf("failed to convert env map")
		}
		settings["env"] = envMap
	}

	if !model.Hooks.IsNull() {
		var hooksMap map[string]map[string]string
		d := model.Hooks.ElementsAs(ctx, &hooksMap, false)
		if d.HasError() {
			diags.Append(d...)
			return nil, fmt.Errorf("failed to convert hooks map")
		}
		settings["hooks"] = hooksMap
	}

	// Handle permissions
	if model.Permissions != nil {
		perms := make(map[string]interface{})

		if !model.Permissions.DefaultMode.IsNull() {
			perms["defaultMode"] = model.Permissions.DefaultMode.ValueString()
		}
		if !model.Permissions.DisableBypassPermissionsMode.IsNull() {
			perms["disableBypassPermissionsMode"] = model.Permissions.DisableBypassPermissionsMode.ValueString()
		}

		if !model.Permissions.Allow.IsNull() {
			var allow []string
			d := model.Permissions.Allow.ElementsAs(ctx, &allow, false)
			if d.HasError() {
				diags.Append(d...)
				return nil, fmt.Errorf("failed to convert allow list")
			}
			perms["allow"] = allow
		}

		if !model.Permissions.Ask.IsNull() {
			var ask []string
			d := model.Permissions.Ask.ElementsAs(ctx, &ask, false)
			if d.HasError() {
				diags.Append(d...)
				return nil, fmt.Errorf("failed to convert ask list")
			}
			perms["ask"] = ask
		}

		if !model.Permissions.Deny.IsNull() {
			var deny []string
			d := model.Permissions.Deny.ElementsAs(ctx, &deny, false)
			if d.HasError() {
				diags.Append(d...)
				return nil, fmt.Errorf("failed to convert deny list")
			}
			perms["deny"] = deny
		}

		if !model.Permissions.AdditionalDirectories.IsNull() {
			var dirs []string
			d := model.Permissions.AdditionalDirectories.ElementsAs(ctx, &dirs, false)
			if d.HasError() {
				diags.Append(d...)
				return nil, fmt.Errorf("failed to convert additional directories")
			}
			perms["additionalDirectories"] = dirs
		}

		if len(perms) > 0 {
			settings["permissions"] = perms
		}
	}

	// Handle status line
	if model.StatusLine != nil {
		statusLine := make(map[string]interface{})
		if !model.StatusLine.Type.IsNull() {
			statusLine["type"] = model.StatusLine.Type.ValueString()
		}
		if !model.StatusLine.Command.IsNull() {
			statusLine["command"] = model.StatusLine.Command.ValueString()
		}
		if len(statusLine) > 0 {
			settings["statusLine"] = statusLine
		}
	}

	return settings, nil
}

func (r *claudeSettingsResource) settingsToModel(ctx context.Context, diags *diag.Diagnostics, model *claudeSettingsResourceModel, settings map[string]interface{}) {
	// Implementation would populate model from settings map
	// This is a simplified version - full implementation would handle all fields

	if val, ok := settings["apiKeyHelper"].(string); ok {
		model.APIKeyHelper = types.StringValue(val)
	} else {
		model.APIKeyHelper = types.StringNull()
	}

	if val, ok := settings["model"].(string); ok {
		model.Model = types.StringValue(val)
	} else {
		model.Model = types.StringNull()
	}

	// Add implementation for remaining fields...
}
