# Gemini Configuration Inspector Implementation Plan

## Overview
This plan outlines the implementation of a Gemini CLI configuration inspector as a Terraform data source, following the pattern established by the existing Claude data source in the agentsmith provider.

## Analysis of Existing Implementation

### Claude Data Source Pattern
The existing `claude_data_source.go` provides a comprehensive template:
- **Schema Definition**: Nested attributes for settings, environment variables, global config
- **File Discovery**: Multi-location settings file reading with proper precedence
- **Environment Variables**: Extensive mapping of all Claude-related environment variables  
- **Hierarchical Loading**: Merging configurations from multiple sources
- **Discovery Logic**: Finding subagents, commands, and hook files in `.claude` directories

### Key Architecture Components
1. **Data Source Models**: Struct definitions with `tfsdk` tags for Terraform integration
2. **Schema Definition**: Comprehensive attribute mapping with descriptions
3. **File Reading Logic**: JSON parsing with error handling and diagnostics
4. **Environment Variable Reading**: Type-safe conversion for strings, booleans, and integers
5. **Provider Integration**: Registration in `provider.go` DataSources method

## Implementation Plan for Gemini Data Source

### 1. Schema Design (Based on Gemini CLI Configuration)

From the gemini_config_helper.md documentation, Gemini CLI uses a different configuration structure than Claude:

#### Core Configuration Categories:
- **`general`**: Basic settings (preferredEditor, vimMode, disableAutoUpdate, etc.)
- **`ui`**: User interface settings (theme, customThemes, hideWindowTitle, etc.)  
- **`ide`**: IDE integration settings
- **`privacy`**: Usage statistics and privacy controls
- **`model`**: Model configuration (name, maxSessionTurns, chatCompression, etc.)
- **`context`**: Context file and memory settings
- **`tools`**: Tool configuration (sandbox, usePty, core tools, etc.)
- **`mcp`**: Model Context Protocol server settings
- **`security`**: Authentication and folder trust settings  
- **`advanced`**: Advanced configuration options
- **`mcpServers`**: Individual MCP server configurations
- **`telemetry`**: Logging and metrics configuration

#### Environment Variables:
- `GEMINI_API_KEY`, `GOOGLE_API_KEY`, `GOOGLE_CLOUD_PROJECT`
- `GOOGLE_APPLICATION_CREDENTIALS`, `GOOGLE_CLOUD_LOCATION`
- `GEMINI_SANDBOX`, `GEMINI_MODEL`
- Debug and proxy settings

### 2. File Discovery Logic

Gemini CLI uses a different hierarchy than Claude:

#### Settings File Locations:
1. **System defaults**: `/etc/gemini-cli/system-defaults.json` (Linux), similar paths for other OS
2. **User settings**: `~/.gemini/settings.json` 
3. **Project settings**: `.gemini/settings.json` in project root
4. **System overrides**: `/etc/gemini-cli/settings.json` (Linux), similar paths for other OS

#### Configuration Precedence:
1. Default values (hardcoded)
2. System defaults file  
3. User settings file
4. Project settings file
5. System settings file (overrides)
6. Environment variables
7. Command-line arguments

### 3. Implementation Structure

#### File: `internal/provider/gemini_data_source.go`

```go
// Core data models following Terraform plugin framework patterns
type geminiDataSource struct {
    client *FileClient
}

type geminiDataSourceModel struct {
    ID                   types.String                    `tfsdk:"id"`
    Settings            *geminiSettingsModel            `tfsdk:"settings"`
    EnvironmentVariables *geminiEnvironmentVariablesModel `tfsdk:"environment_variables"`
    // Additional discovery data like MCP servers, context files, etc.
}

// Nested models for each configuration category
type geminiSettingsModel struct {
    General   *generalSettingsModel   `tfsdk:"general"`
    UI        *uiSettingsModel        `tfsdk:"ui"`  
    Model     *modelSettingsModel     `tfsdk:"model"`
    Context   *contextSettingsModel   `tfsdk:"context"`
    Tools     *toolsSettingsModel     `tfsdk:"tools"`
    Security  *securitySettingsModel  `tfsdk:"security"`
    // ... additional categories
}
```

#### Schema Definition Strategy:
- Map each Gemini configuration category to a nested attribute
- Preserve the hierarchical structure from the Gemini documentation
- Use appropriate Terraform types (String, Bool, Int64, Map, List)
- Include comprehensive descriptions from the Gemini docs

#### File Reading Implementation:
- Support all four settings file locations with proper OS-specific paths
- Implement hierarchical merging with correct precedence
- Handle JSON parsing errors gracefully with diagnostics
- Support environment variable substitution in settings files (`$VAR_NAME` syntax)

### 4. Key Implementation Details

#### Settings File Paths (OS-Specific):
```go
func (d *geminiDataSource) getSettingsFilePaths(diagnostics *diag.Diagnostics) []string {
    // System defaults
    // User settings: ~/.gemini/settings.json  
    // Project settings: .gemini/settings.json
    // System overrides (OS-specific paths)
}
```

#### Environment Variable Mapping:
Comprehensive mapping of all Gemini-related environment variables:
- API keys and authentication
- Model and project configuration  
- Debugging and operational settings
- Proxy and network configuration

#### MCP Server Discovery:
- Parse `mcpServers` configuration section
- Discover `.gemini` directory files related to MCP
- Map server configurations with connection details

### 5. Integration Steps

#### Provider Registration:
Update `internal/provider/provider.go`:
```go  
func (p *agentsmithProvider) DataSources(_ context.Context) []func() datasource.DataSource {
    return []func() datasource.DataSource{
        NewClaudeDataSource,
        NewGeminiDataSource,  // Add this line
    }
}
```

#### Testing Strategy:
1. Unit tests for settings file parsing
2. Integration tests for multi-file configuration merging
3. Environment variable precedence testing  
4. OS-specific path resolution testing

### 6. Configuration Validation

#### File Format Validation:
- JSON schema validation for settings files
- Graceful handling of missing or malformed files
- Warning diagnostics for deprecated settings

#### Environment Variable Validation:
- Type checking for numeric values
- Boolean parsing for string representations  
- Path validation for file-based settings

### 7. Documentation and Examples

#### Terraform Usage Examples:
```hcl
data "agentsmith_gemini" "config" {
}

output "gemini_model" {
  value = data.agentsmith_gemini.config.settings.model.name
}

output "ui_theme" {
  value = data.agentsmith_gemini.config.settings.ui.theme  
}
```

## Benefits of This Implementation

1. **Comprehensive Configuration Inspection**: Full visibility into Gemini CLI configuration
2. **Terraform Integration**: Native Terraform data source for infrastructure-as-code workflows
3. **Multi-Environment Support**: Handles user, project, and system-level configurations
4. **Type Safety**: Proper type mapping for all configuration values
5. **Error Handling**: Robust error reporting and diagnostics
6. **Cross-Platform**: OS-specific path handling for Windows, macOS, and Linux

## Implementation Priority

1. ✅ **Analysis Phase**: Understanding existing patterns (COMPLETED)
2. 🔄 **Schema Design**: Define comprehensive Terraform schema for Gemini settings  
3. 📝 **Core Implementation**: Build gemini_data_source.go with file reading logic
4. 🔧 **Environment Integration**: Add environment variable mapping
5. 🔗 **Provider Registration**: Integrate with existing provider
6. ✅ **Testing**: Comprehensive test coverage
7. 📚 **Documentation**: Usage examples and configuration reference

This implementation will provide a powerful tool for inspecting and managing Gemini CLI configurations through Terraform, enabling infrastructure teams to standardize and validate Gemini setups across development environments.

## Gemini Configuration Resource Implementation Plan

### Overview
This plan details the creation of a new Terraform resource, `agentsmith_gemini_settings_file`, designed to write and manage Gemini CLI `settings.json` files on the local filesystem. This complements the `agentsmith_gemini` data source by providing write capabilities, allowing for declarative configuration of the Gemini CLI environment.

### 1. Resource Naming and Scope
- **Resource Name**: `agentsmith_gemini_settings_file`
- **Functionality**: Manages a single `settings.json` file at a specified scope (`project`, `user`, or `system_override`).

### 2. Schema Design
The resource schema will largely mirror the `geminiDataSourceModel`, but with attributes configured as inputs (`Optional` or `Required`) rather than `Computed`.

#### Key Resource Arguments:
- **`scope`** (String, Required): Defines the target configuration scope. Must be one of `project`, `user`, or `system_override`. This determines the file's location.
- **`project_dir`** (String, Optional): The absolute path to the project's root directory. Required only when `scope` is `project`.
- **`settings`** (Block, Required): A block containing all the Gemini settings, mirroring the structure in the data source. All nested attributes will be `Optional`.

#### Computed Attributes:
- **`id`** (String): The unique ID of the resource, which will be the absolute path to the managed `settings.json` file.
- **`path`** (String): The absolute path to the managed `settings.json` file.

### 3. Implementation Structure

#### File: `internal/provider/gemini_settings_file_resource.go`
```go
// Core resource and model definitions
type geminiSettingsFileResource struct {
    client *FileClient
}

type geminiSettingsFileResourceModel struct {
    Scope      types.String         `tfsdk:"scope"`
    ProjectDir types.String         `tfsdk:"project_dir"`
    Settings   *geminiSettingsModel `tfsdk:"settings"` // Re-use the same settings model
    ID         types.String         `tfsdk:"id"`
    Path       types.String         `tfsdk:"path"`
}

// The geminiSettingsModel from the data source can be reused here,
// but the schema definition will mark its attributes as Optional.
```

### 4. Core Logic (CRUD Operations)

#### `Create` / `Update`:
1.  **Determine Path**: Based on the `scope` and `project_dir` attributes, determine the target file path (e.g., `/path/to/project/.gemini/settings.json` for `project` scope).
2.  **Ensure Directory**: For `project` scope, ensure the `.gemini` directory exists. For other scopes, ensure the parent directory exists.
3.  **Build JSON**: Convert the `settings` block from the Terraform configuration into a Go struct.
4.  **Marshal to JSON**: Marshal the Go struct into a well-formatted JSON string.
5.  **Write File**: Write the JSON content to the target path, overwriting any existing file.
6.  **Set State**: Set the `id` and `path` computed attributes in the Terraform state.

#### `Read`:
1.  Read the file from the path stored in the Terraform state.
2.  If the file doesn't exist, signal to Terraform that the resource needs to be recreated.
3.  If the file exists, unmarshal its JSON content.
4.  Compare the content with the current state. If there's a drift (i.e., the file was changed manually), the plan will show a diff.

#### `Delete`:
1.  Read the file path from the state.
2.  Delete the `settings.json` file. The resource should not delete the parent `.gemini` directory.

### 5. Provider Integration
Update `internal/provider/provider.go` to include the new resource:
```go
func (p *agentsmithProvider) Resources(_ context.Context) []func() resource.Resource {
    return []func() resource.Resource{
        NewClaudeSubagentResource,
        NewGeminiSettingsFileResource, // Add this line
    }
}
```

### 6. Example HCL Usage
This example demonstrates how to create a project-level `settings.json` file.

```hcl
resource "agentsmith_gemini_settings_file" "project_config" {
  scope       = "project"
  project_dir = "/path/to/my/gemini-project"

  settings {
    general {
      vim_mode = true
    }
    
    ui {
      theme      = "GitHub"
      hide_banner = true
    }

    model {
      name = "gemini-1.5-pro-latest"
    }

    tools {
      sandbox = "docker"
      allowed = [
        "run_shell_command(git)",
        "run_shell_command(npm test)"
      ]
    }
  }
}

output "gemini_project_settings_path" {
  value = agentsmith_gemini_settings_file.project_config.path
}
```


### 7. Benefits of This Implementation
1.  **Declarative Setup**: Manage Gemini CLI configuration using Infrastructure as Code.
2.  **Standardization**: Enforce consistent Gemini settings across teams and projects.
3.  **Automation**: Easily bootstrap new development environments with the correct Gemini configuration.
4.  **Lifecycle Management**: Terraform will handle the creation, update, and deletion of the configuration file, preventing orphaned settings.

### 8. Managing Other Configuration Files (`GEMINI.md`, `.env`)

The `agentsmith_gemini_settings_file` resource is specifically designed for the structured `settings.json` file. Other configuration files, which are typically unstructured plain text, should be managed using the standard `hashicorp/local` provider.

This approach avoids duplicating existing Terraform functionality and provides a flexible way to manage all aspects of the Gemini CLI environment.

#### Example: Managing `GEMINI.md` with `local_file`

To declaratively manage a project-specific context file, you can use the `local_file` resource. This ensures your instructional context is version-controlled and applied consistently.

```hcl
resource "local_file" "gemini_context" {
  content  = "# Project: My Awesome TypeScript Library\n\n- Follow existing coding style.\n- Ensure all new functions have JSDoc comments."
  filename = "${path.root}/.gemini/GEMINI.md"
}
```