package provider

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCodexConfigResource_HomeBasicAndDataSourceRoundTrip(t *testing.T) {
	homeDir := t.TempDir()
	workDir := t.TempDir()
	t.Setenv("CODEX_HOME", homeDir)
	t.Setenv("OPENAI_API_KEY", "test")

	cfg := fmt.Sprintf(`
provider "agentsmith" {
  workdir = %q
}

resource "agentsmith_codex_config" "home" {
  scope           = "home"
  merge_strategy  = "preserve_unknown"
  model           = "gpt-4o"
  model_provider  = "openai"

  model_providers {
    id       = "openai"
    name     = "OpenAI"
    base_url = "https://api.openai.com/v1"
    env_key  = "OPENAI_API_KEY"
  }
}

data "agentsmith_codex" "this" {}
`, workDir)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: cfg,
				Check: resource.ComposeTestCheckFunc(
					// Resource basics
					resource.TestCheckResourceAttrSet("agentsmith_codex_config.home", "id"),
					resource.TestCheckResourceAttr("agentsmith_codex_config.home", "scope", "home"),
					resource.TestCheckResourceAttr("agentsmith_codex_config.home", "resolved_path", filepath.Join(homeDir, "config.toml")),

					// Data source round-trip
					resource.TestCheckResourceAttr("data.agentsmith_codex.this", "effective_config.model", "gpt-4o"),
					resource.TestCheckResourceAttr("data.agentsmith_codex.this", "effective_config.model_providers.#", "1"),
					resource.TestCheckResourceAttr("data.agentsmith_codex.this", "effective_config.model_providers.0.id", "openai"),
				),
			},
		},
	})

	// Verify file exists
	if _, err := os.Stat(filepath.Join(homeDir, "config.toml")); err != nil {
		t.Fatalf("expected config.toml to exist: %v", err)
	}
}
