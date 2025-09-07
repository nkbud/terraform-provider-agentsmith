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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"gopkg.in/yaml.v3"
)

var (
	_ resource.Resource                = &claudeSubagentResource{}
	_ resource.ResourceWithConfigure   = &claudeSubagentResource{}
	_ resource.ResourceWithImportState = &claudeSubagentResource{}
)

func NewClaudeSubagentResource() resource.Resource {
	return &claudeSubagentResource{}
}

type claudeSubagentResource struct {
	client *FileClient
}

type claudeSubagentResourceModel struct {
	ID          types.String `tfsdk:"id"`
	Scope       types.String `tfsdk:"scope"`
	Name        types.String `tfsdk:"name"`
	Model       types.String `tfsdk:"model"`
	Description types.String `tfsdk:"description"`
	Color       types.String `tfsdk:"color"`
	Prompt      types.String `tfsdk:"prompt"`
}

type subagentFrontmatter struct {
	Name        string `yaml:"name"`
	Model       string `yaml:"model,omitempty"`
	Description string `yaml:"description,omitempty"`
	Color       string `yaml:"color,omitempty"`
}

func (r *claudeSubagentResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_claude_subagent"
}

func (r *claudeSubagentResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a Claude Code subagent. Subagents are specialized AI assistants with custom prompts and tool permissions, defined in Markdown files with YAML frontmatter.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "A unique identifier for this subagent resource, composed of the scope and name.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"scope": schema.StringAttribute{
				Description: "The scope of the subagent file. Must be either `user` (for `~/.claude/agents`) or `project` (for `<workdir>/.claude/agents`).",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"name": schema.StringAttribute{
				Description: "The name of the subagent. This is used as the filename (e.g., `my-agent.md`) and is specified in the YAML frontmatter.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"model": schema.StringAttribute{
				Description: "An optional model override for this specific subagent.",
				Optional:    true,
			},
			"description": schema.StringAttribute{
				Description: "A human-readable description of the subagent's purpose, included in the YAML frontmatter.",
				Optional:    true,
			},
			"color": schema.StringAttribute{
				Description: "A display color for the subagent in the UI, included in the YAML frontmatter.",
				Optional:    true,
			},
			"prompt": schema.StringAttribute{
				Description: "The main instructional prompt for the subagent, which forms the body of the Markdown file.",
				Required:    true,
			},
		},
	}
}

func (r *claudeSubagentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan claudeSubagentResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	filePath := r.getSubagentFilePath(plan.Scope.ValueString(), plan.Name.ValueString())
	if filePath == "" {
		resp.Diagnostics.AddError("Invalid scope", "Scope must be 'user' or 'project'")
		return
	}

	// Ensure directory exists
	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		resp.Diagnostics.AddError("Failed to create directory", err.Error())
		return
	}

	// Create subagent file content
	content, err := r.modelToMarkdown(&plan)
	if err != nil {
		resp.Diagnostics.AddError("Failed to create subagent content", err.Error())
		return
	}

	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		resp.Diagnostics.AddError("Failed to write subagent file", err.Error())
		return
	}

	// Set computed attributes
	plan.ID = types.StringValue(fmt.Sprintf("claude-subagent-%s-%s", plan.Scope.ValueString(), plan.Name.ValueString()))

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *claudeSubagentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state claudeSubagentResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	filePath := r.getSubagentFilePath(state.Scope.ValueString(), state.Name.ValueString())
	if filePath == "" {
		resp.Diagnostics.AddError("Invalid scope", "Scope must be 'user' or 'project'")
		return
	}

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		resp.State.RemoveResource(ctx)
		return
	}

	// Read and parse subagent file
	data, err := os.ReadFile(filePath)
	if err != nil {
		resp.Diagnostics.AddError("Failed to read subagent file", err.Error())
		return
	}

	// Update state with current file contents
	if err := r.markdownToModel(&state, string(data)); err != nil {
		resp.Diagnostics.AddError("Failed to parse subagent file", err.Error())
		return
	}

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *claudeSubagentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan claudeSubagentResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	filePath := r.getSubagentFilePath(plan.Scope.ValueString(), plan.Name.ValueString())
	if filePath == "" {
		resp.Diagnostics.AddError("Invalid scope", "Scope must be 'user' or 'project'")
		return
	}

	// Create subagent file content
	content, err := r.modelToMarkdown(&plan)
	if err != nil {
		resp.Diagnostics.AddError("Failed to create subagent content", err.Error())
		return
	}

	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		resp.Diagnostics.AddError("Failed to write subagent file", err.Error())
		return
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *claudeSubagentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state claudeSubagentResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	filePath := r.getSubagentFilePath(state.Scope.ValueString(), state.Name.ValueString())
	if filePath == "" {
		resp.Diagnostics.AddError("Invalid scope", "Scope must be 'user' or 'project'")
		return
	}

	if err := os.Remove(filePath); err != nil && !os.IsNotExist(err) {
		resp.Diagnostics.AddError("Failed to delete subagent file", err.Error())
		return
	}
}

func (r *claudeSubagentResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: scope:name (e.g., "user:my-agent" or "project:helper")
	parts := strings.SplitN(req.ID, ":", 2)
	if len(parts) != 2 {
		resp.Diagnostics.AddError("Invalid import ID", "Import ID must be in format 'scope:name' (e.g., 'user:my-agent')")
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
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), fmt.Sprintf("claude-subagent-%s-%s", scope, name))...)
}

func (r *claudeSubagentResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*FileClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *FileClient, got: %%T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = client
}

func (r *claudeSubagentResource) getSubagentFilePath(scope, name string) string {
	switch scope {
	case "user":
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return ""
		}
		return filepath.Join(homeDir, ".claude", "agents", name+".md")
	case "project":
		if r.client == nil {
			return ""
		}
		return filepath.Join(r.client.workDir, ".claude", "agents", name+".md")
	default:
		return ""
	}
}

func (r *claudeSubagentResource) modelToMarkdown(model *claudeSubagentResourceModel) (string, error) {
	frontmatter := subagentFrontmatter{
		Name: model.Name.ValueString(),
	}

	if !model.Model.IsNull() {
		frontmatter.Model = model.Model.ValueString()
	}
	if !model.Description.IsNull() {
		frontmatter.Description = model.Description.ValueString()
	}
	if !model.Color.IsNull() {
		frontmatter.Color = model.Color.ValueString()
	}

	yamlData, err := yaml.Marshal(frontmatter)
	if err != nil {
		return "", fmt.Errorf("failed to marshal YAML frontmatter: %%w", err)
	}

	prompt := model.Prompt.ValueString()
	content := fmt.Sprintf("---\n%%s---\n%%s", yamlData, prompt)

	return content, nil
}

func (r *claudeSubagentResource) markdownToModel(model *claudeSubagentResourceModel, content string) error {
	parts := strings.SplitN(content, "---", 3)
	if len(parts) < 3 {
		return fmt.Errorf("invalid format: missing YAML frontmatter separators")
	}

	var frontmatter subagentFrontmatter
	if err := yaml.Unmarshal([]byte(parts[1]), &frontmatter); err != nil {
		return fmt.Errorf("failed to parse YAML frontmatter: %%w", err)
	}

	// Update model with parsed data
	model.Name = types.StringValue(frontmatter.Name)

	if frontmatter.Model != "" {
		model.Model = types.StringValue(frontmatter.Model)
	} else {
		model.Model = types.StringNull()
	}

	if frontmatter.Description != "" {
		model.Description = types.StringValue(frontmatter.Description)
	} else {
		model.Description = types.StringNull()
	}

	if frontmatter.Color != "" {
		model.Color = types.StringValue(frontmatter.Color)
	} else {
		model.Color = types.StringNull()
	}

	model.Prompt = types.StringValue(strings.TrimSpace(parts[2]))

	return nil
}
