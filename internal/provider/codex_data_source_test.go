package provider

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCodexDataSource_EffectiveConfigAndTables(t *testing.T) {

	// Prepare temp CODEX_HOME and project workdir
	homeDir := t.TempDir()
	workDir := t.TempDir()

	t.Setenv("CODEX_HOME", homeDir)
	t.Setenv("OPENAI_BASE_URL", "http://localhost:11434/v1")
	t.Setenv("OPENAI_API_KEY", "test")

	// Home config (lower precedence)
	homeCfg := `
model = "gpt-5"
approval_policy = "untrusted"
profile = "o3_high"

[model_providers.openai]
name = "OpenAI"
base_url = "https://api.openai.com/v1"
env_key = "OPENAI_API_KEY"

[profiles.o3_high]
model = "o3"
approval_policy = "never"
`
	if err := os.WriteFile(filepath.Join(homeDir, "config.toml"), []byte(homeCfg), 0o644); err != nil {
		t.Fatalf("write home config: %v", err)
	}

	// Project config (higher precedence)
	projectCfg := `
sandbox_mode = "workspace-write"

[sandbox_workspace_write]
network_access = true
writable_roots = ["/tmp"]

[mcp_servers.test]
command = "npx"
args = ["-y", "mcp"]
env = { API_KEY = "xyz" }

[shell_environment_policy]
inherit = "core"
`
	if err := os.MkdirAll(filepath.Join(workDir, ".codex"), 0o755); err != nil {
		t.Fatalf("mkdir project .codex: %v", err)
	}
	if err := os.WriteFile(filepath.Join(workDir, ".codex", "config.toml"), []byte(projectCfg), 0o644); err != nil {
		t.Fatalf("write project config: %v", err)
	}

	cfg := fmt.Sprintf(`
provider "agentsmith" {
  workdir = %q
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
					resource.TestCheckResourceAttr("data.agentsmith_codex.this", "id", "codex-config"),
					resource.TestCheckResourceAttr("data.agentsmith_codex.this", "source_files.#", "2"),
					resource.TestCheckResourceAttr("data.agentsmith_codex.this", "active_profile", "o3_high"),

					// Effective config scalars
					resource.TestCheckResourceAttr("data.agentsmith_codex.this", "effective_config.model", "o3"),
					resource.TestCheckResourceAttr("data.agentsmith_codex.this", "effective_config.approval_policy", "never"),
					resource.TestCheckResourceAttr("data.agentsmith_codex.this", "effective_config.sandbox_mode", "workspace-write"),
					resource.TestCheckResourceAttr("data.agentsmith_codex.this", "effective_config.sandbox_workspace_write.network_access", "true"),

					// Model providers (from env override and home config)
					resource.TestCheckResourceAttr("data.agentsmith_codex.this", "effective_config.model_providers.#", "1"),
					resource.TestCheckResourceAttr("data.agentsmith_codex.this", "effective_config.model_providers.0.id", "openai"),
					resource.TestCheckResourceAttr("data.agentsmith_codex.this", "effective_config.model_providers.0.base_url", "http://localhost:11434/v1"),
					resource.TestCheckResourceAttr("data.agentsmith_codex.this", "effective_config.model_providers.0.env_key", "OPENAI_API_KEY"),
					resource.TestCheckResourceAttr("data.agentsmith_codex.this", "effective_config.model_providers.0.env_key_is_set", "true"),

					// MCP servers
					resource.TestCheckResourceAttr("data.agentsmith_codex.this", "effective_config.mcp_servers.#", "1"),
					resource.TestCheckResourceAttr("data.agentsmith_codex.this", "effective_config.mcp_servers.0.id", "test"),
					resource.TestCheckResourceAttr("data.agentsmith_codex.this", "effective_config.mcp_servers.0.command", "npx"),

					// Profiles listing
					resource.TestCheckResourceAttr("data.agentsmith_codex.this", "effective_config.profiles.#", "1"),
					resource.TestCheckResourceAttr("data.agentsmith_codex.this", "effective_config.profiles.0.name", "o3_high"),
				),
			},
		},
	})
}
