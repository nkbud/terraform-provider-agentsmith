package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = &claudeGlobalConfigResource{}
	_ resource.ResourceWithConfigure   = &claudeGlobalConfigResource{}
	_ resource.ResourceWithImportState = &claudeGlobalConfigResource{}
)

func NewClaudeGlobalConfigResource() resource.Resource {
	return &claudeGlobalConfigResource{}
}

type claudeGlobalConfigResource struct {
	client *FileClient
}

type claudeGlobalConfigResourceModel struct {
	ID                    types.String `tfsdk:"id"`
	AutoUpdates           types.Bool   `tfsdk:"auto_updates"`
	PreferredNotifChannel types.String `tfsdk:"preferred_notif_channel"`
	Theme                 types.String `tfsdk:"theme"`
	Verbose               types.Bool   `tfsdk:"verbose"`
}

func (r *claudeGlobalConfigResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_claude_global_config"
}

func (r *claudeGlobalConfigResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages the Claude Code global configuration file at `~/.claude/config.json`. This resource is a singleton; only one should be defined.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "A static identifier for this singleton resource.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"auto_updates": schema.BoolAttribute{
				Description: "DEPRECATED. This field is no longer used. Use the `DISABLE_AUTOUPDATER` environment variable instead.",
				Optional:    true,
				Default:     booldefault.StaticBool(true),
				Computed:    true,
			},
			"preferred_notif_channel": schema.StringAttribute{
				Description: "The preferred channel for receiving notifications. Valid values are `iterm2`, `iterm2_with_bell`, `terminal_bell`, or `notifications_disabled`.",
				Optional:    true,
			},
			"theme": schema.StringAttribute{
				Description: "The color theme for the UI. Valid values are `dark`, `light`, `light-daltonized`, or `dark-daltonized`.",
				Optional:    true,
			},
			"verbose": schema.BoolAttribute{
				Description: "If true, shows full bash and command outputs. Defaults to `false`.",
				Optional:    true,
				Default:     booldefault.StaticBool(false),
				Computed:    true,
			},
		},
	}
}

func (r *claudeGlobalConfigResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan claudeGlobalConfigResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	filePath := r.getGlobalConfigFilePath()

	// Ensure directory exists
	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		resp.Diagnostics.AddError("Failed to create directory", err.Error())
		return
	}

	// Create global config JSON
	configData := r.modelToConfig(&plan)
	jsonData, err := json.MarshalIndent(configData, "", "  ")
	if err != nil {
		resp.Diagnostics.AddError("Failed to marshal JSON", err.Error())
		return
	}

	if err := os.WriteFile(filePath, jsonData, 0644); err != nil {
		resp.Diagnostics.AddError("Failed to write global config file", err.Error())
		return
	}

	// Set computed attributes
	plan.ID = types.StringValue("claude-global-config")

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *claudeGlobalConfigResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state claudeGlobalConfigResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	filePath := r.getGlobalConfigFilePath()

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		resp.State.RemoveResource(ctx)
		return
	}

	// Read and parse global config file
	data, err := os.ReadFile(filePath)
	if err != nil {
		resp.Diagnostics.AddError("Failed to read global config file", err.Error())
		return
	}

	var configData map[string]interface{}
	if err := json.Unmarshal(data, &configData); err != nil {
		resp.Diagnostics.AddError("Failed to parse global config JSON", err.Error())
		return
	}

	// Update state with current file contents
	r.configToModel(&state, configData)

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *claudeGlobalConfigResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan claudeGlobalConfigResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	filePath := r.getGlobalConfigFilePath()

	// Create global config JSON
	configData := r.modelToConfig(&plan)
	jsonData, err := json.MarshalIndent(configData, "", "  ")
	if err != nil {
		resp.Diagnostics.AddError("Failed to marshal JSON", err.Error())
		return
	}

	if err := os.WriteFile(filePath, jsonData, 0644); err != nil {
		resp.Diagnostics.AddError("Failed to write global config file", err.Error())
		return
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *claudeGlobalConfigResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	filePath := r.getGlobalConfigFilePath()

	if err := os.Remove(filePath); err != nil && !os.IsNotExist(err) {
		resp.Diagnostics.AddError("Failed to delete global config file", err.Error())
		return
	}
}

func (r *claudeGlobalConfigResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import uses a fixed ID since there's only one global config
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), "claude-global-config")...)
}

func (r *claudeGlobalConfigResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *claudeGlobalConfigResource) getGlobalConfigFilePath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(homeDir, ".claude", "config.json")
}

func (r *claudeGlobalConfigResource) modelToConfig(model *claudeGlobalConfigResourceModel) map[string]interface{} {
	config := make(map[string]interface{})

	if !model.AutoUpdates.IsNull() {
		config["autoUpdates"] = model.AutoUpdates.ValueBool()
	}
	if !model.PreferredNotifChannel.IsNull() {
		config["preferredNotifChannel"] = model.PreferredNotifChannel.ValueString()
	}
	if !model.Theme.IsNull() {
		config["theme"] = model.Theme.ValueString()
	}
	if !model.Verbose.IsNull() {
		config["verbose"] = model.Verbose.ValueBool()
	}

	return config
}

func (r *claudeGlobalConfigResource) configToModel(model *claudeGlobalConfigResourceModel, config map[string]interface{}) {
	if val, ok := config["autoUpdates"].(bool); ok {
		model.AutoUpdates = types.BoolValue(val)
	} else {
		model.AutoUpdates = types.BoolNull()
	}

	if val, ok := config["preferredNotifChannel"].(string); ok {
		model.PreferredNotifChannel = types.StringValue(val)
	} else {
		model.PreferredNotifChannel = types.StringNull()
	}

	if val, ok := config["theme"].(string); ok {
		model.Theme = types.StringValue(val)
	} else {
		model.Theme = types.StringNull()
	}

	if val, ok := config["verbose"].(bool); ok {
		model.Verbose = types.BoolValue(val)
	} else {
		model.Verbose = types.BoolNull()
	}
}
