# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Common Commands

### Build and Install
- `go install` - Build and install the provider
- `make build` - Build the provider without installing
- `make install` - Build and install (runs `make build` first)

### Development Workflow
- `make` or `make default` - Run fmt, lint, install, and generate
- `make fmt` - Format Go code using gofmt
- `make lint` - Run golangci-lint
- `make generate` - Generate documentation and other auto-generated files (runs from tools/ directory)

### Testing
- `make test` - Run unit tests with coverage (timeout 120s, parallel=10)
- `make testacc` - Run acceptance tests (sets TF_ACC=1, timeout 120m)
- `go test -v -cover ./internal/provider/` - Run provider tests specifically

### Direct Go Commands
- `go build -v ./...` - Build all packages
- `go test -v -cover -timeout=120s -parallel=10 ./...` - Run all tests
- `TF_ACC=1 go test -v -cover -timeout 120m ./...` - Run acceptance tests

## Architecture

This is a Terraform provider built with the Terraform Plugin Framework that manages AI agent configurations (Claude, Codex, Gemini). The provider operates on local files within a configurable working directory.

### Core Components

#### Provider (`internal/provider/provider.go`)
- Main provider implementation: `agentsmithProvider`
- Configurable working directory via `workdir` attribute or `AGENTSMITH_WORKDIR` env var
- Defaults to user home directory if not specified
- Uses `FileClient` for file system operations

#### File Client (`internal/provider/provider_file_client.go`) 
- Provides CRUD operations on files within the working directory
- Methods: `Create()`, `Read()`, `Update()`, `Delete()`
- Handles path resolution and directory creation

#### Data Sources
- `claude_data_source.go` - Reads Claude AI configuration files
- `codex_data_source.go` - Reads Codex configuration files  
- `gemini_data_source.go` - Reads Gemini configuration files
- All data sources parse YAML/JSON configuration files and expose structured data

#### Resources
- `claude_subagent_resource.go` - Manages Claude subagent configurations
- Resources handle CRUD operations on configuration files via the FileClient

### File Structure Patterns
- Provider code: `internal/provider/`
- Examples: `examples/complete/` - Contains working Terraform configurations
- Tests: `*_test.go` files alongside implementation files
- Documentation generation: `tools/` directory with Go generate commands

### Configuration Management
The provider manages AI agent configurations stored as files (YAML/JSON format) in the specified working directory. Each data source and resource corresponds to specific configuration file types and structures.

### Testing Strategy
- Unit tests: Standard Go tests for individual components
- Acceptance tests: Full Terraform provider tests using the Plugin Framework testing utilities
- Tests use `TF_ACC=1` environment variable to enable acceptance testing mode

## Claude Configuration Resource Implementation Plan

Based on the existing `claude_data_source.go` implementation and Claude configuration documentation, the following resources need to be implemented to provide full CRUD management of Claude configurations on the host filesystem.

### Resource Overview

The Claude configuration system consists of hierarchical JSON settings files that control Claude Code behavior, permissions, environment variables, and tool access. The resources will manage these files at different scopes and locations.

### Required Resources

#### 1. `agentsmith_claude_settings` Resource
**Purpose**: Manages the main `settings.json` files at user, project, and local project levels.

**File Locations**:
- User: `~/.claude/settings.json`
- Project: `{workdir}/.claude/settings.json` 
- Local: `{workdir}/.claude/settings.local.json`

**Key Attributes**:
- `scope` (Required) - Enum: "user", "project", "local"
- `api_key_helper` - Custom script for auth value generation
- `cleanup_period_days` - Chat transcript retention period
- `env` - Map of environment variables for sessions
- `include_co_authored_by` - Boolean for git commit bylines
- `model` - Override default model
- `output_style` - Configure output style
- `force_login_method` - Restrict login method ("claudeai", "console")
- `force_login_org_uuid` - Auto-select organization UUID
- `disable_all_hooks` - Boolean to disable hooks
- `aws_auth_refresh` - Custom AWS auth refresh script
- `aws_credential_export` - Custom AWS credential export script
- `enable_all_project_mcp_servers` - Boolean for MCP server approval
- `enabled_mcpjson_servers` - List of approved MCP servers
- `disabled_mcpjson_servers` - List of rejected MCP servers

**Nested Blocks**:
- `permissions` block:
  - `allow` - List of permission rules to allow
  - `ask` - List of permission rules requiring confirmation
  - `deny` - List of permission rules to deny
  - `additional_directories` - List of additional working directories
  - `default_mode` - Default permission mode
  - `disable_bypass_permissions_mode` - Prevent bypass mode
- `status_line` block:
  - `type` - Status line type
  - `command` - Command for status line
- `hooks` - Map of hook configurations by tool name

#### 2. `agentsmith_claude_global_config` Resource  
**Purpose**: Manages global configuration file `~/.claude/global-config.json`.

**File Location**: `~/.claude/global-config.json`

**Key Attributes**:
- `auto_updates` - Boolean for automatic updates (deprecated)
- `preferred_notif_channel` - Notification channel preference
- `theme` - Color theme selection
- `verbose` - Boolean for full command output display

#### 3. `agentsmith_claude_subagent` Resource
**Purpose**: Manages individual subagent Markdown files with YAML frontmatter.

**File Locations**:
- User: `~/.claude/agents/{name}.md`
- Project: `{workdir}/.claude/agents/{name}.md`

**Key Attributes**:
- `scope` (Required) - Enum: "user", "project" 
- `name` (Required) - Subagent name (also used as filename)
- `model` - Model override for subagent
- `description` - Human-readable description
- `color` - Display color for subagent
- `prompt` (Required) - The main prompt content (Markdown body)

**File Format**: 
```
---
name: subagent-name
model: claude-3-5-sonnet
description: Description here
color: blue
---
Prompt content goes here...
```

#### 4. `agentsmith_claude_command` Resource
**Purpose**: Manages custom slash command files.

**File Locations**:
- User: `~/.claude/commands/{name}`
- Project: `{workdir}/.claude/commands/{name}`

**Key Attributes**:
- `scope` (Required) - Enum: "user", "project"
- `name` (Required) - Command name (used as filename)  
- `content` (Required) - Command script content
- `executable` - Boolean to make file executable (default: true)

#### 5. `agentsmith_claude_hook` Resource
**Purpose**: Manages hook definition files.

**File Locations**:
- User: `~/.claude/hooks/{name}`
- Project: `{workdir}/.claude/hooks/{name}`

**Key Attributes**:
- `scope` (Required) - Enum: "user", "project"
- `name` (Required) - Hook name (used as filename)
- `content` (Required) - Hook script content
- `executable` - Boolean to make file executable (default: true)

### Implementation Details

#### File Path Resolution Strategy
Resources must resolve file paths based on scope and provider working directory:
- **User scope**: Always use `os.UserHomeDir()` + `.claude/`
- **Project scope**: Use provider `workdir` + `.claude/`
- **Local scope**: Use provider `workdir` + `.claude/` (for local settings)

#### JSON Schema Validation
Settings resources should validate JSON structure against expected Claude configuration schema to prevent invalid configurations.

#### Directory Creation
Resources must ensure `.claude/` and subdirectories (`agents/`, `commands/`, `hooks/`) exist before writing files.

#### State Management
- Use file modification time and content hash for detecting external changes
- Implement proper diff detection for updates
- Handle missing files gracefully (treat as resource deletion)

#### Error Handling
- Provide clear error messages for permission issues
- Validate file paths are within expected directories
- Handle JSON parsing errors with context

#### Import Support
Resources should support importing existing configuration files using standard Terraform import workflow.

### Testing Requirements

#### Unit Tests
- Test file path resolution for each scope
- Test JSON marshaling/unmarshaling of settings
- Test YAML frontmatter parsing for subagents
- Test directory creation logic
- Test error handling for invalid configurations

#### Acceptance Tests  
- Create, read, update, delete operations for each resource
- Test different scope combinations
- Test import functionality
- Test provider workdir configuration impact
- Test file permission handling

### Security Considerations

#### File Permissions
- Settings files: 0644 (readable by owner/group, writable by owner)
- Executable files (commands/hooks): 0755
- Directories: 0755

#### Path Validation
- Ensure file paths are within expected `.claude/` directories
- Prevent directory traversal attacks
- Validate file names don't contain dangerous characters

#### Sensitive Data Handling
- Mark sensitive attributes (API keys, tokens) as sensitive in schema
- Don't log sensitive configuration values
- Handle credential helper scripts securely