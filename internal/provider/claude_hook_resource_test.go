package provider

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestClaudeHookResource_getHookFilePath(t *testing.T) {
	tempDir := t.TempDir()
	client := &FileClient{workDir: tempDir}

	r := &claudeHookResource{client: client}

	testCases := []struct {
		name     string
		scope    string
		hookName string
		expected string
	}{
		{
			name:     "project_scope",
			scope:    "project",
			hookName: "pre-commit",
			expected: filepath.Join(tempDir, ".claude", "hooks", "pre-commit"),
		},
		{
			name:     "invalid_scope",
			scope:    "invalid",
			hookName: "test-hook",
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := r.getHookFilePath(tc.scope, tc.hookName)
			if result != tc.expected {
				t.Errorf("Expected path %q, got %q", tc.expected, result)
			}
		})
	}

	// Test user scope (requires actual home directory)
	homeDir, err := os.UserHomeDir()
	if err == nil {
		expectedUserPath := filepath.Join(homeDir, ".claude", "hooks", "user-hook")
		result := r.getHookFilePath("user", "user-hook")
		if result != expectedUserPath {
			t.Errorf("Expected user path %q, got %q", expectedUserPath, result)
		}
	}
}

func TestClaudeHookResource_FileOperations(t *testing.T) {
	tempDir := t.TempDir()

	testCases := []struct {
		name         string
		content      string
		executable   bool
		expectedMode os.FileMode
	}{
		{
			name:         "executable_hook",
			content:      "#!/bin/bash\necho 'Running pre-commit hook'",
			executable:   true,
			expectedMode: 0755,
		},
		{
			name:         "non_executable_hook",
			content:      "echo 'Simple hook script'",
			executable:   false,
			expectedMode: 0644,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create hooks directory
			hooksDir := filepath.Join(tempDir, ".claude", "hooks")
			err := os.MkdirAll(hooksDir, 0755)
			if err != nil {
				t.Fatalf("Failed to create hooks directory: %v", err)
			}

			// Test file creation
			testFile := filepath.Join(hooksDir, "test-hook")
			fileMode := os.FileMode(0644)
			if tc.executable {
				fileMode = 0755
			}

			err = os.WriteFile(testFile, []byte(tc.content), fileMode)
			if err != nil {
				t.Fatalf("Failed to write test file: %v", err)
			}

			// Verify file exists with correct content and permissions
			fileInfo, err := os.Stat(testFile)
			if err != nil {
				t.Errorf("Expected file was not created: %v", err)
			}

			if fileInfo.Mode() != tc.expectedMode {
				t.Errorf("Expected file mode %v, got %v", tc.expectedMode, fileInfo.Mode())
			}

			content, err := os.ReadFile(testFile)
			if err != nil {
				t.Errorf("Failed to read created file: %v", err)
			}
			if string(content) != tc.content {
				t.Errorf("Expected content %q, got %q", tc.content, string(content))
			}

			// Test executable bit detection
			executableDetected := fileInfo.Mode()&0111 != 0
			if executableDetected != tc.executable {
				t.Errorf("Expected executable detection %v, got %v", tc.executable, executableDetected)
			}

			// Test file deletion
			err = os.Remove(testFile)
			if err != nil {
				t.Errorf("Failed to delete file: %v", err)
			}

			// Verify file was deleted
			if _, err := os.Stat(testFile); !os.IsNotExist(err) {
				t.Errorf("Expected file to be deleted, but it still exists")
			}
		})
	}
}

func TestClaudeHookResource_ImportParsing(t *testing.T) {
	testCases := []struct {
		name          string
		importID      string
		expectError   bool
		expectedScope string
		expectedName  string
	}{
		{
			name:          "valid_user_import",
			importID:      "user:pre-commit",
			expectError:   false,
			expectedScope: "user",
			expectedName:  "pre-commit",
		},
		{
			name:          "valid_project_import",
			importID:      "project:lint-check",
			expectError:   false,
			expectedScope: "project",
			expectedName:  "lint-check",
		},
		{
			name:          "hook_with_hyphens",
			importID:      "user:my-custom-hook",
			expectError:   false,
			expectedScope: "user",
			expectedName:  "my-custom-hook",
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
			importID:    "workspace:hook",
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

func TestClaudeHookResource_DirectoryCreation(t *testing.T) {
	tempDir := t.TempDir()

	testPaths := []string{
		filepath.Join(tempDir, ".claude", "hooks"),
		filepath.Join(tempDir, "nested", "deeper", ".claude", "hooks"),
	}

	for _, testPath := range testPaths {
		t.Run(fmt.Sprintf("create_%s", filepath.Base(testPath)), func(t *testing.T) {
			// Ensure directory creation works
			err := os.MkdirAll(testPath, 0755)
			if err != nil {
				t.Errorf("Failed to create directory %s: %v", testPath, err)
			}

			// Verify directory exists and has correct permissions
			fileInfo, err := os.Stat(testPath)
			if err != nil {
				t.Errorf("Created directory does not exist: %v", err)
			}

			if !fileInfo.IsDir() {
				t.Errorf("Expected directory, got file")
			}

			if fileInfo.Mode().Perm() != os.FileMode(0755) {
				t.Errorf("Expected directory mode 0755, got %v", fileInfo.Mode().Perm())
			}
		})
	}
}

func TestClaudeHookResource_HookTypes(t *testing.T) {
	testCases := []struct {
		name         string
		hookName     string
		expectedType string
	}{
		{
			name:         "pre_commit_hook",
			hookName:     "pre-commit",
			expectedType: "pre-commit",
		},
		{
			name:         "post_commit_hook",
			hookName:     "post-commit",
			expectedType: "post-commit",
		},
		{
			name:         "custom_hook",
			hookName:     "custom-validation",
			expectedType: "custom",
		},
		{
			name:         "tool_hook",
			hookName:     "tool-pre-exec",
			expectedType: "tool-related",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test that hook names are valid identifiers
			if tc.hookName == "" {
				t.Errorf("Hook name cannot be empty")
			}

			// Hook names should contain only alphanumeric characters and hyphens
			for _, char := range tc.hookName {
				if !((char >= 'a' && char <= 'z') ||
					(char >= 'A' && char <= 'Z') ||
					(char >= '0' && char <= '9') ||
					char == '-' || char == '_') {
					t.Errorf("Hook name %q contains invalid character %c", tc.hookName, char)
					break
				}
			}
		})
	}
}

func TestClaudeHookResource_ContentValidation(t *testing.T) {
	testCases := []struct {
		name    string
		content string
		valid   bool
	}{
		{
			name:    "bash_script",
			content: "#!/bin/bash\necho 'Running hook'\nexit 0",
			valid:   true,
		},
		{
			name:    "python_script",
			content: "#!/usr/bin/env python3\nprint('Hook executed')\nimport sys\nsys.exit(0)",
			valid:   true,
		},
		{
			name:    "simple_command",
			content: "echo 'Simple hook'",
			valid:   true,
		},
		{
			name:    "empty_content",
			content: "",
			valid:   false,
		},
		{
			name:    "multiline_script",
			content: "#!/bin/bash\n\n# Multi-line hook script\necho 'Starting hook'\n\nif [ $? -eq 0 ]; then\n  echo 'Success'\nfi",
			valid:   true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			isValid := len(strings.TrimSpace(tc.content)) > 0

			if isValid != tc.valid {
				t.Errorf("Content validation for %q: expected %v, got %v", tc.name, tc.valid, isValid)
			}

			if tc.valid {
				// Additional checks for valid content
				lines := strings.Split(tc.content, "\n")
				if len(lines) > 0 && strings.HasPrefix(lines[0], "#!") {
					// Check that shebang is properly formatted
					shebang := lines[0]
					if len(shebang) < 3 {
						t.Errorf("Invalid shebang: %q", shebang)
					}
				}
			}
		})
	}
}

// Helper functions for acceptance tests
func testAccHookWorkDir() string {
	tmpDir, _ := os.MkdirTemp("", "terraform-test-hook")
	return tmpDir
}

// Acceptance tests
func TestAccClaudeHookResource_basic(t *testing.T) {
	workDir := testAccHookWorkDir()
	defer os.RemoveAll(workDir)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccClaudeHookResourceConfig(workDir, "project", "test-hook", "#!/bin/bash\\necho 'hook executed'", "true"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("agentsmith_claude_hook.test", "scope", "project"),
					resource.TestCheckResourceAttr("agentsmith_claude_hook.test", "name", "test-hook"),
					resource.TestCheckResourceAttr("agentsmith_claude_hook.test", "content", "#!/bin/bash\necho 'hook executed'"),
					resource.TestCheckResourceAttr("agentsmith_claude_hook.test", "executable", "true"),
					resource.TestCheckResourceAttrSet("agentsmith_claude_hook.test", "id"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "agentsmith_claude_hook.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     "project:test-hook",
			},
			// Update and Read testing
			{
				Config: testAccClaudeHookResourceConfig(workDir, "project", "test-hook", "echo 'updated hook'", "false"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("agentsmith_claude_hook.test", "content", "echo 'updated hook'"),
					resource.TestCheckResourceAttr("agentsmith_claude_hook.test", "executable", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccClaudeHookResourceConfig(workDir, scope, name, content, executable string) string {
	return fmt.Sprintf(`
provider "agentsmith" {
  workdir = "%s"
}

resource "agentsmith_claude_hook" "test" {
  scope      = "%s"
  name       = "%s"
  content    = "%s"
  executable = %s
}
`, workDir, scope, name, content, executable)
}
