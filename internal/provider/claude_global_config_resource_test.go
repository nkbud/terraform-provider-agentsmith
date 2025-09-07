package provider

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestClaudeGlobalConfigResource_getGlobalConfigFilePath(t *testing.T) {
	r := &claudeGlobalConfigResource{}

	result := r.getGlobalConfigFilePath()

	// Should return a valid path to ~/.claude/global-config.json
	homeDir, err := os.UserHomeDir()
	if err != nil {
		t.Skip("Cannot get user home directory, skipping test")
		return
	}

	expected := filepath.Join(homeDir, ".claude", "global-config.json")
	if result != expected {
		t.Errorf("Expected path %q, got %q", expected, result)
	}
}

func TestClaudeGlobalConfigResource_modelToConfig(t *testing.T) {
	r := &claudeGlobalConfigResource{}

	testCases := []struct {
		name     string
		model    claudeGlobalConfigResourceModel
		expected map[string]interface{}
	}{
		{
			name: "all_fields_set",
			model: claudeGlobalConfigResourceModel{
				AutoUpdates:           types.BoolValue(false),
				PreferredNotifChannel: types.StringValue("iterm2"),
				Theme:                 types.StringValue("dark"),
				Verbose:               types.BoolValue(true),
			},
			expected: map[string]interface{}{
				"autoUpdates":           false,
				"preferredNotifChannel": "iterm2",
				"theme":                 "dark",
				"verbose":               true,
			},
		},
		{
			name: "partial_fields_set",
			model: claudeGlobalConfigResourceModel{
				Theme:   types.StringValue("light"),
				Verbose: types.BoolValue(false),
			},
			expected: map[string]interface{}{
				"theme":   "light",
				"verbose": false,
			},
		},
		{
			name:     "no_fields_set",
			model:    claudeGlobalConfigResourceModel{},
			expected: map[string]interface{}{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := r.modelToConfig(&tc.model)

			// Compare results
			if len(result) != len(tc.expected) {
				t.Errorf("Expected %d fields, got %d", len(tc.expected), len(result))
			}

			for key, expectedValue := range tc.expected {
				actualValue, exists := result[key]
				if !exists {
					t.Errorf("Expected key %q not found in result", key)
					continue
				}
				if actualValue != expectedValue {
					t.Errorf("For key %q: expected %v, got %v", key, expectedValue, actualValue)
				}
			}
		})
	}
}

func TestClaudeGlobalConfigResource_configToModel(t *testing.T) {
	r := &claudeGlobalConfigResource{}

	testCases := []struct {
		name     string
		config   map[string]interface{}
		validate func(*testing.T, *claudeGlobalConfigResourceModel)
	}{
		{
			name: "all_fields_present",
			config: map[string]interface{}{
				"autoUpdates":           true,
				"preferredNotifChannel": "terminal_bell",
				"theme":                 "dark-daltonized",
				"verbose":               false,
			},
			validate: func(t *testing.T, model *claudeGlobalConfigResourceModel) {
				if !model.AutoUpdates.ValueBool() {
					t.Errorf("Expected autoUpdates to be true")
				}
				if model.PreferredNotifChannel.ValueString() != "terminal_bell" {
					t.Errorf("Expected preferredNotifChannel 'terminal_bell', got %q", model.PreferredNotifChannel.ValueString())
				}
				if model.Theme.ValueString() != "dark-daltonized" {
					t.Errorf("Expected theme 'dark-daltonized', got %q", model.Theme.ValueString())
				}
				if model.Verbose.ValueBool() {
					t.Errorf("Expected verbose to be false")
				}
			},
		},
		{
			name: "partial_fields_present",
			config: map[string]interface{}{
				"theme": "light",
			},
			validate: func(t *testing.T, model *claudeGlobalConfigResourceModel) {
				if model.Theme.ValueString() != "light" {
					t.Errorf("Expected theme 'light', got %q", model.Theme.ValueString())
				}
				if !model.AutoUpdates.IsNull() {
					t.Errorf("Expected autoUpdates to be null")
				}
				if !model.PreferredNotifChannel.IsNull() {
					t.Errorf("Expected preferredNotifChannel to be null")
				}
				if !model.Verbose.IsNull() {
					t.Errorf("Expected verbose to be null")
				}
			},
		},
		{
			name:   "empty_config",
			config: map[string]interface{}{},
			validate: func(t *testing.T, model *claudeGlobalConfigResourceModel) {
				if !model.AutoUpdates.IsNull() {
					t.Errorf("Expected autoUpdates to be null")
				}
				if !model.PreferredNotifChannel.IsNull() {
					t.Errorf("Expected preferredNotifChannel to be null")
				}
				if !model.Theme.IsNull() {
					t.Errorf("Expected theme to be null")
				}
				if !model.Verbose.IsNull() {
					t.Errorf("Expected verbose to be null")
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var model claudeGlobalConfigResourceModel
			r.configToModel(&model, tc.config)
			tc.validate(t, &model)
		})
	}
}

func TestClaudeGlobalConfigResource_JSONOperations(t *testing.T) {
	tempDir := t.TempDir()
	configPath := filepath.Join(tempDir, "global-config.json")

	testConfig := map[string]interface{}{
		"theme":                 "dark",
		"verbose":               true,
		"preferredNotifChannel": "iterm2_with_bell",
	}

	// Test JSON marshaling and file creation
	jsonData, err := json.MarshalIndent(testConfig, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal test config: %v", err)
	}

	err = os.WriteFile(configPath, jsonData, 0644)
	if err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	// Test reading and unmarshaling
	data, err := os.ReadFile(configPath)
	if err != nil {
		t.Fatalf("Failed to read config file: %v", err)
	}

	var parsed map[string]interface{}
	err = json.Unmarshal(data, &parsed)
	if err != nil {
		t.Fatalf("Failed to unmarshal config: %v", err)
	}

	// Verify parsed data matches original
	for key, expectedValue := range testConfig {
		actualValue, exists := parsed[key]
		if !exists {
			t.Errorf("Expected key %q not found in parsed data", key)
			continue
		}
		if actualValue != expectedValue {
			t.Errorf("For key %q: expected %v, got %v", key, expectedValue, actualValue)
		}
	}

	// Test file deletion
	err = os.Remove(configPath)
	if err != nil {
		t.Errorf("Failed to delete config file: %v", err)
	}

	if _, err := os.Stat(configPath); !os.IsNotExist(err) {
		t.Errorf("Expected config file to be deleted, but it still exists")
	}
}

func TestClaudeGlobalConfigResource_DirectoryCreation(t *testing.T) {
	tempDir := t.TempDir()
	configDir := filepath.Join(tempDir, ".claude")

	// Test directory creation
	err := os.MkdirAll(configDir, 0755)
	if err != nil {
		t.Errorf("Failed to create config directory: %v", err)
	}

	// Verify directory exists and has correct permissions
	fileInfo, err := os.Stat(configDir)
	if err != nil {
		t.Errorf("Config directory does not exist: %v", err)
	}

	if !fileInfo.IsDir() {
		t.Errorf("Expected directory, got file")
	}

	if fileInfo.Mode().Perm() != os.FileMode(0755) {
		t.Errorf("Expected directory mode 0755, got %v", fileInfo.Mode().Perm())
	}
}

func TestClaudeGlobalConfigResource_ConfigValidation(t *testing.T) {
	testCases := []struct {
		name             string
		theme            string
		notifChannel     string
		expectValidTheme bool
		expectValidNotif bool
	}{
		{
			name:             "valid_values",
			theme:            "dark",
			notifChannel:     "iterm2",
			expectValidTheme: true,
			expectValidNotif: true,
		},
		{
			name:             "valid_daltonized_theme",
			theme:            "light-daltonized",
			notifChannel:     "terminal_bell",
			expectValidTheme: true,
			expectValidNotif: true,
		},
		{
			name:             "invalid_theme",
			theme:            "rainbow",
			notifChannel:     "notifications_disabled",
			expectValidTheme: false,
			expectValidNotif: true,
		},
		{
			name:             "invalid_notif_channel",
			theme:            "light",
			notifChannel:     "carrier_pigeon",
			expectValidTheme: true,
			expectValidNotif: false,
		},
	}

	validThemes := []string{"dark", "light", "light-daltonized", "dark-daltonized"}
	validNotifChannels := []string{"iterm2", "iterm2_with_bell", "terminal_bell", "notifications_disabled"}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Check theme validation
			themeValid := false
			for _, validTheme := range validThemes {
				if tc.theme == validTheme {
					themeValid = true
					break
				}
			}

			if themeValid != tc.expectValidTheme {
				t.Errorf("Theme %q validation: expected %v, got %v", tc.theme, tc.expectValidTheme, themeValid)
			}

			// Check notification channel validation
			notifValid := false
			for _, validNotif := range validNotifChannels {
				if tc.notifChannel == validNotif {
					notifValid = true
					break
				}
			}

			if notifValid != tc.expectValidNotif {
				t.Errorf("Notification channel %q validation: expected %v, got %v", tc.notifChannel, tc.expectValidNotif, notifValid)
			}
		})
	}
}

// Helper functions for acceptance tests
func testAccGlobalConfigWorkDir() string {
	tmpDir, _ := os.MkdirTemp("", "terraform-test-global")
	return tmpDir
}

// Acceptance tests
func TestAccClaudeGlobalConfigResource_basic(t *testing.T) {
	workDir := testAccGlobalConfigWorkDir()
	defer os.RemoveAll(workDir)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccClaudeGlobalConfigResourceConfig(workDir, "dark", "iterm2", "true", "false"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("agentsmith_claude_global_config.test", "theme", "dark"),
					resource.TestCheckResourceAttr("agentsmith_claude_global_config.test", "preferred_notif_channel", "iterm2"),
					resource.TestCheckResourceAttr("agentsmith_claude_global_config.test", "verbose", "true"),
					resource.TestCheckResourceAttr("agentsmith_claude_global_config.test", "auto_updates", "false"),
					resource.TestCheckResourceAttrSet("agentsmith_claude_global_config.test", "id"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "agentsmith_claude_global_config.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: testAccClaudeGlobalConfigResourceConfig(workDir, "light", "terminal_bell", "false", "true"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("agentsmith_claude_global_config.test", "theme", "light"),
					resource.TestCheckResourceAttr("agentsmith_claude_global_config.test", "preferred_notif_channel", "terminal_bell"),
					resource.TestCheckResourceAttr("agentsmith_claude_global_config.test", "verbose", "false"),
					resource.TestCheckResourceAttr("agentsmith_claude_global_config.test", "auto_updates", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccClaudeGlobalConfigResourceConfig(workDir, theme, notifChannel, verbose, autoUpdates string) string {
	return fmt.Sprintf(`
provider "agentsmith" {
  workdir = "%s"
}

resource "agentsmith_claude_global_config" "test" {
  theme                     = "%s"
  preferred_notif_channel   = "%s"
  verbose                   = %s
  auto_updates              = %s
}
`, workDir, theme, notifChannel, verbose, autoUpdates)
}
