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

# Read Gemini CLI configuration and environment
data "agentsmith_gemini" "example" {
  # The data source will automatically discover Gemini CLI configuration files
  # in the following locations with precedence order:
  # 1. System defaults (e.g., /Library/Application Support/GeminiCli/system-defaults.json)
  # 2. User settings (~/.gemini/settings.json)
  # 3. Project settings (.gemini/settings.json)
  # 4. System settings (system overrides)
}

# Output the discovered configuration
output "gemini_config" {
  description = "Gemini CLI configuration details"
  value = {
    # Environment variables
    environment = {
      gemini_api_key                   = data.agentsmith_gemini.example.environment_variables.gemini_api_key
      gemini_model                     = data.agentsmith_gemini.example.environment_variables.gemini_model
      google_api_key                   = data.agentsmith_gemini.example.environment_variables.google_api_key
      google_cloud_project             = data.agentsmith_gemini.example.environment_variables.google_cloud_project
      google_application_credentials   = data.agentsmith_gemini.example.environment_variables.google_application_credentials
      debug                            = data.agentsmith_gemini.example.environment_variables.debug
      debug_mode                       = data.agentsmith_gemini.example.environment_variables.debug_mode
    }
    
    # General settings
    general = {
      preferred_editor      = data.agentsmith_gemini.example.settings.general.preferred_editor
      vim_mode             = data.agentsmith_gemini.example.settings.general.vim_mode
      disable_auto_update  = data.agentsmith_gemini.example.settings.general.disable_auto_update
      disable_update_nag   = data.agentsmith_gemini.example.settings.general.disable_update_nag
      checkpointing_enabled = data.agentsmith_gemini.example.settings.general.checkpointing_enabled
    }
    
    # UI settings
    ui = {
      theme                 = data.agentsmith_gemini.example.settings.ui.theme
      custom_themes         = data.agentsmith_gemini.example.settings.ui.custom_themes
      hide_window_title     = data.agentsmith_gemini.example.settings.ui.hide_window_title
      hide_tips            = data.agentsmith_gemini.example.settings.ui.hide_tips
      hide_banner          = data.agentsmith_gemini.example.settings.ui.hide_banner
      show_memory_usage    = data.agentsmith_gemini.example.settings.ui.show_memory_usage
      show_line_numbers    = data.agentsmith_gemini.example.settings.ui.show_line_numbers
    }
    
    # Model settings
    model = {
      name               = data.agentsmith_gemini.example.settings.model.name
      max_session_turns  = data.agentsmith_gemini.example.settings.model.max_session_turns
    }
    
    # Privacy settings
    privacy = {
      usage_statistics_enabled = data.agentsmith_gemini.example.settings.privacy.usage_statistics_enabled
    }
    
    # Tools configuration
    tools = {
      core     = data.agentsmith_gemini.example.settings.tools.core
      exclude  = data.agentsmith_gemini.example.settings.tools.exclude
      allowed  = data.agentsmith_gemini.example.settings.tools.allowed
    }
    
    # Context settings
    context = {
      file_name            = data.agentsmith_gemini.example.settings.context.file_name
      include_directories  = data.agentsmith_gemini.example.settings.context.include_directories
    }
  }
}

# Example: Use Gemini configuration values in other resources
resource "local_file" "gemini_summary" {
  filename = "gemini-config-summary.txt"
  content  = <<-EOT
    Gemini CLI Configuration Summary
    ===============================
    
    Model Configuration:
    - Model Name: ${data.agentsmith_gemini.example.settings.model.name}
    - Max Session Turns: ${data.agentsmith_gemini.example.settings.model.max_session_turns}
    
    UI Settings:
    - Theme: ${data.agentsmith_gemini.example.settings.ui.theme}
    - Hide Window Title: ${data.agentsmith_gemini.example.settings.ui.hide_window_title}
    - Show Memory Usage: ${data.agentsmith_gemini.example.settings.ui.show_memory_usage}
    - Show Line Numbers: ${data.agentsmith_gemini.example.settings.ui.show_line_numbers}
    
    General Settings:
    - Preferred Editor: ${data.agentsmith_gemini.example.settings.general.preferred_editor}
    - Vim Mode: ${data.agentsmith_gemini.example.settings.general.vim_mode}
    - Auto Update Disabled: ${data.agentsmith_gemini.example.settings.general.disable_auto_update}
    - Checkpointing Enabled: ${data.agentsmith_gemini.example.settings.general.checkpointing_enabled}
    
    Privacy Settings:
    - Usage Statistics: ${data.agentsmith_gemini.example.settings.privacy.usage_statistics_enabled}
    
    Environment Variables:
    - GEMINI_API_KEY: ${data.agentsmith_gemini.example.environment_variables.gemini_api_key != "" ? "Set" : "Not set"}
    - GOOGLE_CLOUD_PROJECT: ${data.agentsmith_gemini.example.environment_variables.google_cloud_project}
    - DEBUG: ${data.agentsmith_gemini.example.environment_variables.debug}
    - DEBUG_MODE: ${data.agentsmith_gemini.example.environment_variables.debug_mode}
  EOT
}