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
	_ resource.Resource                = &claudeCommandResource{}
	_ resource.ResourceWithConfigure   = &claudeCommandResource{}
	_ resource.ResourceWithImportState = &claudeCommandResource{}
)

func NewClaudeCommandResource() resource.Resource {
	return &claudeCommandResource{}
}

type claudeCommandResource struct {
	client *FileClient
}

type claudeCommandResourceModel struct {
	ID         types.String `tfsdk:"id"`
	Scope      types.String `tfsdk:"scope"`
	Name       types.String `tfsdk:"name"`
	Content    types.String `tfsdk:"content"`
	Executable types.Bool   `tfsdk:"executable"`
}

func (r *claudeCommandResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_claude_command"
}

func (r *claudeCommandResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a Claude Code custom slash command. These commands are scripts that can be invoked within a Claude session using `/name`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "A unique identifier for this command resource, composed of the scope and name.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"scope": schema.StringAttribute{
				Description: "The scope of the command file. Must be either `user` (for `~/.claude/commands`) or `project` (for `<workdir>/.claude/commands`).",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"name": schema.StringAttribute{
				Description: "The name of the command, which will also be its filename. This is the name you will use to invoke the command in Claude (e.g., `/my-command`).",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"content": schema.StringAttribute{
				Description: "The script content of the command file. This can be any executable script content, such as a shell script.",
				Required:    true,
			},
			"executable": schema.BoolAttribute{
				Description: "Whether to set the executable bit on the command file. Defaults to `true`.",
				Optional:    true,
				Default:     booldefault.StaticBool(true),
				Computed:    true,
			},
		},
	}
}

func (r *claudeCommandResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan claudeCommandResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	filePath := r.getCommandFilePath(plan.Scope.ValueString(), plan.Name.ValueString())
	if filePath == "" {
		resp.Diagnostics.AddError("Invalid scope", "Scope must be 'user' or 'project'")
		return
	}

	// Ensure directory exists
	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		resp.Diagnostics.AddError("Failed to create directory", err.Error())
		return
	}

	// Write command file
	fileMode := os.FileMode(0644)
	if plan.Executable.ValueBool() {
		fileMode = 0755
	}

	if err := os.WriteFile(filePath, []byte(plan.Content.ValueString()), fileMode); err != nil {
		resp.Diagnostics.AddError("Failed to write command file", err.Error())
		return
	}

	// Set computed attributes
	plan.ID = types.StringValue(fmt.Sprintf("claude-command-%s-%s", plan.Scope.ValueString(), plan.Name.ValueString()))

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *claudeCommandResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state claudeCommandResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	filePath := r.getCommandFilePath(state.Scope.ValueString(), state.Name.ValueString())
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
		resp.Diagnostics.AddError("Failed to stat command file", err.Error())
		return
	}

	// Read command file content
	data, err := os.ReadFile(filePath)
	if err != nil {
		resp.Diagnostics.AddError("Failed to read command file", err.Error())
		return
	}

	// Update state with current file contents and permissions
	state.Content = types.StringValue(string(data))
	state.Executable = types.BoolValue(fileInfo.Mode()&0111 != 0) // Check if any execute bit is set

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *claudeCommandResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan claudeCommandResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	filePath := r.getCommandFilePath(plan.Scope.ValueString(), plan.Name.ValueString())
	if filePath == "" {
		resp.Diagnostics.AddError("Invalid scope", "Scope must be 'user' or 'project'")
		return
	}

	// Write command file
	fileMode := os.FileMode(0644)
	if plan.Executable.ValueBool() {
		fileMode = 0755
	}

	if err := os.WriteFile(filePath, []byte(plan.Content.ValueString()), fileMode); err != nil {
		resp.Diagnostics.AddError("Failed to write command file", err.Error())
		return
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *claudeCommandResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state claudeCommandResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	filePath := r.getCommandFilePath(state.Scope.ValueString(), state.Name.ValueString())
	if filePath == "" {
		resp.Diagnostics.AddError("Invalid scope", "Scope must be 'user' or 'project'")
		return
	}

	if err := os.Remove(filePath); err != nil && !os.IsNotExist(err) {
		resp.Diagnostics.AddError("Failed to delete command file", err.Error())
		return
	}
}

func (r *claudeCommandResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: scope:name (e.g., "user:my-command" or "project:build")
	parts := strings.SplitN(req.ID, ":", 2)
	if len(parts) != 2 {
		resp.Diagnostics.AddError("Invalid import ID", "Import ID must be in format 'scope:name' (e.g., 'user:my-command')")
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
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), fmt.Sprintf("claude-command-%s-%s", scope, name))...)
}

func (r *claudeCommandResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *claudeCommandResource) getCommandFilePath(scope, name string) string {
	switch scope {
	case "user":
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return ""
		}
		return filepath.Join(homeDir, ".claude", "commands", name)
	case "project":
		if r.client == nil {
			return ""
		}
		return filepath.Join(r.client.workDir, ".claude", "commands", name)
	default:
		return ""
	}
}
