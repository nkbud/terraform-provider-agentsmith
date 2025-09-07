resource "agentsmith_claude_settings" "project" {
  # Manages the .claude/settings.json file.
  # Scope can be "project", "user", or "local".
  scope = "project"

  model        = "claude-3-5-sonnet-20240620"
  output_style = "concise"

  permissions = {
    default_mode = "ask"
    allow        = ["read", "write"]
    deny         = ["network"]
  }
}
