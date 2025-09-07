package provider

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	tf "github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccGeminiSettingsFileResource(t *testing.T) {
	tempDir := t.TempDir()
	projectDir := filepath.Join(tempDir, "project")
	err := os.MkdirAll(projectDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create project directory: %s", err)
	}

	settingsFilePath := filepath.Join(projectDir, ".gemini", "settings.json")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccGeminiSettingsFileResourceConfig(projectDir, "gemini-1.5-pro-latest", true),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("agentsmith_gemini_settings_file.test", "path", settingsFilePath),
					testAccCheckFileContent(settingsFilePath, `{
  "model": {
    "name": "gemini-1.5-pro-latest"
  },
  "ui": {
    "hideBanner": true
  }
}`),
				),
			},
			// Update and Read testing
			{
				Config: testAccGeminiSettingsFileResourceConfig(projectDir, "gemini-1.0-ultra", false),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("agentsmith_gemini_settings_file.test", "path", settingsFilePath),
					testAccCheckFileContent(settingsFilePath, `{
  "model": {
    "name": "gemini-1.0-ultra"
  }
}`),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccGeminiSettingsFileResourceConfig(projectDir, modelName string, hideBanner bool) string {
	return fmt.Sprintf(`
provider "agentsmith" {
  workdir = "%s"
}

resource "agentsmith_gemini_settings_file" "test" {
  scope       = "project"
  project_dir = "%s"

  settings = {
    model = {
      name = "%s"
    }
    ui = {
      hide_banner = %t
    }
  }
}
`, projectDir, projectDir, modelName, hideBanner)
}

func testAccCheckFileContent(path, expectedContent string) resource.TestCheckFunc {
	return func(s *tf.State) error {
		content, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read file %s: %w", path, err)
		}
		if string(content) != expectedContent {
			return fmt.Errorf("file content mismatch at %s\nExpected: %s\nGot: %s", path, expectedContent, string(content))
		}
		return nil
	}
}
