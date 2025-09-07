resource "agentsmith_claude_command" "example" {
  name        = "my-custom-command"
  description = "This is a custom command for Claude."
  command     = "echo 'Hello from my-custom-command'"
  subagent    = "default" # or the name of a specific subagent
}
