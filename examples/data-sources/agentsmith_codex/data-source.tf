terraform {
  required_providers {
    agentsmith = {
      source = "hashicorp.com/edu/agentsmith"
    }
  }
}

provider "agentsmith" {
  # Optional: Configure working directory
  # workdir = "/path/to/your/project"
}

# Read Codex configuration and environment
data "agentsmith_codex" "example" {
  # The data source will automatically discover Codex configuration files
  # in the following locations:
  # - System configuration
  # - User home directory (~/.codex/config.toml)
  # - Project directory (.codex/config.toml)
}

# Output the discovered configuration
output "codex_config" {
  description = "Codex configuration details"
  value = {
    active_profile = data.agentsmith_codex.example.active_profile
    source_files   = data.agentsmith_codex.example.source_files
    environment = {
      codex_home                   = data.agentsmith_codex.example.environment.codex_home
      openai_base_url              = data.agentsmith_codex.example.environment.openai_base_url
      sandbox_network_disabled     = data.agentsmith_codex.example.environment.codex_sandbox_network_disabled
    }
    effective_config = {
      # Configuration merged from all sources with precedence order:
      # project > user > system
      provider         = data.agentsmith_codex.example.effective_config.provider
      model           = data.agentsmith_codex.example.effective_config.model
      temperature     = data.agentsmith_codex.example.effective_config.temperature
      max_tokens      = data.agentsmith_codex.example.effective_config.max_tokens
      system_prompt   = data.agentsmith_codex.example.effective_config.system_prompt
      auto_commit     = data.agentsmith_codex.example.effective_config.auto_commit
      auto_test       = data.agentsmith_codex.example.effective_config.auto_test
      editor          = data.agentsmith_codex.example.effective_config.editor
      shell           = data.agentsmith_codex.example.effective_config.shell
    }
  }
}

# Example: Use Codex configuration values in other resources
resource "local_file" "codex_summary" {
  filename = "codex-config-summary.txt"
  content  = <<-EOT
    Codex Configuration Summary
    ==========================
    Active Profile: ${data.agentsmith_codex.example.active_profile}
    Provider: ${data.agentsmith_codex.example.effective_config.provider}
    Model: ${data.agentsmith_codex.example.effective_config.model}
    Temperature: ${data.agentsmith_codex.example.effective_config.temperature}
    Max Tokens: ${data.agentsmith_codex.example.effective_config.max_tokens}
    Auto Commit: ${data.agentsmith_codex.example.effective_config.auto_commit}
    Auto Test: ${data.agentsmith_codex.example.effective_config.auto_test}
    Editor: ${data.agentsmith_codex.example.effective_config.editor}
    Shell: ${data.agentsmith_codex.example.effective_config.shell}
    
    Environment Variables:
    - CODEX_HOME: ${data.agentsmith_codex.example.environment.codex_home}
    - OPENAI_BASE_URL: ${data.agentsmith_codex.example.environment.openai_base_url}
    - CODEX_SANDBOX_NETWORK_DISABLED: ${data.agentsmith_codex.example.environment.codex_sandbox_network_disabled}
    
    Configuration Files Loaded:
    ${join("\n", data.agentsmith_codex.example.source_files)}
  EOT
}