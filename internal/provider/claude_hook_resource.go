package provider

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = &claudeHookResource{}
	_ resource.ResourceWithConfigure   = &claudeHookResource{}
	_ resource.ResourceWithImportState = &claudeHookResource{}
)

func NewClaudeHookResource() resource.Resource {
	return &claudeHookResource{}
}

type claudeHookResource struct {
	client *FileClient
}

type claudeHookResourceModel struct {
	ID         types.String `tfsdk:"id"`
	Scope      types.String `tfsdk:"scope"`
	Name       types.String `tfsdk:"name"`
	Content    types.String `tfsdk:"content"`
	Executable types.Bool   `tfsdk:"executable"`
}

func (r *claudeHookResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_claude_hook"
}

func (r *claudeHookResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a Claude Code hook script. Hooks are custom commands that can be configured to run before or after tool executions.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "A unique identifier for this hook resource, composed of the scope and name.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"scope": schema.StringAttribute{
				Description: "The scope of the hook file. Must be either `user` (for `~/.claude/hooks`) or `project` (for `<workdir>/.claude/hooks`).",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"name": schema.StringAttribute{
				Description: "The name of the hook, which will also be its filename (e.g., `pre-tool-use`, `post-tool-use`).",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"content": schema.StringAttribute{
				Description: "The script content of the hook file.",
				Required:    true,
			},
			"executable": schema.BoolAttribute{
				Description: "Whether to set the executable bit on the hook file. Defaults to `true`.",
				Optional:    true,
				Default:     booldefault.StaticBool(true),
				Computed:    true,
			},
		},
	}
}

func (r *claudeHookResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan claudeHookResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	filePath := r.getHookFilePath(plan.Scope.ValueString(), plan.Name.ValueString())
	if filePath == "" {
		resp.Diagnostics.AddError("Invalid scope", "Scope must be 'user' or 'project'")
		return
	}

	// Ensure directory exists
	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		resp.Diagnostics.AddError("Failed to create directory", err.Error())
		return
	}

	// Write hook file
	fileMode := os.FileMode(0644)
	if plan.Executable.ValueBool() {
		fileMode = 0755
	}

	if err := os.WriteFile(filePath, []byte(plan.Content.ValueString()), fileMode); err != nil {
		resp.Diagnostics.AddError("Failed to write hook file", err.Error())
		return
	}

	// Set computed attributes
	plan.ID = types.StringValue(fmt.Sprintf("claude-hook-%s-%s", plan.Scope.ValueString(), plan.Name.ValueString()))

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *claudeHookResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state claudeHookResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	filePath := r.getHookFilePath(state.Scope.ValueString(), state.Name.ValueString())
	if filePath == "" {
		resp.Diagnostics.AddError("Invalid scope", "Scope must be 'user' or 'project'")
		return
	}

	// Check if file exists
	fileInfo, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("Failed to stat hook file", err.Error())
		return
	}

	// Read hook file content
	data, err := os.ReadFile(filePath)
	if err != nil {
		resp.Diagnostics.AddError("Failed to read hook file", err.Error())
		return
	}

	// Update state with current file contents and permissions
	state.Content = types.StringValue(string(data))
	state.Executable = types.BoolValue(fileInfo.Mode()&0111 != 0) // Check if any execute bit is set

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *claudeHookResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan claudeHookResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	filePath := r.getHookFilePath(plan.Scope.ValueString(), plan.Name.ValueString())
	if filePath == "" {
		resp.Diagnostics.AddError("Invalid scope", "Scope must be 'user' or 'project'")
		return
	}

	// Write hook file
	fileMode := os.FileMode(0644)
	if plan.Executable.ValueBool() {
		fileMode = 0755
	}

	if err := os.WriteFile(filePath, []byte(plan.Content.ValueString()), fileMode); err != nil {
		resp.Diagnostics.AddError("Failed to write hook file", err.Error())
		return
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *claudeHookResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state claudeHookResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	filePath := r.getHookFilePath(state.Scope.ValueString(), state.Name.ValueString())
	if filePath == "" {
		resp.Diagnostics.AddError("Invalid scope", "Scope must be 'user' or 'project'")
		return
	}

	if err := os.Remove(filePath); err != nil && !os.IsNotExist(err) {
		resp.Diagnostics.AddError("Failed to delete hook file", err.Error())
		return
	}
}

func (r *claudeHookResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: scope:name (e.g., "user:pre-commit" or "project:lint")
	parts := strings.SplitN(req.ID, ":", 2)
	if len(parts) != 2 {
		resp.Diagnostics.AddError("Invalid import ID", "Import ID must be in format 'scope:name' (e.g., 'user:pre-commit')")
		return
	}

	scope := strings.TrimSpace(parts[0])
	name := strings.TrimSpace(parts[1])

	if scope != "user" && scope != "project" {
		resp.Diagnostics.AddError("Invalid scope", "Scope must be 'user' or 'project'")
		return
	}

	if name == "" {
		resp.Diagnostics.AddError("Invalid name", "Name cannot be empty")
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("scope"), scope)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("name"), name)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), fmt.Sprintf("claude-hook-%s-%s", scope, name))...)
}

func (r *claudeHookResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *claudeHookResource) getHookFilePath(scope, name string) string {
	switch scope {
	case "user":
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return ""
		}
		return filepath.Join(homeDir, ".claude", "hooks", name)
	case "project":
		if r.client == nil {
			return ""
		}
		return filepath.Join(r.client.workDir, ".claude", "hooks", name)
	default:
		return ""
	}
}
