data "agentsmith_claude" "current" {
  # This data source reads the merged Claude CLI configuration from
  # the environment, including settings files and environment variables.
}

output "claude_model" {
  description = "The currently configured Claude model."
  value       = data.agentsmith_claude.current.settings.model
}

output "claude_subagents" {
  description = "A list of discovered Claude subagents."
  value       = data.agentsmith_claude.current.subagents
}
