package provider

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestClaudeCommandResource_getCommandFilePath(t *testing.T) {
	tempDir := t.TempDir()
	client := &FileClient{workDir: tempDir}

	r := &claudeCommandResource{client: client}

	testCases := []struct {
		name     string
		scope    string
		cmdName  string
		expected string
	}{
		{
			name:     "project_scope",
			scope:    "project",
			cmdName:  "test-cmd",
			expected: filepath.Join(tempDir, ".claude", "commands", "test-cmd"),
		},
		{
			name:     "invalid_scope",
			scope:    "invalid",
			cmdName:  "test-cmd",
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := r.getCommandFilePath(tc.scope, tc.cmdName)
			if result != tc.expected {
				t.Errorf("Expected path %q, got %q", tc.expected, result)
			}
		})
	}

	// Test user scope (requires actual home directory)
	homeDir, err := os.UserHomeDir()
	if err == nil {
		expectedUserPath := filepath.Join(homeDir, ".claude", "commands", "user-cmd")
		result := r.getCommandFilePath("user", "user-cmd")
		if result != expectedUserPath {
			t.Errorf("Expected user path %q, got %q", expectedUserPath, result)
		}
	}
}

func TestClaudeCommandResource_FileOperations(t *testing.T) {
	tempDir := t.TempDir()

	testCases := []struct {
		name         string
		content      string
		executable   bool
		expectedMode os.FileMode
	}{
		{
			name:         "executable_file",
			content:      "#!/bin/bash\necho 'test'",
			executable:   true,
			expectedMode: 0755,
		},
		{
			name:         "non_executable_file",
			content:      "npm run build",
			executable:   false,
			expectedMode: 0644,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create command directory
			commandDir := filepath.Join(tempDir, ".claude", "commands")
			err := os.MkdirAll(commandDir, 0755)
			if err != nil {
				t.Fatalf("Failed to create command directory: %v", err)
			}

			// Test file creation
			testFile := filepath.Join(commandDir, "test-file")
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

func TestClaudeCommandResource_ImportParsing(t *testing.T) {
	testCases := []struct {
		name          string
		importID      string
		expectError   bool
		expectedScope string
		expectedName  string
	}{
		{
			name:          "valid_user_import",
			importID:      "user:my-command",
			expectError:   false,
			expectedScope: "user",
			expectedName:  "my-command",
		},
		{
			name:          "valid_project_import",
			importID:      "project:build-script",
			expectError:   false,
			expectedScope: "project",
			expectedName:  "build-script",
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
			name:          "too_many_parts",
			importID:      "user:name:extra",
			expectError:   false, // SplitN with 2 should handle this
			expectedScope: "user",
			expectedName:  "name:extra",
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

func TestClaudeCommandResource_DirectoryCreation(t *testing.T) {
	tempDir := t.TempDir()

	testPaths := []string{
		filepath.Join(tempDir, ".claude", "commands"),
		filepath.Join(tempDir, "nested", "deeper", ".claude", "commands"),
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

// Helper functions for acceptance tests
func testWorkDir() string {
	tmpDir, _ := os.MkdirTemp("", "terraform-test")
	return tmpDir
}

// Acceptance tests
func TestAccClaudeCommandResource_basic(t *testing.T) {
	workDir := testWorkDir()
	defer os.RemoveAll(workDir)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccClaudeCommandResourceConfig(workDir, "project", "test-acc", "echo 'hello'", "true"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("agentsmith_claude_command.test", "scope", "project"),
					resource.TestCheckResourceAttr("agentsmith_claude_command.test", "name", "test-acc"),
					resource.TestCheckResourceAttr("agentsmith_claude_command.test", "content", "echo 'hello'"),
					resource.TestCheckResourceAttr("agentsmith_claude_command.test", "executable", "true"),
					resource.TestCheckResourceAttrSet("agentsmith_claude_command.test", "id"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "agentsmith_claude_command.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     "project:test-acc",
			},
			// Update and Read testing
			{
				Config: testAccClaudeCommandResourceConfig(workDir, "project", "test-acc", "echo 'updated'", "false"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("agentsmith_claude_command.test", "content", "echo 'updated'"),
					resource.TestCheckResourceAttr("agentsmith_claude_command.test", "executable", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccClaudeCommandResourceConfig(workDir, scope, name, content, executable string) string {
	return fmt.Sprintf(`
provider "agentsmith" {
  workdir = "%s"
}

resource "agentsmith_claude_command" "test" {
  scope      = "%s"
  name       = "%s"
  content    = "%s"
  executable = %s
}
`, workDir, scope, name, content, executable)
}
