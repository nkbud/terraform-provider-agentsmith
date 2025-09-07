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

# Example 1: User-level Gemini settings file
resource "agentsmith_gemini_settings_file" "user_settings" {
  scope = "user"
  
  settings {
    general {
      preferred_editor      = "code"
      vim_mode             = false
      disable_auto_update  = false
      disable_update_nag   = false
      checkpointing_enabled = true
    }
    
    ui {
      theme                 = "dark"
      hide_window_title    = false
      hide_tips           = false
      hide_banner         = false
      show_memory_usage   = true
      show_line_numbers   = true
      show_citations      = true
      accessibility_disable_loading_phrases = false
      
      # Custom themes configuration
      custom_themes = {
        "my_theme" = "{ \"background\": \"#1e1e1e\", \"foreground\": \"#ffffff\" }"
      }
    }
    
    model {
      name               = "gemini-1.5-flash"
      max_session_turns  = 100
      
      summarize_tool_output = {
        "default" = "auto"
      }
    }
    
    privacy {
      usage_statistics_enabled = true
    }
    
    ide {
      enabled      = true
      has_seen_nudge = false
    }
    
    context {
      file_name = [
        "*.py",
        "*.js",
        "*.ts",
        "*.go"
      ]
      include_directories = [
        "src",
        "lib"
      ]
    }
    
    tools {
      core = [
        "file_edit",
        "bash",
        "str_replace"
      ]
      exclude = [
        "dangerous_tool"
      ]
      allowed = [
        "safe_tool"
      ]
    }
    
    mcp {
      allowed = [
        "trusted_mcp_server"
      ]
      excluded = [
        "untrusted_mcp_server"
      ]
    }
    
    security {
      require_explicit_permission_for_tools = true
      auto_accept_low_risk_tools = false
    }
    
    advanced {
      excluded_env_vars = [
        "SECRET_KEY",
        "API_TOKEN"
      ]
      
      bug_command = {
        "linux"   = "reportbug"
        "darwin"  = "open https://github.com/project/issues"
        "windows" = "start https://github.com/project/issues"
      }
    }
    
    telemetry {
      enabled = true
      endpoint = "https://telemetry.example.com"
      level = "info"
    }
  }
  
  # MCP servers configuration
  mcp_servers = {
    "filesystem" = "{\"command\": \"mcp-filesystem-server\", \"args\": []}"
    "database"   = "{\"command\": \"mcp-database-server\", \"args\": [\"--host\", \"localhost\"]}"
  }
}

# Example 2: Project-level Gemini settings file  
resource "agentsmith_gemini_settings_file" "project_settings" {
  scope       = "project"
  project_dir = "/path/to/my/project"
  
  settings {
    general {
      preferred_editor = "vim"
      vim_mode        = true
    }
    
    ui {
      theme = "light"
      hide_banner = true
      show_memory_usage = false
    }
    
    model {
      name = "gemini-1.5-pro"
      max_session_turns = 50
    }
    
    context {
      file_name = [
        "*.tf",
        "*.hcl",
        "*.yaml",
        "*.yml"
      ]
      include_directories = [
        "modules",
        "environments"
      ]
    }
    
    tools {
      core = [
        "file_edit",
        "bash",
        "terraform"
      ]
    }
    
    security {
      require_explicit_permission_for_tools = false
      auto_accept_low_risk_tools = true
    }
  }
  
  mcp_servers = {
    "terraform" = "{\"command\": \"mcp-terraform-server\", \"args\": [\"--workspace\", \"dev\"]}"
  }
}

# Example 3: System override settings
resource "agentsmith_gemini_settings_file" "system_override" {
  scope = "system_override"
  
  settings {
    security {
      require_explicit_permission_for_tools = true
      auto_accept_low_risk_tools = false
    }
    
    advanced {
      excluded_env_vars = [
        "AWS_SECRET_ACCESS_KEY",
        "GOOGLE_APPLICATION_CREDENTIALS",
        "GITHUB_TOKEN",
        "OPENAI_API_KEY"
      ]
    }
    
    telemetry {
      enabled = false
    }
  }
}

# Output the file paths
output "gemini_settings_paths" {
  description = "Paths to the created Gemini settings files"
  value = {
    user_settings     = agentsmith_gemini_settings_file.user_settings.path
    project_settings  = agentsmith_gemini_settings_file.project_settings.path
    system_override   = agentsmith_gemini_settings_file.system_override.path
  }
}