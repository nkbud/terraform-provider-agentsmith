package provider

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccClaudeDataSource_fileDiscovery(t *testing.T) {
	homeDir := t.TempDir()
	t.Setenv("HOME", homeDir)

	projectDir := t.TempDir()

	// --- Create User-level files ---
	userAgentsDir := filepath.Join(homeDir, ".claude", "agents")
	userCommandsDir := filepath.Join(homeDir, ".claude", "commands")
	os.MkdirAll(userAgentsDir, 0755)
	os.MkdirAll(userCommandsDir, 0755)

	userSubagentContent := "---\nname: User Agent\nmodel: test-model\n---\nUser agent prompt."
	os.WriteFile(filepath.Join(userAgentsDir, "user_agent.md"), []byte(userSubagentContent), 0644)
	os.WriteFile(filepath.Join(userCommandsDir, "user_command"), []byte("user command content"), 0644)

	// --- Create Project-level files ---
	projectAgentsDir := filepath.Join(projectDir, ".claude", "agents")
	projectHooksDir := filepath.Join(projectDir, ".claude", "hooks")
	os.MkdirAll(projectAgentsDir, 0755)
	os.MkdirAll(projectHooksDir, 0755)

	projectSubagentContent := "---\nname: Project Agent\n---\nProject agent prompt."
	os.WriteFile(filepath.Join(projectAgentsDir, "project_agent.md"), []byte(projectSubagentContent), 0644)
	os.WriteFile(filepath.Join(projectHooksDir, "project_hook"), []byte("project hook content"), 0644)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
					provider "agentsmith" {
						workdir = "%s"
					}
					data "agentsmith_claude" "test" {}
				`, projectDir),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Check counts
					resource.TestCheckResourceAttr("data.agentsmith_claude.test", "subagents.#", "2"),
					resource.TestCheckResourceAttr("data.agentsmith_claude.test", "commands.#", "1"),
					resource.TestCheckResourceAttr("data.agentsmith_claude.test", "hook_files.#", "1"),

					// Check specific subagent content
					resource.TestCheckResourceAttr("data.agentsmith_claude.test", "subagents.0.name", "User Agent"),
					resource.TestCheckResourceAttr("data.agentsmith_claude.test", "subagents.0.model", "test-model"),
					resource.TestCheckResourceAttr("data.agentsmith_claude.test", "subagents.1.name", "Project Agent"),
				),
			},
		},
	})
}
