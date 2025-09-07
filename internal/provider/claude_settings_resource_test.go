package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestClaudeSettingsResource_getSettingsFilePath(t *testing.T) {
	tempDir := t.TempDir()
	client := &FileClient{workDir: tempDir}

	r := &claudeSettingsResource{client: client}

	testCases := []struct {
		name     string
		scope    string
		expected string
	}{
		{
			name:     "project_scope",
			scope:    "project",
			expected: filepath.Join(tempDir, ".claude", "settings.json"),
		},
		{
			name:     "local_scope",
			scope:    "local",
			expected: filepath.Join(tempDir, ".claude", "settings.local.json"),
		},
		{
			name:     "invalid_scope",
			scope:    "invalid",
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := r.getSettingsFilePath(tc.scope)
			if result != tc.expected {
				t.Errorf("Expected path %q, got %q", tc.expected, result)
			}
		})
	}

	// Test user scope (requires actual home directory)
	homeDir, err := os.UserHomeDir()
	if err == nil {
		expectedUserPath := filepath.Join(homeDir, ".claude", "settings.json")
		result := r.getSettingsFilePath("user")
		if result != expectedUserPath {
			t.Errorf("Expected user path %q, got %q", expectedUserPath, result)
		}
	}
}

func TestClaudeSettingsResource_modelToSettings(t *testing.T) {
	r := &claudeSettingsResource{}
	ctx := context.Background()

	testCases := []struct {
		name     string
		model    claudeSettingsResourceModel
		validate func(*testing.T, map[string]interface{})
	}{
		{
			name: "basic_settings",
			model: claudeSettingsResourceModel{
				APIKeyHelper:      types.StringValue("/path/to/key-helper.sh"),
				CleanupPeriodDays: types.Int64Value(30),
				Model:             types.StringValue("claude-3-5-sonnet"),
				OutputStyle:       types.StringValue("concise"),
			},
			validate: func(t *testing.T, settings map[string]interface{}) {
				if settings["apiKeyHelper"] != "/path/to/key-helper.sh" {
					t.Errorf("Expected apiKeyHelper '/path/to/key-helper.sh', got %v", settings["apiKeyHelper"])
				}
				if settings["cleanupPeriodDays"] != int64(30) {
					t.Errorf("Expected cleanupPeriodDays 30, got %v", settings["cleanupPeriodDays"])
				}
				if settings["model"] != "claude-3-5-sonnet" {
					t.Errorf("Expected model 'claude-3-5-sonnet', got %v", settings["model"])
				}
				if settings["outputStyle"] != "concise" {
					t.Errorf("Expected outputStyle 'concise', got %v", settings["outputStyle"])
				}
			},
		},
		{
			name: "settings_with_permissions",
			model: claudeSettingsResourceModel{
				Model: types.StringValue("claude-3-opus"),
				Permissions: &claudeSettingsPermissionsModel{
					DefaultMode: types.StringValue("ask"),
					Allow: func() types.List {
						list, _ := types.ListValueFrom(ctx, types.StringType, []string{"read", "write"})
						return list
					}(),
					Ask: func() types.List {
						list, _ := types.ListValueFrom(ctx, types.StringType, []string{"exec", "network"})
						return list
					}(),
				},
			},
			validate: func(t *testing.T, settings map[string]interface{}) {
				perms, exists := settings["permissions"].(map[string]interface{})
				if !exists {
					t.Error("Expected permissions object")
					return
				}

				if perms["defaultMode"] != "ask" {
					t.Errorf("Expected defaultMode 'ask', got %v", perms["defaultMode"])
				}

				allowList, ok := perms["allow"].([]string)
				if !ok {
					t.Error("Expected allow to be []string")
				} else if len(allowList) != 2 || allowList[0] != "read" || allowList[1] != "write" {
					t.Errorf("Expected allow ['read', 'write'], got %v", allowList)
				}

				askList, ok := perms["ask"].([]string)
				if !ok {
					t.Error("Expected ask to be []string")
				} else if len(askList) != 2 || askList[0] != "exec" || askList[1] != "network" {
					t.Errorf("Expected ask ['exec', 'network'], got %v", askList)
				}
			},
		},
		{
			name: "settings_with_status_line",
			model: claudeSettingsResourceModel{
				Model: types.StringValue("claude-3-5-sonnet"),
				StatusLine: &claudeSettingsStatusLineModel{
					Type:    types.StringValue("command"),
					Command: types.StringValue("git branch --show-current"),
				},
			},
			validate: func(t *testing.T, settings map[string]interface{}) {
				statusLine, exists := settings["statusLine"].(map[string]interface{})
				if !exists {
					t.Error("Expected statusLine object")
					return
				}

				if statusLine["type"] != "command" {
					t.Errorf("Expected statusLine type 'command', got %v", statusLine["type"])
				}
				if statusLine["command"] != "git branch --show-current" {
					t.Errorf("Expected statusLine command 'git branch --show-current', got %v", statusLine["command"])
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var diags diag.Diagnostics
			result, err := r.modelToSettings(ctx, &diags, &tc.model)

			if err != nil {
				t.Errorf("Failed to convert model to settings: %v", err)
				return
			}

			if diags.HasError() {
				t.Errorf("Diagnostics has errors: %v", diags.Errors())
				return
			}

			tc.validate(t, result)
		})
	}
}

func TestClaudeSettingsResource_JSONOperations(t *testing.T) {
	tempDir := t.TempDir()
	settingsPath := filepath.Join(tempDir, "settings.json")

	testSettings := map[string]interface{}{
		"model":             "claude-3-5-sonnet",
		"outputStyle":       "verbose",
		"cleanupPeriodDays": float64(14), // JSON numbers are float64
		"env": map[string]interface{}{
			"NODE_ENV": "development",
			"DEBUG":    "true",
		},
		"permissions": map[string]interface{}{
			"defaultMode": "ask",
			"allow":       []interface{}{"read", "write"},
			"deny":        []interface{}{"network", "exec"},
		},
	}

	// Test JSON marshaling and file creation
	jsonData, err := json.MarshalIndent(testSettings, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal test settings: %v", err)
	}

	err = os.WriteFile(settingsPath, jsonData, 0644)
	if err != nil {
		t.Fatalf("Failed to write settings file: %v", err)
	}

	// Test reading and unmarshaling
	data, err := os.ReadFile(settingsPath)
	if err != nil {
		t.Fatalf("Failed to read settings file: %v", err)
	}

	var parsed map[string]interface{}
	err = json.Unmarshal(data, &parsed)
	if err != nil {
		t.Fatalf("Failed to unmarshal settings: %v", err)
	}

	// Validate key fields
	if parsed["model"] != "claude-3-5-sonnet" {
		t.Errorf("Expected model 'claude-3-5-sonnet', got %v", parsed["model"])
	}

	if parsed["outputStyle"] != "verbose" {
		t.Errorf("Expected outputStyle 'verbose', got %v", parsed["outputStyle"])
	}

	// Test nested objects
	env, ok := parsed["env"].(map[string]interface{})
	if !ok {
		t.Error("Expected env to be a map")
	} else {
		if env["NODE_ENV"] != "development" {
			t.Errorf("Expected NODE_ENV 'development', got %v", env["NODE_ENV"])
		}
	}

	permissions, ok := parsed["permissions"].(map[string]interface{})
	if !ok {
		t.Error("Expected permissions to be a map")
	} else {
		if permissions["defaultMode"] != "ask" {
			t.Errorf("Expected defaultMode 'ask', got %v", permissions["defaultMode"])
		}
	}

	// Test file deletion
	err = os.Remove(settingsPath)
	if err != nil {
		t.Errorf("Failed to delete settings file: %v", err)
	}

	if _, err := os.Stat(settingsPath); !os.IsNotExist(err) {
		t.Errorf("Expected settings file to be deleted, but it still exists")
	}
}

func TestClaudeSettingsResource_ImportParsing(t *testing.T) {
	testCases := []struct {
		name          string
		importID      string
		expectError   bool
		expectedScope string
	}{
		{
			name:          "valid_user_scope",
			importID:      "user",
			expectError:   false,
			expectedScope: "user",
		},
		{
			name:          "valid_project_scope",
			importID:      "project",
			expectError:   false,
			expectedScope: "project",
		},
		{
			name:          "valid_local_scope",
			importID:      "local",
			expectError:   false,
			expectedScope: "local",
		},
		{
			name:        "invalid_scope",
			importID:    "global",
			expectError: true,
		},
		{
			name:        "empty_scope",
			importID:    "",
			expectError: true,
		},
		{
			name:        "whitespace_only",
			importID:    "  ",
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			scope := strings.TrimSpace(tc.importID)

			validScopes := []string{"user", "project", "local"}
			isValid := false
			for _, validScope := range validScopes {
				if scope == validScope {
					isValid = true
					break
				}
			}

			if tc.expectError {
				if isValid && scope != "" {
					t.Errorf("Expected error but got valid scope: %s", scope)
				}
			} else {
				if !isValid {
					t.Errorf("Expected valid scope but got invalid: %s", scope)
				}
				if scope != tc.expectedScope {
					t.Errorf("Expected scope %q, got %q", tc.expectedScope, scope)
				}
			}
		})
	}
}

func TestClaudeSettingsResource_DirectoryCreation(t *testing.T) {
	tempDir := t.TempDir()

	testPaths := []string{
		filepath.Join(tempDir, ".claude"),
		filepath.Join(tempDir, "nested", "deeper", ".claude"),
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

func TestClaudeSettingsResource_ComplexSettings(t *testing.T) {
	testCases := []struct {
		name     string
		settings map[string]interface{}
		valid    bool
	}{
		{
			name: "comprehensive_settings",
			settings: map[string]interface{}{
				"model":                      "claude-3-5-sonnet",
				"outputStyle":                "concise",
				"cleanupPeriodDays":          float64(30),
				"includeCoAuthoredBy":        true,
				"forceLoginMethod":           "claudeai",
				"disableAllHooks":            false,
				"enableAllProjectMcpServers": true,
				"enabledMcpjsonServers":      []interface{}{"server1", "server2"},
				"disabledMcpjsonServers":     []interface{}{"server3"},
				"env": map[string]interface{}{
					"NODE_ENV":     "production",
					"API_TIMEOUT":  "30000",
					"ENABLE_DEBUG": "false",
				},
				"permissions": map[string]interface{}{
					"defaultMode":                  "ask",
					"disableBypassPermissionsMode": "disable",
					"allow":                        []interface{}{"read", "write"},
					"ask":                          []interface{}{"exec", "network"},
					"deny":                         []interface{}{"system"},
					"additionalDirectories":        []interface{}{"/home/user/projects"},
				},
				"statusLine": map[string]interface{}{
					"type":    "command",
					"command": "pwd && git branch --show-current",
				},
				"hooks": map[string]interface{}{
					"pre-commit": map[string]interface{}{
						"command": "./scripts/pre-commit.sh",
						"timeout": "30",
					},
				},
			},
			valid: true,
		},
		{
			name: "minimal_settings",
			settings: map[string]interface{}{
				"model": "claude-3-5-sonnet",
			},
			valid: true,
		},
		{
			name:     "empty_settings",
			settings: map[string]interface{}{},
			valid:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test JSON marshaling
			jsonData, err := json.MarshalIndent(tc.settings, "", "  ")
			if err != nil {
				if tc.valid {
					t.Errorf("Failed to marshal valid settings: %v", err)
				}
				return
			}

			// Test JSON unmarshaling
			var parsed map[string]interface{}
			err = json.Unmarshal(jsonData, &parsed)
			if err != nil {
				if tc.valid {
					t.Errorf("Failed to unmarshal settings: %v", err)
				}
				return
			}

			if !tc.valid {
				t.Error("Expected invalid settings but got valid")
			}

			// Validate specific fields for comprehensive settings
			if tc.name == "comprehensive_settings" {
				if parsed["model"] != "claude-3-5-sonnet" {
					t.Errorf("Expected model 'claude-3-5-sonnet', got %v", parsed["model"])
				}

				permissions, ok := parsed["permissions"].(map[string]interface{})
				if !ok {
					t.Error("Expected permissions to be a map")
				} else {
					allow, ok := permissions["allow"].([]interface{})
					if !ok || len(allow) != 2 {
						t.Errorf("Expected allow to have 2 items, got %v", allow)
					}
				}

				env, ok := parsed["env"].(map[string]interface{})
				if !ok {
					t.Error("Expected env to be a map")
				} else if env["NODE_ENV"] != "production" {
					t.Errorf("Expected NODE_ENV 'production', got %v", env["NODE_ENV"])
				}
			}
		})
	}
}

// Helper functions for acceptance tests
func testAccSettingsWorkDir() string {
	tmpDir, _ := os.MkdirTemp("", "terraform-test-settings")
	return tmpDir
}

// Acceptance tests
func TestAccClaudeSettingsResource_basic(t *testing.T) {
	workDir := testAccSettingsWorkDir()
	defer os.RemoveAll(workDir)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccClaudeSettingsResourceConfig(workDir, "project", "claude-3-5-sonnet", "concise"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("agentsmith_claude_settings.test", "scope", "project"),
					resource.TestCheckResourceAttr("agentsmith_claude_settings.test", "model", "claude-3-5-sonnet"),
					resource.TestCheckResourceAttr("agentsmith_claude_settings.test", "output_style", "concise"),
					resource.TestCheckResourceAttrSet("agentsmith_claude_settings.test", "id"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "agentsmith_claude_settings.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     "project",
			},
			// Update and Read testing
			{
				Config: testAccClaudeSettingsResourceConfig(workDir, "project", "claude-3-opus", "verbose"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("agentsmith_claude_settings.test", "model", "claude-3-opus"),
					resource.TestCheckResourceAttr("agentsmith_claude_settings.test", "output_style", "verbose"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccClaudeSettingsResourceConfig(workDir, scope, model, outputStyle string) string {
	return fmt.Sprintf(`
provider "agentsmith" {
  workdir = "%s"
}

resource "agentsmith_claude_settings" "test" {
  scope        = "%s"
  model        = "%s"
  output_style = "%s"
}
`, workDir, scope, model, outputStyle)
}
