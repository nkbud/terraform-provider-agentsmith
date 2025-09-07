package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/diag"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestGeminiDataSource_getGeminiSettingsFilePaths(t *testing.T) {
	tempDir := t.TempDir()
	client := &FileClient{workDir: tempDir}

	d := &geminiDataSource{client: client}
	var diagnostics diag.Diagnostics

	paths := d.getGeminiSettingsFilePaths(&diagnostics)

	if diagnostics.HasError() {
		t.Errorf("Unexpected diagnostics errors: %v", diagnostics.Errors())
	}

	// Should include at least user settings and project settings
	expectedMinPaths := 2
	if len(paths) < expectedMinPaths {
		t.Errorf("Expected at least %d paths, got %d: %v", expectedMinPaths, len(paths), paths)
	}

	// Test that paths are constructed correctly based on OS
	homeDir, err := os.UserHomeDir()
	if err != nil {
		t.Skip("Cannot get user home directory, skipping path validation")
		return
	}

	expectedUserPath := filepath.Join(homeDir, ".gemini", "settings.json")
	expectedProjectPath := filepath.Join(tempDir, ".gemini", "settings.json")

	foundUser := false
	foundProject := false
	for _, path := range paths {
		if path == expectedUserPath {
			foundUser = true
		}
		if path == expectedProjectPath {
			foundProject = true
		}
	}

	if !foundUser {
		t.Errorf("Expected user settings path %q not found in paths: %v", expectedUserPath, paths)
	}
	if !foundProject {
		t.Errorf("Expected project settings path %q not found in paths: %v", expectedProjectPath, paths)
	}
}

func TestGeminiDataSource_getGeminiSettingsFilePaths_CrossPlatform(t *testing.T) {
	tempDir := t.TempDir()
	client := &FileClient{workDir: tempDir}

	d := &geminiDataSource{client: client}
	var diagnostics diag.Diagnostics

	paths := d.getGeminiSettingsFilePaths(&diagnostics)

	if diagnostics.HasError() {
		t.Errorf("Unexpected diagnostics errors: %v", diagnostics.Errors())
	}

	// Check that system paths are OS-appropriate
	var expectedSystemDefaults, expectedSystemSettings string
	switch runtime.GOOS {
	case "darwin":
		expectedSystemDefaults = "/Library/Application Support/GeminiCli/system-defaults.json"
		expectedSystemSettings = "/Library/Application Support/GeminiCli/settings.json"
	case "linux":
		expectedSystemDefaults = "/etc/gemini-cli/system-defaults.json"
		expectedSystemSettings = "/etc/gemini-cli/settings.json"
	case "windows":
		expectedSystemDefaults = "C:\\ProgramData\\gemini-cli\\system-defaults.json"
		expectedSystemSettings = "C:\\ProgramData\\gemini-cli\\settings.json"
	default:
		t.Skipf("Unsupported OS: %s", runtime.GOOS)
	}

	foundSystemDefaults := false
	foundSystemSettings := false
	for _, path := range paths {
		if path == expectedSystemDefaults {
			foundSystemDefaults = true
		}
		if path == expectedSystemSettings {
			foundSystemSettings = true
		}
	}

	if !foundSystemDefaults {
		t.Errorf("Expected system defaults path %q not found in paths: %v", expectedSystemDefaults, paths)
	}
	if !foundSystemSettings {
		t.Errorf("Expected system settings path %q not found in paths: %v", expectedSystemSettings, paths)
	}
}

func TestGeminiDataSource_readGeminiSettingsFile(t *testing.T) {
	tempDir := t.TempDir()
	d := &geminiDataSource{}

	testCases := []struct {
		name        string
		content     map[string]interface{}
		expectError bool
	}{
		{
			name: "valid_settings",
			content: map[string]interface{}{
				"general": map[string]interface{}{
					"vimMode":           true,
					"preferredEditor":   "vim",
					"disableAutoUpdate": false,
				},
				"ui": map[string]interface{}{
					"theme":           "dark",
					"hideWindowTitle": false,
					"showMemoryUsage": true,
				},
				"model": map[string]interface{}{
					"name":            "gemini-1.5-pro",
					"maxSessionTurns": 100,
				},
				"advanced": map[string]interface{}{
					"excludedEnvVars":     []string{"DEBUG", "DEBUG_MODE"},
					"autoConfigureMemory": true,
				},
			},
			expectError: false,
		},
		{
			name: "minimal_settings",
			content: map[string]interface{}{
				"model": map[string]interface{}{
					"name": "gemini-1.5-pro",
				},
			},
			expectError: false,
		},
		{
			name:        "empty_settings",
			content:     map[string]interface{}{},
			expectError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create test file
			testFile := filepath.Join(tempDir, fmt.Sprintf("%s.json", tc.name))

			jsonData, err := json.MarshalIndent(tc.content, "", "  ")
			if err != nil {
				t.Fatalf("Failed to marshal test content: %v", err)
			}

			err = os.WriteFile(testFile, jsonData, 0644)
			if err != nil {
				t.Fatalf("Failed to write test file: %v", err)
			}

			// Test reading
			result, err := d.readGeminiSettingsFile(testFile)

			if tc.expectError {
				if err == nil {
					t.Error("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}

				// Validate key fields
				for key, expectedValue := range tc.content {
					actualValue, exists := result[key]
					if !exists {
						t.Errorf("Expected key %q not found in result", key)
						continue
					}

					// For nested maps, do a basic existence check
					if expectedMap, ok := expectedValue.(map[string]interface{}); ok {
						actualMap, ok := actualValue.(map[string]interface{})
						if !ok {
							t.Errorf("Expected nested map for key %q", key)
							continue
						}
						if len(actualMap) != len(expectedMap) {
							t.Errorf("Map size mismatch for key %q: expected %d, got %d", key, len(expectedMap), len(actualMap))
						}
					}
				}
			}

			// Clean up
			os.Remove(testFile)
		})
	}
}

func TestGeminiDataSource_readGeminiSettingsFile_EdgeCases(t *testing.T) {
	tempDir := t.TempDir()
	d := &geminiDataSource{}

	testCases := []struct {
		name        string
		fileContent string
		expectError bool
	}{
		{
			name:        "empty_file",
			fileContent: "",
			expectError: false,
		},
		{
			name:        "invalid_json",
			fileContent: `{"invalid": json syntax}`,
			expectError: true,
		},
		{
			name:        "null_json",
			fileContent: "null",
			expectError: false,
		},
		{
			name:        "whitespace_only",
			fileContent: "   \n  \t  ",
			expectError: false,
		},
		{
			name:        "json_array",
			fileContent: `["not", "an", "object"]`,
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testFile := filepath.Join(tempDir, fmt.Sprintf("%s.json", tc.name))

			err := os.WriteFile(testFile, []byte(tc.fileContent), 0644)
			if err != nil {
				t.Fatalf("Failed to write test file: %v", err)
			}

			result, err := d.readGeminiSettingsFile(testFile)

			if tc.expectError {
				if err == nil {
					t.Error("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if result == nil {
					t.Error("Expected non-nil result")
				}
			}

			os.Remove(testFile)
		})
	}
}

func TestGeminiDataSource_getMergedGeminiSettings(t *testing.T) {
	tempDir := t.TempDir()

	// Create .gemini directory
	geminiDir := filepath.Join(tempDir, ".gemini")
	err := os.MkdirAll(geminiDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create gemini directory: %v", err)
	}

	// Create test settings file
	projectSettings := map[string]interface{}{
		"general": map[string]interface{}{
			"vimMode":         true,
			"preferredEditor": "nvim",
		},
		"model": map[string]interface{}{
			"name": "gemini-1.5-pro",
		},
		"advanced": map[string]interface{}{
			"excludedEnvVars": []string{"SECRET_KEY", "API_TOKEN"},
		},
	}

	projectFile := filepath.Join(geminiDir, "settings.json")
	jsonData, err := json.MarshalIndent(projectSettings, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal project settings: %v", err)
	}

	err = os.WriteFile(projectFile, jsonData, 0644)
	if err != nil {
		t.Fatalf("Failed to write project settings: %v", err)
	}

	client := &FileClient{workDir: tempDir}
	d := &geminiDataSource{client: client}

	ctx := context.Background()
	var diagnostics diag.Diagnostics

	result := d.getMergedGeminiSettings(ctx, &diagnostics)

	// Should have warnings about missing system files, but no errors
	if diagnostics.HasError() {
		t.Errorf("Unexpected errors: %v", diagnostics.Errors())
	}

	// Should contain merged settings from the project file
	general, ok := result["general"].(map[string]interface{})
	if !ok {
		t.Error("Expected general settings to be present")
	} else {
		if general["vimMode"] != true {
			t.Errorf("Expected vimMode true, got %v", general["vimMode"])
		}
		if general["preferredEditor"] != "nvim" {
			t.Errorf("Expected preferredEditor 'nvim', got %v", general["preferredEditor"])
		}
	}

	model, ok := result["model"].(map[string]interface{})
	if !ok {
		t.Error("Expected model settings to be present")
	} else {
		if model["name"] != "gemini-1.5-pro" {
			t.Errorf("Expected model name 'gemini-1.5-pro', got %v", model["name"])
		}
	}

	// Clean up
	os.RemoveAll(geminiDir)
}

func TestGeminiDataSource_populateGeminiSettingsFromMap(t *testing.T) {
	d := &geminiDataSource{}
	ctx := context.Background()
	var diagnostics diag.Diagnostics

	testData := map[string]interface{}{
		"general": map[string]interface{}{
			"vimMode":           true,
			"preferredEditor":   "vim",
			"disableAutoUpdate": false,
			"checkpointing": map[string]interface{}{
				"enabled": true,
			},
		},
		"ui": map[string]interface{}{
			"theme":           "dark",
			"hideWindowTitle": false,
			"showMemoryUsage": true,
			"customThemes": map[string]interface{}{
				"myTheme": "#ff0000",
			},
			"accessibility": map[string]interface{}{
				"disableLoadingPhrases": true,
			},
		},
		"model": map[string]interface{}{
			"name":            "gemini-1.5-pro",
			"maxSessionTurns": float64(100), // JSON numbers are float64
			"summarizeToolOutput": map[string]interface{}{
				"enabled": "true",
			},
			"chatCompression": map[string]interface{}{
				"contextPercentageThreshold": "80",
			},
		},
		"context": map[string]interface{}{
			"fileName":           []interface{}{"GEMINI.md", "context.md"},
			"includeDirectories": []interface{}{"/path/to/include"},
			"fileFiltering": map[string]interface{}{
				"respectGitIgnore":    true,
				"respectGeminiIgnore": false,
			},
		},
		"advanced": map[string]interface{}{
			"excludedEnvVars":     []interface{}{"DEBUG", "SECRET"},
			"autoConfigureMemory": true,
		},
		"tools": map[string]interface{}{
			"sandbox": "docker",
			"core":    []interface{}{"bash", "python"},
			"exclude": []interface{}{"dangerous-tool"},
		},
		"security": map[string]interface{}{
			"folderTrust": map[string]interface{}{
				"enabled": true,
			},
			"auth": map[string]interface{}{
				"selectedType": "oauth",
				"enforcedType": "oauth",
				"useExternal":  false,
			},
		},
	}

	settingsModel := &geminiSettingsModel{}
	d.populateGeminiSettingsFromMap(ctx, &diagnostics, settingsModel, testData)

	if diagnostics.HasError() {
		t.Errorf("Unexpected diagnostics errors: %v", diagnostics.Errors())
	}

	// Validate general settings
	if settingsModel.General == nil {
		t.Error("Expected general settings to be populated")
	} else {
		if !settingsModel.General.VimMode.ValueBool() {
			t.Error("Expected vimMode to be true")
		}
		if settingsModel.General.PreferredEditor.ValueString() != "vim" {
			t.Errorf("Expected preferredEditor 'vim', got %q", settingsModel.General.PreferredEditor.ValueString())
		}
		if !settingsModel.General.CheckpointingEnabled.ValueBool() {
			t.Error("Expected checkpointing to be enabled")
		}
	}

	// Validate UI settings
	if settingsModel.UI == nil {
		t.Error("Expected UI settings to be populated")
	} else {
		if settingsModel.UI.Theme.ValueString() != "dark" {
			t.Errorf("Expected theme 'dark', got %q", settingsModel.UI.Theme.ValueString())
		}
		if settingsModel.UI.ShowMemoryUsage.IsNull() || !settingsModel.UI.ShowMemoryUsage.ValueBool() {
			t.Error("Expected showMemoryUsage to be true")
		}
		if settingsModel.UI.AccessibilityDisableLoadingPhrases.IsNull() || !settingsModel.UI.AccessibilityDisableLoadingPhrases.ValueBool() {
			t.Error("Expected accessibility.disableLoadingPhrases to be true")
		}
	}

	// Validate model settings
	if settingsModel.Model == nil {
		t.Error("Expected model settings to be populated")
	} else {
		if settingsModel.Model.Name.ValueString() != "gemini-1.5-pro" {
			t.Errorf("Expected model name 'gemini-1.5-pro', got %q", settingsModel.Model.Name.ValueString())
		}
		if settingsModel.Model.MaxSessionTurns.ValueInt64() != 100 {
			t.Errorf("Expected maxSessionTurns 100, got %d", settingsModel.Model.MaxSessionTurns.ValueInt64())
		}
	}

	// Validate context settings with list fields
	if settingsModel.Context == nil {
		t.Error("Expected context settings to be populated")
	} else {
		if settingsModel.Context.FileName.IsNull() {
			t.Error("Expected fileName list to be populated")
		}
		if settingsModel.Context.IncludeDirectories.IsNull() {
			t.Error("Expected includeDirectories list to be populated")
		}
	}

	// Validate advanced settings
	if settingsModel.Advanced == nil {
		t.Error("Expected advanced settings to be populated")
	} else {
		if settingsModel.Advanced.ExcludedEnvVars.IsNull() {
			t.Error("Expected excludedEnvVars list to be populated")
		}
		if !settingsModel.Advanced.AutoConfigureMemory.ValueBool() {
			t.Error("Expected autoConfigureMemory to be true")
		}
	}

	// Validate tools settings
	if settingsModel.Tools == nil {
		t.Error("Expected tools settings to be populated")
	} else {
		if settingsModel.Tools.Sandbox.ValueString() != "docker" {
			t.Errorf("Expected sandbox 'docker', got %q", settingsModel.Tools.Sandbox.ValueString())
		}
		if settingsModel.Tools.Core.IsNull() {
			t.Error("Expected core tools list to be populated")
		}
	}

	// Validate security settings
	if settingsModel.Security == nil {
		t.Error("Expected security settings to be populated")
	} else {
		if !settingsModel.Security.FolderTrustEnabled.ValueBool() {
			t.Error("Expected folderTrust to be enabled")
		}
		if settingsModel.Security.AuthSelectedType.ValueString() != "oauth" {
			t.Errorf("Expected authSelectedType 'oauth', got %q", settingsModel.Security.AuthSelectedType.ValueString())
		}
	}
}

func TestGeminiDataSource_readGeminiEnvironmentVariables(t *testing.T) {
	// Set test environment variables
	testEnvVars := map[string]string{
		"GEMINI_API_KEY":                 "test-api-key",
		"GEMINI_MODEL":                   "gemini-1.5-pro",
		"GOOGLE_CLOUD_PROJECT":           "my-project",
		"GOOGLE_APPLICATION_CREDENTIALS": "/path/to/creds.json",
		"DEBUG":                          "1",
		"DEBUG_MODE":                     "true",
		"NO_COLOR":                       "1",
	}

	// Set environment variables
	for key, value := range testEnvVars {
		err := os.Setenv(key, value)
		if err != nil {
			t.Fatalf("Failed to set environment variable %s: %v", key, err)
		}
	}

	// Clean up after test
	defer func() {
		for key := range testEnvVars {
			os.Unsetenv(key)
		}
	}()

	envModel := &geminiEnvironmentVariablesModel{}
	readGeminiEnvironmentVariables(envModel)

	// Validate environment variables were read correctly
	if envModel.GeminiAPIKey.ValueString() != "test-api-key" {
		t.Errorf("Expected GeminiAPIKey 'test-api-key', got %q", envModel.GeminiAPIKey.ValueString())
	}
	if envModel.GeminiModel.ValueString() != "gemini-1.5-pro" {
		t.Errorf("Expected GeminiModel 'gemini-1.5-pro', got %q", envModel.GeminiModel.ValueString())
	}
	if envModel.GoogleCloudProject.ValueString() != "my-project" {
		t.Errorf("Expected GoogleCloudProject 'my-project', got %q", envModel.GoogleCloudProject.ValueString())
	}
	if envModel.GoogleApplicationCredentials.ValueString() != "/path/to/creds.json" {
		t.Errorf("Expected GoogleApplicationCredentials '/path/to/creds.json', got %q", envModel.GoogleApplicationCredentials.ValueString())
	}
	if !envModel.Debug.ValueBool() {
		t.Error("Expected Debug to be true")
	}
	if !envModel.DebugMode.ValueBool() {
		t.Error("Expected DebugMode to be true")
	}
	if envModel.NoColor.ValueString() != "1" {
		t.Errorf("Expected NoColor '1', got %q", envModel.NoColor.ValueString())
	}
}

func TestGeminiDataSource_listAndMapInitialization(t *testing.T) {
	d := &geminiDataSource{}
	ctx := context.Background()
	var diagnostics diag.Diagnostics

	// Test with empty data to ensure proper initialization
	emptyData := map[string]interface{}{}
	settingsModel := &geminiSettingsModel{}

	d.populateGeminiSettingsFromMap(ctx, &diagnostics, settingsModel, emptyData)

	if diagnostics.HasError() {
		t.Errorf("Unexpected diagnostics errors: %v", diagnostics.Errors())
	}

	// All nested models should be initialized
	if settingsModel.General == nil {
		t.Error("Expected general settings to be initialized")
	}
	if settingsModel.UI == nil {
		t.Error("Expected UI settings to be initialized")
	}
	if settingsModel.Model == nil {
		t.Error("Expected model settings to be initialized")
	}
	if settingsModel.Context == nil {
		t.Error("Expected context settings to be initialized")
	}
	if settingsModel.Tools == nil {
		t.Error("Expected tools settings to be initialized")
	}
	if settingsModel.Advanced == nil {
		t.Error("Expected advanced settings to be initialized")
	}

	// List and map fields should be properly typed null values
	if settingsModel.Context.FileName.IsUnknown() {
		t.Error("Expected FileName to be null, not unknown")
	}
	if settingsModel.Advanced.ExcludedEnvVars.IsUnknown() {
		t.Error("Expected ExcludedEnvVars to be null, not unknown")
	}
	if settingsModel.UI.CustomThemes.IsUnknown() {
		t.Error("Expected CustomThemes to be null, not unknown")
	}
}

func TestGeminiDataSource_ErrorHandling_MissingFiles(t *testing.T) {
	tempDir := t.TempDir()

	// Override HOME environment variable to use temp directory
	// This prevents the test from finding the real user's ~/.gemini/settings.json
	originalHome := os.Getenv("HOME")
	defer os.Setenv("HOME", originalHome)
	os.Setenv("HOME", tempDir)

	client := &FileClient{workDir: tempDir}
	d := &geminiDataSource{client: client}

	ctx := context.Background()
	var diagnostics diag.Diagnostics

	// Test with no settings files present
	result := d.getMergedGeminiSettings(ctx, &diagnostics)

	// Should have warnings but no errors
	if diagnostics.HasError() {
		t.Errorf("Unexpected errors: %v", diagnostics.Errors())
	}

	warnings := diagnostics.Warnings()
	if len(warnings) == 0 {
		t.Error("Expected warnings about missing settings files")
	}

	// Should return empty result
	if len(result) != 0 {
		t.Errorf("Expected empty result with no config files, got %v", result)
	}
}

// Helper functions for acceptance tests
func testAccGeminiWorkDir() string {
	tmpDir, _ := os.MkdirTemp("", "terraform-test-gemini")
	return tmpDir
}

// Acceptance tests
func TestAccGeminiDataSource_basic(t *testing.T) {
	workDir := testAccGeminiWorkDir()
	defer os.RemoveAll(workDir)

	// Create a test settings file
	geminiDir := filepath.Join(workDir, ".gemini")
	err := os.MkdirAll(geminiDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create gemini directory: %v", err)
	}

	testSettings := map[string]interface{}{
		"general": map[string]interface{}{
			"vimMode":         true,
			"preferredEditor": "vim",
		},
		"model": map[string]interface{}{
			"name": "gemini-1.5-pro",
		},
	}

	settingsFile := filepath.Join(geminiDir, "settings.json")
	jsonData, err := json.MarshalIndent(testSettings, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal test settings: %v", err)
	}

	err = os.WriteFile(settingsFile, jsonData, 0644)
	if err != nil {
		t.Fatalf("Failed to write test settings: %v", err)
	}

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccGeminiDataSourceConfig(workDir),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.agentsmith_gemini.test", "id"),
					resource.TestCheckResourceAttr("data.agentsmith_gemini.test", "settings.general.vim_mode", "true"),
					resource.TestCheckResourceAttr("data.agentsmith_gemini.test", "settings.general.preferred_editor", "vim"),
					resource.TestCheckResourceAttr("data.agentsmith_gemini.test", "settings.model.name", "gemini-1.5-pro"),
				),
			},
		},
	})
}

func testAccGeminiDataSourceConfig(workDir string) string {
	return fmt.Sprintf(`
provider "agentsmith" {
  workdir = "%s"
}

data "agentsmith_gemini" "test" {}
`, workDir)
}
