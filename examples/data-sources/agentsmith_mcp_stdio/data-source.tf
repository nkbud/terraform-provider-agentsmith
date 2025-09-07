data "agentsmith_mcp_stdio" "example" {
  command     = "python"
  args        = ["-m", "my_mcp_server"]
  description = "An example stdio MCP server."
  timeout     = 5000

  env = {
    "LOG_LEVEL" = "info"
  }
}

output "mcp_stdio_server_json" {
  description = "The JSON representation of the stdio MCP server."
  value       = data.agentsmith_mcp_stdio.example.json
}
