package provider

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"gopkg.in/yaml.v3"
)

func TestClaudeSubagentResource_getSubagentFilePath(t *testing.T) {
	tempDir := t.TempDir()
	client := &FileClient{workDir: tempDir}

	r := &claudeSubagentResource{client: client}

	testCases := []struct {
		name         string
		scope        string
		subagentName string
		expected     string
	}{
		{
			name:         "project_scope",
			scope:        "project",
			subagentName: "helper",
			expected:     filepath.Join(tempDir, ".claude", "agents", "helper.md"),
		},
		{
			name:         "invalid_scope",
			scope:        "invalid",
			subagentName: "test-agent",
			expected:     "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := r.getSubagentFilePath(tc.scope, tc.subagentName)
			if result != tc.expected {
				t.Errorf("Expected path %q, got %q", tc.expected, result)
			}
		})
	}

	// Test user scope (requires actual home directory)
	homeDir, err := os.UserHomeDir()
	if err == nil {
		expectedUserPath := filepath.Join(homeDir, ".claude", "agents", "user-agent.md")
		result := r.getSubagentFilePath("user", "user-agent")
		if result != expectedUserPath {
			t.Errorf("Expected user path %q, got %q", expectedUserPath, result)
		}
	}
}

func TestClaudeSubagentResource_modelToMarkdown(t *testing.T) {
	r := &claudeSubagentResource{}

	testCases := []struct {
		name     string
		model    claudeSubagentResourceModel
		validate func(*testing.T, string)
	}{
		{
			name: "complete_subagent",
			model: claudeSubagentResourceModel{
				Name:        types.StringValue("test-agent"),
				Model:       types.StringValue("claude-3-5-sonnet"),
				Description: types.StringValue("A helpful test agent"),
				Color:       types.StringValue("blue"),
				Prompt:      types.StringValue("You are a helpful assistant that helps with testing."),
			},
			validate: func(t *testing.T, content string) {
				// Should have YAML frontmatter
				if !strings.HasPrefix(content, "---\n") {
					t.Error("Expected content to start with YAML frontmatter")
				}

				parts := strings.SplitN(content, "---", 3)
				if len(parts) != 3 {
					t.Error("Expected content to have proper frontmatter structure")
					return
				}

				// Parse frontmatter
				var frontmatter subagentFrontmatter
				err := yaml.Unmarshal([]byte(parts[1]), &frontmatter)
				if err != nil {
					t.Errorf("Failed to parse frontmatter: %v", err)
					return
				}

				// Validate frontmatter fields
				if frontmatter.Name != "test-agent" {
					t.Errorf("Expected name 'test-agent', got %q", frontmatter.Name)
				}
				if frontmatter.Model != "claude-3-5-sonnet" {
					t.Errorf("Expected model 'claude-3-5-sonnet', got %q", frontmatter.Model)
				}
				if frontmatter.Description != "A helpful test agent" {
					t.Errorf("Expected description 'A helpful test agent', got %q", frontmatter.Description)
				}
				if frontmatter.Color != "blue" {
					t.Errorf("Expected color 'blue', got %q", frontmatter.Color)
				}

				// Validate prompt content
				prompt := strings.TrimSpace(parts[2])
				if prompt != "You are a helpful assistant that helps with testing." {
					t.Errorf("Expected prompt content mismatch, got %q", prompt)
				}
			},
		},
		{
			name: "minimal_subagent",
			model: claudeSubagentResourceModel{
				Name:   types.StringValue("minimal"),
				Prompt: types.StringValue("Simple prompt."),
			},
			validate: func(t *testing.T, content string) {
				parts := strings.SplitN(content, "---", 3)
				if len(parts) != 3 {
					t.Error("Expected content to have proper frontmatter structure")
					return
				}

				var frontmatter subagentFrontmatter
				err := yaml.Unmarshal([]byte(parts[1]), &frontmatter)
				if err != nil {
					t.Errorf("Failed to parse frontmatter: %v", err)
					return
				}

				if frontmatter.Name != "minimal" {
					t.Errorf("Expected name 'minimal', got %q", frontmatter.Name)
				}
				if frontmatter.Model != "" {
					t.Errorf("Expected empty model, got %q", frontmatter.Model)
				}
				if frontmatter.Description != "" {
					t.Errorf("Expected empty description, got %q", frontmatter.Description)
				}
				if frontmatter.Color != "" {
					t.Errorf("Expected empty color, got %q", frontmatter.Color)
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := r.modelToMarkdown(&tc.model)
			if err != nil {
				t.Errorf("Failed to convert model to markdown: %v", err)
				return
			}
			tc.validate(t, result)
		})
	}
}

func TestClaudeSubagentResource_markdownToModel(t *testing.T) {
	r := &claudeSubagentResource{}

	testCases := []struct {
		name     string
		content  string
		validate func(*testing.T, *claudeSubagentResourceModel)
	}{
		{
			name: "complete_markdown",
			content: `---
name: test-agent
model: claude-3-5-sonnet
description: A helpful test agent
color: blue
---
You are a helpful assistant that helps with testing.

This is a multi-line prompt with detailed instructions.`,
			validate: func(t *testing.T, model *claudeSubagentResourceModel) {
				if model.Name.ValueString() != "test-agent" {
					t.Errorf("Expected name 'test-agent', got %q", model.Name.ValueString())
				}
				if model.Model.ValueString() != "claude-3-5-sonnet" {
					t.Errorf("Expected model 'claude-3-5-sonnet', got %q", model.Model.ValueString())
				}
				if model.Description.ValueString() != "A helpful test agent" {
					t.Errorf("Expected description 'A helpful test agent', got %q", model.Description.ValueString())
				}
				if model.Color.ValueString() != "blue" {
					t.Errorf("Expected color 'blue', got %q", model.Color.ValueString())
				}

				expectedPrompt := "You are a helpful assistant that helps with testing.\n\nThis is a multi-line prompt with detailed instructions."
				if model.Prompt.ValueString() != expectedPrompt {
					t.Errorf("Prompt content mismatch:\nExpected: %q\nGot: %q", expectedPrompt, model.Prompt.ValueString())
				}
			},
		},
		{
			name: "minimal_markdown",
			content: `---
name: minimal
---
Simple prompt.`,
			validate: func(t *testing.T, model *claudeSubagentResourceModel) {
				if model.Name.ValueString() != "minimal" {
					t.Errorf("Expected name 'minimal', got %q", model.Name.ValueString())
				}
				if !model.Model.IsNull() {
					t.Errorf("Expected model to be null, got %q", model.Model.ValueString())
				}
				if !model.Description.IsNull() {
					t.Errorf("Expected description to be null, got %q", model.Description.ValueString())
				}
				if !model.Color.IsNull() {
					t.Errorf("Expected color to be null, got %q", model.Color.ValueString())
				}
				if model.Prompt.ValueString() != "Simple prompt." {
					t.Errorf("Expected prompt 'Simple prompt.', got %q", model.Prompt.ValueString())
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var model claudeSubagentResourceModel
			err := r.markdownToModel(&model, tc.content)
			if err != nil {
				t.Errorf("Failed to parse markdown: %v", err)
				return
			}
			tc.validate(t, &model)
		})
	}
}

func TestClaudeSubagentResource_YAMLFrontmatterParsing(t *testing.T) {
	testCases := []struct {
		name        string
		content     string
		expectError bool
	}{
		{
			name: "valid_frontmatter",
			content: `---
name: test
model: claude-3-5-sonnet
description: Test agent
color: red
---
Prompt content here.`,
			expectError: false,
		},
		{
			name: "invalid_yaml",
			content: `---
name: test
invalid: yaml: structure: here
---
Prompt content.`,
			expectError: true,
		},
		{
			name: "missing_separators",
			content: `name: test
model: claude-3-5-sonnet
Prompt without proper frontmatter.`,
			expectError: true,
		},
		{
			name: "empty_frontmatter",
			content: `---
---
Just prompt content.`,
			expectError: false,
		},
		{
			name: "extra_separators",
			content: `---
name: test
---
Prompt content.
---
Extra separator should be part of content.`,
			expectError: false,
		},
	}

	r := &claudeSubagentResource{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var model claudeSubagentResourceModel
			err := r.markdownToModel(&model, tc.content)

			if tc.expectError {
				if err == nil {
					t.Error("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}
		})
	}
}

func TestClaudeSubagentResource_FileOperations(t *testing.T) {
	tempDir := t.TempDir()

	// Create agents directory
	agentsDir := filepath.Join(tempDir, ".claude", "agents")
	err := os.MkdirAll(agentsDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create agents directory: %v", err)
	}

	testContent := `---
name: test-agent
model: claude-3-5-sonnet
description: Test subagent
color: green
---
You are a test assistant.

Help users with testing tasks.`

	testFile := filepath.Join(agentsDir, "test-agent.md")

	// Test file creation
	err = os.WriteFile(testFile, []byte(testContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	// Test file reading
	content, err := os.ReadFile(testFile)
	if err != nil {
		t.Errorf("Failed to read test file: %v", err)
	}

	if string(content) != testContent {
		t.Errorf("File content mismatch:\nExpected: %q\nGot: %q", testContent, string(content))
	}

	// Test file permissions
	fileInfo, err := os.Stat(testFile)
	if err != nil {
		t.Errorf("Failed to stat test file: %v", err)
	}

	if fileInfo.Mode() != os.FileMode(0644) {
		t.Errorf("Expected file mode 0644, got %v", fileInfo.Mode())
	}

	// Test file deletion
	err = os.Remove(testFile)
	if err != nil {
		t.Errorf("Failed to delete file: %v", err)
	}

	if _, err := os.Stat(testFile); !os.IsNotExist(err) {
		t.Errorf("Expected file to be deleted, but it still exists")
	}
}

func TestClaudeSubagentResource_ImportParsing(t *testing.T) {
	testCases := []struct {
		name          string
		importID      string
		expectError   bool
		expectedScope string
		expectedName  string
	}{
		{
			name:          "valid_user_import",
			importID:      "user:helper-agent",
			expectError:   false,
			expectedScope: "user",
			expectedName:  "helper-agent",
		},
		{
			name:          "valid_project_import",
			importID:      "project:code-reviewer",
			expectError:   false,
			expectedScope: "project",
			expectedName:  "code-reviewer",
		},
		{
			name:        "invalid_format",
			importID:    "invalid-format",
			expectError: true,
		},
		{
			name:        "empty_name",
			importID:    "user:",
			expectError: true,
		},
		{
			name:        "invalid_scope",
			importID:    "workspace:agent",
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			parts := strings.SplitN(tc.importID, ":", 2)

			if len(parts) != 2 {
				if !tc.expectError {
					t.Errorf("Expected valid format but got invalid")
				}
				return
			}

			scope := strings.TrimSpace(parts[0])
			name := strings.TrimSpace(parts[1])

			if scope != "user" && scope != "project" {
				if !tc.expectError {
					t.Errorf("Expected valid scope but got invalid: %s", scope)
				}
				return
			}

			if name == "" {
				if !tc.expectError {
					t.Errorf("Expected valid name but got empty")
				}
				return
			}

			if tc.expectError {
				t.Errorf("Expected error but got valid parsing")
			}

			if scope != tc.expectedScope {
				t.Errorf("Expected scope %q, got %q", tc.expectedScope, scope)
			}
			if name != tc.expectedName {
				t.Errorf("Expected name %q, got %q", tc.expectedName, name)
			}
		})
	}
}

func TestClaudeSubagentResource_SubagentValidation(t *testing.T) {
	testCases := []struct {
		name      string
		agentName string
		model     string
		color     string
		valid     bool
	}{
		{
			name:      "valid_agent_basic",
			agentName: "helper",
			model:     "claude-3-5-sonnet",
			color:     "blue",
			valid:     true,
		},
		{
			name:      "valid_agent_with_hyphens",
			agentName: "code-reviewer",
			model:     "claude-3-opus",
			color:     "green",
			valid:     true,
		},
		{
			name:      "invalid_empty_name",
			agentName: "",
			model:     "claude-3-5-sonnet",
			color:     "blue",
			valid:     false,
		},
		{
			name:      "valid_minimal",
			agentName: "minimal",
			model:     "",
			color:     "",
			valid:     true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Basic name validation
			nameValid := len(strings.TrimSpace(tc.agentName)) > 0

			if nameValid != tc.valid {
				t.Errorf("Name validation for %q: expected %v, got %v", tc.agentName, tc.valid, nameValid)
			}

			if tc.valid && tc.agentName != "" {
				// Agent names should be valid identifiers
				for _, char := range tc.agentName {
					if !((char >= 'a' && char <= 'z') ||
						(char >= 'A' && char <= 'Z') ||
						(char >= '0' && char <= '9') ||
						char == '-' || char == '_') {
						t.Errorf("Agent name %q contains invalid character %c", tc.agentName, char)
						break
					}
				}
			}
		})
	}
}

// Helper functions for acceptance tests
func testAccSubagentWorkDir() string {
	tmpDir, _ := os.MkdirTemp("", "terraform-test-subagent")
	return tmpDir
}

// Acceptance tests
func TestAccClaudeSubagentResource_basic(t *testing.T) {
	workDir := testAccSubagentWorkDir()
	defer os.RemoveAll(workDir)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccClaudeSubagentResourceConfig(workDir, "project", "test-agent", "claude-3-5-sonnet", "A test agent", "blue", "You are a helpful test assistant."),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("agentsmith_claude_subagent.test", "scope", "project"),
					resource.TestCheckResourceAttr("agentsmith_claude_subagent.test", "name", "test-agent"),
					resource.TestCheckResourceAttr("agentsmith_claude_subagent.test", "model", "claude-3-5-sonnet"),
					resource.TestCheckResourceAttr("agentsmith_claude_subagent.test", "description", "A test agent"),
					resource.TestCheckResourceAttr("agentsmith_claude_subagent.test", "color", "blue"),
					resource.TestCheckResourceAttr("agentsmith_claude_subagent.test", "prompt", "You are a helpful test assistant."),
					resource.TestCheckResourceAttrSet("agentsmith_claude_subagent.test", "id"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "agentsmith_claude_subagent.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     "project:test-agent",
			},
			// Update and Read testing
			{
				Config: testAccClaudeSubagentResourceConfig(workDir, "project", "test-agent", "claude-3-opus", "Updated test agent", "red", "You are an updated test assistant."),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("agentsmith_claude_subagent.test", "model", "claude-3-opus"),
					resource.TestCheckResourceAttr("agentsmith_claude_subagent.test", "description", "Updated test agent"),
					resource.TestCheckResourceAttr("agentsmith_claude_subagent.test", "color", "red"),
					resource.TestCheckResourceAttr("agentsmith_claude_subagent.test", "prompt", "You are an updated test assistant."),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccClaudeSubagentResourceConfig(workDir, scope, name, model, description, color, prompt string) string {
	return fmt.Sprintf(`
provider "agentsmith" {
  workdir = "%s"
}

resource "agentsmith_claude_subagent" "test" {
  scope       = "%s"
  name        = "%s"
  model       = "%s"
  description = "%s"
  color       = "%s"
  prompt      = "%s"
}
`, workDir, scope, name, model, description, color, prompt)
}
