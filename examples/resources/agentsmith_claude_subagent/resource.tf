resource "agentsmith_claude_subagent" "example" {
  name = "my-subagent"

  # The persona is a markdown-formatted string that defines
  # the subagent's personality and instructions.
  persona = <<-EOT
    # My Subagent Persona

    You are a helpful assistant that specializes in writing Terraform code.
    Always provide clear explanations for the code you generate.
  EOT

  # Tools are other commands or agents this subagent can use.
  tools = ["default:read-file", "default:write-file"]
}
