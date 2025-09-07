resource "agentsmith_claude_global_config" "example" {
  # This resource manages the global ~/.claude/config.json file.
  # It is a singleton resource; only one should be defined.

  model             = "claude-3-opus-20240229"
  output_style      = "pretty"
  cleanup_period_days = 30
}
