resource "agentsmith_claude_hook" "example" {
  name    = "pre-prompt"
  command = "echo 'A new prompt is about to be sent.'"
  subagent = "default" # or the name of a specific subagent
}
