data "agentsmith_mcp_remote" "example" {
  url         = "http://localhost:8888/mcp"
  transport   = "sse"
  description = "An example remote MCP server."
  timeout     = 15000

  headers = {
    "X-API-Key" = "my-secret-key"
  }
}

output "mcp_remote_server_json" {
  description = "The JSON representation of the remote MCP server."
  value       = data.agentsmith_mcp_remote.example.json
  sensitive   = true
}
