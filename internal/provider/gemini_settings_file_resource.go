package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"os"
	"path/filepath"
)

var (
	_ resource.Resource              = &geminiSettingsFileResource{}
	_ resource.ResourceWithConfigure = &geminiSettingsFileResource{}
)

func NewGeminiSettingsFileResource() resource.Resource {
	return &geminiSettingsFileResource{}
}

type geminiSettingsFileResource struct {
	client *FileClient
}

type geminiSettingsFileResourceModel struct {
	Scope      types.String         `tfsdk:"scope"`
	ProjectDir types.String         `tfsdk:"project_dir"`
	Settings   *geminiSettingsModel `tfsdk:"settings"`
	ID         types.String         `tfsdk:"id"`
	Path       types.String         `tfsdk:"path"`
}

func (r *geminiSettingsFileResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gemini_settings_file"
}

func (r *geminiSettingsFileResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a Gemini CLI settings.json file.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The absolute path to the managed settings.json file, used as the resource ID.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"path": schema.StringAttribute{
				Description: "The absolute path to the managed settings.json file.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"scope": schema.StringAttribute{
				Description: "Defines the target configuration scope. Must be one of `project`, `user`, or `system_override`.",
				Required:    true,
			},
			"project_dir": schema.StringAttribute{
				Description: "The absolute path to the project's root directory. Required only when scope is `project`.",
				Optional:    true,
			},
			"settings": schema.SingleNestedAttribute{
				Description: "A block containing all the Gemini settings.",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"general": schema.SingleNestedAttribute{
						Description: "General settings.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"preferred_editor":      schema.StringAttribute{Optional: true},
							"vim_mode":              schema.BoolAttribute{Optional: true},
							"disable_auto_update":   schema.BoolAttribute{Optional: true},
							"disable_update_nag":    schema.BoolAttribute{Optional: true},
							"checkpointing_enabled": schema.BoolAttribute{Optional: true},
						},
					},
					"ui": schema.SingleNestedAttribute{
						Description: "User interface settings.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"theme":                                 schema.StringAttribute{Optional: true},
							"custom_themes":                         schema.MapAttribute{ElementType: types.StringType, Optional: true},
							"hide_window_title":                     schema.BoolAttribute{Optional: true},
							"hide_tips":                             schema.BoolAttribute{Optional: true},
							"hide_banner":                           schema.BoolAttribute{Optional: true},
							"hide_footer":                           schema.BoolAttribute{Optional: true},
							"show_memory_usage":                     schema.BoolAttribute{Optional: true},
							"show_line_numbers":                     schema.BoolAttribute{Optional: true},
							"show_citations":                        schema.BoolAttribute{Optional: true},
							"accessibility_disable_loading_phrases": schema.BoolAttribute{Optional: true},
						},
					},
					"ide": schema.SingleNestedAttribute{
						Description: "IDE integration settings.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"enabled":        schema.BoolAttribute{Optional: true},
							"has_seen_nudge": schema.BoolAttribute{Optional: true},
						},
					},
					"privacy": schema.SingleNestedAttribute{
						Description: "Privacy settings.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"usage_statistics_enabled": schema.BoolAttribute{Optional: true},
						},
					},
					"model": schema.SingleNestedAttribute{
						Description: "Model configuration settings.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"name":                  schema.StringAttribute{Optional: true},
							"max_session_turns":     schema.Int64Attribute{Optional: true},
							"summarize_tool_output": schema.MapAttribute{ElementType: types.StringType, Optional: true},
							"chat_compression_context_percentage_threshold": schema.StringAttribute{Optional: true},
							"skip_next_speaker_check":                       schema.BoolAttribute{Optional: true},
						},
					},
					"context": schema.SingleNestedAttribute{
						Description: "Context file and memory settings.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"file_name":                                   schema.ListAttribute{ElementType: types.StringType, Optional: true},
							"import_format":                               schema.StringAttribute{Optional: true},
							"discovery_max_dirs":                          schema.Int64Attribute{Optional: true},
							"include_directories":                         schema.ListAttribute{ElementType: types.StringType, Optional: true},
							"load_from_include_directories":               schema.BoolAttribute{Optional: true},
							"file_filtering_respect_git_ignore":           schema.BoolAttribute{Optional: true},
							"file_filtering_respect_gemini_ignore":        schema.BoolAttribute{Optional: true},
							"file_filtering_enable_recursive_file_search": schema.BoolAttribute{Optional: true},
						},
					},
					"tools": schema.SingleNestedAttribute{
						Description: "Tool configuration settings.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"sandbox":           schema.StringAttribute{Optional: true},
							"use_pty":           schema.BoolAttribute{Optional: true},
							"core":              schema.ListAttribute{ElementType: types.StringType, Optional: true},
							"exclude":           schema.ListAttribute{ElementType: types.StringType, Optional: true},
							"allowed":           schema.ListAttribute{ElementType: types.StringType, Optional: true},
							"discovery_command": schema.StringAttribute{Optional: true},
							"call_command":      schema.StringAttribute{Optional: true},
						},
					},
					"mcp": schema.SingleNestedAttribute{
						Description: "Model Context Protocol settings.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"server_command": schema.StringAttribute{Optional: true},
							"allowed":        schema.ListAttribute{ElementType: types.StringType, Optional: true},
							"excluded":       schema.ListAttribute{ElementType: types.StringType, Optional: true},
						},
					},
					"security": schema.SingleNestedAttribute{
						Description: "Security settings.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"folder_trust_enabled": schema.BoolAttribute{Optional: true},
							"auth_selected_type":   schema.StringAttribute{Optional: true},
							"auth_enforced_type":   schema.StringAttribute{Optional: true},
							"auth_use_external":    schema.BoolAttribute{Optional: true},
						},
					},
					"advanced": schema.SingleNestedAttribute{
						Description: "Advanced configuration settings.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"auto_configure_memory": schema.BoolAttribute{Optional: true},
							"dns_resolution_order":  schema.StringAttribute{Optional: true},
							"excluded_env_vars":     schema.ListAttribute{ElementType: types.StringType, Optional: true},
							"bug_command":           schema.MapAttribute{ElementType: types.StringType, Optional: true},
						},
					},
					"mcp_servers": schema.MapAttribute{
						Description: "Individual MCP server configurations.",
						ElementType: types.StringType,
						Optional:    true,
					},
					"telemetry": schema.SingleNestedAttribute{
						Description: "Logging and metrics configuration.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"enabled":       schema.BoolAttribute{Optional: true},
							"target":        schema.StringAttribute{Optional: true},
							"otlp_endpoint": schema.StringAttribute{Optional: true},
							"otlp_protocol": schema.StringAttribute{Optional: true},
							"log_prompts":   schema.BoolAttribute{Optional: true},
							"outfile":       schema.StringAttribute{Optional: true},
						},
					},
				},
			},
		},
	}
}

func (r *geminiSettingsFileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan geminiSettingsFileResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	path, err := r.getSettingsPath(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error determining settings path", err.Error())
		return
	}

	plan.ID = types.StringValue(path)
	plan.Path = types.StringValue(path)

	if err := r.writeSettingsFile(ctx, path, plan.Settings); err != nil {
		resp.Diagnostics.AddError("Error writing settings file", err.Error())
		return
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *geminiSettingsFileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state geminiSettingsFileResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	path := state.ID.ValueString()
	if _, err := os.Stat(path); os.IsNotExist(err) {
		resp.State.RemoveResource(ctx)
		return
	}

	bytes, err := os.ReadFile(path)
	if err != nil {
		resp.Diagnostics.AddError("Error reading settings file", fmt.Sprintf("Could not read %s: %s", path, err.Error()))
		return
	}

	var data map[string]interface{}
	if len(bytes) > 0 {
		if err := json.Unmarshal(bytes, &data); err != nil {
			resp.Diagnostics.AddError("Error parsing settings file", fmt.Sprintf("Could not parse JSON from %s: %s", path, err.Error()))
			return
		}
	}

	// Populate state from file content
	state.Settings = &geminiSettingsModel{}
	populateGeminiSettingsFromMap(ctx, &resp.Diagnostics, state.Settings, data)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *geminiSettingsFileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan geminiSettingsFileResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	path, err := r.getSettingsPath(plan)
	if err != nil {
		resp.Diagnostics.AddError("Error determining settings path", err.Error())
		return
	}

	plan.ID = types.StringValue(path)
	plan.Path = types.StringValue(path)

	if err := r.writeSettingsFile(ctx, path, plan.Settings); err != nil {
		resp.Diagnostics.AddError("Error writing settings file", err.Error())
		return
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *geminiSettingsFileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state geminiSettingsFileResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	path := state.ID.ValueString()
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// File already gone
		return
	}

	err := os.Remove(path)
	if err != nil {
		resp.Diagnostics.AddError("Error deleting settings file", err.Error())
		return
	}
}

func (r *geminiSettingsFileResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.client = client
}

func (r *geminiSettingsFileResource) getSettingsPath(plan geminiSettingsFileResourceModel) (string, error) {
	scope := plan.Scope.ValueString()
	switch scope {
	case "project":
		projectDir := plan.ProjectDir.ValueString()
		if projectDir == "" {
			return "", fmt.Errorf("'project_dir' must be set when scope is 'project'")
		}
		return filepath.Join(projectDir, ".gemini", "settings.json"), nil
	case "user":
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("failed to get user home directory: %w", err)
		}
		return filepath.Join(homeDir, ".gemini", "settings.json"), nil
	// system_override is omitted for safety, as it requires elevated permissions.
	default:
		return "", fmt.Errorf("invalid scope: %s. Must be one of 'project' or 'user'", scope)
	}
}

func (r *geminiSettingsFileResource) writeSettingsFile(ctx context.Context, path string, settings *geminiSettingsModel) error {
	settingsMap, err := r.settingsModelToMap(ctx, settings)
	if err != nil {
		return err
	}

	if len(settingsMap) == 0 {
		// If the file exists, remove it. Otherwise, do nothing.
		if _, err := os.Stat(path); err == nil {
			return os.Remove(path)
		}
		return nil
	}

	bytes, err := json.MarshalIndent(settingsMap, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal settings to JSON: %w", err)
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}

	return os.WriteFile(path, bytes, 0644)
}

// settingsModelToMap converts the Terraform model to a map for JSON marshaling.
func (r *geminiSettingsFileResource) settingsModelToMap(ctx context.Context, settings *geminiSettingsModel) (map[string]interface{}, error) {
	output := make(map[string]interface{})
	if settings == nil {
		return output, nil
	}

	if settings.General != nil {
		general := make(map[string]interface{})
		if !settings.General.PreferredEditor.IsNull() && !settings.General.PreferredEditor.IsUnknown() {
			general["preferredEditor"] = settings.General.PreferredEditor.ValueString()
		}
		if !settings.General.VimMode.IsNull() && !settings.General.VimMode.IsUnknown() {
			general["vimMode"] = settings.General.VimMode.ValueBool()
		}
		if !settings.General.DisableAutoUpdate.IsNull() && !settings.General.DisableAutoUpdate.IsUnknown() {
			general["disableAutoUpdate"] = settings.General.DisableAutoUpdate.ValueBool()
		}
		if !settings.General.DisableUpdateNag.IsNull() && !settings.General.DisableUpdateNag.IsUnknown() {
			general["disableUpdateNag"] = settings.General.DisableUpdateNag.ValueBool()
		}
		if !settings.General.CheckpointingEnabled.IsNull() && !settings.General.CheckpointingEnabled.IsUnknown() {
			general["checkpointing"] = map[string]interface{}{"enabled": settings.General.CheckpointingEnabled.ValueBool()}
		}
		if len(general) > 0 {
			output["general"] = general
		}
	}

	if settings.UI != nil {
		ui := make(map[string]interface{})
		if !settings.UI.Theme.IsNull() && !settings.UI.Theme.IsUnknown() {
			ui["theme"] = settings.UI.Theme.ValueString()
		}
		if !settings.UI.HideWindowTitle.IsNull() && !settings.UI.HideWindowTitle.IsUnknown() {
			ui["hideWindowTitle"] = settings.UI.HideWindowTitle.ValueBool()
		}
		if !settings.UI.HideTips.IsNull() && !settings.UI.HideTips.IsUnknown() {
			ui["hideTips"] = settings.UI.HideTips.ValueBool()
		}
		if !settings.UI.HideBanner.IsNull() && !settings.UI.HideBanner.IsUnknown() {
			ui["hideBanner"] = settings.UI.HideBanner.ValueBool()
		}
		if !settings.UI.HideFooter.IsNull() && !settings.UI.HideFooter.IsUnknown() {
			ui["hideFooter"] = settings.UI.HideFooter.ValueBool()
		}
		if !settings.UI.ShowMemoryUsage.IsNull() && !settings.UI.ShowMemoryUsage.IsUnknown() {
			ui["showMemoryUsage"] = settings.UI.ShowMemoryUsage.ValueBool()
		}
		if !settings.UI.ShowLineNumbers.IsNull() && !settings.UI.ShowLineNumbers.IsUnknown() {
			ui["showLineNumbers"] = settings.UI.ShowLineNumbers.ValueBool()
		}
		if !settings.UI.ShowCitations.IsNull() && !settings.UI.ShowCitations.IsUnknown() {
			ui["showCitations"] = settings.UI.ShowCitations.ValueBool()
		}
		if !settings.UI.AccessibilityDisableLoadingPhrases.IsNull() && !settings.UI.AccessibilityDisableLoadingPhrases.IsUnknown() {
			ui["accessibility"] = map[string]interface{}{"disableLoadingPhrases": settings.UI.AccessibilityDisableLoadingPhrases.ValueBool()}
		}
		if !settings.UI.CustomThemes.IsNull() && !settings.UI.CustomThemes.IsUnknown() {
			var themes map[string]string
			diags := settings.UI.CustomThemes.ElementsAs(ctx, &themes, false)
			if diags.HasError() {
				return nil, fmt.Errorf("failed to read ui.custom_themes")
			}
			ui["customThemes"] = themes
		}
		if len(ui) > 0 {
			output["ui"] = ui
		}
	}

	if settings.IDE != nil {
		ide := make(map[string]interface{})
		if !settings.IDE.Enabled.IsNull() && !settings.IDE.Enabled.IsUnknown() {
			ide["enabled"] = settings.IDE.Enabled.ValueBool()
		}
		if !settings.IDE.HasSeenNudge.IsNull() && !settings.IDE.HasSeenNudge.IsUnknown() {
			ide["hasSeenNudge"] = settings.IDE.HasSeenNudge.ValueBool()
		}
		if len(ide) > 0 {
			output["ide"] = ide
		}
	}

	if settings.Privacy != nil {
		privacy := make(map[string]interface{})
		if !settings.Privacy.UsageStatisticsEnabled.IsNull() && !settings.Privacy.UsageStatisticsEnabled.IsUnknown() {
			privacy["usageStatisticsEnabled"] = settings.Privacy.UsageStatisticsEnabled.ValueBool()
		}
		if len(privacy) > 0 {
			output["privacy"] = privacy
		}
	}

	if settings.Model != nil {
		model := make(map[string]interface{})
		if !settings.Model.Name.IsNull() && !settings.Model.Name.IsUnknown() {
			model["name"] = settings.Model.Name.ValueString()
		}
		if !settings.Model.MaxSessionTurns.IsNull() && !settings.Model.MaxSessionTurns.IsUnknown() {
			model["maxSessionTurns"] = settings.Model.MaxSessionTurns.ValueInt64()
		}
		if !settings.Model.SkipNextSpeakerCheck.IsNull() && !settings.Model.SkipNextSpeakerCheck.IsUnknown() {
			model["skipNextSpeakerCheck"] = settings.Model.SkipNextSpeakerCheck.ValueBool()
		}
		if !settings.Model.SummarizeToolOutput.IsNull() && !settings.Model.SummarizeToolOutput.IsUnknown() {
			var summary map[string]string
			diags := settings.Model.SummarizeToolOutput.ElementsAs(ctx, &summary, false)
			if diags.HasError() {
				return nil, fmt.Errorf("failed to read model.summarize_tool_output")
			}
			model["summarizeToolOutput"] = summary
		}
		if !settings.Model.ChatCompressionContextPercentageThreshold.IsNull() && !settings.Model.ChatCompressionContextPercentageThreshold.IsUnknown() {
			model["chatCompression"] = map[string]interface{}{"contextPercentageThreshold": settings.Model.ChatCompressionContextPercentageThreshold.ValueString()}
		}
		if len(model) > 0 {
			output["model"] = model
		}
	}

	if settings.Context != nil {
		contextSettings := make(map[string]interface{})
		if !settings.Context.ImportFormat.IsNull() && !settings.Context.ImportFormat.IsUnknown() {
			contextSettings["importFormat"] = settings.Context.ImportFormat.ValueString()
		}
		if !settings.Context.DiscoveryMaxDirs.IsNull() && !settings.Context.DiscoveryMaxDirs.IsUnknown() {
			contextSettings["discoveryMaxDirs"] = settings.Context.DiscoveryMaxDirs.ValueInt64()
		}
		if !settings.Context.LoadFromIncludeDirectories.IsNull() && !settings.Context.LoadFromIncludeDirectories.IsUnknown() {
			contextSettings["loadFromIncludeDirectories"] = settings.Context.LoadFromIncludeDirectories.ValueBool()
		}
		if !settings.Context.FileName.IsNull() && !settings.Context.FileName.IsUnknown() {
			var fileNames []string
			diags := settings.Context.FileName.ElementsAs(ctx, &fileNames, false)
			if diags.HasError() {
				return nil, fmt.Errorf("failed to read context.file_name")
			}
			contextSettings["fileName"] = fileNames
		}
		if !settings.Context.IncludeDirectories.IsNull() && !settings.Context.IncludeDirectories.IsUnknown() {
			var dirs []string
			diags := settings.Context.IncludeDirectories.ElementsAs(ctx, &dirs, false)
			if diags.HasError() {
				return nil, fmt.Errorf("failed to read context.include_directories")
			}
			contextSettings["includeDirectories"] = dirs
		}
		fileFiltering := make(map[string]interface{})
		if !settings.Context.FileFilteringRespectGitIgnore.IsNull() && !settings.Context.FileFilteringRespectGitIgnore.IsUnknown() {
			fileFiltering["respectGitIgnore"] = settings.Context.FileFilteringRespectGitIgnore.ValueBool()
		}
		if !settings.Context.FileFilteringRespectGeminiIgnore.IsNull() && !settings.Context.FileFilteringRespectGeminiIgnore.IsUnknown() {
			fileFiltering["respectGeminiIgnore"] = settings.Context.FileFilteringRespectGeminiIgnore.ValueBool()
		}
		if !settings.Context.FileFilteringEnableRecursiveFileSearch.IsNull() && !settings.Context.FileFilteringEnableRecursiveFileSearch.IsUnknown() {
			fileFiltering["enableRecursiveFileSearch"] = settings.Context.FileFilteringEnableRecursiveFileSearch.ValueBool()
		}
		if len(fileFiltering) > 0 {
			contextSettings["fileFiltering"] = fileFiltering
		}
		if len(contextSettings) > 0 {
			output["context"] = contextSettings
		}
	}

	if settings.Tools != nil {
		tools := make(map[string]interface{})
		if !settings.Tools.Sandbox.IsNull() && !settings.Tools.Sandbox.IsUnknown() {
			tools["sandbox"] = settings.Tools.Sandbox.ValueString()
		}
		if !settings.Tools.UsePty.IsNull() && !settings.Tools.UsePty.IsUnknown() {
			tools["usePty"] = settings.Tools.UsePty.ValueBool()
		}
		if !settings.Tools.DiscoveryCommand.IsNull() && !settings.Tools.DiscoveryCommand.IsUnknown() {
			tools["discoveryCommand"] = settings.Tools.DiscoveryCommand.ValueString()
		}
		if !settings.Tools.CallCommand.IsNull() && !settings.Tools.CallCommand.IsUnknown() {
			tools["callCommand"] = settings.Tools.CallCommand.ValueString()
		}
		if !settings.Tools.Core.IsNull() && !settings.Tools.Core.IsUnknown() {
			var core []string
			diags := settings.Tools.Core.ElementsAs(ctx, &core, false)
			if diags.HasError() {
				return nil, fmt.Errorf("failed to read tools.core")
			}
			tools["core"] = core
		}
		if !settings.Tools.Exclude.IsNull() && !settings.Tools.Exclude.IsUnknown() {
			var exclude []string
			diags := settings.Tools.Exclude.ElementsAs(ctx, &exclude, false)
			if diags.HasError() {
				return nil, fmt.Errorf("failed to read tools.exclude")
			}
			tools["exclude"] = exclude
		}
		if !settings.Tools.Allowed.IsNull() && !settings.Tools.Allowed.IsUnknown() {
			var allowed []string
			diags := settings.Tools.Allowed.ElementsAs(ctx, &allowed, false)
			if diags.HasError() {
				return nil, fmt.Errorf("failed to read tools.allowed")
			}
			tools["allowed"] = allowed
		}
		if len(tools) > 0 {
			output["tools"] = tools
		}
	}

	if settings.MCP != nil {
		mcp := make(map[string]interface{})
		if !settings.MCP.ServerCommand.IsNull() && !settings.MCP.ServerCommand.IsUnknown() {
			mcp["serverCommand"] = settings.MCP.ServerCommand.ValueString()
		}
		if !settings.MCP.Allowed.IsNull() && !settings.MCP.Allowed.IsUnknown() {
			var allowed []string
			diags := settings.MCP.Allowed.ElementsAs(ctx, &allowed, false)
			if diags.HasError() {
				return nil, fmt.Errorf("failed to read mcp.allowed")
			}
			mcp["allowed"] = allowed
		}
		if !settings.MCP.Excluded.IsNull() && !settings.MCP.Excluded.IsUnknown() {
			var excluded []string
			diags := settings.MCP.Excluded.ElementsAs(ctx, &excluded, false)
			if diags.HasError() {
				return nil, fmt.Errorf("failed to read mcp.excluded")
			}
			mcp["excluded"] = excluded
		}
		if len(mcp) > 0 {
			output["mcp"] = mcp
		}
	}

	if settings.Security != nil {
		security := make(map[string]interface{})
		if !settings.Security.FolderTrustEnabled.IsNull() && !settings.Security.FolderTrustEnabled.IsUnknown() {
			security["folderTrust"] = map[string]interface{}{"enabled": settings.Security.FolderTrustEnabled.ValueBool()}
		}
		auth := make(map[string]interface{})
		if !settings.Security.AuthSelectedType.IsNull() && !settings.Security.AuthSelectedType.IsUnknown() {
			auth["selectedType"] = settings.Security.AuthSelectedType.ValueString()
		}
		if !settings.Security.AuthEnforcedType.IsNull() && !settings.Security.AuthEnforcedType.IsUnknown() {
			auth["enforcedType"] = settings.Security.AuthEnforcedType.ValueString()
		}
		if !settings.Security.AuthUseExternal.IsNull() && !settings.Security.AuthUseExternal.IsUnknown() {
			auth["useExternal"] = settings.Security.AuthUseExternal.ValueBool()
		}
		if len(auth) > 0 {
			security["auth"] = auth
		}
		if len(security) > 0 {
			output["security"] = security
		}
	}

	if settings.Advanced != nil {
		advanced := make(map[string]interface{})
		if !settings.Advanced.AutoConfigureMemory.IsNull() && !settings.Advanced.AutoConfigureMemory.IsUnknown() {
			advanced["autoConfigureMemory"] = settings.Advanced.AutoConfigureMemory.ValueBool()
		}
		if !settings.Advanced.DNSResolutionOrder.IsNull() && !settings.Advanced.DNSResolutionOrder.IsUnknown() {
			advanced["dnsResolutionOrder"] = settings.Advanced.DNSResolutionOrder.ValueString()
		}
		if !settings.Advanced.ExcludedEnvVars.IsNull() && !settings.Advanced.ExcludedEnvVars.IsUnknown() {
			var vars []string
			diags := settings.Advanced.ExcludedEnvVars.ElementsAs(ctx, &vars, false)
			if diags.HasError() {
				return nil, fmt.Errorf("failed to read advanced.excluded_env_vars")
			}
			advanced["excludedEnvVars"] = vars
		}
		if !settings.Advanced.BugCommand.IsNull() && !settings.Advanced.BugCommand.IsUnknown() {
			var cmd map[string]string
			diags := settings.Advanced.BugCommand.ElementsAs(ctx, &cmd, false)
			if diags.HasError() {
				return nil, fmt.Errorf("failed to read advanced.bug_command")
			}
			advanced["bugCommand"] = cmd
		}
		if len(advanced) > 0 {
			output["advanced"] = advanced
		}
	}

	if !settings.MCPServers.IsNull() && !settings.MCPServers.IsUnknown() {
		var servers map[string]string
		diags := settings.MCPServers.ElementsAs(ctx, &servers, false)
		if diags.HasError() {
			return nil, fmt.Errorf("failed to read mcp_servers")
		}
		output["mcpServers"] = servers
	}

	if settings.Telemetry != nil {
		telemetry := make(map[string]interface{})
		if !settings.Telemetry.Enabled.IsNull() && !settings.Telemetry.Enabled.IsUnknown() {
			telemetry["enabled"] = settings.Telemetry.Enabled.ValueBool()
		}
		if !settings.Telemetry.Target.IsNull() && !settings.Telemetry.Target.IsUnknown() {
			telemetry["target"] = settings.Telemetry.Target.ValueString()
		}
		if !settings.Telemetry.OTLPEndpoint.IsNull() && !settings.Telemetry.OTLPEndpoint.IsUnknown() {
			telemetry["otlpEndpoint"] = settings.Telemetry.OTLPEndpoint.ValueString()
		}
		if !settings.Telemetry.OTLPProtocol.IsNull() && !settings.Telemetry.OTLPProtocol.IsUnknown() {
			telemetry["otlpProtocol"] = settings.Telemetry.OTLPProtocol.ValueString()
		}
		if !settings.Telemetry.LogPrompts.IsNull() && !settings.Telemetry.LogPrompts.IsUnknown() {
			telemetry["logPrompts"] = settings.Telemetry.LogPrompts.ValueBool()
		}
		if !settings.Telemetry.Outfile.IsNull() && !settings.Telemetry.Outfile.IsUnknown() {
			telemetry["outfile"] = settings.Telemetry.Outfile.ValueString()
		}
		if len(telemetry) > 0 {
			output["telemetry"] = telemetry
		}
	}

	return output, nil
}

func populateGeminiSettingsFromMap(ctx context.Context, diags *diag.Diagnostics, settingsModel *geminiSettingsModel, data map[string]interface{}) {
	// This function is a copy of the one in gemini_data_source.go
	// It is duplicated here to avoid circular dependencies.
	// In a real-world scenario, this might be refactored into a shared package.
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
			themesMap, d := types.MapValueFrom(ctx, types.StringType, customThemes)
			if d.HasError() {
				diags.Append(d...)
				return
			}
			settingsModel.UI.CustomThemes = themesMap
		} else {
			settingsModel.UI.CustomThemes = types.MapNull(types.StringType)
		}
		if accessibilityData, ok := uiData["accessibility"].(map[string]interface{}); ok {
			if val, ok := accessibilityData["disableLoadingPhrases"].(bool); ok {
				settingsModel.UI.AccessibilityDisableLoadingPhrases = types.BoolValue(val)
			}
		}
	}

	if ideData, ok := data["ide"].(map[string]interface{}); ok {
		if val, ok := ideData["enabled"].(bool); ok {
			settingsModel.IDE.Enabled = types.BoolValue(val)
		}
		if val, ok := ideData["hasSeenNudge"].(bool); ok {
			settingsModel.IDE.HasSeenNudge = types.BoolValue(val)
		}
	}

	if privacyData, ok := data["privacy"].(map[string]interface{}); ok {
		if val, ok := privacyData["usageStatisticsEnabled"].(bool); ok {
			settingsModel.Privacy.UsageStatisticsEnabled = types.BoolValue(val)
		}
	}

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
			summaryMap, d := types.MapValueFrom(ctx, types.StringType, summaryData)
			if d.HasError() {
				diags.Append(d...)
				return
			}
			settingsModel.Model.SummarizeToolOutput = summaryMap
		} else {
			settingsModel.Model.SummarizeToolOutput = types.MapNull(types.StringType)
		}
		if chatCompressionData, ok := modelData["chatCompression"].(map[string]interface{}); ok {
			if val, ok := chatCompressionData["contextPercentageThreshold"].(string); ok {
				settingsModel.Model.ChatCompressionContextPercentageThreshold = types.StringValue(val)
			}
		}
	}

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
			var nameList []string
			for _, name := range fileNames {
				if s, ok := name.(string); ok {
					nameList = append(nameList, s)
				}
			}
			fileNamesList, d := types.ListValueFrom(ctx, types.StringType, nameList)
			if d.HasError() {
				diags.Append(d...)
				return
			}
			settingsModel.Context.FileName = fileNamesList
		} else {
			settingsModel.Context.FileName = types.ListNull(types.StringType)
		}
		if includeDirectories, ok := contextData["includeDirectories"].([]interface{}); ok {
			var dirList []string
			for _, dir := range includeDirectories {
				if s, ok := dir.(string); ok {
					dirList = append(dirList, s)
				}
			}
			dirsList, d := types.ListValueFrom(ctx, types.StringType, dirList)
			if d.HasError() {
				diags.Append(d...)
				return
			}
			settingsModel.Context.IncludeDirectories = dirsList
		} else {
			settingsModel.Context.IncludeDirectories = types.ListNull(types.StringType)
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
			var coreList []string
			for _, tool := range core {
				if s, ok := tool.(string); ok {
					coreList = append(coreList, s)
				}
			}
			coresList, d := types.ListValueFrom(ctx, types.StringType, coreList)
			if d.HasError() {
				diags.Append(d...)
				return
			}
			settingsModel.Tools.Core = coresList
		} else {
			settingsModel.Tools.Core = types.ListNull(types.StringType)
		}
		if exclude, ok := toolsData["exclude"].([]interface{}); ok {
			var excludeList []string
			for _, tool := range exclude {
				if s, ok := tool.(string); ok {
					excludeList = append(excludeList, s)
				}
			}
			excludesList, d := types.ListValueFrom(ctx, types.StringType, excludeList)
			if d.HasError() {
				diags.Append(d...)
				return
			}
			settingsModel.Tools.Exclude = excludesList
		} else {
			settingsModel.Tools.Exclude = types.ListNull(types.StringType)
		}
		if allowed, ok := toolsData["allowed"].([]interface{}); ok {
			var allowedList []string
			for _, tool := range allowed {
				if s, ok := tool.(string); ok {
					allowedList = append(allowedList, s)
				}
			}
			allowedsList, d := types.ListValueFrom(ctx, types.StringType, allowedList)
			if d.HasError() {
				diags.Append(d...)
				return
			}
			settingsModel.Tools.Allowed = allowedsList
		} else {
			settingsModel.Tools.Allowed = types.ListNull(types.StringType)
		}
	}

	if mcpData, ok := data["mcp"].(map[string]interface{}); ok {
		if val, ok := mcpData["serverCommand"].(string); ok {
			settingsModel.MCP.ServerCommand = types.StringValue(val)
		}
		if allowed, ok := mcpData["allowed"].([]interface{}); ok {
			var allowedList []string
			for _, server := range allowed {
				if s, ok := server.(string); ok {
					allowedList = append(allowedList, s)
				}
			}
			allowedsList, d := types.ListValueFrom(ctx, types.StringType, allowedList)
			if d.HasError() {
				diags.Append(d...)
				return
			}
			settingsModel.MCP.Allowed = allowedsList
		} else {
			settingsModel.MCP.Allowed = types.ListNull(types.StringType)
		}
		if excluded, ok := mcpData["excluded"].([]interface{}); ok {
			var excludedList []string
			for _, server := range excluded {
				if s, ok := server.(string); ok {
					excludedList = append(excludedList, s)
				}
			}
			excludedsList, d := types.ListValueFrom(ctx, types.StringType, excludedList)
			if d.HasError() {
				diags.Append(d...)
				return
			}
			settingsModel.MCP.Excluded = excludedsList
		} else {
			settingsModel.MCP.Excluded = types.ListNull(types.StringType)
		}
	}

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

	if advancedData, ok := data["advanced"].(map[string]interface{}); ok {
		if val, ok := advancedData["autoConfigureMemory"].(bool); ok {
			settingsModel.Advanced.AutoConfigureMemory = types.BoolValue(val)
		}
		if val, ok := advancedData["dnsResolutionOrder"].(string); ok {
			settingsModel.Advanced.DNSResolutionOrder = types.StringValue(val)
		}
		if excludedVars, ok := advancedData["excludedEnvVars"].([]interface{}); ok {
			var varsList []string
			for _, envVar := range excludedVars {
				if s, ok := envVar.(string); ok {
					varsList = append(varsList, s)
				}
			}
			varsList_, d := types.ListValueFrom(ctx, types.StringType, varsList)
			if d.HasError() {
				diags.Append(d...)
				return
			}
			settingsModel.Advanced.ExcludedEnvVars = varsList_
		} else {
			settingsModel.Advanced.ExcludedEnvVars = types.ListNull(types.StringType)
		}
		if bugCommand, ok := advancedData["bugCommand"].(map[string]interface{}); ok {
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

	if mcpServers, ok := data["mcpServers"].(map[string]interface{}); ok {
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
