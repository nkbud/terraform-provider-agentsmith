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

# Example 1: User-level Codex configuration file
resource "agentsmith_codex_config" "user_config" {
  scope                       = "user"
  create_directories          = true
  backup_on_write            = true
  allow_sensitive_env_writes = false
  validate_strict            = true
  
  # Basic configuration
  profile                    = "default"
  model                     = "gpt-4o"
  model_provider           = "openai"
  model_context_window     = 128000
  model_max_output_tokens  = 4096
  approval_policy          = "ask"
  sandbox_mode             = "safe"
  file_opener              = "code"
  hide_agent_reasoning     = false
  show_raw_agent_reasoning = false
  model_reasoning_effort   = "medium"
  model_reasoning_summary  = "auto"
  model_verbosity         = "normal"
  model_supports_reasoning_summaries = true
  project_doc_max_bytes   = 1048576
  
  # Notification settings
  notify = ["desktop", "sound"]
  
  # TUI settings
  tui = {
    theme = "dark"
    show_line_numbers = "true"
    tab_size = "2"
  }
  
  # Sandbox workspace configuration
  sandbox_workspace_write {
    writable_roots = [
      "~/Projects",
      "~/temp"
    ]
    network_access = true
    exclude_patterns = [
      "*.secret",
      ".env",
      "id_rsa*"
    ]
  }
  
  # History settings
  history {
    keep_last_n = 100
    auto_save   = true
  }
  
  # Shell environment policy
  shell_environment_policy {
    inherit_from_parent = true
    allowed_env_vars = [
      "PATH",
      "HOME",
      "USER",
      "SHELL"
    ]
    blocked_env_vars = [
      "AWS_SECRET_ACCESS_KEY",
      "OPENAI_API_KEY"
    ]
  }
  
  # Model providers configuration
  model_providers {
    name        = "openai"
    description = "OpenAI GPT models"
    
    models {
      name              = "gpt-4o"
      supports_tools    = true
      supports_vision   = true
      context_window    = 128000
      max_output_tokens = 4096
      reasoning_effort_levels = ["low", "medium", "high"]
    }
    
    models {
      name              = "gpt-4o-mini"
      supports_tools    = true
      supports_vision   = true
      context_window    = 128000
      max_output_tokens = 16384
      reasoning_effort_levels = ["low", "medium"]
    }
    
    authentication {
      type      = "api_key"
      env_var   = "OPENAI_API_KEY"
      required  = true
    }
    
    request_options {
      timeout = 60
      retries = 3
    }
  }
  
  # MCP servers configuration
  mcp_servers {
    name        = "filesystem"
    description = "File system operations"
    
    transport {
      type    = "stdio"
      command = "mcp-filesystem-server"
      args    = ["--safe-mode"]
    }
  }
  
  mcp_servers {
    name        = "database"
    description = "Database operations"
    
    transport {
      type    = "stdio" 
      command = "mcp-database-server"
      args    = ["--host", "localhost", "--port", "5432"]
    }
    
    environment {
      DB_PASSWORD = "sensitive_value"
    }
  }
  
  # Profiles configuration
  profiles {
    name = "development"
    
    model                = "gpt-4o-mini"
    approval_policy      = "auto"
    sandbox_mode         = "permissive"
    model_reasoning_effort = "low"
    
    # Override some settings for development
    sandbox_workspace_write {
      writable_roots = [
        "~/dev",
        "~/Projects"
      ]
      network_access = true
    }
  }
  
  profiles {
    name = "production"
    
    model           = "gpt-4o"
    approval_policy = "strict"
    sandbox_mode   = "restricted"
    model_reasoning_effort = "high"
    
    # Strict settings for production
    sandbox_workspace_write {
      writable_roots = [
        "~/prod-workspace"
      ]
      network_access = false
      exclude_patterns = [
        "*",
        "!*.txt",
        "!*.md"
      ]
    }
    
    shell_environment_policy {
      inherit_from_parent = false
      allowed_env_vars = [
        "PATH"
      ]
    }
  }
}

# Example 2: Project-level Codex configuration
resource "agentsmith_codex_config" "project_config" {
  scope              = "project" 
  path               = ".codex/config.toml"
  create_directories = true
  merge_strategy     = "merge"
  
  # Override settings for this project
  profile         = "terraform"
  model          = "gpt-4o"
  approval_policy = "ask"
  
  # Project-specific TUI settings
  tui = {
    theme           = "solarized-dark"
    show_line_numbers = "true"
    tab_size        = "4"
    word_wrap       = "true"
  }
  
  # Project-specific sandbox settings
  sandbox_workspace_write {
    writable_roots = [
      ".",
      "./modules",
      "./environments"
    ]
    network_access = false
    exclude_patterns = [
      "*.tfstate",
      "*.tfstate.backup",
      ".terraform/",
      "*.secret"
    ]
  }
  
  # Project-specific MCP servers
  mcp_servers {
    name = "terraform"
    
    transport {
      type    = "stdio"
      command = "mcp-terraform-server"
      args    = ["--workspace", "development"]
    }
  }
  
  # Terraform-specific profile
  profiles {
    name = "terraform"
    
    model               = "gpt-4o"
    approval_policy     = "ask"
    project_doc_max_bytes = 2097152 # 2MB for larger Terraform docs
    
    sandbox_workspace_write {
      writable_roots = [
        ".",
        "./modules"
      ]
      exclude_patterns = [
        "*.tfstate*",
        ".terraform/"
      ]
    }
  }
}

# Example 3: System-level configuration
resource "agentsmith_codex_config" "system_config" {
  scope                      = "system"
  path                      = "/etc/codex/config.toml"
  create_directories        = true
  allow_sensitive_env_writes = false
  validate_strict           = true
  keep_file_on_destroy      = true
  
  # System-wide defaults
  approval_policy    = "strict"
  sandbox_mode      = "safe"
  model_verbosity   = "quiet"
  
  # System-wide security settings
  sandbox_workspace_write {
    writable_roots = [
      "/tmp/codex-workspace"
    ]
    network_access = false
    exclude_patterns = [
      "/etc/*",
      "/var/*",
      "/usr/*",
      "*.key",
      "*.pem",
      "*.p12"
    ]
  }
  
  # Restrictive shell environment
  shell_environment_policy {
    inherit_from_parent = false
    allowed_env_vars = [
      "PATH",
      "HOME",
      "USER",
      "LANG",
      "LC_*"
    ]
    blocked_env_vars = [
      "*_KEY",
      "*_SECRET",
      "*_TOKEN",
      "*_PASSWORD"
    ]
  }
  
  # System-wide history settings
  history {
    keep_last_n = 50
    auto_save   = false
  }
}

# Output the configuration file paths
output "codex_config_paths" {
  description = "Paths to the created Codex configuration files"
  value = {
    user_config    = agentsmith_codex_config.user_config.resolved_path
    project_config = agentsmith_codex_config.project_config.resolved_path  
    system_config  = agentsmith_codex_config.system_config.resolved_path
  }
}