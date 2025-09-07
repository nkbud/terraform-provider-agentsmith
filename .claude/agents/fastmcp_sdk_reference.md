
Directory structure:
└── python-sdk/
├── fastmcp-cli-__init__.mdx
├── fastmcp-cli-claude.mdx
├── fastmcp-cli-cli.mdx
├── fastmcp-cli-install-__init__.mdx
├── fastmcp-cli-install-claude_code.mdx
├── fastmcp-cli-install-claude_desktop.mdx
├── fastmcp-cli-install-cursor.mdx
├── fastmcp-cli-install-mcp_json.mdx
├── fastmcp-cli-install-shared.mdx
├── fastmcp-cli-run.mdx
├── fastmcp-client-__init__.mdx
├── fastmcp-client-auth-__init__.mdx
├── fastmcp-client-auth-bearer.mdx
├── fastmcp-client-auth-oauth.mdx
├── fastmcp-client-client.mdx
├── fastmcp-client-elicitation.mdx
├── fastmcp-client-logging.mdx
├── fastmcp-client-messages.mdx
├── fastmcp-client-oauth_callback.mdx
├── fastmcp-client-progress.mdx
├── fastmcp-client-roots.mdx
├── fastmcp-client-sampling.mdx
├── fastmcp-client-transports.mdx
├── fastmcp-exceptions.mdx
├── fastmcp-mcp_config.mdx
├── fastmcp-prompts-__init__.mdx
├── fastmcp-prompts-prompt.mdx
├── fastmcp-prompts-prompt_manager.mdx
├── fastmcp-resources-__init__.mdx
├── fastmcp-resources-resource.mdx
├── fastmcp-resources-resource_manager.mdx
├── fastmcp-resources-template.mdx
├── fastmcp-resources-types.mdx
├── fastmcp-server-__init__.mdx
├── fastmcp-server-auth-__init__.mdx
├── fastmcp-server-auth-auth.mdx
├── fastmcp-server-auth-oauth_proxy.mdx
├── fastmcp-server-auth-providers-__init__.mdx
├── fastmcp-server-auth-providers-azure.mdx
├── fastmcp-server-auth-providers-bearer.mdx
├── fastmcp-server-auth-providers-github.mdx
├── fastmcp-server-auth-providers-google.mdx
├── fastmcp-server-auth-providers-in_memory.mdx
├── fastmcp-server-auth-providers-jwt.mdx
├── fastmcp-server-auth-providers-workos.mdx
├── fastmcp-server-auth-redirect_validation.mdx
├── fastmcp-server-context.mdx
├── fastmcp-server-dependencies.mdx
├── fastmcp-server-elicitation.mdx
├── fastmcp-server-http.mdx
├── fastmcp-server-low_level.mdx
├── fastmcp-server-middleware-__init__.mdx
├── fastmcp-server-middleware-error_handling.mdx
├── fastmcp-server-middleware-logging.mdx
├── fastmcp-server-middleware-middleware.mdx
├── fastmcp-server-middleware-rate_limiting.mdx
├── fastmcp-server-middleware-timing.mdx
├── fastmcp-server-openapi.mdx
├── fastmcp-server-proxy.mdx
├── fastmcp-server-server.mdx
├── fastmcp-settings.mdx
├── fastmcp-tools-__init__.mdx
├── fastmcp-tools-tool.mdx
├── fastmcp-tools-tool_manager.mdx
├── fastmcp-tools-tool_transform.mdx
├── fastmcp-utilities-__init__.mdx
├── fastmcp-utilities-auth.mdx
├── fastmcp-utilities-cli.mdx
├── fastmcp-utilities-components.mdx
├── fastmcp-utilities-exceptions.mdx
├── fastmcp-utilities-http.mdx
├── fastmcp-utilities-inspect.mdx
├── fastmcp-utilities-json_schema.mdx
├── fastmcp-utilities-json_schema_type.mdx
├── fastmcp-utilities-logging.mdx
├── fastmcp-utilities-mcp_config.mdx
├── fastmcp-utilities-mcp_server_config-__init__.mdx
├── fastmcp-utilities-mcp_server_config-v1-__init__.mdx
├── fastmcp-utilities-mcp_server_config-v1-environments-__init__.mdx
├── fastmcp-utilities-mcp_server_config-v1-environments-base.mdx
├── fastmcp-utilities-mcp_server_config-v1-environments-uv.mdx
├── fastmcp-utilities-mcp_server_config-v1-mcp_server_config.mdx
├── fastmcp-utilities-mcp_server_config-v1-sources-__init__.mdx
├── fastmcp-utilities-mcp_server_config-v1-sources-base.mdx
├── fastmcp-utilities-mcp_server_config-v1-sources-filesystem.mdx
├── fastmcp-utilities-openapi.mdx
├── fastmcp-utilities-tests.mdx
└── fastmcp-utilities-types.mdx

================================================
FILE: docs/python-sdk/fastmcp-cli-__init__.mdx
================================================
---
title: __init__
sidebarTitle: __init__
---

# `fastmcp.cli`


FastMCP CLI package.



================================================
FILE: docs/python-sdk/fastmcp-cli-claude.mdx
================================================
---
title: claude
sidebarTitle: claude
---

# `fastmcp.cli.claude`


Claude app integration utilities.

## Functions

### `get_claude_config_path` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/cli/claude.py#L15" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
get_claude_config_path() -> Path | None
```


Get the Claude config directory based on platform.


### `update_claude_config` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/cli/claude.py#L33" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
update_claude_config(file_spec: str, server_name: str) -> bool
```


Add or update a FastMCP server in Claude's configuration.

**Args:**
- `file_spec`: Path to the server file, optionally with \:object suffix
- `server_name`: Name for the server in Claude's config
- `with_editable`: Optional list of directories to install in editable mode
- `with_packages`: Optional list of additional packages to install
- `env_vars`: Optional dictionary of environment variables. These are merged with
  any existing variables, with new values taking precedence.

**Raises:**
- `RuntimeError`: If Claude Desktop's config directory is not found, indicating
  Claude Desktop may not be installed or properly set up.




================================================
FILE: docs/python-sdk/fastmcp-cli-cli.mdx
================================================
---
title: cli
sidebarTitle: cli
---

# `fastmcp.cli.cli`


FastMCP CLI tools using Cyclopts.

## Functions

### `with_argv` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/cli/cli.py#L68" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
with_argv(args: list[str] | None)
```


Temporarily replace sys.argv if args provided.

This context manager is used at the CLI boundary to inject
server arguments when needed, without mutating sys.argv deep
in the source loading logic.

Args are provided without the script name, so we preserve sys.argv[0]
and replace the rest.


### `version` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/cli/cli.py#L91" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
version()
```


Display version information and platform details.


### `dev` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/cli/cli.py#L129" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
dev(server_spec: str | None = None) -> None
```


Run an MCP server with the MCP Inspector for development.

**Args:**
- `server_spec`: Python file to run, optionally with \:object suffix, or None to auto-detect fastmcp.json


### `run` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/cli/cli.py#L310" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
run(server_spec: str | None = None, *server_args: str) -> None
```


Run an MCP server or connect to a remote one.

The server can be specified in several ways:
1. Module approach: "server.py" - runs the module directly, looking for an object named 'mcp', 'server', or 'app'
2. Import approach: "server.py:app" - imports and runs the specified server object
3. URL approach: "http://server-url" - connects to a remote server and creates a proxy
4. MCPConfig file: "mcp.json" - runs as a proxy server for the MCP Servers in the MCPConfig file
5. FastMCP config: "fastmcp.json" - runs server using FastMCP configuration
6. No argument: looks for fastmcp.json in current directory

Server arguments can be passed after -- :
fastmcp run server.py -- --config config.json --debug

**Args:**
- `server_spec`: Python file, object specification (file\:obj), config file, URL, or None to auto-detect


### `inspect` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/cli/cli.py#L524" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
inspect(server_spec: str | None = None) -> None
```


Inspect an MCP server and display information or generate a JSON report.

This command analyzes an MCP server. Without flags, it displays a text summary.
Use --format to output complete JSON data.

**Examples:**

# Show text summary
fastmcp inspect server.py

# Output FastMCP format JSON to stdout
fastmcp inspect server.py --format fastmcp

# Save MCP protocol format to file (format required with -o)
fastmcp inspect server.py --format mcp -o manifest.json

# Inspect from fastmcp.json configuration
fastmcp inspect fastmcp.json
fastmcp inspect  # auto-detect fastmcp.json

**Args:**
- `server_spec`: Python file to inspect, optionally with \:object suffix, or fastmcp.json


### `prepare` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/cli/cli.py#L765" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
prepare(config_path: Annotated[str | None, cyclopts.Parameter(help='Path to fastmcp.json configuration file')] = None, output_dir: Annotated[str | None, cyclopts.Parameter(help='Directory to create the persistent environment in')] = None, skip_source: Annotated[bool, cyclopts.Parameter(help='Skip source preparation (e.g., git clone)')] = False) -> None
```


Prepare a FastMCP project by creating a persistent uv environment.

This command creates a persistent uv project with all dependencies installed:
- Creates a pyproject.toml with dependencies from the config
- Installs all Python packages into a .venv
- Prepares the source (git clone, download, etc.) unless --skip-source

After running this command, you can use:
fastmcp run &lt;config&gt; --project &lt;output-dir&gt;

This is useful for:
- CI/CD pipelines with separate build and run stages
- Docker images where you prepare during build
- Production deployments where you want fast startup times




================================================
FILE: docs/python-sdk/fastmcp-cli-install-__init__.mdx
================================================
---
title: __init__
sidebarTitle: __init__
---

# `fastmcp.cli.install`


Install subcommands for FastMCP CLI using Cyclopts.



================================================
FILE: docs/python-sdk/fastmcp-cli-install-claude_code.mdx
================================================
---
title: claude_code
sidebarTitle: claude_code
---

# `fastmcp.cli.install.claude_code`


Claude Code integration for FastMCP install using Cyclopts.

## Functions

### `find_claude_command` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/cli/install/claude_code.py#L20" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
find_claude_command() -> str | None
```


Find the Claude Code CLI command.

Checks common installation locations since 'claude' is often a shell alias
that doesn't work with subprocess calls.


### `check_claude_code_available` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/cli/install/claude_code.py#L68" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
check_claude_code_available() -> bool
```


Check if Claude Code CLI is available.


### `install_claude_code` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/cli/install/claude_code.py#L73" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
install_claude_code(file: Path, server_object: str | None, name: str) -> bool
```


Install FastMCP server in Claude Code.

**Args:**
- `file`: Path to the server file
- `server_object`: Optional server object name (for \:object suffix)
- `name`: Name for the server in Claude Code
- `with_editable`: Optional list of directories to install in editable mode
- `with_packages`: Optional list of additional packages to install
- `env_vars`: Optional dictionary of environment variables
- `python_version`: Optional Python version to use
- `with_requirements`: Optional requirements file to install from
- `project`: Optional project directory to run within

**Returns:**
- True if installation was successful, False otherwise


### `claude_code_command` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/cli/install/claude_code.py#L162" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
claude_code_command(server_spec: str) -> None
```


Install an MCP server in Claude Code.

**Args:**
- `server_spec`: Python file to install, optionally with \:object suffix




================================================
FILE: docs/python-sdk/fastmcp-cli-install-claude_desktop.mdx
================================================
---
title: claude_desktop
sidebarTitle: claude_desktop
---

# `fastmcp.cli.install.claude_desktop`


Claude Desktop integration for FastMCP install using Cyclopts.

## Functions

### `get_claude_config_path` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/cli/install/claude_desktop.py#L20" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
get_claude_config_path() -> Path | None
```


Get the Claude config directory based on platform.


### `install_claude_desktop` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/cli/install/claude_desktop.py#L38" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
install_claude_desktop(file: Path, server_object: str | None, name: str) -> bool
```


Install FastMCP server in Claude Desktop.

**Args:**
- `file`: Path to the server file
- `server_object`: Optional server object name (for \:object suffix)
- `name`: Name for the server in Claude's config
- `with_editable`: Optional list of directories to install in editable mode
- `with_packages`: Optional list of additional packages to install
- `env_vars`: Optional dictionary of environment variables
- `python_version`: Optional Python version to use
- `with_requirements`: Optional requirements file to install from
- `project`: Optional project directory to run within

**Returns:**
- True if installation was successful, False otherwise


### `claude_desktop_command` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/cli/install/claude_desktop.py#L133" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
claude_desktop_command(server_spec: str) -> None
```


Install an MCP server in Claude Desktop.

**Args:**
- `server_spec`: Python file to install, optionally with \:object suffix




================================================
FILE: docs/python-sdk/fastmcp-cli-install-cursor.mdx
================================================
---
title: cursor
sidebarTitle: cursor
---

# `fastmcp.cli.install.cursor`


Cursor integration for FastMCP install using Cyclopts.

## Functions

### `generate_cursor_deeplink` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/cli/install/cursor.py#L21" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
generate_cursor_deeplink(server_name: str, server_config: StdioMCPServer) -> str
```


Generate a Cursor deeplink for installing the MCP server.

**Args:**
- `server_name`: Name of the server
- `server_config`: Server configuration

**Returns:**
- Deeplink URL that can be clicked to install the server


### `open_deeplink` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/cli/install/cursor.py#L45" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
open_deeplink(deeplink: str) -> bool
```


Attempt to open a deeplink URL using the system's default handler.

**Args:**
- `deeplink`: The deeplink URL to open

**Returns:**
- True if the command succeeded, False otherwise


### `install_cursor_workspace` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/cli/install/cursor.py#L68" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
install_cursor_workspace(file: Path, server_object: str | None, name: str, workspace_path: Path) -> bool
```


Install FastMCP server to workspace-specific Cursor configuration.

**Args:**
- `file`: Path to the server file
- `server_object`: Optional server object name (for \:object suffix)
- `name`: Name for the server in Cursor
- `workspace_path`: Path to the workspace directory
- `with_editable`: Optional list of directories to install in editable mode
- `with_packages`: Optional list of additional packages to install
- `env_vars`: Optional dictionary of environment variables
- `python_version`: Optional Python version to use
- `with_requirements`: Optional requirements file to install from
- `project`: Optional project directory to run within

**Returns:**
- True if installation was successful, False otherwise


### `install_cursor` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/cli/install/cursor.py#L157" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
install_cursor(file: Path, server_object: str | None, name: str) -> bool
```


Install FastMCP server in Cursor.

**Args:**
- `file`: Path to the server file
- `server_object`: Optional server object name (for \:object suffix)
- `name`: Name for the server in Cursor
- `with_editable`: Optional list of directories to install in editable mode
- `with_packages`: Optional list of additional packages to install
- `env_vars`: Optional dictionary of environment variables
- `python_version`: Optional Python version to use
- `with_requirements`: Optional requirements file to install from
- `project`: Optional project directory to run within
- `workspace`: Optional workspace directory for project-specific installation

**Returns:**
- True if installation was successful, False otherwise


### `cursor_command` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/cli/install/cursor.py#L250" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
cursor_command(server_spec: str) -> None
```


Install an MCP server in Cursor.

**Args:**
- `server_spec`: Python file to install, optionally with \:object suffix




================================================
FILE: docs/python-sdk/fastmcp-cli-install-mcp_json.mdx
================================================
---
title: mcp_json
sidebarTitle: mcp_json
---

# `fastmcp.cli.install.mcp_json`


MCP configuration JSON generation for FastMCP install using Cyclopts.

## Functions

### `install_mcp_json` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/cli/install/mcp_json.py#L20" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
install_mcp_json(file: Path, server_object: str | None, name: str) -> bool
```


Generate MCP configuration JSON for manual installation.

**Args:**
- `file`: Path to the server file
- `server_object`: Optional server object name (for \:object suffix)
- `name`: Name for the server in MCP config
- `with_editable`: Optional list of directories to install in editable mode
- `with_packages`: Optional list of additional packages to install
- `env_vars`: Optional dictionary of environment variables
- `copy`: If True, copy to clipboard instead of printing to stdout
- `python_version`: Optional Python version to use
- `with_requirements`: Optional requirements file to install from
- `project`: Optional project directory to run within

**Returns:**
- True if generation was successful, False otherwise


### `mcp_json_command` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/cli/install/mcp_json.py#L106" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
mcp_json_command(server_spec: str) -> None
```


Generate MCP configuration JSON for manual installation.

**Args:**
- `server_spec`: Python file to install, optionally with \:object suffix




================================================
FILE: docs/python-sdk/fastmcp-cli-install-shared.mdx
================================================
---
title: shared
sidebarTitle: shared
---

# `fastmcp.cli.install.shared`


Shared utilities for install commands.

## Functions

### `parse_env_var` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/cli/install/shared.py#L18" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
parse_env_var(env_var: str) -> tuple[str, str]
```


Parse environment variable string in format KEY=VALUE.


### `process_common_args` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/cli/install/shared.py#L29" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
process_common_args(server_spec: str, server_name: str | None, with_packages: list[str] | None, env_vars: list[str] | None, env_file: Path | None) -> tuple[Path, str | None, str, list[str], dict[str, str] | None]
```


Process common arguments shared by all install commands.

Handles both fastmcp.json config files and traditional file.py:object syntax.




================================================
FILE: docs/python-sdk/fastmcp-cli-run.mdx
================================================
---
title: run
sidebarTitle: run
---

# `fastmcp.cli.run`


FastMCP run command implementation with enhanced type hints.

## Functions

### `is_url` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/cli/run.py#L28" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
is_url(path: str) -> bool
```


Check if a string is a URL.


### `run_with_uv` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/cli/run.py#L34" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
run_with_uv(server_spec: str, python_version: str | None = None, with_packages: list[str] | None = None, with_requirements: Path | None = None, project: Path | None = None, transport: TransportType | None = None, host: str | None = None, port: int | None = None, path: str | None = None, log_level: LogLevelType | None = None, show_banner: bool = True, editable: str | list[str] | None = None) -> None
```


Run a MCP server using uv run subprocess.

This function is called when we need to set up a Python environment with specific
dependencies before running the server. The config parsing and merging should already
be done by the caller.

**Args:**
- `server_spec`: Python file, object specification (file\:obj), config file, or URL
- `python_version`: Python version to use (e.g. "3.10")
- `with_packages`: Additional packages to install
- `with_requirements`: Requirements file to use
- `project`: Run the command within the given project directory
- `transport`: Transport protocol to use
- `host`: Host to bind to when using http transport
- `port`: Port to bind to when using http transport
- `path`: Path to bind to when using http transport
- `log_level`: Log level
- `show_banner`: Whether to show the server banner
- `editable`: Editable package paths


### `create_client_server` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/cli/run.py#L115" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
create_client_server(url: str) -> Any
```


Create a FastMCP server from a client URL.

**Args:**
- `url`: The URL to connect to

**Returns:**
- A FastMCP server instance


### `create_mcp_config_server` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/cli/run.py#L135" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
create_mcp_config_server(mcp_config_path: Path) -> FastMCP[None]
```


Create a FastMCP server from a MCPConfig.


### `load_mcp_server_config` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/cli/run.py#L146" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
load_mcp_server_config(config_path: Path) -> MCPServerConfig
```


Load a FastMCP configuration from a fastmcp.json file.

**Args:**
- `config_path`: Path to fastmcp.json file

**Returns:**
- MCPServerConfig object


### `run_command` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/cli/run.py#L163" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
run_command(server_spec: str, transport: TransportType | None = None, host: str | None = None, port: int | None = None, path: str | None = None, log_level: LogLevelType | None = None, server_args: list[str] | None = None, show_banner: bool = True, use_direct_import: bool = False, skip_source: bool = False) -> None
```


Run a MCP server or connect to a remote one.

**Args:**
- `server_spec`: Python file, object specification (file\:obj), config file, or URL
- `transport`: Transport protocol to use
- `host`: Host to bind to when using http transport
- `port`: Port to bind to when using http transport
- `path`: Path to bind to when using http transport
- `log_level`: Log level
- `server_args`: Additional arguments to pass to the server
- `show_banner`: Whether to show the server banner
- `use_direct_import`: Whether to use direct import instead of subprocess
- `skip_source`: Whether to skip source preparation step


### `run_v1_server` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/cli/run.py#L284" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
run_v1_server(server: FastMCP1x, host: str | None = None, port: int | None = None, transport: TransportType | None = None) -> None
```



================================================
FILE: docs/python-sdk/fastmcp-client-__init__.mdx
================================================
---
title: __init__
sidebarTitle: __init__
---

# `fastmcp.client`

*This module is empty or contains only private/internal implementations.*



================================================
FILE: docs/python-sdk/fastmcp-client-auth-__init__.mdx
================================================
---
title: __init__
sidebarTitle: __init__
---

# `fastmcp.client.auth`

*This module is empty or contains only private/internal implementations.*



================================================
FILE: docs/python-sdk/fastmcp-client-auth-bearer.mdx
================================================
---
title: bearer
sidebarTitle: bearer
---

# `fastmcp.client.auth.bearer`

## Classes

### `BearerAuth` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/auth/bearer.py#L11" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

**Methods:**

#### `auth_flow` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/auth/bearer.py#L15" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
auth_flow(self, request)
```



================================================
FILE: docs/python-sdk/fastmcp-client-auth-oauth.mdx
================================================
---
title: oauth
sidebarTitle: oauth
---

# `fastmcp.client.auth.oauth`

## Functions

### `default_cache_dir` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/auth/oauth.py#L55" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
default_cache_dir() -> Path
```

### `check_if_auth_required` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/auth/oauth.py#L199" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
check_if_auth_required(mcp_url: str, httpx_kwargs: dict[str, Any] | None = None) -> bool
```


Check if the MCP endpoint requires authentication by making a test request.

**Returns:**
- True if auth appears to be required, False otherwise


## Classes

### `ClientNotFoundError` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/auth/oauth.py#L38" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Raised when OAuth client credentials are not found on the server.


### `StoredToken` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/auth/oauth.py#L44" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Token storage format with absolute expiry time.


### `FileTokenStorage` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/auth/oauth.py#L59" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


File-based token storage implementation for OAuth credentials and tokens.
Implements the mcp.client.auth.TokenStorage protocol.

Each instance is tied to a specific server URL for proper token isolation.


**Methods:**

#### `get_base_url` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/auth/oauth.py#L74" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
get_base_url(url: str) -> str
```

Extract the base URL (scheme + host) from a URL.


#### `get_cache_key` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/auth/oauth.py#L79" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
get_cache_key(self) -> str
```

Generate a safe filesystem key from the server's base URL.


#### `get_tokens` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/auth/oauth.py#L94" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
get_tokens(self) -> OAuthToken | None
```

Load tokens from file storage.


#### `set_tokens` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/auth/oauth.py#L126" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
set_tokens(self, tokens: OAuthToken) -> None
```

Save tokens to file storage.


#### `get_client_info` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/auth/oauth.py#L143" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
get_client_info(self) -> OAuthClientInformationFull | None
```

Load client information from file storage.


#### `set_client_info` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/auth/oauth.py#L171" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
set_client_info(self, client_info: OAuthClientInformationFull) -> None
```

Save client information to file storage.


#### `clear` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/auth/oauth.py#L177" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
clear(self) -> None
```

Clear all cached data for this server.


#### `clear_all` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/auth/oauth.py#L186" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
clear_all(cls, cache_dir: Path | None = None) -> None
```

Clear all cached data for all servers.


### `OAuth` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/auth/oauth.py#L229" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


OAuth client provider for MCP servers with browser-based authentication.

This class provides OAuth authentication for FastMCP clients by opening
a browser for user authorization and running a local callback server.


**Methods:**

#### `redirect_handler` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/auth/oauth.py#L309" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
redirect_handler(self, authorization_url: str) -> None
```

Open browser for authorization, with pre-flight check for invalid client.


#### `callback_handler` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/auth/oauth.py#L330" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
callback_handler(self) -> tuple[str, str | None]
```

Handle OAuth callback and return (auth_code, state).


#### `async_auth_flow` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/auth/oauth.py#L363" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
async_auth_flow(self, request: httpx.Request) -> AsyncGenerator[httpx.Request, httpx.Response]
```

HTTPX auth flow with automatic retry on stale cached credentials.

If the OAuth flow fails due to invalid/stale client credentials,
clears the cache and retries once with fresh registration.




================================================
FILE: docs/python-sdk/fastmcp-client-client.mdx
================================================
---
title: client
sidebarTitle: client
---

# `fastmcp.client.client`

## Classes

### `ClientSessionState` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/client.py#L81" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Holds all session-related state for a Client instance.

This allows clean separation of configuration (which is copied) from
session state (which should be fresh for each new client instance).


### `Client` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/client.py#L97" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


MCP client that delegates connection management to a Transport instance.

The Client class is responsible for MCP protocol logic, while the Transport
handles connection establishment and management. Client provides methods for
working with resources, prompts, tools and other MCP capabilities.

This client supports reentrant context managers (multiple concurrent
`async with client:` blocks) using reference counting and background session
management. This allows efficient session reuse in any scenario with
nested or concurrent client usage.

MCP SDK 1.10 introduced automatic list_tools() calls during call_tool()
execution. This created a race condition where events could be reset while
other tasks were waiting on them, causing deadlocks. The issue was exposed
in proxy scenarios but affects any reentrant usage.

The solution uses reference counting to track active context managers,
a background task to manage the session lifecycle, events to coordinate
between tasks, and ensures all session state changes happen within a lock.
Events are only created when needed, never reset outside locks.

This design prevents race conditions where tasks wait on events that get
replaced by other tasks, ensuring reliable coordination in concurrent scenarios.

**Args:**
- `transport`:
  Connection source specification, which can be\:

    - ClientTransport\: Direct transport instance
    - FastMCP\: In-process FastMCP server
    - AnyUrl or str\: URL to connect to
    - Path\: File path for local socket
    - MCPConfig\: MCP server configuration
    - dict\: Transport configuration
- `roots`: Optional RootsList or RootsHandler for filesystem access
- `sampling_handler`: Optional handler for sampling requests
- `log_handler`: Optional handler for log messages
- `message_handler`: Optional handler for protocol messages
- `progress_handler`: Optional handler for progress notifications
- `timeout`: Optional timeout for requests (seconds or timedelta)
- `init_timeout`: Optional timeout for initial connection (seconds or timedelta).
  Set to 0 to disable. If None, uses the value in the FastMCP global settings.

**Examples:**

```python
# Connect to FastMCP server
client = Client("http://localhost:8080")

async with client:
    # List available resources
    resources = await client.list_resources()

    # Call a tool
    result = await client.call_tool("my_tool", {"param": "value"})
```


**Methods:**

#### `session` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/client.py#L283" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
session(self) -> ClientSession
```

Get the current active session. Raises RuntimeError if not connected.


#### `initialize_result` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/client.py#L293" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
initialize_result(self) -> mcp.types.InitializeResult
```

Get the result of the initialization request.


#### `set_roots` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/client.py#L301" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
set_roots(self, roots: RootsList | RootsHandler) -> None
```

Set the roots for the client. This does not automatically call `send_roots_list_changed`.


#### `set_sampling_callback` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/client.py#L305" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
set_sampling_callback(self, sampling_callback: ClientSamplingHandler) -> None
```

Set the sampling callback for the client.


#### `set_elicitation_callback` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/client.py#L311" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
set_elicitation_callback(self, elicitation_callback: ElicitationHandler) -> None
```

Set the elicitation callback for the client.


#### `is_connected` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/client.py#L319" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
is_connected(self) -> bool
```

Check if the client is currently connected.


#### `new` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/client.py#L323" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
new(self) -> Client[ClientTransportT]
```

Create a new client instance with the same configuration but fresh session state.

This creates a new client with the same transport, handlers, and configuration,
but with no active session. Useful for creating independent sessions that don't
share state with the original client.

**Returns:**
- A new Client instance with the same configuration but disconnected state.


#### `close` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/client.py#L490" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
close(self)
```

#### `ping` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/client.py#L496" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
ping(self) -> bool
```

Send a ping request.


#### `cancel` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/client.py#L501" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
cancel(self, request_id: str | int, reason: str | None = None) -> None
```

Send a cancellation notification for an in-progress request.


#### `progress` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/client.py#L518" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
progress(self, progress_token: str | int, progress: float, total: float | None = None, message: str | None = None) -> None
```

Send a progress notification.


#### `set_logging_level` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/client.py#L530" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
set_logging_level(self, level: mcp.types.LoggingLevel) -> None
```

Send a logging/setLevel request.


#### `send_roots_list_changed` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/client.py#L534" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
send_roots_list_changed(self) -> None
```

Send a roots/list_changed notification.


#### `list_resources_mcp` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/client.py#L540" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
list_resources_mcp(self) -> mcp.types.ListResourcesResult
```

Send a resources/list request and return the complete MCP protocol result.

**Returns:**
- mcp.types.ListResourcesResult: The complete response object from the protocol,
  containing the list of resources and any additional metadata.

**Raises:**
- `RuntimeError`: If called while the client is not connected.


#### `list_resources` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/client.py#L555" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
list_resources(self) -> list[mcp.types.Resource]
```

Retrieve a list of resources available on the server.

**Returns:**
- list\[mcp.types.Resource]: A list of Resource objects.

**Raises:**
- `RuntimeError`: If called while the client is not connected.


#### `list_resource_templates_mcp` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/client.py#L567" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
list_resource_templates_mcp(self) -> mcp.types.ListResourceTemplatesResult
```

Send a resources/listResourceTemplates request and return the complete MCP protocol result.

**Returns:**
- mcp.types.ListResourceTemplatesResult: The complete response object from the protocol,
  containing the list of resource templates and any additional metadata.

**Raises:**
- `RuntimeError`: If called while the client is not connected.


#### `list_resource_templates` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/client.py#L584" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
list_resource_templates(self) -> list[mcp.types.ResourceTemplate]
```

Retrieve a list of resource templates available on the server.

**Returns:**
- list\[mcp.types.ResourceTemplate]: A list of ResourceTemplate objects.

**Raises:**
- `RuntimeError`: If called while the client is not connected.


#### `read_resource_mcp` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/client.py#L598" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
read_resource_mcp(self, uri: AnyUrl | str) -> mcp.types.ReadResourceResult
```

Send a resources/read request and return the complete MCP protocol result.

**Args:**
- `uri`: The URI of the resource to read. Can be a string or an AnyUrl object.

**Returns:**
- mcp.types.ReadResourceResult: The complete response object from the protocol,
  containing the resource contents and any additional metadata.

**Raises:**
- `RuntimeError`: If called while the client is not connected.


#### `read_resource` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/client.py#L620" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
read_resource(self, uri: AnyUrl | str) -> list[mcp.types.TextResourceContents | mcp.types.BlobResourceContents]
```

Read the contents of a resource or resolved template.

**Args:**
- `uri`: The URI of the resource to read. Can be a string or an AnyUrl object.

**Returns:**
- list\[mcp.types.TextResourceContents | mcp.types.BlobResourceContents]: A list of content
  objects, typically containing either text or binary data.

**Raises:**
- `RuntimeError`: If called while the client is not connected.


#### `list_prompts_mcp` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/client.py#L659" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
list_prompts_mcp(self) -> mcp.types.ListPromptsResult
```

Send a prompts/list request and return the complete MCP protocol result.

**Returns:**
- mcp.types.ListPromptsResult: The complete response object from the protocol,
  containing the list of prompts and any additional metadata.

**Raises:**
- `RuntimeError`: If called while the client is not connected.


#### `list_prompts` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/client.py#L674" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
list_prompts(self) -> list[mcp.types.Prompt]
```

Retrieve a list of prompts available on the server.

**Returns:**
- list\[mcp.types.Prompt]: A list of Prompt objects.

**Raises:**
- `RuntimeError`: If called while the client is not connected.


#### `get_prompt_mcp` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/client.py#L687" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
get_prompt_mcp(self, name: str, arguments: dict[str, Any] | None = None) -> mcp.types.GetPromptResult
```

Send a prompts/get request and return the complete MCP protocol result.

**Args:**
- `name`: The name of the prompt to retrieve.
- `arguments`: Arguments to pass to the prompt. Defaults to None.

**Returns:**
- mcp.types.GetPromptResult: The complete response object from the protocol,
  containing the prompt messages and any additional metadata.

**Raises:**
- `RuntimeError`: If called while the client is not connected.


#### `get_prompt` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/client.py#L723" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
get_prompt(self, name: str, arguments: dict[str, Any] | None = None) -> mcp.types.GetPromptResult
```

Retrieve a rendered prompt message list from the server.

**Args:**
- `name`: The name of the prompt to retrieve.
- `arguments`: Arguments to pass to the prompt. Defaults to None.

**Returns:**
- mcp.types.GetPromptResult: The complete response object from the protocol,
  containing the prompt messages and any additional metadata.

**Raises:**
- `RuntimeError`: If called while the client is not connected.


#### `complete_mcp` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/client.py#L744" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
complete_mcp(self, ref: mcp.types.ResourceReference | mcp.types.PromptReference, argument: dict[str, str]) -> mcp.types.CompleteResult
```

Send a completion request and return the complete MCP protocol result.

**Args:**
- `ref`: The reference to complete.
- `argument`: Arguments to pass to the completion request.

**Returns:**
- mcp.types.CompleteResult: The complete response object from the protocol,
  containing the completion and any additional metadata.

**Raises:**
- `RuntimeError`: If called while the client is not connected.


#### `complete` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/client.py#L767" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
complete(self, ref: mcp.types.ResourceReference | mcp.types.PromptReference, argument: dict[str, str]) -> mcp.types.Completion
```

Send a completion request to the server.

**Args:**
- `ref`: The reference to complete.
- `argument`: Arguments to pass to the completion request.

**Returns:**
- mcp.types.Completion: The completion object.

**Raises:**
- `RuntimeError`: If called while the client is not connected.


#### `list_tools_mcp` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/client.py#L789" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
list_tools_mcp(self) -> mcp.types.ListToolsResult
```

Send a tools/list request and return the complete MCP protocol result.

**Returns:**
- mcp.types.ListToolsResult: The complete response object from the protocol,
  containing the list of tools and any additional metadata.

**Raises:**
- `RuntimeError`: If called while the client is not connected.


#### `list_tools` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/client.py#L804" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
list_tools(self) -> list[mcp.types.Tool]
```

Retrieve a list of tools available on the server.

**Returns:**
- list\[mcp.types.Tool]: A list of Tool objects.

**Raises:**
- `RuntimeError`: If called while the client is not connected.


#### `call_tool_mcp` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/client.py#L818" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
call_tool_mcp(self, name: str, arguments: dict[str, Any], progress_handler: ProgressHandler | None = None, timeout: datetime.timedelta | float | int | None = None) -> mcp.types.CallToolResult
```

Send a tools/call request and return the complete MCP protocol result.

This method returns the raw CallToolResult object, which includes an isError flag
and other metadata. It does not raise an exception if the tool call results in an error.

**Args:**
- `name`: The name of the tool to call.
- `arguments`: Arguments to pass to the tool.
- `timeout`: The timeout for the tool call. Defaults to None.
- `progress_handler`: The progress handler to use for the tool call. Defaults to None.

**Returns:**
- mcp.types.CallToolResult: The complete response object from the protocol,
  containing the tool result and any additional metadata.

**Raises:**
- `RuntimeError`: If called while the client is not connected.


#### `call_tool` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/client.py#L855" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
call_tool(self, name: str, arguments: dict[str, Any] | None = None, timeout: datetime.timedelta | float | int | None = None, progress_handler: ProgressHandler | None = None, raise_on_error: bool = True) -> CallToolResult
```

Call a tool on the server.

Unlike call_tool_mcp, this method raises a ToolError if the tool call results in an error.

**Args:**
- `name`: The name of the tool to call.
- `arguments`: Arguments to pass to the tool. Defaults to None.
- `timeout`: The timeout for the tool call. Defaults to None.
- `progress_handler`: The progress handler to use for the tool call. Defaults to None.

**Returns:**
- 
The content returned by the tool. If the tool returns structured
outputs, they are returned as a dataclass (if an output schema
is available) or a dictionary; otherwise, a list of content
blocks is returned. Note: to receive both structured and
unstructured outputs, use call_tool_mcp instead and access the
raw result object.

**Raises:**
- `ToolError`: If the tool call results in an error.
- `RuntimeError`: If called while the client is not connected.


#### `generate_name` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/client.py#L926" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
generate_name(cls, name: str | None = None) -> str
```

### `CallToolResult` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/client.py#L935" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>



================================================
FILE: docs/python-sdk/fastmcp-client-elicitation.mdx
================================================
---
title: elicitation
sidebarTitle: elicitation
---

# `fastmcp.client.elicitation`

## Functions

### `create_elicitation_callback` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/elicitation.py#L37" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
create_elicitation_callback(elicitation_handler: ElicitationHandler) -> ElicitationFnT
```

## Classes

### `ElicitResult` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/elicitation.py#L22" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>



================================================
FILE: docs/python-sdk/fastmcp-client-logging.mdx
================================================
---
title: logging
sidebarTitle: logging
---

# `fastmcp.client.logging`

## Functions

### `default_log_handler` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/logging.py#L15" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
default_log_handler(message: LogMessage) -> None
```


Default handler that properly routes server log messages to appropriate log levels.


### `create_log_callback` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/logging.py#L43" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
create_log_callback(handler: LogHandler | None = None) -> LoggingFnT
```



================================================
FILE: docs/python-sdk/fastmcp-client-messages.mdx
================================================
---
title: messages
sidebarTitle: messages
---

# `fastmcp.client.messages`

## Classes

### `MessageHandler` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/messages.py#L16" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


This class is used to handle MCP messages sent to the client. It is used to handle all messages,
requests, notifications, and exceptions. Users can override any of the hooks


**Methods:**

#### `dispatch` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/messages.py#L30" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
dispatch(self, message: Message) -> None
```

#### `on_message` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/messages.py#L74" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
on_message(self, message: Message) -> None
```

#### `on_request` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/messages.py#L77" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
on_request(self, message: RequestResponder[mcp.types.ServerRequest, mcp.types.ClientResult]) -> None
```

#### `on_ping` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/messages.py#L82" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
on_ping(self, message: mcp.types.PingRequest) -> None
```

#### `on_list_roots` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/messages.py#L85" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
on_list_roots(self, message: mcp.types.ListRootsRequest) -> None
```

#### `on_create_message` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/messages.py#L88" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
on_create_message(self, message: mcp.types.CreateMessageRequest) -> None
```

#### `on_notification` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/messages.py#L91" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
on_notification(self, message: mcp.types.ServerNotification) -> None
```

#### `on_exception` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/messages.py#L94" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
on_exception(self, message: Exception) -> None
```

#### `on_progress` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/messages.py#L97" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
on_progress(self, message: mcp.types.ProgressNotification) -> None
```

#### `on_logging_message` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/messages.py#L100" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
on_logging_message(self, message: mcp.types.LoggingMessageNotification) -> None
```

#### `on_tool_list_changed` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/messages.py#L105" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
on_tool_list_changed(self, message: mcp.types.ToolListChangedNotification) -> None
```

#### `on_resource_list_changed` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/messages.py#L110" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
on_resource_list_changed(self, message: mcp.types.ResourceListChangedNotification) -> None
```

#### `on_prompt_list_changed` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/messages.py#L115" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
on_prompt_list_changed(self, message: mcp.types.PromptListChangedNotification) -> None
```

#### `on_resource_updated` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/messages.py#L120" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
on_resource_updated(self, message: mcp.types.ResourceUpdatedNotification) -> None
```

#### `on_cancelled` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/messages.py#L125" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
on_cancelled(self, message: mcp.types.CancelledNotification) -> None
```



================================================
FILE: docs/python-sdk/fastmcp-client-oauth_callback.mdx
================================================
---
title: oauth_callback
sidebarTitle: oauth_callback
---

# `fastmcp.client.oauth_callback`



OAuth callback server for handling authorization code flows.

This module provides a reusable callback server that can handle OAuth redirects
and display styled responses to users.


## Functions

### `create_callback_html` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/oauth_callback.py#L25" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
create_callback_html(message: str, is_success: bool = True, title: str = 'FastMCP OAuth', server_url: str | None = None) -> str
```


Create a styled HTML response for OAuth callbacks.


### `create_oauth_callback_server` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/oauth_callback.py#L197" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
create_oauth_callback_server(port: int, callback_path: str = '/callback', server_url: str | None = None, response_future: asyncio.Future | None = None) -> Server
```


Create an OAuth callback server.

**Args:**
- `port`: The port to run the server on
- `callback_path`: The path to listen for OAuth redirects on
- `server_url`: Optional server URL to display in success messages
- `response_future`: Optional future to resolve when OAuth callback is received

**Returns:**
- Configured uvicorn Server instance (not yet running)


## Classes

### `CallbackResponse` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/oauth_callback.py#L183" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

**Methods:**

#### `from_dict` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/oauth_callback.py#L190" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
from_dict(cls, data: dict[str, str]) -> CallbackResponse
```

#### `to_dict` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/oauth_callback.py#L193" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
to_dict(self) -> dict[str, str]
```



================================================
FILE: docs/python-sdk/fastmcp-client-progress.mdx
================================================
---
title: progress
sidebarTitle: progress
---

# `fastmcp.client.progress`

## Functions

### `default_progress_handler` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/progress.py#L12" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
default_progress_handler(progress: float, total: float | None, message: str | None) -> None
```


Default handler for progress notifications.

Logs progress updates at debug level, properly handling missing total or message values.

**Args:**
- `progress`: Current progress value
- `total`: Optional total expected value
- `message`: Optional status message




================================================
FILE: docs/python-sdk/fastmcp-client-roots.mdx
================================================
---
title: roots
sidebarTitle: roots
---

# `fastmcp.client.roots`

## Functions

### `convert_roots_list` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/roots.py#L19" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
convert_roots_list(roots: RootsList) -> list[mcp.types.Root]
```

### `create_roots_callback` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/roots.py#L33" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
create_roots_callback(handler: RootsList | RootsHandler) -> ListRootsFnT
```



================================================
FILE: docs/python-sdk/fastmcp-client-sampling.mdx
================================================
---
title: sampling
sidebarTitle: sampling
---

# `fastmcp.client.sampling`

## Functions

### `create_sampling_callback` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/sampling.py#L31" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
create_sampling_callback(sampling_handler: ClientSamplingHandler[LifespanContextT]) -> SamplingFnT
```



================================================
FILE: docs/python-sdk/fastmcp-client-transports.mdx
================================================
---
title: transports
sidebarTitle: transports
---

# `fastmcp.client.transports`

## Functions

### `infer_transport` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/transports.py#L971" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
infer_transport(transport: ClientTransport | FastMCP | FastMCP1Server | AnyUrl | Path | MCPConfig | dict[str, Any] | str) -> ClientTransport
```


Infer the appropriate transport type from the given transport argument.

This function attempts to infer the correct transport type from the provided
argument, handling various input types and converting them to the appropriate
ClientTransport subclass.

The function supports these input types:
- ClientTransport: Used directly without modification
- FastMCP or FastMCP1Server: Creates an in-memory FastMCPTransport
- Path or str (file path): Creates PythonStdioTransport (.py) or NodeStdioTransport (.js)
- AnyUrl or str (URL): Creates StreamableHttpTransport (default) or SSETransport (for /sse endpoints)
- MCPConfig or dict: Creates MCPConfigTransport, potentially connecting to multiple servers

For HTTP URLs, they are assumed to be Streamable HTTP URLs unless they end in `/sse`.

For MCPConfig with multiple servers, a composite client is created where each server
is mounted with its name as prefix. This allows accessing tools and resources from multiple
servers through a single unified client interface, using naming patterns like
`servername_toolname` for tools and `protocol://servername/path` for resources.
If the MCPConfig contains only one server, a direct connection is established without prefixing.

**Examples:**

```python
# Connect to a local Python script
transport = infer_transport("my_script.py")

# Connect to a remote server via HTTP
transport = infer_transport("http://example.com/mcp")

# Connect to multiple servers using MCPConfig
config = {
    "mcpServers": {
        "weather": {"url": "http://weather.example.com/mcp"},
        "calendar": {"url": "http://calendar.example.com/mcp"}
    }
}
transport = infer_transport(config)
```


## Classes

### `SessionKwargs` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/transports.py#L63" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Keyword arguments for the MCP ClientSession constructor.


### `ClientTransport` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/transports.py#L75" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Abstract base class for different MCP client transport mechanisms.

A Transport is responsible for establishing and managing connections
to an MCP server, and providing a ClientSession within an async context.


**Methods:**

#### `connect_session` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/transports.py#L86" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
connect_session(self, **session_kwargs: Unpack[SessionKwargs]) -> AsyncIterator[ClientSession]
```

Establishes a connection and yields an active ClientSession.

The ClientSession is *not* expected to be initialized in this context manager.

The session is guaranteed to be valid only within the scope of the
async context manager. Connection setup and teardown are handled
within this context.

**Args:**
- `**session_kwargs`: Keyword arguments to pass to the ClientSession
  constructor (e.g., callbacks, timeouts).


#### `close` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/transports.py#L112" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
close(self)
```

Close the transport.


### `WSTransport` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/transports.py#L121" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Transport implementation that connects to an MCP server via WebSockets.


**Methods:**

#### `connect_session` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/transports.py#L139" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
connect_session(self, **session_kwargs: Unpack[SessionKwargs]) -> AsyncIterator[ClientSession]
```

### `SSETransport` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/transports.py#L160" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Transport implementation that connects to an MCP server via Server-Sent Events.


**Methods:**

#### `connect_session` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/transports.py#L196" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
connect_session(self, **session_kwargs: Unpack[SessionKwargs]) -> AsyncIterator[ClientSession]
```

### `StreamableHttpTransport` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/transports.py#L230" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Transport implementation that connects to an MCP server via Streamable HTTP Requests.


**Methods:**

#### `connect_session` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/transports.py#L266" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
connect_session(self, **session_kwargs: Unpack[SessionKwargs]) -> AsyncIterator[ClientSession]
```

### `StdioTransport` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/transports.py#L301" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Base transport for connecting to an MCP server via subprocess with stdio.

This is a base class that can be subclassed for specific command-based
transports like Python, Node, Uvx, etc.


**Methods:**

#### `connect_session` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/transports.py#L344" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
connect_session(self, **session_kwargs: Unpack[SessionKwargs]) -> AsyncIterator[ClientSession]
```

#### `connect` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/transports.py#L356" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
connect(self, **session_kwargs: Unpack[SessionKwargs]) -> ClientSession | None
```

#### `disconnect` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/transports.py#L390" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
disconnect(self)
```

#### `close` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/transports.py#L405" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
close(self)
```

### `PythonStdioTransport` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/transports.py#L463" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Transport for running Python scripts.


### `FastMCPStdioTransport` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/transports.py#L509" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Transport for running FastMCP servers using the FastMCP CLI.


### `NodeStdioTransport` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/transports.py#L536" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Transport for running Node.js scripts.


### `UvStdioTransport` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/transports.py#L578" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Transport for running commands via the uv tool.


### `UvxStdioTransport` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/transports.py#L657" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Transport for running commands via the uvx tool.


### `NpxStdioTransport` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/transports.py#L721" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Transport for running commands via the npx tool.


### `FastMCPTransport` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/transports.py#L783" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


In-memory transport for FastMCP servers.

This transport connects directly to a FastMCP server instance in the same
Python process. It works with both FastMCP 2.x servers and FastMCP 1.0
servers from the low-level MCP SDK. This is particularly useful for unit
tests or scenarios where client and server run in the same runtime.


**Methods:**

#### `connect_session` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/transports.py#L802" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
connect_session(self, **session_kwargs: Unpack[SessionKwargs]) -> AsyncIterator[ClientSession]
```

### `MCPConfigTransport` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/transports.py#L837" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Transport for connecting to one or more MCP servers defined in an MCPConfig.

This transport provides a unified interface to multiple MCP servers defined in an MCPConfig
object or dictionary matching the MCPConfig schema. It supports two key scenarios:

1. If the MCPConfig contains exactly one server, it creates a direct transport to that server.
2. If the MCPConfig contains multiple servers, it creates a composite client by mounting
   all servers on a single FastMCP instance, with each server's name, by default, used as its mounting prefix.

In the multi-server case, tools are accessible with the prefix pattern `{server_name}_{tool_name}`
and resources with the pattern `protocol://{server_name}/path/to/resource`.

This is particularly useful for creating clients that need to interact with multiple specialized
MCP servers through a single interface, simplifying client code.

**Examples:**

```python
from fastmcp import Client
from fastmcp.utilities.mcp_config import MCPConfig

# Create a config with multiple servers
config = {
    "mcpServers": {
        "weather": {
            "url": "https://weather-api.example.com/mcp",
            "transport": "http"
        },
        "calendar": {
            "url": "https://calendar-api.example.com/mcp",
            "transport": "http"
        }
    }
}

# Create a client with the config
client = Client(config)

async with client:
    # Access tools with prefixes
    weather = await client.call_tool("weather_get_forecast", {"city": "London"})
    events = await client.call_tool("calendar_list_events", {"date": "2023-06-01"})

    # Access resources with prefixed URIs
    icons = await client.read_resource("weather://weather/icons/sunny")
```


**Methods:**

#### `connect_session` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/transports.py#L919" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
connect_session(self, **session_kwargs: Unpack[SessionKwargs]) -> AsyncIterator[ClientSession]
```

#### `close` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/client/transports.py#L925" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
close(self)
```



================================================
FILE: docs/python-sdk/fastmcp-exceptions.mdx
================================================
---
title: exceptions
sidebarTitle: exceptions
---

# `fastmcp.exceptions`


Custom exceptions for FastMCP.

## Classes

### `FastMCPError` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/exceptions.py#L6" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Base error for FastMCP.


### `ValidationError` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/exceptions.py#L10" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Error in validating parameters or return values.


### `ResourceError` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/exceptions.py#L14" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Error in resource operations.


### `ToolError` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/exceptions.py#L18" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Error in tool operations.


### `PromptError` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/exceptions.py#L22" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Error in prompt operations.


### `InvalidSignature` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/exceptions.py#L26" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Invalid signature for use with FastMCP.


### `ClientError` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/exceptions.py#L30" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Error in client operations.


### `NotFoundError` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/exceptions.py#L34" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Object not found.


### `DisabledError` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/exceptions.py#L38" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Object is disabled.




================================================
FILE: docs/python-sdk/fastmcp-mcp_config.mdx
================================================
---
title: mcp_config
sidebarTitle: mcp_config
---

# `fastmcp.mcp_config`


Canonical MCP Configuration Format.

This module defines the standard configuration format for Model Context Protocol (MCP) servers.
It provides a client-agnostic, extensible format that can be used across all MCP implementations.

The configuration format supports both stdio and remote (HTTP/SSE) transports, with comprehensive
field definitions for server metadata, authentication, and execution parameters.

Example configuration:
```json
{
    "mcpServers": {
        "my-server": {
            "command": "npx",
            "args": ["-y", "@my/mcp-server"],
            "env": {"API_KEY": "secret"},
            "timeout": 30000,
            "description": "My MCP server"
        }
    }
}
```


## Functions

### `infer_transport_type_from_url` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/mcp_config.py#L56" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
infer_transport_type_from_url(url: str | AnyUrl) -> Literal['http', 'sse']
```


Infer the appropriate transport type from the given URL.


### `update_config_file` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/mcp_config.py#L313" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
update_config_file(file_path: Path, server_name: str, server_config: CanonicalMCPServerTypes) -> None
```


Update an MCP configuration file from a server object, preserving existing fields.

This is used for updating the mcpServer configurations of third-party tools so we do not
worry about transforming server objects here.


## Classes

### `StdioMCPServer` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/mcp_config.py#L126" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


MCP server configuration for stdio transport.

This is the canonical configuration format for MCP servers using stdio transport.


**Methods:**

#### `to_transport` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/mcp_config.py#L156" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
to_transport(self) -> StdioTransport
```

### `TransformingStdioMCPServer` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/mcp_config.py#L167" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


A Stdio server with tool transforms.


### `RemoteMCPServer` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/mcp_config.py#L171" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


MCP server configuration for HTTP/SSE transport.

This is the canonical configuration format for MCP servers using remote transports.


**Methods:**

#### `to_transport` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/mcp_config.py#L207" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
to_transport(self) -> StreamableHttpTransport | SSETransport
```

### `TransformingRemoteMCPServer` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/mcp_config.py#L232" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


A Remote server with tool transforms.


### `MCPConfig` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/mcp_config.py#L243" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


A configuration object for MCP Servers that conforms to the canonical MCP configuration format
while adding additional fields for enabling FastMCP-specific features like tool transformations
and filtering by tags.

For an MCPConfig that is strictly canonical, see the `CanonicalMCPConfig` class.


**Methods:**

#### `wrap_servers_at_root` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/mcp_config.py#L257" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
wrap_servers_at_root(cls, values: dict[str, Any]) -> dict[str, Any]
```

If there's no mcpServers key but there are server configs at root, wrap them.


#### `add_server` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/mcp_config.py#L270" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
add_server(self, name: str, server: MCPServerTypes) -> None
```

Add or update a server in the configuration.


#### `from_dict` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/mcp_config.py#L275" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
from_dict(cls, config: dict[str, Any]) -> Self
```

Parse MCP configuration from dictionary format.


#### `to_dict` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/mcp_config.py#L279" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
to_dict(self) -> dict[str, Any]
```

Convert MCPConfig to dictionary format, preserving all fields.


#### `write_to_file` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/mcp_config.py#L283" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
write_to_file(self, file_path: Path) -> None
```

Write configuration to JSON file.


#### `from_file` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/mcp_config.py#L289" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
from_file(cls, file_path: Path) -> Self
```

Load configuration from JSON file.


### `CanonicalMCPConfig` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/mcp_config.py#L298" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Canonical MCP configuration format.

This defines the standard configuration format for Model Context Protocol servers.
The format is designed to be client-agnostic and extensible for future use cases.


**Methods:**

#### `add_server` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/mcp_config.py#L308" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
add_server(self, name: str, server: CanonicalMCPServerTypes) -> None
```

Add or update a server in the configuration.




================================================
FILE: docs/python-sdk/fastmcp-prompts-__init__.mdx
================================================
---
title: __init__
sidebarTitle: __init__
---

# `fastmcp.prompts`

*This module is empty or contains only private/internal implementations.*



================================================
FILE: docs/python-sdk/fastmcp-prompts-prompt.mdx
================================================
---
title: prompt
sidebarTitle: prompt
---

# `fastmcp.prompts.prompt`


Base classes for FastMCP prompts.

## Functions

### `Message` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/prompts/prompt.py#L31" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
Message(content: str | ContentBlock, role: Role | None = None, **kwargs: Any) -> PromptMessage
```


A user-friendly constructor for PromptMessage.


## Classes

### `PromptArgument` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/prompts/prompt.py#L53" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


An argument that can be passed to a prompt.


### `Prompt` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/prompts/prompt.py#L65" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


A prompt template that can be rendered with parameters.


**Methods:**

#### `enable` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/prompts/prompt.py#L72" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
enable(self) -> None
```

#### `disable` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/prompts/prompt.py#L80" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
disable(self) -> None
```

#### `to_mcp_prompt` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/prompts/prompt.py#L88" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
to_mcp_prompt(self, **overrides: Any) -> MCPPrompt
```

Convert the prompt to an MCP prompt.


#### `from_function` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/prompts/prompt.py#L113" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
from_function(fn: Callable[..., PromptResult | Awaitable[PromptResult]], name: str | None = None, title: str | None = None, description: str | None = None, tags: set[str] | None = None, enabled: bool | None = None, meta: dict[str, Any] | None = None) -> FunctionPrompt
```

Create a Prompt from a function.

The function can return:
- A string (converted to a message)
- A Message object
- A dict (converted to a message)
- A sequence of any of the above


#### `render` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/prompts/prompt.py#L141" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
render(self, arguments: dict[str, Any] | None = None) -> list[PromptMessage]
```

Render the prompt with arguments.


### `FunctionPrompt` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/prompts/prompt.py#L149" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


A prompt that is a function.


**Methods:**

#### `from_function` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/prompts/prompt.py#L155" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
from_function(cls, fn: Callable[..., PromptResult | Awaitable[PromptResult]], name: str | None = None, title: str | None = None, description: str | None = None, tags: set[str] | None = None, enabled: bool | None = None, meta: dict[str, Any] | None = None) -> FunctionPrompt
```

Create a Prompt from a function.

The function can return:
- A string (converted to a message)
- A Message object
- A dict (converted to a message)
- A sequence of any of the above


#### `render` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/prompts/prompt.py#L315" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
render(self, arguments: dict[str, Any] | None = None) -> list[PromptMessage]
```

Render the prompt with arguments.




================================================
FILE: docs/python-sdk/fastmcp-prompts-prompt_manager.mdx
================================================
---
title: prompt_manager
sidebarTitle: prompt_manager
---

# `fastmcp.prompts.prompt_manager`

## Classes

### `PromptManager` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/prompts/prompt_manager.py#L21" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Manages FastMCP prompts.


**Methods:**

#### `mount` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/prompts/prompt_manager.py#L45" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
mount(self, server: MountedServer) -> None
```

Adds a mounted server as a source for prompts.


#### `has_prompt` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/prompts/prompt_manager.py#L91" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
has_prompt(self, key: str) -> bool
```

Check if a prompt exists.


#### `get_prompt` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/prompts/prompt_manager.py#L96" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
get_prompt(self, key: str) -> Prompt
```

Get prompt by key.


#### `get_prompts` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/prompts/prompt_manager.py#L103" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
get_prompts(self) -> dict[str, Prompt]
```

Gets the complete, unfiltered inventory of all prompts.


#### `list_prompts` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/prompts/prompt_manager.py#L109" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
list_prompts(self) -> list[Prompt]
```

Lists all prompts, applying protocol filtering.


#### `add_prompt_from_fn` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/prompts/prompt_manager.py#L116" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
add_prompt_from_fn(self, fn: Callable[..., PromptResult | Awaitable[PromptResult]], name: str | None = None, description: str | None = None, tags: set[str] | None = None) -> FunctionPrompt
```

Create a prompt from a function.


#### `add_prompt` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/prompts/prompt_manager.py#L136" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
add_prompt(self, prompt: Prompt) -> Prompt
```

Add a prompt to the manager.


#### `render_prompt` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/prompts/prompt_manager.py#L154" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
render_prompt(self, name: str, arguments: dict[str, Any] | None = None) -> GetPromptResult
```

Internal API for servers: Finds and renders a prompt, respecting the
filtered protocol path.




================================================
FILE: docs/python-sdk/fastmcp-resources-__init__.mdx
================================================
---
title: __init__
sidebarTitle: __init__
---

# `fastmcp.resources`

*This module is empty or contains only private/internal implementations.*



================================================
FILE: docs/python-sdk/fastmcp-resources-resource.mdx
================================================
---
title: resource
sidebarTitle: resource
---

# `fastmcp.resources.resource`


Base classes and interfaces for FastMCP resources.

## Classes

### `Resource` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/resource.py#L33" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Base class for all resources.


**Methods:**

#### `enable` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/resource.py#L52" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
enable(self) -> None
```

#### `disable` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/resource.py#L60" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
disable(self) -> None
```

#### `from_function` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/resource.py#L69" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
from_function(fn: Callable[..., Any], uri: str | AnyUrl, name: str | None = None, title: str | None = None, description: str | None = None, mime_type: str | None = None, tags: set[str] | None = None, enabled: bool | None = None, annotations: Annotations | None = None, meta: dict[str, Any] | None = None) -> FunctionResource
```

#### `set_default_mime_type` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/resource.py#L96" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
set_default_mime_type(cls, mime_type: str | None) -> str
```

Set default MIME type if not provided.


#### `set_default_name` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/resource.py#L103" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
set_default_name(self) -> Self
```

Set default name from URI if not provided.


#### `read` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/resource.py#L114" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
read(self) -> str | bytes
```

Read the resource content.


#### `to_mcp_resource` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/resource.py#L118" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
to_mcp_resource(self, **overrides: Any) -> MCPResource
```

Convert the resource to an MCPResource.


#### `key` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/resource.py#L140" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
key(self) -> str
```

The key of the component. This is used for internal bookkeeping
and may reflect e.g. prefixes or other identifiers. You should not depend on
keys having a certain value, as the same tool loaded from different
hierarchies of servers may have different keys.


### `FunctionResource` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/resource.py#L150" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


A resource that defers data loading by wrapping a function.

The function is only called when the resource is read, allowing for lazy loading
of potentially expensive data. This is particularly useful when listing resources,
as the function won't be called until the resource is actually accessed.

The function can return:
- str for text content (default)
- bytes for binary content
- other types will be converted to JSON


**Methods:**

#### `from_function` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/resource.py#L166" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
from_function(cls, fn: Callable[..., Any], uri: str | AnyUrl, name: str | None = None, title: str | None = None, description: str | None = None, mime_type: str | None = None, tags: set[str] | None = None, enabled: bool | None = None, annotations: Annotations | None = None, meta: dict[str, Any] | None = None) -> FunctionResource
```

Create a FunctionResource from a function.


#### `read` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/resource.py#L195" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
read(self) -> str | bytes
```

Read the resource by calling the wrapped function.




================================================
FILE: docs/python-sdk/fastmcp-resources-resource_manager.mdx
================================================
---
title: resource_manager
sidebarTitle: resource_manager
---

# `fastmcp.resources.resource_manager`


Resource manager functionality.

## Classes

### `ResourceManager` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/resource_manager.py#L28" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Manages FastMCP resources.


**Methods:**

#### `mount` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/resource_manager.py#L60" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
mount(self, server: MountedServer) -> None
```

Adds a mounted server as a source for resources and templates.


#### `get_resources` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/resource_manager.py#L64" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
get_resources(self) -> dict[str, Resource]
```

Get all registered resources, keyed by URI.


#### `get_resource_templates` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/resource_manager.py#L68" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
get_resource_templates(self) -> dict[str, ResourceTemplate]
```

Get all registered templates, keyed by URI template.


#### `list_resources` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/resource_manager.py#L178" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
list_resources(self) -> list[Resource]
```

Lists all resources, applying protocol filtering.


#### `list_resource_templates` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/resource_manager.py#L185" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
list_resource_templates(self) -> list[ResourceTemplate]
```

Lists all templates, applying protocol filtering.


#### `add_resource_or_template_from_fn` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/resource_manager.py#L192" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
add_resource_or_template_from_fn(self, fn: Callable[..., Any], uri: str, name: str | None = None, description: str | None = None, mime_type: str | None = None, tags: set[str] | None = None) -> Resource | ResourceTemplate
```

Add a resource or template to the manager from a function.

**Args:**
- `fn`: The function to register as a resource or template
- `uri`: The URI for the resource or template
- `name`: Optional name for the resource or template
- `description`: Optional description of the resource or template
- `mime_type`: Optional MIME type for the resource or template
- `tags`: Optional set of tags for categorizing the resource or template

**Returns:**
- The added resource or template. If a resource or template with the same URI already exists,
- returns the existing resource or template.


#### `add_resource_from_fn` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/resource_manager.py#L240" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
add_resource_from_fn(self, fn: Callable[..., Any], uri: str, name: str | None = None, description: str | None = None, mime_type: str | None = None, tags: set[str] | None = None) -> Resource
```

Add a resource to the manager from a function.

**Args:**
- `fn`: The function to register as a resource
- `uri`: The URI for the resource
- `name`: Optional name for the resource
- `description`: Optional description of the resource
- `mime_type`: Optional MIME type for the resource
- `tags`: Optional set of tags for categorizing the resource

**Returns:**
- The added resource. If a resource with the same URI already exists,
- returns the existing resource.


#### `add_resource` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/resource_manager.py#L280" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
add_resource(self, resource: Resource) -> Resource
```

Add a resource to the manager.

**Args:**
- `resource`: A Resource instance to add. The resource's .key attribute
  will be used as the storage key. To overwrite it, call
  Resource.model_copy(key=new_key) before calling this method.


#### `add_template_from_fn` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/resource_manager.py#L302" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
add_template_from_fn(self, fn: Callable[..., Any], uri_template: str, name: str | None = None, description: str | None = None, mime_type: str | None = None, tags: set[str] | None = None) -> ResourceTemplate
```

Create a template from a function.


#### `add_template` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/resource_manager.py#L329" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
add_template(self, template: ResourceTemplate) -> ResourceTemplate
```

Add a template to the manager.

**Args:**
- `template`: A ResourceTemplate instance to add. The template's .key attribute
  will be used as the storage key. To overwrite it, call
  ResourceTemplate.model_copy(key=new_key) before calling this method.

**Returns:**
- The added template. If a template with the same URI already exists,
- returns the existing template.


#### `has_resource` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/resource_manager.py#L355" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
has_resource(self, uri: AnyUrl | str) -> bool
```

Check if a resource exists.


#### `get_resource` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/resource_manager.py#L372" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
get_resource(self, uri: AnyUrl | str) -> Resource
```

Get resource by URI, checking concrete resources first, then templates.

**Args:**
- `uri`: The URI of the resource to get

**Raises:**
- `NotFoundError`: If no resource or template matching the URI is found.


#### `read_resource` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/resource_manager.py#L417" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
read_resource(self, uri: AnyUrl | str) -> str | bytes
```

Internal API for servers: Finds and reads a resource, respecting the
filtered protocol path.




================================================
FILE: docs/python-sdk/fastmcp-resources-template.mdx
================================================
---
title: template
sidebarTitle: template
---

# `fastmcp.resources.template`


Resource template functionality.

## Functions

### `build_regex` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/template.py#L29" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
build_regex(template: str) -> re.Pattern
```

### `match_uri_template` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/template.py#L45" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
match_uri_template(uri: str, uri_template: str) -> dict[str, str] | None
```

## Classes

### `ResourceTemplate` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/template.py#L53" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


A template for dynamically creating resources.


**Methods:**

#### `enable` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/template.py#L72" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
enable(self) -> None
```

#### `disable` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/template.py#L80" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
disable(self) -> None
```

#### `from_function` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/template.py#L89" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
from_function(fn: Callable[..., Any], uri_template: str, name: str | None = None, title: str | None = None, description: str | None = None, mime_type: str | None = None, tags: set[str] | None = None, enabled: bool | None = None, annotations: Annotations | None = None, meta: dict[str, Any] | None = None) -> FunctionResourceTemplate
```

#### `set_default_mime_type` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/template.py#L116" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
set_default_mime_type(cls, mime_type: str | None) -> str
```

Set default MIME type if not provided.


#### `matches` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/template.py#L122" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
matches(self, uri: str) -> dict[str, Any] | None
```

Check if URI matches template and extract parameters.


#### `read` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/template.py#L126" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
read(self, arguments: dict[str, Any]) -> str | bytes
```

Read the resource content.


#### `create_resource` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/template.py#L132" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
create_resource(self, uri: str, params: dict[str, Any]) -> Resource
```

Create a resource from the template with the given parameters.


#### `to_mcp_template` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/template.py#L150" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
to_mcp_template(self, **overrides: Any) -> MCPResourceTemplate
```

Convert the resource template to an MCPResourceTemplate.


#### `from_mcp_template` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/template.py#L169" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
from_mcp_template(cls, mcp_template: MCPResourceTemplate) -> ResourceTemplate
```

Creates a FastMCP ResourceTemplate from a raw MCP ResourceTemplate object.


#### `key` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/template.py#L182" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
key(self) -> str
```

The key of the component. This is used for internal bookkeeping
and may reflect e.g. prefixes or other identifiers. You should not depend on
keys having a certain value, as the same tool loaded from different
hierarchies of servers may have different keys.


### `FunctionResourceTemplate` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/template.py#L192" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


A template for dynamically creating resources.


**Methods:**

#### `read` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/template.py#L197" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
read(self, arguments: dict[str, Any]) -> str | bytes
```

Read the resource content.


#### `from_function` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/template.py#L213" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
from_function(cls, fn: Callable[..., Any], uri_template: str, name: str | None = None, title: str | None = None, description: str | None = None, mime_type: str | None = None, tags: set[str] | None = None, enabled: bool | None = None, annotations: Annotations | None = None, meta: dict[str, Any] | None = None) -> FunctionResourceTemplate
```

Create a template from a function.




================================================
FILE: docs/python-sdk/fastmcp-resources-types.mdx
================================================
---
title: types
sidebarTitle: types
---

# `fastmcp.resources.types`


Concrete resource implementations.

## Classes

### `TextResource` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/types.py#L21" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


A resource that reads from a string.


**Methods:**

#### `read` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/types.py#L26" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
read(self) -> str
```

Read the text content.


### `BinaryResource` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/types.py#L31" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


A resource that reads from bytes.


**Methods:**

#### `read` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/types.py#L36" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
read(self) -> bytes
```

Read the binary content.


### `FileResource` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/types.py#L41" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


A resource that reads from a file.

Set is_binary=True to read file as binary data instead of text.


**Methods:**

#### `validate_absolute_path` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/types.py#L59" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
validate_absolute_path(cls, path: Path) -> Path
```

Ensure path is absolute.


#### `set_binary_from_mime_type` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/types.py#L67" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
set_binary_from_mime_type(cls, is_binary: bool, info: ValidationInfo) -> bool
```

Set is_binary based on mime_type if not explicitly set.


#### `read` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/types.py#L74" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
read(self) -> str | bytes
```

Read the file content.


### `HttpResource` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/types.py#L84" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


A resource that reads from an HTTP endpoint.


**Methods:**

#### `read` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/types.py#L92" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
read(self) -> str | bytes
```

Read the HTTP content.


### `DirectoryResource` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/types.py#L100" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


A resource that lists files in a directory.


**Methods:**

#### `validate_absolute_path` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/types.py#L116" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
validate_absolute_path(cls, path: Path) -> Path
```

Ensure path is absolute.


#### `list_files` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/types.py#L122" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
list_files(self) -> list[Path]
```

List files in the directory.


#### `read` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/resources/types.py#L144" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
read(self) -> str
```

Read the directory listing.




================================================
FILE: docs/python-sdk/fastmcp-server-__init__.mdx
================================================
---
title: __init__
sidebarTitle: __init__
---

# `fastmcp.server`

*This module is empty or contains only private/internal implementations.*



================================================
FILE: docs/python-sdk/fastmcp-server-auth-__init__.mdx
================================================
---
title: __init__
sidebarTitle: __init__
---

# `fastmcp.server.auth`

*This module is empty or contains only private/internal implementations.*



================================================
FILE: docs/python-sdk/fastmcp-server-auth-auth.mdx
================================================
---
title: auth
sidebarTitle: auth
---

# `fastmcp.server.auth.auth`

## Classes

### `AccessToken` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/auth.py#L36" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


AccessToken that includes all JWT claims.


### `AuthProvider` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/auth.py#L42" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Base class for all FastMCP authentication providers.

This class provides a unified interface for all authentication providers,
whether they are simple token verifiers or full OAuth authorization servers.
All providers must be able to verify tokens and can optionally provide
custom authentication routes.


**Methods:**

#### `verify_token` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/auth.py#L69" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
verify_token(self, token: str) -> AccessToken | None
```

Verify a bearer token and return access info if valid.

All auth providers must implement token verification.

**Args:**
- `token`: The token string to validate

**Returns:**
- AccessToken object if valid, None if invalid or expired


#### `get_routes` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/auth.py#L82" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
get_routes(self, mcp_path: str | None = None, mcp_endpoint: Any | None = None) -> list[Route]
```

Get the routes for this authentication provider.

Each provider is responsible for creating whatever routes it needs:
- TokenVerifier: typically no routes (default implementation)
- RemoteAuthProvider: protected resource metadata routes
- OAuthProvider: full OAuth authorization server routes
- Custom providers: whatever routes they need

**Args:**
- `mcp_path`: The path where the MCP endpoint is mounted (e.g., "/mcp")
- `mcp_endpoint`: The MCP endpoint handler to protect with auth

**Returns:**
- List of routes for this provider, including protected MCP endpoints if provided


#### `get_middleware` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/auth.py#L122" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
get_middleware(self) -> list
```

Get HTTP application-level middleware for this auth provider.

**Returns:**
- List of Starlette Middleware instances to apply to the HTTP app


### `TokenVerifier` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/auth.py#L154" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Base class for token verifiers (Resource Servers).

This class provides token verification capability without OAuth server functionality.
Token verifiers typically don't provide authentication routes by default.


**Methods:**

#### `verify_token` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/auth.py#L175" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
verify_token(self, token: str) -> AccessToken | None
```

Verify a bearer token and return access info if valid.


### `RemoteAuthProvider` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/auth.py#L180" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Authentication provider for resource servers that verify tokens from known authorization servers.

This provider composes a TokenVerifier with authorization server metadata to create
standardized OAuth 2.0 Protected Resource endpoints (RFC 9728). Perfect for:
- JWT verification with known issuers
- Remote token introspection services
- Any resource server that knows where its tokens come from

Use this when you have token verification logic and want to advertise
the authorization servers that issue valid tokens.


**Methods:**

#### `verify_token` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/auth.py#L221" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
verify_token(self, token: str) -> AccessToken | None
```

Verify token using the configured token verifier.


#### `get_routes` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/auth.py#L225" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
get_routes(self, mcp_path: str | None = None, mcp_endpoint: Any | None = None) -> list[Route]
```

Get OAuth routes for this provider.

Creates protected resource metadata routes and optionally wraps MCP endpoints with auth.


### `OAuthProvider` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/auth.py#L255" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


OAuth Authorization Server provider.

This class provides full OAuth server functionality including client registration,
authorization flows, token issuance, and token verification.


**Methods:**

#### `verify_token` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/auth.py#L311" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
verify_token(self, token: str) -> AccessToken | None
```

Verify a bearer token and return access info if valid.

This method implements the TokenVerifier protocol by delegating
to our existing load_access_token method.

**Args:**
- `token`: The token string to validate

**Returns:**
- AccessToken object if valid, None if invalid or expired


#### `get_routes` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/auth.py#L326" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
get_routes(self, mcp_path: str | None = None, mcp_endpoint: Any | None = None) -> list[Route]
```

Get OAuth authorization server routes and optional protected resource routes.

This method creates the full set of OAuth routes including:
- Standard OAuth authorization server routes (/.well-known/oauth-authorization-server, /authorize, /token, etc.)
- Optional protected resource routes
- Protected MCP endpoints if provided

**Returns:**
- List of OAuth routes




================================================
FILE: docs/python-sdk/fastmcp-server-auth-oauth_proxy.mdx
================================================
---
title: oauth_proxy
sidebarTitle: oauth_proxy
---

# `fastmcp.server.auth.oauth_proxy`


OAuth Proxy Provider for FastMCP.

This provider acts as a transparent proxy to an upstream OAuth Authorization Server,
handling Dynamic Client Registration locally while forwarding all other OAuth flows.
This enables authentication with upstream providers that don't support DCR or have
restricted client registration policies.

Key features:
- Proxies authorization and token endpoints to upstream server
- Implements local Dynamic Client Registration with fixed upstream credentials
- Validates tokens using upstream JWKS
- Maintains minimal local state for bookkeeping
- Enhanced logging with request correlation

This implementation is based on the OAuth 2.1 specification and is designed for
production use with enterprise identity providers.


## Classes

### `ProxyDCRClient` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/oauth_proxy.py#L58" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Client for DCR proxy with configurable redirect URI validation.

This special client class is critical for the OAuth proxy to work correctly
with Dynamic Client Registration (DCR). Here's why it exists:

Problem:
--------
When MCP clients use OAuth, they dynamically register with random localhost
ports (e.g., http://localhost:55454/callback). The OAuth proxy needs to:
1. Accept these dynamic redirect URIs from clients based on configured patterns
2. Use its own fixed redirect URI with the upstream provider (Google, GitHub, etc.)
3. Forward the authorization code back to the client's dynamic URI

Solution:
---------
This class validates redirect URIs against configurable patterns,
while the proxy internally uses its own fixed redirect URI with the upstream
provider. This allows the flow to work even when clients reconnect with
different ports or when tokens are cached.

Without proper validation, clients could get "Redirect URI not registered" errors
when trying to authenticate with cached tokens, or security vulnerabilities could
arise from accepting arbitrary redirect URIs.


**Methods:**

#### `validate_redirect_uri` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/oauth_proxy.py#L97" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
validate_redirect_uri(self, redirect_uri: AnyUrl | None) -> AnyUrl
```

Validate redirect URI against allowed patterns.

Since we're acting as a proxy and clients register dynamically,
we validate their redirect URIs against configurable patterns.
This is essential for cached token scenarios where the client may
reconnect with a different port.


### `OAuthProxy` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/oauth_proxy.py#L123" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


OAuth provider that presents a DCR-compliant interface while proxying to non-DCR IDPs.

Purpose
-------
MCP clients expect OAuth providers to support Dynamic Client Registration (DCR),
where clients can register themselves dynamically and receive unique credentials.
Most enterprise IDPs (Google, GitHub, Azure AD, etc.) don't support DCR and require
pre-registered OAuth applications with fixed credentials.

This proxy bridges that gap by:
- Presenting a full DCR-compliant OAuth interface to MCP clients
- Translating DCR registration requests to use pre-configured upstream credentials
- Proxying all OAuth flows to the upstream IDP with appropriate translations
- Managing the state and security requirements of both protocols

Architecture Overview
--------------------
The proxy maintains a single OAuth app registration with the upstream provider
while allowing unlimited MCP clients to register and authenticate dynamically.
It implements the complete OAuth 2.1 + DCR specification for clients while
translating to whatever OAuth variant the upstream provider requires.

Key Translation Challenges Solved
---------------------------------
1. Dynamic Client Registration:
    - MCP clients expect to register dynamically and get unique credentials
    - Upstream IDPs require pre-registered apps with fixed credentials
    - Solution: Accept DCR requests, return shared upstream credentials

2. Dynamic Redirect URIs:
    - MCP clients use random localhost ports that change between sessions
    - Upstream IDPs require fixed, pre-registered redirect URIs
    - Solution: Use proxy's fixed callback URL with upstream, forward to client's dynamic URI

3. Authorization Code Mapping:
    - Upstream returns codes for the proxy's redirect URI
    - Clients expect codes for their own redirect URIs
    - Solution: Exchange upstream code server-side, issue new code to client

4. State Parameter Collision:
    - Both client and proxy need to maintain state through the flow
    - Only one state parameter available in OAuth
    - Solution: Use transaction ID as state with upstream, preserve client's state

5. Token Management:
    - Clients may expect different token formats/claims than upstream provides
    - Need to track tokens for revocation and refresh
    - Solution: Store token relationships, forward upstream tokens transparently

OAuth Flow Implementation
------------------------
1. Client Registration (DCR):
    - Accept any client registration request
    - Store ProxyDCRClient that accepts dynamic redirect URIs

2. Authorization:
    - Store transaction mapping client details to proxy flow
    - Redirect to upstream with proxy's fixed redirect URI
    - Use transaction ID as state parameter with upstream

3. Upstream Callback:
    - Exchange upstream authorization code for tokens (server-side)
    - Generate new authorization code bound to client's PKCE challenge
    - Redirect to client's original dynamic redirect URI

4. Token Exchange:
    - Validate client's code and PKCE verifier
    - Return previously obtained upstream tokens
    - Clean up one-time use authorization code

5. Token Refresh:
    - Forward refresh requests to upstream using authlib
    - Handle token rotation if upstream issues new refresh token
    - Update local token mappings

State Management
---------------
The proxy maintains minimal but crucial state:
- _clients: DCR registrations (all use ProxyDCRClient for flexibility)
- _oauth_transactions: Active authorization flows with client context
- _client_codes: Authorization codes with PKCE challenges and upstream tokens
- _access_tokens, _refresh_tokens: Token storage for revocation
- Token relationship mappings for cleanup and rotation

Security Considerations
----------------------
- PKCE enforced end-to-end (client to proxy, proxy to upstream)
- Authorization codes are single-use with short expiry
- Transaction IDs are cryptographically random
- All state is cleaned up after use to prevent replay
- Token validation delegates to upstream provider

Provider Compatibility
---------------------
Works with any OAuth 2.0 provider that supports:
- Authorization code flow
- Fixed redirect URI (configured in provider's app settings)
- Standard token endpoint

Handles provider-specific requirements:
- Google: Ensures minimum scope requirements
- GitHub: Compatible with OAuth Apps and GitHub Apps
- Azure AD: Handles tenant-specific endpoints
- Generic: Works with any spec-compliant provider


**Methods:**

#### `get_client` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/oauth_proxy.py#L368" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
get_client(self, client_id: str) -> OAuthClientInformationFull | None
```

Get client information by ID. This is generally the random ID
provided to the DCR client during registration, not the upstream client ID.

For unregistered clients, returns None (which will raise an error in the SDK).


#### `register_client` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/oauth_proxy.py#L378" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
register_client(self, client_info: OAuthClientInformationFull) -> None
```

Register a client locally

When a client registers, we create a ProxyDCRClient that is more
forgiving about validating redirect URIs, since the DCR client's
redirect URI will likely be localhost or unknown to the proxied IDP. The
proxied IDP only knows about this server's fixed redirect URI.


#### `authorize` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/oauth_proxy.py#L421" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
authorize(self, client: OAuthClientInformationFull, params: AuthorizationParams) -> str
```

Start OAuth transaction and redirect to upstream IdP.

This implements the DCR-compliant proxy pattern:
1. Store transaction with client details and PKCE challenge
2. Generate proxy's own PKCE parameters if forwarding is enabled
3. Use transaction ID as state for IdP
4. Redirect to IdP with our fixed callback URL and proxy's PKCE


#### `load_authorization_code` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/oauth_proxy.py#L504" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
load_authorization_code(self, client: OAuthClientInformationFull, authorization_code: str) -> AuthorizationCode | None
```

Load authorization code for validation.

Look up our client code and return authorization code object
with PKCE challenge for validation.


#### `exchange_authorization_code` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/oauth_proxy.py#L546" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
exchange_authorization_code(self, client: OAuthClientInformationFull, authorization_code: AuthorizationCode) -> OAuthToken
```

Exchange authorization code for stored IdP tokens.

For the DCR-compliant proxy flow, we return the IdP tokens that were obtained
during the IdP callback exchange. PKCE validation is handled by the MCP framework.


#### `load_refresh_token` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/oauth_proxy.py#L613" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
load_refresh_token(self, client: OAuthClientInformationFull, refresh_token: str) -> RefreshToken | None
```

Load refresh token from local storage.


#### `exchange_refresh_token` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/oauth_proxy.py#L621" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
exchange_refresh_token(self, client: OAuthClientInformationFull, refresh_token: RefreshToken, scopes: list[str]) -> OAuthToken
```

Exchange refresh token for new access token using authlib.


#### `load_access_token` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/oauth_proxy.py#L697" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
load_access_token(self, token: str) -> AccessToken | None
```

Validate access token using upstream JWKS.

Delegates to the JWT verifier which handles signature validation,
expiration checking, and claims validation using the upstream JWKS.


#### `revoke_token` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/oauth_proxy.py#L714" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
revoke_token(self, token: AccessToken | RefreshToken) -> None
```

Revoke token locally and with upstream server if supported.

Removes tokens from local storage and attempts to revoke them with
the upstream server if a revocation endpoint is configured.


#### `get_routes` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/oauth_proxy.py#L758" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
get_routes(self, mcp_path: str | None = None, mcp_endpoint: Any | None = None) -> list[Route]
```

Get OAuth routes with custom proxy token handler.

This method creates standard OAuth routes and replaces the token endpoint
with our proxy handler that forwards requests to the upstream OAuth server.

**Args:**
- `mcp_path`: The path where the MCP endpoint is mounted (e.g., "/mcp")
- `mcp_endpoint`: The MCP endpoint handler to protect with auth




================================================
FILE: docs/python-sdk/fastmcp-server-auth-providers-__init__.mdx
================================================
---
title: __init__
sidebarTitle: __init__
---

# `fastmcp.server.auth.providers`

*This module is empty or contains only private/internal implementations.*



================================================
FILE: docs/python-sdk/fastmcp-server-auth-providers-azure.mdx
================================================
---
title: azure
sidebarTitle: azure
---

# `fastmcp.server.auth.providers.azure`


Azure (Microsoft Entra) OAuth provider for FastMCP.

This provider implements Azure/Microsoft Entra ID OAuth authentication
using the OAuth Proxy pattern for non-DCR OAuth flows.


## Classes

### `AzureProviderSettings` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/providers/azure.py#L22" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Settings for Azure OAuth provider.


### `AzureTokenVerifier` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/providers/azure.py#L46" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Token verifier for Azure OAuth tokens.

Azure tokens are JWTs, but we verify them by calling the Microsoft Graph API
to get user information and validate the token.


**Methods:**

#### `verify_token` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/providers/azure.py#L68" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
verify_token(self, token: str) -> AccessToken | None
```

Verify Azure OAuth token by calling Microsoft Graph API.


### `AzureProvider` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/providers/azure.py#L117" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Azure (Microsoft Entra) OAuth provider for FastMCP.

This provider implements Azure/Microsoft Entra ID authentication using the
OAuth Proxy pattern. It supports both organizational accounts and personal
Microsoft accounts depending on the tenant configuration.

Features:
- Transparent OAuth proxy to Azure/Microsoft identity platform
- Automatic token validation via Microsoft Graph API
- User information extraction
- Support for different tenant configurations (common, organizations, consumers)

Setup Requirements:
1. Register an application in Azure Portal (portal.azure.com)
2. Configure redirect URI as: http://localhost:8000/auth/callback
3. Note your Application (client) ID and create a client secret
4. Optionally note your Directory (tenant) ID for single-tenant apps




================================================
FILE: docs/python-sdk/fastmcp-server-auth-providers-bearer.mdx
================================================
---
title: bearer
sidebarTitle: bearer
---

# `fastmcp.server.auth.providers.bearer`


Backwards compatibility shim for BearerAuthProvider.

The BearerAuthProvider class has been moved to fastmcp.server.auth.providers.jwt.JWTVerifier
for better organization. This module provides a backwards-compatible import.




================================================
FILE: docs/python-sdk/fastmcp-server-auth-providers-github.mdx
================================================
---
title: github
sidebarTitle: github
---

# `fastmcp.server.auth.providers.github`


GitHub OAuth provider for FastMCP.

This module provides a complete GitHub OAuth integration that's ready to use
with just a client ID and client secret. It handles all the complexity of
GitHub's OAuth flow, token validation, and user management.

Example:
```python
from fastmcp import FastMCP
from fastmcp.server.auth.providers.github import GitHubProvider

    # Simple GitHub OAuth protection
    auth = GitHubProvider(
        client_id="your-github-client-id",
        client_secret="your-github-client-secret"
    )

    mcp = FastMCP("My Protected Server", auth=auth)
    ```


## Classes

### `GitHubProviderSettings` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/providers/github.py#L38" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Settings for GitHub OAuth provider.


### `GitHubTokenVerifier` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/providers/github.py#L61" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Token verifier for GitHub OAuth tokens.

GitHub OAuth tokens are opaque (not JWTs), so we verify them
by calling GitHub's API to check if they're valid and get user info.


**Methods:**

#### `verify_token` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/providers/github.py#L83" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
verify_token(self, token: str) -> AccessToken | None
```

Verify GitHub OAuth token by calling GitHub API.


### `GitHubProvider` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/providers/github.py#L166" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Complete GitHub OAuth provider for FastMCP.

This provider makes it trivial to add GitHub OAuth protection to any
FastMCP server. Just provide your GitHub OAuth app credentials and
a base URL, and you're ready to go.

Features:
- Transparent OAuth proxy to GitHub
- Automatic token validation via GitHub API
- User information extraction
- Minimal configuration required




================================================
FILE: docs/python-sdk/fastmcp-server-auth-providers-google.mdx
================================================
---
title: google
sidebarTitle: google
---

# `fastmcp.server.auth.providers.google`


Google OAuth provider for FastMCP.

This module provides a complete Google OAuth integration that's ready to use
with just a client ID and client secret. It handles all the complexity of
Google's OAuth flow, token validation, and user management.

Example:
```python
from fastmcp import FastMCP
from fastmcp.server.auth.providers.google import GoogleProvider

    # Simple Google OAuth protection
    auth = GoogleProvider(
        client_id="your-google-client-id.apps.googleusercontent.com",
        client_secret="your-google-client-secret"
    )

    mcp = FastMCP("My Protected Server", auth=auth)
    ```


## Classes

### `GoogleProviderSettings` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/providers/google.py#L40" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Settings for Google OAuth provider.


### `GoogleTokenVerifier` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/providers/google.py#L63" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Token verifier for Google OAuth tokens.

Google OAuth tokens are opaque (not JWTs), so we verify them
by calling Google's tokeninfo API to check if they're valid and get user info.


**Methods:**

#### `verify_token` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/providers/google.py#L85" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
verify_token(self, token: str) -> AccessToken | None
```

Verify Google OAuth token by calling Google's tokeninfo API.


### `GoogleProvider` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/providers/google.py#L182" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Complete Google OAuth provider for FastMCP.

This provider makes it trivial to add Google OAuth protection to any
FastMCP server. Just provide your Google OAuth app credentials and
a base URL, and you're ready to go.

Features:
- Transparent OAuth proxy to Google
- Automatic token validation via Google's tokeninfo API
- User information extraction from Google APIs
- Minimal configuration required




================================================
FILE: docs/python-sdk/fastmcp-server-auth-providers-in_memory.mdx
================================================
---
title: in_memory
sidebarTitle: in_memory
---

# `fastmcp.server.auth.providers.in_memory`

## Classes

### `InMemoryOAuthProvider` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/providers/in_memory.py#L31" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


An in-memory OAuth provider for testing purposes.
It simulates the OAuth 2.1 flow locally without external calls.


**Methods:**

#### `get_client` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/providers/in_memory.py#L65" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
get_client(self, client_id: str) -> OAuthClientInformationFull | None
```

#### `register_client` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/providers/in_memory.py#L68" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
register_client(self, client_info: OAuthClientInformationFull) -> None
```

#### `authorize` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/providers/in_memory.py#L76" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
authorize(self, client: OAuthClientInformationFull, params: AuthorizationParams) -> str
```

Simulates user authorization and generates an authorization code.
Returns a redirect URI with the code and state.


#### `load_authorization_code` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/providers/in_memory.py#L129" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
load_authorization_code(self, client: OAuthClientInformationFull, authorization_code: str) -> AuthorizationCode | None
```

#### `exchange_authorization_code` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/providers/in_memory.py#L142" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
exchange_authorization_code(self, client: OAuthClientInformationFull, authorization_code: AuthorizationCode) -> OAuthToken
```

#### `load_refresh_token` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/providers/in_memory.py#L193" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
load_refresh_token(self, client: OAuthClientInformationFull, refresh_token: str) -> RefreshToken | None
```

#### `exchange_refresh_token` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/providers/in_memory.py#L208" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
exchange_refresh_token(self, client: OAuthClientInformationFull, refresh_token: RefreshToken, scopes: list[str]) -> OAuthToken
```

#### `load_access_token` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/providers/in_memory.py#L263" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
load_access_token(self, token: str) -> AccessToken | None
```

#### `verify_token` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/providers/in_memory.py#L274" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
verify_token(self, token: str) -> AccessToken | None
```

Verify a bearer token and return access info if valid.

This method implements the TokenVerifier protocol by delegating
to our existing load_access_token method.

**Args:**
- `token`: The token string to validate

**Returns:**
- AccessToken object if valid, None if invalid or expired


#### `revoke_token` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/providers/in_memory.py#L331" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
revoke_token(self, token: AccessToken | RefreshToken) -> None
```

Revokes an access or refresh token and its counterpart.




================================================
FILE: docs/python-sdk/fastmcp-server-auth-providers-jwt.mdx
================================================
---
title: jwt
sidebarTitle: jwt
---

# `fastmcp.server.auth.providers.jwt`


TokenVerifier implementations for FastMCP.

## Classes

### `JWKData` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/providers/jwt.py#L26" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


JSON Web Key data structure.


### `JWKSData` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/providers/jwt.py#L39" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


JSON Web Key Set data structure.


### `RSAKeyPair` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/providers/jwt.py#L46" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


RSA key pair for JWT testing.


**Methods:**

#### `generate` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/providers/jwt.py#L53" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
generate(cls) -> RSAKeyPair
```

Generate an RSA key pair for testing.

**Returns:**
- Generated key pair


#### `create_token` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/providers/jwt.py#L88" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
create_token(self, subject: str = 'fastmcp-user', issuer: str = 'https://fastmcp.example.com', audience: str | list[str] | None = None, scopes: list[str] | None = None, expires_in_seconds: int = 3600, additional_claims: dict[str, Any] | None = None, kid: str | None = None) -> str
```

Generate a test JWT token for testing purposes.

**Args:**
- `subject`: Subject claim (usually user ID)
- `issuer`: Issuer claim
- `audience`: Audience claim - can be a string or list of strings (optional)
- `scopes`: List of scopes to include
- `expires_in_seconds`: Token expiration time in seconds
- `additional_claims`: Any additional claims to include
- `kid`: Key ID to include in header


### `JWTVerifierSettings` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/providers/jwt.py#L141" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Settings for JWT token verification.


### `JWTVerifier` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/providers/jwt.py#L164" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


JWT token verifier supporting both asymmetric (RSA/ECDSA) and symmetric (HMAC) algorithms.

This verifier validates JWT tokens using various signing algorithms:
- **Asymmetric algorithms** (RS256/384/512, ES256/384/512, PS256/384/512):
  Uses public/private key pairs. Ideal for external clients and services where
  only the authorization server has the private key.
- **Symmetric algorithms** (HS256/384/512): Uses a shared secret for both
  signing and verification. Perfect for internal microservices and trusted
  environments where the secret can be securely shared.

Use this when:
- You have JWT tokens issued by an external service (asymmetric)
- You need JWKS support for automatic key rotation (asymmetric)
- You have internal microservices sharing a secret key (symmetric)
- Your tokens contain standard OAuth scopes and claims


**Methods:**

#### `load_access_token` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/providers/jwt.py#L366" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
load_access_token(self, token: str) -> AccessToken | None
```

Validates the provided JWT bearer token.

**Args:**
- `token`: The JWT token string to validate

**Returns:**
- AccessToken object if valid, None if invalid or expired


#### `verify_token` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/providers/jwt.py#L468" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
verify_token(self, token: str) -> AccessToken | None
```

Verify a bearer token and return access info if valid.

This method implements the TokenVerifier protocol by delegating
to our existing load_access_token method.

**Args:**
- `token`: The JWT token string to validate

**Returns:**
- AccessToken object if valid, None if invalid or expired


### `StaticTokenVerifier` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/providers/jwt.py#L484" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Simple static token verifier for testing and development.

This verifier validates tokens against a predefined dictionary of valid token
strings and their associated claims. When a token string matches a key in the
dictionary, the verifier returns the corresponding claims as if the token was
validated by a real authorization server.

Use this when:
- You're developing or testing locally without a real OAuth server
- You need predictable tokens for automated testing
- You want to simulate different users/scopes without complex setup
- You're prototyping and need simple API key-style authentication

WARNING: Never use this in production - tokens are stored in plain text!


**Methods:**

#### `verify_token` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/providers/jwt.py#L518" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
verify_token(self, token: str) -> AccessToken | None
```

Verify token against static token dictionary.




================================================
FILE: docs/python-sdk/fastmcp-server-auth-providers-workos.mdx
================================================
---
title: workos
sidebarTitle: workos
---

# `fastmcp.server.auth.providers.workos`


WorkOS authentication providers for FastMCP.

This module provides two WorkOS authentication strategies:

1. WorkOSProvider - OAuth proxy for WorkOS Connect applications (non-DCR)
2. AuthKitProvider - DCR-compliant provider for WorkOS AuthKit

Choose based on your WorkOS setup and authentication requirements.


## Classes

### `WorkOSProviderSettings` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/providers/workos.py#L31" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Settings for WorkOS OAuth provider.


### `WorkOSTokenVerifier` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/providers/workos.py#L55" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Token verifier for WorkOS OAuth tokens.

WorkOS AuthKit tokens are opaque, so we verify them by calling
the /oauth2/userinfo endpoint to check validity and get user info.


**Methods:**

#### `verify_token` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/providers/workos.py#L80" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
verify_token(self, token: str) -> AccessToken | None
```

Verify WorkOS OAuth token by calling userinfo endpoint.


### `WorkOSProvider` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/providers/workos.py#L127" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Complete WorkOS OAuth provider for FastMCP.

This provider implements WorkOS AuthKit OAuth using the OAuth Proxy pattern.
It provides OAuth2 authentication for users through WorkOS Connect applications.

Features:
- Transparent OAuth proxy to WorkOS AuthKit
- Automatic token validation via userinfo endpoint
- User information extraction from ID tokens
- Support for standard OAuth scopes (openid, profile, email)

Setup Requirements:
1. Create a WorkOS Connect application in your dashboard
2. Note your AuthKit domain (e.g., "https://your-app.authkit.app")
3. Configure redirect URI as: http://localhost:8000/auth/callback
4. Note your Client ID and Client Secret


### `AuthKitProviderSettings` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/providers/workos.py#L260" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

### `AuthKitProvider` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/providers/workos.py#L277" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


AuthKit metadata provider for DCR (Dynamic Client Registration).

This provider implements AuthKit integration using metadata forwarding
instead of OAuth proxying. This is the recommended approach for WorkOS DCR
as it allows WorkOS to handle the OAuth flow directly while FastMCP acts
as a resource server.

IMPORTANT SETUP REQUIREMENTS:

1. Enable Dynamic Client Registration in WorkOS Dashboard:
    - Go to Applications → Configuration
    - Toggle "Dynamic Client Registration" to enabled

2. Configure your FastMCP server URL as a callback:
    - Add your server URL to the Redirects tab in WorkOS dashboard
    - Example: https://your-fastmcp-server.com/oauth2/callback

For detailed setup instructions, see:
https://workos.com/docs/authkit/mcp/integrating/token-verification


**Methods:**

#### `get_routes` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/providers/workos.py#L360" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
get_routes(self, mcp_path: str | None = None, mcp_endpoint: Any | None = None) -> list[Route]
```

Get OAuth routes including AuthKit authorization server metadata forwarding.

This returns the standard protected resource routes plus an authorization server
metadata endpoint that forwards AuthKit's OAuth metadata to clients.

**Args:**
- `mcp_path`: The path where the MCP endpoint is mounted (e.g., "/mcp")
- `mcp_endpoint`: The MCP endpoint handler to protect with auth




================================================
FILE: docs/python-sdk/fastmcp-server-auth-redirect_validation.mdx
================================================
---
title: redirect_validation
sidebarTitle: redirect_validation
---

# `fastmcp.server.auth.redirect_validation`


Utilities for validating client redirect URIs in OAuth flows.

## Functions

### `matches_allowed_pattern` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/redirect_validation.py#L8" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
matches_allowed_pattern(uri: str, pattern: str) -> bool
```


Check if a URI matches an allowed pattern with wildcard support.

Patterns support * wildcard matching:
- http://localhost:* matches any localhost port
- http://127.0.0.1:* matches any 127.0.0.1 port
- https://*.example.com/* matches any subdomain of example.com
- https://app.example.com/auth/* matches any path under /auth/

**Args:**
- `uri`: The redirect URI to validate
- `pattern`: The allowed pattern (may contain wildcards)

**Returns:**
- True if the URI matches the pattern


### `validate_redirect_uri` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/auth/redirect_validation.py#L28" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
validate_redirect_uri(redirect_uri: str | AnyUrl | None, allowed_patterns: list[str] | None) -> bool
```


Validate a redirect URI against allowed patterns.

**Args:**
- `redirect_uri`: The redirect URI to validate
- `allowed_patterns`: List of allowed patterns. If None, all URIs are allowed (for DCR compatibility).
  If empty list, no URIs are allowed.
  To restrict to localhost only, explicitly pass DEFAULT_LOCALHOST_PATTERNS.

**Returns:**
- True if the redirect URI is allowed




================================================
FILE: docs/python-sdk/fastmcp-server-context.mdx
================================================
---
title: context
sidebarTitle: context
---

# `fastmcp.server.context`

## Functions

### `set_context` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/context.py#L69" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
set_context(context: Context) -> Generator[Context, None, None]
```

## Classes

### `LogData` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/context.py#L57" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Data object for passing log arguments to client-side handlers.

This provides an interface to match the Python standard library logging,
for compatibility with structured logging.


### `Context` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/context.py#L78" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Context object providing access to MCP capabilities.

This provides a cleaner interface to MCP's RequestContext functionality.
It gets injected into tool and resource functions that request it via type hints.

To use context in a tool function, add a parameter with the Context type annotation:

```python
@server.tool
def my_tool(x: int, ctx: Context) -> str:
    # Log messages to the client
    ctx.info(f"Processing {x}")
    ctx.debug("Debug info")
    ctx.warning("Warning message")
    ctx.error("Error message")

    # Report progress
    ctx.report_progress(50, 100, "Processing")

    # Access resources
    data = ctx.read_resource("resource://data")

    # Get request info
    request_id = ctx.request_id
    client_id = ctx.client_id

    # Manage state across the request
    ctx.set_state("key", "value")
    value = ctx.get_state("key")

    return str(x)
```

State Management:
Context objects maintain a state dictionary that can be used to store and share
data across middleware and tool calls within a request. When a new context
is created (nested contexts), it inherits a copy of its parent's state, ensuring
that modifications in child contexts don't affect parent contexts.

The context parameter name can be anything as long as it's annotated with Context.
The context is optional - tools that don't need it can omit the parameter.


**Methods:**

#### `fastmcp` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/context.py#L130" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
fastmcp(self) -> FastMCP
```

Get the FastMCP instance.


#### `request_context` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/context.py#L159" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
request_context(self) -> RequestContext[ServerSession, Any, Request]
```

Access to the underlying request context.

If called outside of a request context, this will raise a ValueError.


#### `report_progress` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/context.py#L169" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
report_progress(self, progress: float, total: float | None = None, message: str | None = None) -> None
```

Report progress for the current operation.

**Args:**
- `progress`: Current progress value e.g. 24
- `total`: Optional total value e.g. 100


#### `read_resource` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/context.py#L196" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
read_resource(self, uri: str | AnyUrl) -> list[ReadResourceContents]
```

Read a resource by URI.

**Args:**
- `uri`: Resource URI to read

**Returns:**
- The resource content as either text or bytes


#### `log` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/context.py#L209" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
log(self, message: str, level: LoggingLevel | None = None, logger_name: str | None = None, extra: Mapping[str, Any] | None = None) -> None
```

Send a log message to the client.

**Args:**
- `message`: Log message
- `level`: Optional log level. One of "debug", "info", "notice", "warning", "error", "critical",
  "alert", or "emergency". Default is "info".
- `logger_name`: Optional logger name
- `extra`: Optional mapping for additional arguments


#### `client_id` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/context.py#L236" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
client_id(self) -> str | None
```

Get the client ID if available.


#### `request_id` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/context.py#L245" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
request_id(self) -> str
```

Get the unique ID for this request.


#### `session_id` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/context.py#L250" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
session_id(self) -> str
```

Get the MCP session ID for ALL transports.

Returns the session ID that can be used as a key for session-based
data storage (e.g., Redis) to share data between tool calls within
the same client session.

**Returns:**
- The session ID for StreamableHTTP transports, or a generated ID
- for other transports.


#### `session` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/context.py#L294" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
session(self) -> ServerSession
```

Access to the underlying session for advanced usage.


#### `debug` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/context.py#L299" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
debug(self, message: str, logger_name: str | None = None, extra: Mapping[str, Any] | None = None) -> None
```

Send a debug log message.


#### `info` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/context.py#L310" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
info(self, message: str, logger_name: str | None = None, extra: Mapping[str, Any] | None = None) -> None
```

Send an info log message.


#### `warning` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/context.py#L321" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
warning(self, message: str, logger_name: str | None = None, extra: Mapping[str, Any] | None = None) -> None
```

Send a warning log message.


#### `error` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/context.py#L332" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
error(self, message: str, logger_name: str | None = None, extra: Mapping[str, Any] | None = None) -> None
```

Send an error log message.


#### `list_roots` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/context.py#L343" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
list_roots(self) -> list[Root]
```

List the roots available to the server, as indicated by the client.


#### `send_tool_list_changed` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/context.py#L348" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
send_tool_list_changed(self) -> None
```

Send a tool list changed notification to the client.


#### `send_resource_list_changed` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/context.py#L352" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
send_resource_list_changed(self) -> None
```

Send a resource list changed notification to the client.


#### `send_prompt_list_changed` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/context.py#L356" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
send_prompt_list_changed(self) -> None
```

Send a prompt list changed notification to the client.


#### `sample` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/context.py#L360" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
sample(self, messages: str | list[str | SamplingMessage], system_prompt: str | None = None, include_context: IncludeContext | None = None, temperature: float | None = None, max_tokens: int | None = None, model_preferences: ModelPreferences | str | list[str] | None = None) -> ContentBlock
```

Send a sampling request to the client and await the response.

Call this method at any time to have the server request an LLM
completion from the client. The client must be appropriately configured,
or the request will error.


#### `elicit` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/context.py#L444" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
elicit(self, message: str, response_type: None) -> AcceptedElicitation[dict[str, Any]] | DeclinedElicitation | CancelledElicitation
```

#### `elicit` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/context.py#L456" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
elicit(self, message: str, response_type: type[T]) -> AcceptedElicitation[T] | DeclinedElicitation | CancelledElicitation
```

#### `elicit` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/context.py#L466" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
elicit(self, message: str, response_type: list[str]) -> AcceptedElicitation[str] | DeclinedElicitation | CancelledElicitation
```

#### `elicit` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/context.py#L475" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
elicit(self, message: str, response_type: type[T] | list[str] | None = None) -> AcceptedElicitation[T] | AcceptedElicitation[dict[str, Any]] | AcceptedElicitation[str] | DeclinedElicitation | CancelledElicitation
```

Send an elicitation request to the client and await the response.

Call this method at any time to request additional information from
the user through the client. The client must support elicitation,
or the request will error.

Note that the MCP protocol only supports simple object schemas with
primitive types. You can provide a dataclass, TypedDict, or BaseModel to
comply. If you provide a primitive type, an object schema with a single
"value" field will be generated for the MCP interaction and
automatically deconstructed into the primitive type upon response.

If the response_type is None, the generated schema will be that of an
empty object in order to comply with the MCP protocol requirements.
Clients must send an empty object ("{}")in response.

**Args:**
- `message`: A human-readable message explaining what information is needed
- `response_type`: The type of the response, which should be a primitive
  type or dataclass or BaseModel. If it is a primitive type, an
  object schema with a single "value" field will be generated.


#### `get_http_request` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/context.py#L568" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
get_http_request(self) -> Request
```

Get the active starlette request.


#### `set_state` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/context.py#L583" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
set_state(self, key: str, value: Any) -> None
```

Set a value in the context state.


#### `get_state` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/context.py#L587" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
get_state(self, key: str) -> Any
```

Get a value from the context state. Returns None if the key is not found.




================================================
FILE: docs/python-sdk/fastmcp-server-dependencies.mdx
================================================
---
title: dependencies
sidebarTitle: dependencies
---

# `fastmcp.server.dependencies`

## Functions

### `get_context` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/dependencies.py#L27" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
get_context() -> Context
```

### `get_http_request` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/dependencies.py#L39" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
get_http_request() -> Request
```

### `get_http_headers` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/dependencies.py#L53" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
get_http_headers(include_all: bool = False) -> dict[str, str]
```


Extract headers from the current HTTP request if available.

Never raises an exception, even if there is no active HTTP request (in which case
an empty dict is returned).

By default, strips problematic headers like `content-length` that cause issues if forwarded to downstream clients.
If `include_all` is True, all headers are returned.


### `get_access_token` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/dependencies.py#L102" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
get_access_token() -> AccessToken | None
```


Get the FastMCP access token from the current context.

**Returns:**
- The access token if an authenticated user is available, None otherwise.




================================================
FILE: docs/python-sdk/fastmcp-server-elicitation.mdx
================================================
---
title: elicitation
sidebarTitle: elicitation
---

# `fastmcp.server.elicitation`

## Functions

### `get_elicitation_schema` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/elicitation.py#L99" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
get_elicitation_schema(response_type: type[T]) -> dict[str, Any]
```


Get the schema for an elicitation response.

**Args:**
- `response_type`: The type of the response


### `validate_elicitation_json_schema` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/elicitation.py#L118" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
validate_elicitation_json_schema(schema: dict[str, Any]) -> None
```


Validate that a JSON schema follows MCP elicitation requirements.

This ensures the schema is compatible with MCP elicitation requirements:
- Must be an object schema
- Must only contain primitive field types (string, number, integer, boolean)
- Must be flat (no nested objects or arrays of objects)
- Allows const fields (for Literal types) and enum fields (for Enum types)
- Only primitive types and their nullable variants are allowed

**Args:**
- `schema`: The JSON schema to validate

**Raises:**
- `TypeError`: If the schema doesn't meet MCP elicitation requirements


## Classes

### `ElicitationJsonSchema` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/elicitation.py#L32" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Custom JSON schema generator for MCP elicitation that always inlines enums.

MCP elicitation requires inline enum schemas without $ref/$defs references.
This generator ensures enums are always generated inline for compatibility.
Optionally adds enumNames for better UI display when available.


**Methods:**

#### `generate_inner` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/elicitation.py#L40" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
generate_inner(self, schema: core_schema.CoreSchema) -> JsonSchemaValue
```

Override to prevent ref generation for enums.


#### `enum_schema` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/elicitation.py#L50" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
enum_schema(self, schema: core_schema.EnumSchema) -> JsonSchemaValue
```

Generate inline enum schema with optional enumNames for better UI.

If enum members have a _display_name_ attribute or custom __str__,
we'll include enumNames for better UI representation.


### `AcceptedElicitation` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/elicitation.py#L87" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Result when user accepts the elicitation.


### `ScalarElicitationType` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/elicitation.py#L95" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>



================================================
FILE: docs/python-sdk/fastmcp-server-http.mdx
================================================
---
title: http
sidebarTitle: http
---

# `fastmcp.server.http`

## Functions

### `set_http_request` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/http.py#L74" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
set_http_request(request: Request) -> Generator[Request, None, None]
```

### `create_base_app` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/http.py#L98" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
create_base_app(routes: list[BaseRoute], middleware: list[Middleware], debug: bool = False, lifespan: Callable | None = None) -> StarletteWithLifespan
```


Create a base Starlette app with common middleware and routes.

**Args:**
- `routes`: List of routes to include in the app
- `middleware`: List of middleware to include in the app
- `debug`: Whether to enable debug mode
- `lifespan`: Optional lifespan manager for the app

**Returns:**
- A Starlette application


### `create_sse_app` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/http.py#L126" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
create_sse_app(server: FastMCP[LifespanResultT], message_path: str, sse_path: str, auth: AuthProvider | None = None, debug: bool = False, routes: list[BaseRoute] | None = None, middleware: list[Middleware] | None = None) -> StarletteWithLifespan
```


Return an instance of the SSE server app.

**Args:**
- `server`: The FastMCP server instance
- `message_path`: Path for SSE messages
- `sse_path`: Path for SSE connections
- `auth`: Optional authentication provider (AuthProvider)
- `debug`: Whether to enable debug mode
- `routes`: Optional list of custom routes
- `middleware`: Optional list of middleware

Returns:
A Starlette application with RequestContextMiddleware


### `create_streamable_http_app` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/http.py#L231" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
create_streamable_http_app(server: FastMCP[LifespanResultT], streamable_http_path: str, event_store: EventStore | None = None, auth: AuthProvider | None = None, json_response: bool = False, stateless_http: bool = False, debug: bool = False, routes: list[BaseRoute] | None = None, middleware: list[Middleware] | None = None) -> StarletteWithLifespan
```


Return an instance of the StreamableHTTP server app.

**Args:**
- `server`: The FastMCP server instance
- `streamable_http_path`: Path for StreamableHTTP connections
- `event_store`: Optional event store for session management
- `auth`: Optional authentication provider (AuthProvider)
- `json_response`: Whether to use JSON response format
- `stateless_http`: Whether to use stateless mode (new transport per request)
- `debug`: Whether to enable debug mode
- `routes`: Optional list of custom routes
- `middleware`: Optional list of middleware

**Returns:**
- A Starlette application with StreamableHTTP support


## Classes

### `StreamableHTTPASGIApp` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/http.py#L29" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


ASGI application wrapper for Streamable HTTP server transport.


### `StarletteWithLifespan` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/http.py#L67" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

**Methods:**

#### `lifespan` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/http.py#L69" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
lifespan(self) -> Lifespan[Starlette]
```

### `RequestContextMiddleware` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/http.py#L82" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Middleware that stores each request in a ContextVar




================================================
FILE: docs/python-sdk/fastmcp-server-low_level.mdx
================================================
---
title: low_level
sidebarTitle: low_level
---

# `fastmcp.server.low_level`

## Classes

### `LowLevelServer` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/low_level.py#L14" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

**Methods:**

#### `create_initialization_options` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/low_level.py#L24" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
create_initialization_options(self, notification_options: NotificationOptions | None = None, experimental_capabilities: dict[str, dict[str, Any]] | None = None, **kwargs: Any) -> InitializationOptions
```



================================================
FILE: docs/python-sdk/fastmcp-server-middleware-__init__.mdx
================================================
---
title: __init__
sidebarTitle: __init__
---

# `fastmcp.server.middleware`

*This module is empty or contains only private/internal implementations.*



================================================
FILE: docs/python-sdk/fastmcp-server-middleware-error_handling.mdx
================================================
---
title: error_handling
sidebarTitle: error_handling
---

# `fastmcp.server.middleware.error_handling`


Error handling middleware for consistent error responses and tracking.

## Classes

### `ErrorHandlingMiddleware` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/middleware/error_handling.py#L15" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Middleware that provides consistent error handling and logging.

Catches exceptions, logs them appropriately, and converts them to
proper MCP error responses. Also tracks error patterns for monitoring.


**Methods:**

#### `on_message` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/middleware/error_handling.py#L110" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
on_message(self, context: MiddlewareContext, call_next: CallNext) -> Any
```

Handle errors for all messages.


#### `get_error_stats` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/middleware/error_handling.py#L121" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
get_error_stats(self) -> dict[str, int]
```

Get error statistics for monitoring.


### `RetryMiddleware` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/middleware/error_handling.py#L126" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Middleware that implements automatic retry logic for failed requests.

Retries requests that fail with transient errors, using exponential
backoff to avoid overwhelming the server or external dependencies.


**Methods:**

#### `on_request` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/middleware/error_handling.py#L182" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
on_request(self, context: MiddlewareContext, call_next: CallNext) -> Any
```

Implement retry logic for requests.




================================================
FILE: docs/python-sdk/fastmcp-server-middleware-logging.mdx
================================================
---
title: logging
sidebarTitle: logging
---

# `fastmcp.server.middleware.logging`


Comprehensive logging middleware for FastMCP servers.

## Functions

### `default_serializer` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/middleware/logging.py#L14" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
default_serializer(data: Any) -> str
```


The default serializer for Payloads in the logging middleware.


## Classes

### `LoggingMiddleware` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/middleware/logging.py#L19" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Middleware that provides comprehensive request and response logging.

Logs all MCP messages with configurable detail levels. Useful for debugging,
monitoring, and understanding server usage patterns.


**Methods:**

#### `on_message` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/middleware/logging.py#L91" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
on_message(self, context: MiddlewareContext[Any], call_next: CallNext[Any, Any]) -> Any
```

Log all messages.


### `StructuredLoggingMiddleware` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/middleware/logging.py#L114" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Middleware that provides structured JSON logging for better log analysis.

Outputs structured logs that are easier to parse and analyze with log
aggregation tools like ELK stack, Splunk, or cloud logging services.


**Methods:**

#### `on_message` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/middleware/logging.py#L185" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
on_message(self, context: MiddlewareContext[Any], call_next: CallNext[Any, Any]) -> Any
```

Log structured message information.




================================================
FILE: docs/python-sdk/fastmcp-server-middleware-middleware.mdx
================================================
---
title: middleware
sidebarTitle: middleware
---

# `fastmcp.server.middleware.middleware`

## Functions

### `make_middleware_wrapper` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/middleware/middleware.py#L66" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
make_middleware_wrapper(middleware: Middleware, call_next: CallNext[T, R]) -> CallNext[T, R]
```


Create a wrapper that applies a single middleware to a context. The
closure bakes in the middleware and call_next function, so it can be
passed to other functions that expect a call_next function.


## Classes

### `CallNext` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/middleware/middleware.py#L42" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

### `MiddlewareContext` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/middleware/middleware.py#L47" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Unified context for all middleware operations.


**Methods:**

#### `copy` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/middleware/middleware.py#L62" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
copy(self, **kwargs: Any) -> MiddlewareContext[T]
```

### `Middleware` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/middleware/middleware.py#L79" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Base class for FastMCP middleware with dispatching hooks.


**Methods:**

#### `on_message` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/middleware/middleware.py#L126" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
on_message(self, context: MiddlewareContext[Any], call_next: CallNext[Any, Any]) -> Any
```

#### `on_request` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/middleware/middleware.py#L133" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
on_request(self, context: MiddlewareContext[mt.Request], call_next: CallNext[mt.Request, Any]) -> Any
```

#### `on_notification` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/middleware/middleware.py#L140" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
on_notification(self, context: MiddlewareContext[mt.Notification], call_next: CallNext[mt.Notification, Any]) -> Any
```

#### `on_call_tool` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/middleware/middleware.py#L147" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
on_call_tool(self, context: MiddlewareContext[mt.CallToolRequestParams], call_next: CallNext[mt.CallToolRequestParams, ToolResult]) -> ToolResult
```

#### `on_read_resource` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/middleware/middleware.py#L154" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
on_read_resource(self, context: MiddlewareContext[mt.ReadResourceRequestParams], call_next: CallNext[mt.ReadResourceRequestParams, mt.ReadResourceResult]) -> mt.ReadResourceResult
```

#### `on_get_prompt` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/middleware/middleware.py#L161" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
on_get_prompt(self, context: MiddlewareContext[mt.GetPromptRequestParams], call_next: CallNext[mt.GetPromptRequestParams, mt.GetPromptResult]) -> mt.GetPromptResult
```

#### `on_list_tools` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/middleware/middleware.py#L168" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
on_list_tools(self, context: MiddlewareContext[mt.ListToolsRequest], call_next: CallNext[mt.ListToolsRequest, list[Tool]]) -> list[Tool]
```

#### `on_list_resources` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/middleware/middleware.py#L175" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
on_list_resources(self, context: MiddlewareContext[mt.ListResourcesRequest], call_next: CallNext[mt.ListResourcesRequest, list[Resource]]) -> list[Resource]
```

#### `on_list_resource_templates` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/middleware/middleware.py#L182" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
on_list_resource_templates(self, context: MiddlewareContext[mt.ListResourceTemplatesRequest], call_next: CallNext[mt.ListResourceTemplatesRequest, list[ResourceTemplate]]) -> list[ResourceTemplate]
```

#### `on_list_prompts` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/middleware/middleware.py#L189" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
on_list_prompts(self, context: MiddlewareContext[mt.ListPromptsRequest], call_next: CallNext[mt.ListPromptsRequest, list[Prompt]]) -> list[Prompt]
```



================================================
FILE: docs/python-sdk/fastmcp-server-middleware-rate_limiting.mdx
================================================
---
title: rate_limiting
sidebarTitle: rate_limiting
---

# `fastmcp.server.middleware.rate_limiting`


Rate limiting middleware for protecting FastMCP servers from abuse.

## Classes

### `RateLimitError` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/middleware/rate_limiting.py#L15" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Error raised when rate limit is exceeded.


### `TokenBucketRateLimiter` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/middleware/rate_limiting.py#L22" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Token bucket implementation for rate limiting.


**Methods:**

#### `consume` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/middleware/rate_limiting.py#L38" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
consume(self, tokens: int = 1) -> bool
```

Try to consume tokens from the bucket.

**Args:**
- `tokens`: Number of tokens to consume

**Returns:**
- True if tokens were available and consumed, False otherwise


### `SlidingWindowRateLimiter` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/middleware/rate_limiting.py#L61" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Sliding window rate limiter implementation.


**Methods:**

#### `is_allowed` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/middleware/rate_limiting.py#L76" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
is_allowed(self) -> bool
```

Check if a request is allowed.


### `RateLimitingMiddleware` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/middleware/rate_limiting.py#L92" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Middleware that implements rate limiting to prevent server abuse.

Uses a token bucket algorithm by default, allowing for burst traffic
while maintaining a sustainable long-term rate.


**Methods:**

#### `on_request` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/middleware/rate_limiting.py#L152" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
on_request(self, context: MiddlewareContext, call_next: CallNext) -> Any
```

Apply rate limiting to requests.


### `SlidingWindowRateLimitingMiddleware` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/middleware/rate_limiting.py#L170" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Middleware that implements sliding window rate limiting.

Uses a sliding window approach which provides more precise rate limiting
but uses more memory to track individual request timestamps.


**Methods:**

#### `on_request` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/middleware/rate_limiting.py#L219" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
on_request(self, context: MiddlewareContext, call_next: CallNext) -> Any
```

Apply sliding window rate limiting to requests.




================================================
FILE: docs/python-sdk/fastmcp-server-middleware-timing.mdx
================================================
---
title: timing
sidebarTitle: timing
---

# `fastmcp.server.middleware.timing`


Timing middleware for measuring and logging request performance.

## Classes

### `TimingMiddleware` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/middleware/timing.py#L10" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Middleware that logs the execution time of requests.

Only measures and logs timing for request messages (not notifications).
Provides insights into performance characteristics of your MCP server.


**Methods:**

#### `on_request` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/middleware/timing.py#L39" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
on_request(self, context: MiddlewareContext, call_next: CallNext) -> Any
```

Time request execution and log the results.


### `DetailedTimingMiddleware` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/middleware/timing.py#L60" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Enhanced timing middleware with per-operation breakdowns.

Provides detailed timing information for different types of MCP operations,
allowing you to identify performance bottlenecks in specific operations.


**Methods:**

#### `on_call_tool` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/middleware/timing.py#L111" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
on_call_tool(self, context: MiddlewareContext, call_next: CallNext) -> Any
```

Time tool execution.


#### `on_read_resource` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/middleware/timing.py#L118" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
on_read_resource(self, context: MiddlewareContext, call_next: CallNext) -> Any
```

Time resource reading.


#### `on_get_prompt` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/middleware/timing.py#L127" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
on_get_prompt(self, context: MiddlewareContext, call_next: CallNext) -> Any
```

Time prompt retrieval.


#### `on_list_tools` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/middleware/timing.py#L134" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
on_list_tools(self, context: MiddlewareContext, call_next: CallNext) -> Any
```

Time tool listing.


#### `on_list_resources` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/middleware/timing.py#L140" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
on_list_resources(self, context: MiddlewareContext, call_next: CallNext) -> Any
```

Time resource listing.


#### `on_list_resource_templates` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/middleware/timing.py#L146" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
on_list_resource_templates(self, context: MiddlewareContext, call_next: CallNext) -> Any
```

Time resource template listing.


#### `on_list_prompts` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/middleware/timing.py#L152" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
on_list_prompts(self, context: MiddlewareContext, call_next: CallNext) -> Any
```

Time prompt listing.




================================================
FILE: docs/python-sdk/fastmcp-server-openapi.mdx
================================================
---
title: openapi
sidebarTitle: openapi
---

# `fastmcp.server.openapi`


FastMCP server implementation for OpenAPI integration.

## Classes

### `MCPType` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/openapi.py#L78" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Type of FastMCP component to create from a route.


### `RouteType` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/openapi.py#L97" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Deprecated: Use MCPType instead.

This enum is kept for backward compatibility and will be removed in a future version.


### `RouteMap` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/openapi.py#L111" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Mapping configuration for HTTP routes to FastMCP component types.


### `OpenAPITool` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/openapi.py#L229" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Tool implementation for OpenAPI endpoints.


**Methods:**

#### `run` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/openapi.py#L262" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
run(self, arguments: dict[str, Any]) -> ToolResult
```

Execute the HTTP request based on the route configuration.


### `OpenAPIResource` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/openapi.py#L523" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Resource implementation for OpenAPI endpoints.


**Methods:**

#### `read` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/openapi.py#L552" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
read(self) -> str | bytes
```

Fetch the resource data by making an HTTP request.


### `OpenAPIResourceTemplate` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/openapi.py#L642" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Resource template implementation for OpenAPI endpoints.


**Methods:**

#### `create_resource` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/openapi.py#L671" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
create_resource(self, uri: str, params: dict[str, Any], context: Context | None = None) -> Resource
```

Create a resource with the given parameters.


### `FastMCPOpenAPI` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/openapi.py#L696" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


FastMCP server implementation that creates components from an OpenAPI schema.

This class parses an OpenAPI specification and creates appropriate FastMCP components
(Tools, Resources, ResourceTemplates) based on route mappings.




================================================
FILE: docs/python-sdk/fastmcp-server-proxy.mdx
================================================
---
title: proxy
sidebarTitle: proxy
---

# `fastmcp.server.proxy`

## Functions

### `default_proxy_roots_handler` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/proxy.py#L521" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
default_proxy_roots_handler(context: RequestContext[ClientSession, LifespanContextT]) -> RootsList
```


A handler that forwards the list roots request from the remote server to the proxy's connected clients and relays the response back to the remote server.


## Classes

### `ProxyManagerMixin` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/proxy.py#L56" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


A mixin for proxy managers to provide a unified client retrieval method.


### `ProxyToolManager` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/proxy.py#L69" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


A ToolManager that sources its tools from a remote client in addition to local and mounted tools.


**Methods:**

#### `get_tools` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/proxy.py#L76" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
get_tools(self) -> dict[str, Tool]
```

Gets the unfiltered tool inventory including local, mounted, and proxy tools.


#### `list_tools` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/proxy.py#L102" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
list_tools(self) -> list[Tool]
```

Gets the filtered list of tools including local, mounted, and proxy tools.


#### `call_tool` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/proxy.py#L107" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
call_tool(self, key: str, arguments: dict[str, Any]) -> ToolResult
```

Calls a tool, trying local/mounted first, then proxy if not found.


### `ProxyResourceManager` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/proxy.py#L123" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


A ResourceManager that sources its resources from a remote client in addition to local and mounted resources.


**Methods:**

#### `get_resources` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/proxy.py#L130" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
get_resources(self) -> dict[str, Resource]
```

Gets the unfiltered resource inventory including local, mounted, and proxy resources.


#### `get_resource_templates` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/proxy.py#L153" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
get_resource_templates(self) -> dict[str, ResourceTemplate]
```

Gets the unfiltered template inventory including local, mounted, and proxy templates.


#### `list_resources` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/proxy.py#L176" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
list_resources(self) -> list[Resource]
```

Gets the filtered list of resources including local, mounted, and proxy resources.


#### `list_resource_templates` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/proxy.py#L181" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
list_resource_templates(self) -> list[ResourceTemplate]
```

Gets the filtered list of templates including local, mounted, and proxy templates.


#### `read_resource` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/proxy.py#L186" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
read_resource(self, uri: AnyUrl | str) -> str | bytes
```

Reads a resource, trying local/mounted first, then proxy if not found.


### `ProxyPromptManager` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/proxy.py#L204" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


A PromptManager that sources its prompts from a remote client in addition to local and mounted prompts.


**Methods:**

#### `get_prompts` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/proxy.py#L211" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
get_prompts(self) -> dict[str, Prompt]
```

Gets the unfiltered prompt inventory including local, mounted, and proxy prompts.


#### `list_prompts` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/proxy.py#L234" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
list_prompts(self) -> list[Prompt]
```

Gets the filtered list of prompts including local, mounted, and proxy prompts.


#### `render_prompt` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/proxy.py#L239" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
render_prompt(self, name: str, arguments: dict[str, Any] | None = None) -> GetPromptResult
```

Renders a prompt, trying local/mounted first, then proxy if not found.


### `ProxyTool` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/proxy.py#L256" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


A Tool that represents and executes a tool on a remote server.


**Methods:**

#### `from_mcp_tool` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/proxy.py#L266" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
from_mcp_tool(cls, client: Client, mcp_tool: mcp.types.Tool) -> ProxyTool
```

Factory method to create a ProxyTool from a raw MCP tool schema.


#### `run` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/proxy.py#L280" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
run(self, arguments: dict[str, Any], context: Context | None = None) -> ToolResult
```

Executes the tool by making a call through the client.


### `ProxyResource` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/proxy.py#L299" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


A Resource that represents and reads a resource from a remote server.


**Methods:**

#### `from_mcp_resource` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/proxy.py#L319" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
from_mcp_resource(cls, client: Client, mcp_resource: mcp.types.Resource) -> ProxyResource
```

Factory method to create a ProxyResource from a raw MCP resource schema.


#### `read` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/proxy.py#L337" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
read(self) -> str | bytes
```

Read the resource content from the remote server.


### `ProxyTemplate` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/proxy.py#L352" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


A ResourceTemplate that represents and creates resources from a remote server template.


**Methods:**

#### `from_mcp_template` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/proxy.py#L362" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
from_mcp_template(cls, client: Client, mcp_template: mcp.types.ResourceTemplate) -> ProxyTemplate
```

Factory method to create a ProxyTemplate from a raw MCP template schema.


#### `create_resource` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/proxy.py#L378" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
create_resource(self, uri: str, params: dict[str, Any], context: Context | None = None) -> ProxyResource
```

Create a resource from the template by calling the remote server.


### `ProxyPrompt` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/proxy.py#L413" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


A Prompt that represents and renders a prompt from a remote server.


**Methods:**

#### `from_mcp_prompt` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/proxy.py#L425" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
from_mcp_prompt(cls, client: Client, mcp_prompt: mcp.types.Prompt) -> ProxyPrompt
```

Factory method to create a ProxyPrompt from a raw MCP prompt schema.


#### `render` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/proxy.py#L447" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
render(self, arguments: dict[str, Any]) -> list[PromptMessage]
```

Render the prompt by making a call through the client.


### `FastMCPProxy` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/proxy.py#L454" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


A FastMCP server that acts as a proxy to a remote MCP-compliant server.
It uses specialized managers that fulfill requests via a client factory.


### `ProxyClient` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/proxy.py#L531" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


A proxy client that forwards advanced interactions between a remote MCP server and the proxy's connected clients.
Supports forwarding roots, sampling, elicitation, logging, and progress.


**Methods:**

#### `default_sampling_handler` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/proxy.py#L564" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
default_sampling_handler(cls, messages: list[mcp.types.SamplingMessage], params: mcp.types.CreateMessageRequestParams, context: RequestContext[ClientSession, LifespanContextT]) -> mcp.types.CreateMessageResult
```

A handler that forwards the sampling request from the remote server to the proxy's connected clients and relays the response back to the remote server.


#### `default_elicitation_handler` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/proxy.py#L590" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
default_elicitation_handler(cls, message: str, response_type: type, params: mcp.types.ElicitRequestParams, context: RequestContext[ClientSession, LifespanContextT]) -> ElicitResult
```

A handler that forwards the elicitation request from the remote server to the proxy's connected clients and relays the response back to the remote server.


#### `default_log_handler` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/proxy.py#L609" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
default_log_handler(cls, message: LogMessage) -> None
```

A handler that forwards the log notification from the remote server to the proxy's connected clients.


#### `default_progress_handler` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/proxy.py#L619" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
default_progress_handler(cls, progress: float, total: float | None, message: str | None) -> None
```

A handler that forwards the progress notification from the remote server to the proxy's connected clients.


### `StatefulProxyClient` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/proxy.py#L632" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


A proxy client that provides a stateful client factory for the proxy server.

The stateful proxy client bound its copy to the server session.
And it will be disconnected when the session is exited.

This is useful to proxy a stateful mcp server such as the Playwright MCP server.
Note that it is essential to ensure that the proxy server itself is also stateful.


**Methods:**

#### `clear` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/proxy.py#L654" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
clear(self)
```

Clear all cached clients and force disconnect them.


#### `new_stateful` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/proxy.py#L662" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
new_stateful(self) -> Client[ClientTransportT]
```

Create a new stateful proxy client instance with the same configuration.

Use this method as the client factory for stateful proxy server.




================================================
FILE: docs/python-sdk/fastmcp-server-server.mdx
================================================
---
title: server
sidebarTitle: server
---

# `fastmcp.server.server`


FastMCP - A more ergonomic interface for MCP servers.

## Functions

### `default_lifespan` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/server.py#L98" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
default_lifespan(server: FastMCP[LifespanResultT]) -> AsyncIterator[Any]
```


Default lifespan context manager that does nothing.

**Args:**
- `server`: The server instance this lifespan is managing

**Returns:**
- An empty context object


### `add_resource_prefix` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/server.py#L2222" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
add_resource_prefix(uri: str, prefix: str, prefix_format: Literal['protocol', 'path'] | None = None) -> str
```


Add a prefix to a resource URI.

**Args:**
- `uri`: The original resource URI
- `prefix`: The prefix to add

**Returns:**
- The resource URI with the prefix added

**Examples:**

With new style:
```python
add_resource_prefix("resource://path/to/resource", "prefix")
"resource://prefix/path/to/resource"
```
With legacy style:
```python
add_resource_prefix("resource://path/to/resource", "prefix")
"prefix+resource://path/to/resource"
```
With absolute path:
```python
add_resource_prefix("resource:///absolute/path", "prefix")
"resource://prefix//absolute/path"
```

**Raises:**
- `ValueError`: If the URI doesn't match the expected protocol\://path format


### `remove_resource_prefix` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/server.py#L2282" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
remove_resource_prefix(uri: str, prefix: str, prefix_format: Literal['protocol', 'path'] | None = None) -> str
```


Remove a prefix from a resource URI.

**Args:**
- `uri`: The resource URI with a prefix
- `prefix`: The prefix to remove
- `prefix_format`: The format of the prefix to remove

Returns:
The resource URI with the prefix removed

**Examples:**

With new style:
```python
remove_resource_prefix("resource://prefix/path/to/resource", "prefix")
"resource://path/to/resource"
```
With legacy style:
```python
remove_resource_prefix("prefix+resource://path/to/resource", "prefix")
"resource://path/to/resource"
```
With absolute path:
```python
remove_resource_prefix("resource://prefix//absolute/path", "prefix")
"resource:///absolute/path"
```

**Raises:**
- `ValueError`: If the URI doesn't match the expected protocol\://path format


### `has_resource_prefix` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/server.py#L2349" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
has_resource_prefix(uri: str, prefix: str, prefix_format: Literal['protocol', 'path'] | None = None) -> bool
```


Check if a resource URI has a specific prefix.

**Args:**
- `uri`: The resource URI to check
- `prefix`: The prefix to look for

**Returns:**
- True if the URI has the specified prefix, False otherwise

**Examples:**

With new style:
```python
has_resource_prefix("resource://prefix/path/to/resource", "prefix")
True
```
With legacy style:
```python
has_resource_prefix("prefix+resource://path/to/resource", "prefix")
True
```
With other path:
```python
has_resource_prefix("resource://other/path/to/resource", "prefix")
False
```

**Raises:**
- `ValueError`: If the URI doesn't match the expected protocol\://path format


## Classes

### `FastMCP` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/server.py#L129" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

**Methods:**

#### `settings` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/server.py#L314" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
settings(self) -> Settings
```

#### `name` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/server.py#L325" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
name(self) -> str
```

#### `instructions` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/server.py#L329" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
instructions(self) -> str | None
```

#### `version` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/server.py#L333" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
version(self) -> str | None
```

#### `run_async` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/server.py#L336" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
run_async(self, transport: Transport | None = None, show_banner: bool = True, **transport_kwargs: Any) -> None
```

Run the FastMCP server asynchronously.

**Args:**
- `transport`: Transport protocol to use ("stdio", "sse", or "streamable-http")


#### `run` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/server.py#L366" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
run(self, transport: Transport | None = None, show_banner: bool = True, **transport_kwargs: Any) -> None
```

Run the FastMCP server. Note this is a synchronous function.

**Args:**
- `transport`: Transport protocol to use ("stdio", "sse", or "streamable-http")


#### `add_middleware` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/server.py#L408" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
add_middleware(self, middleware: Middleware) -> None
```

#### `get_tools` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/server.py#L411" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
get_tools(self) -> dict[str, Tool]
```

Get all registered tools, indexed by registered key.


#### `get_tool` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/server.py#L415" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
get_tool(self, key: str) -> Tool
```

#### `get_resources` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/server.py#L421" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
get_resources(self) -> dict[str, Resource]
```

Get all registered resources, indexed by registered key.


#### `get_resource` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/server.py#L425" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
get_resource(self, key: str) -> Resource
```

#### `get_resource_templates` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/server.py#L431" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
get_resource_templates(self) -> dict[str, ResourceTemplate]
```

Get all registered resource templates, indexed by registered key.


#### `get_resource_template` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/server.py#L435" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
get_resource_template(self, key: str) -> ResourceTemplate
```

Get a registered resource template by key.


#### `get_prompts` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/server.py#L442" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
get_prompts(self) -> dict[str, Prompt]
```

List all available prompts.


#### `get_prompt` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/server.py#L448" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
get_prompt(self, key: str) -> Prompt
```

#### `custom_route` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/server.py#L454" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
custom_route(self, path: str, methods: list[str], name: str | None = None, include_in_schema: bool = True) -> Callable[[Callable[[Request], Awaitable[Response]]], Callable[[Request], Awaitable[Response]]]
```

Decorator to register a custom HTTP route on the FastMCP server.

Allows adding arbitrary HTTP endpoints outside the standard MCP protocol,
which can be useful for OAuth callbacks, health checks, or admin APIs.
The handler function must be an async function that accepts a Starlette
Request and returns a Response.

**Args:**
- `path`: URL path for the route (e.g., "/auth/callback")
- `methods`: List of HTTP methods to support (e.g., ["GET", "POST"])
- `name`: Optional name for the route (to reference this route with
  Starlette's reverse URL lookup feature)
- `include_in_schema`: Whether to include in OpenAPI schema, defaults to True


#### `add_tool` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/server.py#L856" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
add_tool(self, tool: Tool) -> Tool
```

Add a tool to the server.

The tool function can optionally request a Context object by adding a parameter
with the Context type annotation. See the @tool decorator for examples.

**Args:**
- `tool`: The Tool instance to register

**Returns:**
- The tool instance that was added to the server.


#### `remove_tool` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/server.py#L881" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
remove_tool(self, name: str) -> None
```

Remove a tool from the server.

**Args:**
- `name`: The name of the tool to remove

**Raises:**
- `NotFoundError`: If the tool is not found


#### `add_tool_transformation` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/server.py#L901" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
add_tool_transformation(self, tool_name: str, transformation: ToolTransformConfig) -> None
```

Add a tool transformation.


#### `remove_tool_transformation` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/server.py#L907" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
remove_tool_transformation(self, tool_name: str) -> None
```

Remove a tool transformation.


#### `tool` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/server.py#L912" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
tool(self, name_or_fn: AnyFunction) -> FunctionTool
```

#### `tool` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/server.py#L928" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
tool(self, name_or_fn: str | None = None) -> Callable[[AnyFunction], FunctionTool]
```

#### `tool` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/server.py#L943" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
tool(self, name_or_fn: str | AnyFunction | None = None) -> Callable[[AnyFunction], FunctionTool] | FunctionTool
```

Decorator to register a tool.

Tools can optionally request a Context object by adding a parameter with the
Context type annotation. The context provides access to MCP capabilities like
logging, progress reporting, and resource access.

This decorator supports multiple calling patterns:
- @server.tool (without parentheses)
- @server.tool (with empty parentheses)
- @server.tool("custom_name") (with name as first argument)
- @server.tool(name="custom_name") (with name as keyword argument)
- server.tool(function, name="custom_name") (direct function call)

**Args:**
- `name_or_fn`: Either a function (when used as @tool), a string name, or None
- `name`: Optional name for the tool (keyword-only, alternative to name_or_fn)
- `description`: Optional description of what the tool does
- `tags`: Optional set of tags for categorizing the tool
- `output_schema`: Optional JSON schema for the tool's output
- `annotations`: Optional annotations about the tool's behavior
- `exclude_args`: Optional list of argument names to exclude from the tool schema
- `meta`: Optional meta information about the tool
- `enabled`: Optional boolean to enable or disable the tool

**Examples:**

Register a tool with a custom name:
```python
@server.tool
def my_tool(x: int) -> str:
    return str(x)

# Register a tool with a custom name
@server.tool
def my_tool(x: int) -> str:
    return str(x)

@server.tool("custom_name")
def my_tool(x: int) -> str:
    return str(x)

@server.tool(name="custom_name")
def my_tool(x: int) -> str:
    return str(x)

# Direct function call
server.tool(my_function, name="custom_name")
```


#### `add_resource` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/server.py#L1074" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
add_resource(self, resource: Resource) -> Resource
```

Add a resource to the server.

**Args:**
- `resource`: A Resource instance to add

**Returns:**
- The resource instance that was added to the server.


#### `add_template` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/server.py#L1096" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
add_template(self, template: ResourceTemplate) -> ResourceTemplate
```

Add a resource template to the server.

**Args:**
- `template`: A ResourceTemplate instance to add

**Returns:**
- The template instance that was added to the server.


#### `add_resource_fn` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/server.py#L1118" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
add_resource_fn(self, fn: AnyFunction, uri: str, name: str | None = None, description: str | None = None, mime_type: str | None = None, tags: set[str] | None = None) -> None
```

Add a resource or template to the server from a function.

If the URI contains parameters (e.g. "resource://{param}") or the function
has parameters, it will be registered as a template resource.

**Args:**
- `fn`: The function to register as a resource
- `uri`: The URI for the resource
- `name`: Optional name for the resource
- `description`: Optional description of the resource
- `mime_type`: Optional MIME type for the resource
- `tags`: Optional set of tags for categorizing the resource


#### `resource` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/server.py#L1156" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
resource(self, uri: str) -> Callable[[AnyFunction], Resource | ResourceTemplate]
```

Decorator to register a function as a resource.

The function will be called when the resource is read to generate its content.
The function can return:
- str for text content
- bytes for binary content
- other types will be converted to JSON

Resources can optionally request a Context object by adding a parameter with the
Context type annotation. The context provides access to MCP capabilities like
logging, progress reporting, and session information.

If the URI contains parameters (e.g. "resource://{param}") or the function
has parameters, it will be registered as a template resource.

**Args:**
- `uri`: URI for the resource (e.g. "resource\://my-resource" or "resource\://{param}")
- `name`: Optional name for the resource
- `description`: Optional description of the resource
- `mime_type`: Optional MIME type for the resource
- `tags`: Optional set of tags for categorizing the resource
- `enabled`: Optional boolean to enable or disable the resource
- `annotations`: Optional annotations about the resource's behavior
- `meta`: Optional meta information about the resource

**Examples:**

Register a resource with a custom name:
```python
@server.resource("resource://my-resource")
def get_data() -> str:
    return "Hello, world!"

@server.resource("resource://my-resource")
async get_data() -> str:
    data = await fetch_data()
    return f"Hello, world! {data}"

@server.resource("resource://{city}/weather")
def get_weather(city: str) -> str:
    return f"Weather for {city}"

@server.resource("resource://{city}/weather")
def get_weather_with_context(city: str, ctx: Context) -> str:
    ctx.info(f"Fetching weather for {city}")
    return f"Weather for {city}"

@server.resource("resource://{city}/weather")
async def get_weather(city: str) -> str:
    data = await fetch_weather(city)
    return f"Weather for {city}: {data}"
```


#### `add_prompt` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/server.py#L1293" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
add_prompt(self, prompt: Prompt) -> Prompt
```

Add a prompt to the server.

**Args:**
- `prompt`: A Prompt instance to add

**Returns:**
- The prompt instance that was added to the server.


#### `prompt` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/server.py#L1316" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
prompt(self, name_or_fn: AnyFunction) -> FunctionPrompt
```

#### `prompt` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/server.py#L1329" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
prompt(self, name_or_fn: str | None = None) -> Callable[[AnyFunction], FunctionPrompt]
```

#### `prompt` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/server.py#L1341" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
prompt(self, name_or_fn: str | AnyFunction | None = None) -> Callable[[AnyFunction], FunctionPrompt] | FunctionPrompt
```

Decorator to register a prompt.

        Prompts can optionally request a Context object by adding a parameter with the
        Context type annotation. The context provides access to MCP capabilities like
        logging, progress reporting, and session information.

        This decorator supports multiple calling patterns:
        - @server.prompt (without parentheses)
        - @server.prompt() (with empty parentheses)
        - @server.prompt("custom_name") (with name as first argument)
        - @server.prompt(name="custom_name") (with name as keyword argument)
        - server.prompt(function, name="custom_name") (direct function call)

        Args:
            name_or_fn: Either a function (when used as @prompt), a string name, or None
            name: Optional name for the prompt (keyword-only, alternative to name_or_fn)
            description: Optional description of what the prompt does
            tags: Optional set of tags for categorizing the prompt
            enabled: Optional boolean to enable or disable the prompt
            meta: Optional meta information about the prompt

        Examples:

            ```python
            @server.prompt
            def analyze_table(table_name: str) -> list[Message]:
                schema = read_table_schema(table_name)
                return [
                    {
                        "role": "user",
                        "content": f"Analyze this schema:
{schema}"
}
]

            @server.prompt()
            def analyze_with_context(table_name: str, ctx: Context) -> list[Message]:
                ctx.info(f"Analyzing table {table_name}")
                schema = read_table_schema(table_name)
                return [
                    {
                        "role": "user",
                        "content": f"Analyze this schema:
{schema}"
}
]

            @server.prompt("custom_name")
            def analyze_file(path: str) -> list[Message]:
                content = await read_file(path)
                return [
                    {
                        "role": "user",
                        "content": {
                            "type": "resource",
                            "resource": {
                                "uri": f"file://{path}",
                                "text": content
                            }
                        }
                    }
                ]

            @server.prompt(name="custom_name")
            def another_prompt(data: str) -> list[Message]:
                return [{"role": "user", "content": data}]

            # Direct function call
            server.prompt(my_function, name="custom_name")
            ```


#### `run_stdio_async` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/server.py#L1482" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
run_stdio_async(self, show_banner: bool = True) -> None
```

Run the server using stdio transport.


#### `run_http_async` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/server.py#L1502" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
run_http_async(self, show_banner: bool = True, transport: Literal['http', 'streamable-http', 'sse'] = 'http', host: str | None = None, port: int | None = None, log_level: str | None = None, path: str | None = None, uvicorn_config: dict[str, Any] | None = None, middleware: list[ASGIMiddleware] | None = None, stateless_http: bool | None = None) -> None
```

Run the server using HTTP transport.

**Args:**
- `transport`: Transport protocol to use - either "streamable-http" (default) or "sse"
- `host`: Host address to bind to (defaults to settings.host)
- `port`: Port to bind to (defaults to settings.port)
- `log_level`: Log level for the server (defaults to settings.log_level)
- `path`: Path for the endpoint (defaults to settings.streamable_http_path or settings.sse_path)
- `uvicorn_config`: Additional configuration for the Uvicorn server
- `middleware`: A list of middleware to apply to the app
- `stateless_http`: Whether to use stateless HTTP (defaults to settings.stateless_http)


#### `run_sse_async` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/server.py#L1576" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
run_sse_async(self, host: str | None = None, port: int | None = None, log_level: str | None = None, path: str | None = None, uvicorn_config: dict[str, Any] | None = None) -> None
```

Run the server using SSE transport.


#### `sse_app` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/server.py#L1604" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
sse_app(self, path: str | None = None, message_path: str | None = None, middleware: list[ASGIMiddleware] | None = None) -> StarletteWithLifespan
```

Create a Starlette app for the SSE server.

**Args:**
- `path`: The path to the SSE endpoint
- `message_path`: The path to the message endpoint
- `middleware`: A list of middleware to apply to the app


#### `streamable_http_app` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/server.py#L1635" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
streamable_http_app(self, path: str | None = None, middleware: list[ASGIMiddleware] | None = None) -> StarletteWithLifespan
```

Create a Starlette app for the StreamableHTTP server.

**Args:**
- `path`: The path to the StreamableHTTP endpoint
- `middleware`: A list of middleware to apply to the app


#### `http_app` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/server.py#L1656" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
http_app(self, path: str | None = None, middleware: list[ASGIMiddleware] | None = None, json_response: bool | None = None, stateless_http: bool | None = None, transport: Literal['http', 'streamable-http', 'sse'] = 'http') -> StarletteWithLifespan
```

Create a Starlette app using the specified HTTP transport.

**Args:**
- `path`: The path for the HTTP endpoint
- `middleware`: A list of middleware to apply to the app
- `transport`: Transport protocol to use - either "streamable-http" (default) or "sse"

**Returns:**
- A Starlette application configured with the specified transport


#### `run_streamable_http_async` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/server.py#L1705" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
run_streamable_http_async(self, host: str | None = None, port: int | None = None, log_level: str | None = None, path: str | None = None, uvicorn_config: dict[str, Any] | None = None) -> None
```

#### `mount` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/server.py#L1730" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
mount(self, server: FastMCP[LifespanResultT], prefix: str | None = None, as_proxy: bool | None = None) -> None
```

Mount another FastMCP server on this server with an optional prefix.

Unlike importing (with import_server), mounting establishes a dynamic connection
between servers. When a client interacts with a mounted server's objects through
the parent server, requests are forwarded to the mounted server in real-time.
This means changes to the mounted server are immediately reflected when accessed
through the parent.

When a server is mounted with a prefix:
- Tools from the mounted server are accessible with prefixed names.
  Example: If server has a tool named "get_weather", it will be available as "prefix_get_weather".
- Resources are accessible with prefixed URIs.
  Example: If server has a resource with URI "weather://forecast", it will be available as
  "weather://prefix/forecast".
- Templates are accessible with prefixed URI templates.
  Example: If server has a template with URI "weather://location/{id}", it will be available
  as "weather://prefix/location/{id}".
- Prompts are accessible with prefixed names.
  Example: If server has a prompt named "weather_prompt", it will be available as
  "prefix_weather_prompt".

When a server is mounted without a prefix (prefix=None), its tools, resources, templates,
and prompts are accessible with their original names. Multiple servers can be mounted
without prefixes, and they will be tried in order until a match is found.

There are two modes for mounting servers:
1. Direct mounting (default when server has no custom lifespan): The parent server
   directly accesses the mounted server's objects in-memory for better performance.
   In this mode, no client lifecycle events occur on the mounted server, including
   lifespan execution.

2. Proxy mounting (default when server has a custom lifespan): The parent server
   treats the mounted server as a separate entity and communicates with it via a
   Client transport. This preserves all client-facing behaviors, including lifespan
   execution, but with slightly higher overhead.

**Args:**
- `server`: The FastMCP server to mount.
- `prefix`: Optional prefix to use for the mounted server's objects. If None,
  the server's objects are accessible with their original names.
- `as_proxy`: Whether to treat the mounted server as a proxy. If None (default),
  automatically determined based on whether the server has a custom lifespan
  (True if it has a custom lifespan, False otherwise).
- `tool_separator`: Deprecated. Separator character for tool names.
- `resource_separator`: Deprecated. Separator character for resource URIs.
- `prompt_separator`: Deprecated. Separator character for prompt names.


#### `import_server` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/server.py#L1852" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
import_server(self, server: FastMCP[LifespanResultT], prefix: str | None = None, tool_separator: str | None = None, resource_separator: str | None = None, prompt_separator: str | None = None) -> None
```

Import the MCP objects from another FastMCP server into this one,
optionally with a given prefix.

Note that when a server is *imported*, its objects are immediately
registered to the importing server. This is a one-time operation and
future changes to the imported server will not be reflected in the
importing server. Server-level configurations and lifespans are not imported.

When a server is imported with a prefix:
- The tools are imported with prefixed names
  Example: If server has a tool named "get_weather", it will be
  available as "prefix_get_weather"
- The resources are imported with prefixed URIs using the new format
  Example: If server has a resource with URI "weather://forecast", it will
  be available as "weather://prefix/forecast"
- The templates are imported with prefixed URI templates using the new format
  Example: If server has a template with URI "weather://location/{id}", it will
  be available as "weather://prefix/location/{id}"
- The prompts are imported with prefixed names
  Example: If server has a prompt named "weather_prompt", it will be available as
  "prefix_weather_prompt"

When a server is imported without a prefix (prefix=None), its tools, resources,
templates, and prompts are imported with their original names.

**Args:**
- `server`: The FastMCP server to import
- `prefix`: Optional prefix to use for the imported server's objects. If None,
  objects are imported with their original names.
- `tool_separator`: Deprecated. Separator for tool names.
- `resource_separator`: Deprecated and ignored. Prefix is now
  applied using the protocol\://prefix/path format
- `prompt_separator`: Deprecated. Separator for prompt names.


#### `from_openapi` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/server.py#L1981" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
from_openapi(cls, openapi_spec: dict[str, Any], client: httpx.AsyncClient, route_maps: list[RouteMap] | list[RouteMapNew] | None = None, route_map_fn: OpenAPIRouteMapFn | OpenAPIRouteMapFnNew | None = None, mcp_component_fn: OpenAPIComponentFn | OpenAPIComponentFnNew | None = None, mcp_names: dict[str, str] | None = None, tags: set[str] | None = None, **settings: Any) -> FastMCPOpenAPI | FastMCPOpenAPINew
```

Create a FastMCP server from an OpenAPI specification.


#### `from_fastapi` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/server.py#L2030" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
from_fastapi(cls, app: Any, name: str | None = None, route_maps: list[RouteMap] | list[RouteMapNew] | None = None, route_map_fn: OpenAPIRouteMapFn | OpenAPIRouteMapFnNew | None = None, mcp_component_fn: OpenAPIComponentFn | OpenAPIComponentFnNew | None = None, mcp_names: dict[str, str] | None = None, httpx_client_kwargs: dict[str, Any] | None = None, tags: set[str] | None = None, **settings: Any) -> FastMCPOpenAPI | FastMCPOpenAPINew
```

Create a FastMCP server from a FastAPI application.


#### `as_proxy` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/server.py#L2093" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
as_proxy(cls, backend: Client[ClientTransportT] | ClientTransport | FastMCP[Any] | AnyUrl | Path | MCPConfig | dict[str, Any] | str, **settings: Any) -> FastMCPProxy
```

Create a FastMCP proxy server for the given backend.

The `backend` argument can be either an existing `fastmcp.client.Client`
instance or any value accepted as the `transport` argument of
`fastmcp.client.Client`. This mirrors the convenience of the
`fastmcp.client.Client` constructor.


#### `from_client` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/server.py#L2154" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
from_client(cls, client: Client[ClientTransportT], **settings: Any) -> FastMCPProxy
```

Create a FastMCP proxy server from a FastMCP client.


#### `generate_name` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/server.py#L2206" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
generate_name(cls, name: str | None = None) -> str
```

### `MountedServer` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/server/server.py#L2216" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>



================================================
FILE: docs/python-sdk/fastmcp-settings.mdx
================================================
---
title: settings
sidebarTitle: settings
---

# `fastmcp.settings`

## Classes

### `ExtendedEnvSettingsSource` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/settings.py#L27" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


A special EnvSettingsSource that allows for multiple env var prefixes to be used.

Raises a deprecation warning if the old `FASTMCP_SERVER_` prefix is used.


**Methods:**

#### `get_field_value` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/settings.py#L34" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
get_field_value(self, field: FieldInfo, field_name: str) -> tuple[Any, str, bool]
```

### `ExtendedSettingsConfigDict` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/settings.py#L54" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

### `ExperimentalSettings` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/settings.py#L58" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

### `Settings` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/settings.py#L77" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


FastMCP settings.


**Methods:**

#### `get_setting` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/settings.py#L89" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
get_setting(self, attr: str) -> Any
```

Get a setting. If the setting contains one or more `__`, it will be
treated as a nested setting.


#### `set_setting` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/settings.py#L102" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
set_setting(self, attr: str, value: Any) -> None
```

Set a setting. If the setting contains one or more `__`, it will be
treated as a nested setting.


#### `settings_customise_sources` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/settings.py#L116" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
settings_customise_sources(cls, settings_cls: type[BaseSettings], init_settings: PydanticBaseSettingsSource, env_settings: PydanticBaseSettingsSource, dotenv_settings: PydanticBaseSettingsSource, file_secret_settings: PydanticBaseSettingsSource) -> tuple[PydanticBaseSettingsSource, ...]
```

#### `settings` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/settings.py#L134" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
settings(self) -> Self
```

This property is for backwards compatibility with FastMCP < 2.8.0,
which accessed fastmcp.settings.settings


#### `normalize_log_level` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/settings.py#L154" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
normalize_log_level(cls, v)
```



================================================
FILE: docs/python-sdk/fastmcp-tools-__init__.mdx
================================================
---
title: __init__
sidebarTitle: __init__
---

# `fastmcp.tools`

*This module is empty or contains only private/internal implementations.*



================================================
FILE: docs/python-sdk/fastmcp-tools-tool.mdx
================================================
---
title: tool
sidebarTitle: tool
---

# `fastmcp.tools.tool`

## Functions

### `default_serializer` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/tools/tool.py#L58" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
default_serializer(data: Any) -> str
```

## Classes

### `ToolResult` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/tools/tool.py#L62" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

**Methods:**

#### `to_mcp_result` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/tools/tool.py#L93" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
to_mcp_result(self) -> list[ContentBlock] | tuple[list[ContentBlock], dict[str, Any]]
```

### `Tool` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/tools/tool.py#L101" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Internal tool registration info.


**Methods:**

#### `enable` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/tools/tool.py#L119" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
enable(self) -> None
```

#### `disable` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/tools/tool.py#L127" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
disable(self) -> None
```

#### `to_mcp_tool` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/tools/tool.py#L135" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
to_mcp_tool(self, **overrides: Any) -> MCPTool
```

#### `from_function` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/tools/tool.py#L160" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
from_function(fn: Callable[..., Any], name: str | None = None, title: str | None = None, description: str | None = None, tags: set[str] | None = None, annotations: ToolAnnotations | None = None, exclude_args: list[str] | None = None, output_schema: dict[str, Any] | None | NotSetT | Literal[False] = NotSet, serializer: Callable[[Any], str] | None = None, meta: dict[str, Any] | None = None, enabled: bool | None = None) -> FunctionTool
```

Create a Tool from a function.


#### `run` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/tools/tool.py#L188" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
run(self, arguments: dict[str, Any]) -> ToolResult
```

Run the tool with arguments.

This method is not implemented in the base Tool class and must be
implemented by subclasses.

`run()` can EITHER return a list of ContentBlocks, or a tuple of
(list of ContentBlocks, dict of structured output).


#### `from_tool` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/tools/tool.py#L201" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
from_tool(cls, tool: Tool) -> TransformedTool
```

### `FunctionTool` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/tools/tool.py#L235" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

**Methods:**

#### `from_function` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/tools/tool.py#L239" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
from_function(cls, fn: Callable[..., Any], name: str | None = None, title: str | None = None, description: str | None = None, tags: set[str] | None = None, annotations: ToolAnnotations | None = None, exclude_args: list[str] | None = None, output_schema: dict[str, Any] | None | NotSetT | Literal[False] = NotSet, serializer: Callable[[Any], str] | None = None, meta: dict[str, Any] | None = None, enabled: bool | None = None) -> FunctionTool
```

Create a Tool from a function.


#### `run` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/tools/tool.py#L297" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
run(self, arguments: dict[str, Any]) -> ToolResult
```

Run the tool with arguments.


### `ParsedFunction` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/tools/tool.py#L343" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

**Methods:**

#### `from_function` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/tools/tool.py#L351" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
from_function(cls, fn: Callable[..., Any], exclude_args: list[str] | None = None, validate: bool = True, wrap_non_object_output_schema: bool = True) -> ParsedFunction
```



================================================
FILE: docs/python-sdk/fastmcp-tools-tool_manager.mdx
================================================
---
title: tool_manager
sidebarTitle: tool_manager
---

# `fastmcp.tools.tool_manager`

## Classes

### `ToolManager` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/tools/tool_manager.py#L25" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Manages FastMCP tools.


**Methods:**

#### `mount` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/tools/tool_manager.py#L51" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
mount(self, server: MountedServer) -> None
```

Adds a mounted server as a source for tools.


#### `has_tool` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/tools/tool_manager.py#L103" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
has_tool(self, key: str) -> bool
```

Check if a tool exists.


#### `get_tool` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/tools/tool_manager.py#L108" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
get_tool(self, key: str) -> Tool
```

Get tool by key.


#### `get_tools` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/tools/tool_manager.py#L115" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
get_tools(self) -> dict[str, Tool]
```

Gets the complete, unfiltered inventory of all tools.


#### `list_tools` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/tools/tool_manager.py#L121" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
list_tools(self) -> list[Tool]
```

Lists all tools, applying protocol filtering.


#### `add_tool_from_fn` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/tools/tool_manager.py#L137" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
add_tool_from_fn(self, fn: Callable[..., Any], name: str | None = None, description: str | None = None, tags: set[str] | None = None, annotations: ToolAnnotations | None = None, serializer: Callable[[Any], str] | None = None, exclude_args: list[str] | None = None) -> Tool
```

Add a tool to the server.


#### `add_tool` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/tools/tool_manager.py#L166" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
add_tool(self, tool: Tool) -> Tool
```

Register a tool with the server.


#### `add_tool_transformation` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/tools/tool_manager.py#L183" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
add_tool_transformation(self, tool_name: str, transformation: ToolTransformConfig) -> None
```

Add a tool transformation.


#### `get_tool_transformation` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/tools/tool_manager.py#L189" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
get_tool_transformation(self, tool_name: str) -> ToolTransformConfig | None
```

Get a tool transformation.


#### `remove_tool_transformation` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/tools/tool_manager.py#L193" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
remove_tool_transformation(self, tool_name: str) -> None
```

Remove a tool transformation.


#### `remove_tool` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/tools/tool_manager.py#L198" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
remove_tool(self, key: str) -> None
```

Remove a tool from the server.

**Args:**
- `key`: The key of the tool to remove

**Raises:**
- `NotFoundError`: If the tool is not found


#### `call_tool` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/tools/tool_manager.py#L212" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
call_tool(self, key: str, arguments: dict[str, Any]) -> ToolResult
```

Internal API for servers: Finds and calls a tool, respecting the
filtered protocol path.




================================================
FILE: docs/python-sdk/fastmcp-tools-tool_transform.mdx
================================================
---
title: tool_transform
sidebarTitle: tool_transform
---

# `fastmcp.tools.tool_transform`

## Functions

### `forward` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/tools/tool_transform.py#L37" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
forward(**kwargs) -> ToolResult
```


Forward to parent tool with argument transformation applied.

This function can only be called from within a transformed tool's custom
function. It applies argument transformation (renaming, validation) before
calling the parent tool.

For example, if the parent tool has args `x` and `y`, but the transformed
tool has args `a` and `b`, and an `transform_args` was provided that maps `x` to
`a` and `y` to `b`, then `forward(a=1, b=2)` will call the parent tool with
`x=1` and `y=2`.

**Args:**
- `**kwargs`: Arguments to forward to the parent tool (using transformed names).

**Returns:**
- The ToolResult from the parent tool execution.

**Raises:**
- `RuntimeError`: If called outside a transformed tool context.
- `TypeError`: If provided arguments don't match the transformed schema.


### `forward_raw` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/tools/tool_transform.py#L67" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
forward_raw(**kwargs) -> ToolResult
```


Forward directly to parent tool without transformation.

This function bypasses all argument transformation and validation, calling the parent
tool directly with the provided arguments. Use this when you need to call the parent
with its original parameter names and structure.

For example, if the parent tool has args `x` and `y`, then `forward_raw(x=1,
y=2)` will call the parent tool with `x=1` and `y=2`.

**Args:**
- `**kwargs`: Arguments to pass directly to the parent tool (using original names).

**Returns:**
- The ToolResult from the parent tool execution.

**Raises:**
- `RuntimeError`: If called outside a transformed tool context.


### `apply_transformations_to_tools` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/tools/tool_transform.py#L933" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
apply_transformations_to_tools(tools: dict[str, Tool], transformations: dict[str, ToolTransformConfig]) -> dict[str, Tool]
```


Apply a list of transformations to a list of tools. Tools that do not have any transforamtions
are left unchanged.


## Classes

### `ArgTransform` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/tools/tool_transform.py#L94" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Configuration for transforming a parent tool's argument.

This class allows fine-grained control over how individual arguments are transformed
when creating a new tool from an existing one. You can rename arguments, change their
descriptions, add default values, or hide them from clients while passing constants.

**Examples:**

Rename argument 'old_name' to 'new_name'
```python
ArgTransform(name="new_name")
```

Change description only
```python
ArgTransform(description="Updated description")
```

Add a default value (makes argument optional)
```python
ArgTransform(default=42)
```

Add a default factory (makes argument optional)
```python
ArgTransform(default_factory=lambda: time.time())
```

Change the type
```python
ArgTransform(type=str)
```

Hide the argument entirely from clients
```python
ArgTransform(hide=True)
```

Hide argument but pass a constant value to parent
```python
ArgTransform(hide=True, default="constant_value")
```

Hide argument but pass a factory-generated value to parent
```python
ArgTransform(hide=True, default_factory=lambda: uuid.uuid4().hex)
```

Make an optional parameter required (removes any default)
```python
ArgTransform(required=True)
```

Combine multiple transformations
```python
ArgTransform(name="new_name", description="New desc", default=None, type=int)
```


### `ArgTransformConfig` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/tools/tool_transform.py#L208" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


A model for requesting a single argument transform.


**Methods:**

#### `to_arg_transform` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/tools/tool_transform.py#L226" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
to_arg_transform(self) -> ArgTransform
```

Convert the argument transform to a FastMCP argument transform.


### `TransformedTool` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/tools/tool_transform.py#L232" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


A tool that is transformed from another tool.

This class represents a tool that has been created by transforming another tool.
It supports argument renaming, schema modification, custom function injection,
structured output control, and provides context for the forward() and forward_raw() functions.

The transformation can be purely schema-based (argument renaming, dropping, etc.)
or can include a custom function that uses forward() to call the parent tool
with transformed arguments. Output schemas and structured outputs are automatically
inherited from the parent tool but can be overridden or disabled.


**Methods:**

#### `run` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/tools/tool_transform.py#L259" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
run(self, arguments: dict[str, Any]) -> ToolResult
```

Run the tool with context set for forward() functions.

This method executes the tool's function while setting up the context
that allows forward() and forward_raw() to work correctly within custom
functions.

**Args:**
- `arguments`: Dictionary of arguments to pass to the tool's function.

**Returns:**
- ToolResult object containing content and optional structured output.


#### `from_tool` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/tools/tool_transform.py#L364" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
from_tool(cls, tool: Tool, name: str | None = None, title: str | None | NotSetT = NotSet, description: str | None | NotSetT = NotSet, tags: set[str] | None = None, transform_fn: Callable[..., Any] | None = None, transform_args: dict[str, ArgTransform] | None = None, annotations: ToolAnnotations | None | NotSetT = NotSet, output_schema: dict[str, Any] | None | NotSetT | Literal[False] = NotSet, serializer: Callable[[Any], str] | None | NotSetT = NotSet, meta: dict[str, Any] | None | NotSetT = NotSet, enabled: bool | None = None) -> TransformedTool
```

Create a transformed tool from a parent tool.

**Args:**
- `tool`: The parent tool to transform.
- `transform_fn`: Optional custom function. Can use forward() and forward_raw()
  to call the parent tool. Functions with **kwargs receive transformed
  argument names.
- `name`: New name for the tool. Defaults to parent tool's name.
- `title`: New title for the tool. Defaults to parent tool's title.
- `transform_args`: Optional transformations for parent tool arguments.
  Only specified arguments are transformed, others pass through unchanged\:
- Simple rename (str)
- Complex transformation (rename/description/default/drop) (ArgTransform)
- Drop the argument (None)
- `description`: New description. Defaults to parent's description.
- `tags`: New tags. Defaults to parent's tags.
- `annotations`: New annotations. Defaults to parent's annotations.
- `output_schema`: Control output schema for structured outputs\:
- None (default)\: Inherit from transform_fn if available, then parent tool
- dict\: Use custom output schema
- False\: Disable output schema and structured outputs
- `serializer`: New serializer. Defaults to parent's serializer.
- `meta`: Control meta information\:
- NotSet (default)\: Inherit from parent tool
- dict\: Use custom meta information
- None\: Remove meta information

**Returns:**
- TransformedTool with the specified transformations.

**Examples:**

# Transform specific arguments only
```python
Tool.from_tool(parent, transform_args={"old": "new"})  # Others unchanged
```

# Custom function with partial transforms
```python
async def custom(x: int, y: int) -> str:
    result = await forward(x=x, y=y)
    return f"Custom: {result}"

Tool.from_tool(parent, transform_fn=custom, transform_args={"a": "x", "b": "y"})
```

# Using **kwargs (gets all args, transformed and untransformed)
```python
async def flexible(**kwargs) -> str:
    result = await forward(**kwargs)
    return f"Got: {kwargs}"

Tool.from_tool(parent, transform_fn=flexible, transform_args={"a": "x"})
```

# Control structured outputs and schemas
```python
# Custom output schema
Tool.from_tool(parent, output_schema={
    "type": "object",
    "properties": {"status": {"type": "string"}}
})

# Disable structured outputs
Tool.from_tool(parent, output_schema=None)

# Return ToolResult for full control
async def custom_output(**kwargs) -> ToolResult:
    result = await forward(**kwargs)
    return ToolResult(
        content=[TextContent(text="Summary")],
        structured_content={"processed": True}
    )
```


### `ToolTransformConfig` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/tools/tool_transform.py#L887" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Provides a way to transform a tool.


**Methods:**

#### `apply` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/tools/tool_transform.py#L919" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
apply(self, tool: Tool) -> TransformedTool
```

Create a TransformedTool from a provided tool and this transformation configuration.




================================================
FILE: docs/python-sdk/fastmcp-utilities-__init__.mdx
================================================
---
title: __init__
sidebarTitle: __init__
---

# `fastmcp.utilities`


FastMCP utility modules.



================================================
FILE: docs/python-sdk/fastmcp-utilities-auth.mdx
================================================
---
title: auth
sidebarTitle: auth
---

# `fastmcp.utilities.auth`


Authentication utility helpers.

## Functions

### `parse_scopes` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/auth.py#L9" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
parse_scopes(value: Any) -> list[str] | None
```


Parse scopes from environment variables or settings values.

Accepts either a JSON array string, a comma- or space-separated string,
a list of strings, or ``None``. Returns a list of scopes or ``None`` if
no value is provided.




================================================
FILE: docs/python-sdk/fastmcp-utilities-cli.mdx
================================================
---
title: cli
sidebarTitle: cli
---

# `fastmcp.utilities.cli`

## Functions

### `is_already_in_uv_subprocess` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/cli.py#L28" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
is_already_in_uv_subprocess() -> bool
```


Check if we're already running in a FastMCP uv subprocess.


### `load_and_merge_config` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/cli.py#L33" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
load_and_merge_config(server_spec: str | None, **cli_overrides) -> tuple[MCPServerConfig, str]
```


Load config from server_spec and apply CLI overrides.

This consolidates the config parsing logic that was duplicated across
run, inspect, and dev commands.

**Args:**
- `server_spec`: Python file, config file, URL, or None to auto-detect
- `cli_overrides`: CLI arguments that override config values

**Returns:**
- Tuple of (MCPServerConfig, resolved_server_spec)


### `log_server_banner` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/cli.py#L151" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
log_server_banner(server: FastMCP[Any], transport: Literal['stdio', 'http', 'sse', 'streamable-http']) -> None
```


Creates and logs a formatted banner with server information and logo.

**Args:**
- `transport`: The transport protocol being used
- `server_name`: Optional server name to display
- `host`: Host address (for HTTP transports)
- `port`: Port number (for HTTP transports)
- `path`: Server path (for HTTP transports)




================================================
FILE: docs/python-sdk/fastmcp-utilities-components.mdx
================================================
---
title: components
sidebarTitle: components
---

# `fastmcp.utilities.components`

## Classes

### `FastMCPMeta` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/components.py#L15" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

### `FastMCPComponent` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/components.py#L28" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Base class for FastMCP tools, prompts, resources, and resource templates.


**Methods:**

#### `key` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/components.py#L61" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
key(self) -> str
```

The key of the component. This is used for internal bookkeeping
and may reflect e.g. prefixes or other identifiers. You should not depend on
keys having a certain value, as the same tool loaded from different
hierarchies of servers may have different keys.


#### `get_meta` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/components.py#L70" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
get_meta(self, include_fastmcp_meta: bool | None = None) -> dict[str, Any] | None
```

Get the meta information about the component.

If include_fastmcp_meta is True, a `_fastmcp` key will be added to the
meta, containing a `tags` field with the tags of the component.


#### `model_copy` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/components.py#L94" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
model_copy(self) -> Self
```

Create a copy of the component.

**Args:**
- `update`: A dictionary of fields to update.
- `deep`: Whether to deep copy the component.
- `key`: The key to use for the copy.


#### `enable` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/components.py#L127" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
enable(self) -> None
```

Enable the component.


#### `disable` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/components.py#L131" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
disable(self) -> None
```

Disable the component.


#### `copy` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/components.py#L135" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
copy(self) -> Self
```

Create a copy of the component.


### `MirroredComponent` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/components.py#L140" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Base class for components that are mirrored from a remote server.

Mirrored components cannot be enabled or disabled directly. Call copy() first
to create a local version you can modify.


**Methods:**

#### `enable` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/components.py#L153" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
enable(self) -> None
```

Enable the component.


#### `disable` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/components.py#L162" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
disable(self) -> None
```

Disable the component.


#### `copy` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/components.py#L171" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
copy(self) -> Self
```

Create a copy of the component that can be modified.




================================================
FILE: docs/python-sdk/fastmcp-utilities-exceptions.mdx
================================================
---
title: exceptions
sidebarTitle: exceptions
---

# `fastmcp.utilities.exceptions`

## Functions

### `iter_exc` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/exceptions.py#L12" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
iter_exc(group: BaseExceptionGroup)
```

### `get_catch_handlers` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/exceptions.py#L42" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
get_catch_handlers() -> Mapping[type[BaseException] | Iterable[type[BaseException]], Callable[[BaseExceptionGroup[Any]], Any]]
```



================================================
FILE: docs/python-sdk/fastmcp-utilities-http.mdx
================================================
---
title: http
sidebarTitle: http
---

# `fastmcp.utilities.http`

## Functions

### `find_available_port` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/http.py#L4" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
find_available_port() -> int
```


Find an available port by letting the OS assign one.




================================================
FILE: docs/python-sdk/fastmcp-utilities-inspect.mdx
================================================
---
title: inspect
sidebarTitle: inspect
---

# `fastmcp.utilities.inspect`


Utilities for inspecting FastMCP instances.

## Functions

### `inspect_fastmcp_v2` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/inspect.py#L98" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
inspect_fastmcp_v2(mcp: FastMCP[Any]) -> FastMCPInfo
```


Extract information from a FastMCP v2.x instance.

**Args:**
- `mcp`: The FastMCP v2.x instance to inspect

**Returns:**
- FastMCPInfo dataclass containing the extracted information


### `inspect_fastmcp_v1` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/inspect.py#L215" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
inspect_fastmcp_v1(mcp: FastMCP1x) -> FastMCPInfo
```


Extract information from a FastMCP v1.x instance using a Client.

**Args:**
- `mcp`: The FastMCP v1.x instance to inspect

**Returns:**
- FastMCPInfo dataclass containing the extracted information


### `inspect_fastmcp` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/inspect.py#L336" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
inspect_fastmcp(mcp: FastMCP[Any] | FastMCP1x) -> FastMCPInfo
```


Extract information from a FastMCP instance into a dataclass.

This function automatically detects whether the instance is FastMCP v1.x or v2.x
and uses the appropriate extraction method.

**Args:**
- `mcp`: The FastMCP instance to inspect (v1.x or v2.x)

**Returns:**
- FastMCPInfo dataclass containing the extracted information


### `format_fastmcp_info` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/inspect.py#L361" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
format_fastmcp_info(info: FastMCPInfo) -> bytes
```


Format FastMCPInfo as FastMCP-specific JSON.

This includes FastMCP-specific fields like tags, enabled, annotations, etc.


### `format_mcp_info` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/inspect.py#L388" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
format_mcp_info(mcp: FastMCP[Any] | FastMCP1x) -> bytes
```


Format server info as standard MCP protocol JSON.

Uses Client to get the standard MCP protocol format with camelCase fields.
Includes version metadata at the top level.


### `format_info` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/inspect.py#L421" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
format_info(mcp: FastMCP[Any] | FastMCP1x, format: InspectFormat | Literal['fastmcp', 'mcp'], info: FastMCPInfo | None = None) -> bytes
```


Format server information according to the specified format.

**Args:**
- `mcp`: The FastMCP instance
- `format`: Output format ("fastmcp" or "mcp")
- `info`: Pre-extracted FastMCPInfo (optional, will be extracted if not provided)

**Returns:**
- JSON bytes in the requested format


## Classes

### `ToolInfo` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/inspect.py#L19" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Information about a tool.


### `PromptInfo` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/inspect.py#L35" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Information about a prompt.


### `ResourceInfo` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/inspect.py#L49" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Information about a resource.


### `TemplateInfo` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/inspect.py#L65" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Information about a resource template.


### `FastMCPInfo` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/inspect.py#L82" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Information extracted from a FastMCP instance.


### `InspectFormat` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/inspect.py#L354" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Output format for inspect command.




================================================
FILE: docs/python-sdk/fastmcp-utilities-json_schema.mdx
================================================
---
title: json_schema
sidebarTitle: json_schema
---

# `fastmcp.utilities.json_schema`

## Functions

### `compress_schema` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/json_schema.py#L183" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
compress_schema(schema: dict, prune_params: list[str] | None = None, prune_defs: bool = True, prune_additional_properties: bool = True, prune_titles: bool = False) -> dict
```


Remove the given parameters from the schema.

**Args:**
- `schema`: The schema to compress
- `prune_params`: List of parameter names to remove from properties
- `prune_defs`: Whether to remove unused definitions
- `prune_additional_properties`: Whether to remove additionalProperties\: false
- `prune_titles`: Whether to remove title fields from the schema




================================================
FILE: docs/python-sdk/fastmcp-utilities-json_schema_type.mdx
================================================
---
title: json_schema_type
sidebarTitle: json_schema_type
---

# `fastmcp.utilities.json_schema_type`


Convert JSON Schema to Python types with validation.

The json_schema_to_type function converts a JSON Schema into a Python type that can be used
for validation with Pydantic. It supports:

- Basic types (string, number, integer, boolean, null)
- Complex types (arrays, objects)
- Format constraints (date-time, email, uri)
- Numeric constraints (minimum, maximum, multipleOf)
- String constraints (minLength, maxLength, pattern)
- Array constraints (minItems, maxItems, uniqueItems)
- Object properties with defaults
- References and recursive schemas
- Enums and constants
- Union types

Example:
```python
schema = {
"type": "object",
"properties": {
"name": {"type": "string", "minLength": 1},
"age": {"type": "integer", "minimum": 0},
"email": {"type": "string", "format": "email"}
},
"required": ["name", "age"]
}

    # Name is optional and will be inferred from schema's "title" property if not provided
    Person = json_schema_to_type(schema)
    # Creates a validated dataclass with name, age, and optional email fields
    ```


## Functions

### `json_schema_to_type` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/json_schema_type.py#L110" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
json_schema_to_type(schema: Mapping[str, Any], name: str | None = None) -> type
```


Convert JSON schema to appropriate Python type with validation.

**Args:**
- `schema`: A JSON Schema dictionary defining the type structure and validation rules
- `name`: Optional name for object schemas. Only allowed when schema type is "object".
  If not provided for objects, name will be inferred from schema's "title"
  property or default to "Root".

**Returns:**
- A Python type (typically a dataclass for objects) with Pydantic validation

**Raises:**
- `ValueError`: If a name is provided for a non-object schema

**Examples:**

Create a dataclass from an object schema:
```python
schema = {
    "type": "object",
    "title": "Person",
    "properties": {
        "name": {"type": "string", "minLength": 1},
        "age": {"type": "integer", "minimum": 0},
        "email": {"type": "string", "format": "email"}
    },
    "required": ["name", "age"]
}

Person = json_schema_to_type(schema)
# Creates a dataclass with name, age, and optional email fields:
# @dataclass
# class Person:
#     name: str
#     age: int
#     email: str | None = None
```
Person(name="John", age=30)

Create a scalar type with constraints:
```python
schema = {
    "type": "string",
    "minLength": 3,
    "pattern": "^[A-Z][a-z]+$"
}

NameType = json_schema_to_type(schema)
# Creates Annotated[str, StringConstraints(min_length=3, pattern="^[A-Z][a-z]+$")]

@dataclass
class Name:
    name: NameType
```


## Classes

### `JSONSchema` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/json_schema_type.py#L77" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>



================================================
FILE: docs/python-sdk/fastmcp-utilities-logging.mdx
================================================
---
title: logging
sidebarTitle: logging
---

# `fastmcp.utilities.logging`


Logging utilities for FastMCP.

## Functions

### `get_logger` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/logging.py#L10" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
get_logger(name: str) -> logging.Logger
```


Get a logger nested under FastMCP namespace.

**Args:**
- `name`: the name of the logger, which will be prefixed with 'FastMCP.'

**Returns:**
- a configured logger instance


### `configure_logging` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/logging.py#L22" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
configure_logging(level: Literal['DEBUG', 'INFO', 'WARNING', 'ERROR', 'CRITICAL'] | int = 'INFO', logger: logging.Logger | None = None, enable_rich_tracebacks: bool = True, **rich_kwargs: Any) -> None
```


Configure logging for FastMCP.

**Args:**
- `logger`: the logger to configure
- `level`: the log level to use
- `rich_kwargs`: the parameters to use for creating RichHandler




================================================
FILE: docs/python-sdk/fastmcp-utilities-mcp_config.mdx
================================================
---
title: mcp_config
sidebarTitle: mcp_config
---

# `fastmcp.utilities.mcp_config`

## Functions

### `mcp_config_to_servers_and_transports` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/mcp_config.py#L17" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
mcp_config_to_servers_and_transports(config: MCPConfig) -> list[tuple[str, FastMCP[Any], ClientTransport]]
```


A utility function to convert each entry of an MCP Config into a transport and server.


### `mcp_server_type_to_servers_and_transports` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/mcp_config.py#L27" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
mcp_server_type_to_servers_and_transports(name: str, mcp_server: MCPServerTypes) -> tuple[str, FastMCP[Any], ClientTransport]
```


A utility function to convert each entry of an MCP Config into a transport and server.




================================================
FILE: docs/python-sdk/fastmcp-utilities-mcp_server_config-__init__.mdx
================================================
---
title: __init__
sidebarTitle: __init__
---

# `fastmcp.utilities.mcp_server_config`


FastMCP Configuration module.

This module provides versioned configuration support for FastMCP servers.
The current version is v1, which is re-exported here for convenience.




================================================
FILE: docs/python-sdk/fastmcp-utilities-mcp_server_config-v1-__init__.mdx
================================================
---
title: __init__
sidebarTitle: __init__
---

# `fastmcp.utilities.mcp_server_config.v1`

*This module is empty or contains only private/internal implementations.*



================================================
FILE: docs/python-sdk/fastmcp-utilities-mcp_server_config-v1-environments-__init__.mdx
================================================
---
title: __init__
sidebarTitle: __init__
---

# `fastmcp.utilities.mcp_server_config.v1.environments`


Environment configuration for MCP servers.



================================================
FILE: docs/python-sdk/fastmcp-utilities-mcp_server_config-v1-environments-base.mdx
================================================
---
title: base
sidebarTitle: base
---

# `fastmcp.utilities.mcp_server_config.v1.environments.base`

## Classes

### `Environment` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/mcp_server_config/v1/environments/base.py#L7" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Base class for environment configuration.


**Methods:**

#### `build_command` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/mcp_server_config/v1/environments/base.py#L13" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
build_command(self, command: list[str]) -> list[str]
```

Build the full command with environment setup.

**Args:**
- `command`: Base command to wrap with environment setup

**Returns:**
- Full command ready for subprocess execution


#### `prepare` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/mcp_server_config/v1/environments/base.py#L24" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
prepare(self, output_dir: Path | None = None) -> None
```

Prepare the environment (optional, can be no-op).

**Args:**
- `output_dir`: Directory for persistent environment setup




================================================
FILE: docs/python-sdk/fastmcp-utilities-mcp_server_config-v1-environments-uv.mdx
================================================
---
title: uv
sidebarTitle: uv
---

# `fastmcp.utilities.mcp_server_config.v1.environments.uv`

## Classes

### `UVEnvironment` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/mcp_server_config/v1/environments/uv.py#L16" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Configuration for Python environment setup.


**Methods:**

#### `build_command` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/mcp_server_config/v1/environments/uv.py#L51" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
build_command(self, command: list[str]) -> list[str]
```

Build complete uv run command with environment args and command to execute.

**Args:**
- `command`: Command to execute (e.g., ["fastmcp", "run", "server.py"])

**Returns:**
- Complete command ready for subprocess.run, including "uv" prefix if needed.
- If no environment configuration is set, returns the command unchanged.


#### `run_with_uv` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/mcp_server_config/v1/environments/uv.py#L95" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
run_with_uv(self, command: list[str]) -> None
```

Execute a command using uv run with this environment configuration.

**Args:**
- `command`: Command and arguments to execute (e.g., ["fastmcp", "run", "server.py"])


#### `needs_uv` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/mcp_server_config/v1/environments/uv.py#L136" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
needs_uv(self) -> bool
```

Deprecated: Use _needs_setup() internally or check if build_command modifies the command.


#### `build_uv_run_command` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/mcp_server_config/v1/environments/uv.py#L140" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
build_uv_run_command(self, command: list[str]) -> list[str]
```

Deprecated: Use build_command() instead.


#### `prepare` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/mcp_server_config/v1/environments/uv.py#L144" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
prepare(self, output_dir: Path | None = None) -> None
```

Prepare the Python environment using uv.

**Args:**
- `output_dir`: Directory where the persistent uv project will be created.
  If None, creates a temporary directory for ephemeral use.




================================================
FILE: docs/python-sdk/fastmcp-utilities-mcp_server_config-v1-mcp_server_config.mdx
================================================
---
title: mcp_server_config
sidebarTitle: mcp_server_config
---

# `fastmcp.utilities.mcp_server_config.v1.mcp_server_config`


FastMCP Configuration File Support.

This module provides support for fastmcp.json configuration files that allow
users to specify server settings in a declarative format instead of using
command-line arguments.


## Functions

### `generate_schema` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/mcp_server_config/v1/mcp_server_config.py#L415" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
generate_schema(output_path: Path | str | None = None) -> dict[str, Any] | None
```


Generate JSON schema for fastmcp.json files.

This is used to create the schema file that IDEs can use for
validation and auto-completion.

**Args:**
- `output_path`: Optional path to write the schema to. If provided,
  writes the schema and returns None. If not provided,
  returns the schema as a dictionary.

**Returns:**
- JSON schema as a dictionary if output_path is None, otherwise None


## Classes

### `Deployment` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/mcp_server_config/v1/mcp_server_config.py#L36" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Configuration for server deployment and runtime settings.


**Methods:**

#### `apply_runtime_settings` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/mcp_server_config/v1/mcp_server_config.py#L85" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
apply_runtime_settings(self, config_path: Path | None = None) -> None
```

Apply runtime settings like environment variables and working directory.

**Args:**
- `config_path`: Path to config file for resolving relative paths

Environment variables support interpolation with ${VAR_NAME} syntax.
For example: "API_URL": "https://api.${ENVIRONMENT}.example.com"
will substitute the value of the ENVIRONMENT variable at runtime.


### `MCPServerConfig` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/mcp_server_config/v1/mcp_server_config.py#L134" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Configuration for a FastMCP server.

This configuration file allows you to specify all settings needed to run
a FastMCP server in a declarative format.


**Methods:**

#### `validate_source` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/mcp_server_config/v1/mcp_server_config.py#L183" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
validate_source(cls, v: dict | Source) -> SourceType
```

Validate and convert source to proper format.

Supports:
- Dict format: `{"path": "server.py", "entrypoint": "app"}`
- FileSystemSource instance (passed through)

No string parsing happens here - that's only at CLI boundaries.
MCPServerConfig works only with properly typed objects.


#### `validate_environment` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/mcp_server_config/v1/mcp_server_config.py#L199" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
validate_environment(cls, v: dict | Any) -> EnvironmentType
```

Ensure environment has a type field for discrimination.

For backward compatibility, if no type is specified, default to "uv".


#### `validate_deployment` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/mcp_server_config/v1/mcp_server_config.py#L210" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
validate_deployment(cls, v: dict | Deployment) -> Deployment
```

Validate and convert deployment to Deployment.

Accepts:
- Deployment instance
- dict that can be converted to Deployment


#### `from_file` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/mcp_server_config/v1/mcp_server_config.py#L223" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
from_file(cls, file_path: Path) -> MCPServerConfig
```

Load configuration from a JSON file.

**Args:**
- `file_path`: Path to the configuration file

**Returns:**
- MCPServerConfig instance

**Raises:**
- `FileNotFoundError`: If the file doesn't exist
- `json.JSONDecodeError`: If the file is not valid JSON
- `pydantic.ValidationError`: If the configuration is invalid


#### `from_cli_args` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/mcp_server_config/v1/mcp_server_config.py#L246" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
from_cli_args(cls, source: FileSystemSource, transport: Literal['stdio', 'http', 'sse', 'streamable-http'] | None = None, host: str | None = None, port: int | None = None, path: str | None = None, log_level: Literal['DEBUG', 'INFO', 'WARNING', 'ERROR', 'CRITICAL'] | None = None, python: str | None = None, dependencies: list[str] | None = None, requirements: str | None = None, project: str | None = None, editable: str | None = None, env: dict[str, str] | None = None, cwd: str | None = None, args: list[str] | None = None) -> MCPServerConfig
```

Create a config from CLI arguments.

This allows us to have a single code path where everything
goes through a config object.

**Args:**
- `source`: Server source (FileSystemSource instance)
- `transport`: Transport protocol
- `host`: Host for HTTP transport
- `port`: Port for HTTP transport
- `path`: URL path for server
- `log_level`: Logging level
- `python`: Python version
- `dependencies`: Python packages to install
- `requirements`: Path to requirements file
- `project`: Path to project directory
- `editable`: Path to install in editable mode
- `env`: Environment variables
- `cwd`: Working directory
- `args`: Server arguments

**Returns:**
- MCPServerConfig instance


#### `find_config` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/mcp_server_config/v1/mcp_server_config.py#L323" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
find_config(cls, start_path: Path | None = None) -> Path | None
```

Find a fastmcp.json file in the specified directory.

**Args:**
- `start_path`: Directory to look in (defaults to current directory)

**Returns:**
- Path to the configuration file, or None if not found


#### `prepare` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/mcp_server_config/v1/mcp_server_config.py#L342" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
prepare(self, skip_source: bool = False, output_dir: Path | None = None) -> None
```

Prepare environment and source for execution.

When output_dir is provided, creates a persistent uv project.
When output_dir is None, does ephemeral caching (for backwards compatibility).

**Args:**
- `skip_source`: Skip source preparation if True
- `output_dir`: Directory to create the persistent uv project in (optional)


#### `prepare_environment` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/mcp_server_config/v1/mcp_server_config.py#L363" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
prepare_environment(self, output_dir: Path | None = None) -> None
```

Prepare the Python environment.

**Args:**
- `output_dir`: If provided, creates a persistent uv project in this directory.
  If None, just populates uv's cache for ephemeral use.

Delegates to the environment's prepare() method


#### `prepare_source` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/mcp_server_config/v1/mcp_server_config.py#L374" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
prepare_source(self) -> None
```

Prepare the source for loading.

Delegates to the source's prepare() method.


#### `run_server` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/mcp_server_config/v1/mcp_server_config.py#L381" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
run_server(self, **kwargs: Any) -> None
```

Load and run the server with this configuration.

**Args:**
- `**kwargs`: Additional arguments to pass to server.run_async()
  These override config settings




================================================
FILE: docs/python-sdk/fastmcp-utilities-mcp_server_config-v1-sources-__init__.mdx
================================================
---
title: __init__
sidebarTitle: __init__
---

# `fastmcp.utilities.mcp_server_config.v1.sources`

*This module is empty or contains only private/internal implementations.*



================================================
FILE: docs/python-sdk/fastmcp-utilities-mcp_server_config-v1-sources-base.mdx
================================================
---
title: base
sidebarTitle: base
---

# `fastmcp.utilities.mcp_server_config.v1.sources.base`

## Classes

### `Source` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/mcp_server_config/v1/sources/base.py#L7" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Abstract base class for all source types.


**Methods:**

#### `prepare` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/mcp_server_config/v1/sources/base.py#L12" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
prepare(self) -> None
```

Prepare the source (download, clone, install, etc).

For sources that need preparation (e.g., git clone, download),
this method performs that preparation. For sources that don't
need preparation (e.g., local files), this is a no-op.


#### `load_server` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/mcp_server_config/v1/sources/base.py#L23" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
load_server(self) -> Any
```

Load and return the FastMCP server instance.

Must be called after prepare() if the source requires preparation.
All information needed to load the server should be available
as attributes on the source instance.




================================================
FILE: docs/python-sdk/fastmcp-utilities-mcp_server_config-v1-sources-filesystem.mdx
================================================
---
title: filesystem
sidebarTitle: filesystem
---

# `fastmcp.utilities.mcp_server_config.v1.sources.filesystem`

## Classes

### `FileSystemSource` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/mcp_server_config/v1/sources/filesystem.py#L15" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Source for local Python files.


**Methods:**

#### `parse_path_with_object` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/mcp_server_config/v1/sources/filesystem.py#L28" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
parse_path_with_object(cls, v: str) -> str
```

Parse path:object syntax and extract the object name.

This validator runs before the model is created, allowing us to
handle the "file.py:object" syntax at the model boundary.


#### `load_server` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/mcp_server_config/v1/sources/filesystem.py#L63" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
load_server(self) -> Any
```

Load server from filesystem.




================================================
FILE: docs/python-sdk/fastmcp-utilities-openapi.mdx
================================================
---
title: openapi
sidebarTitle: openapi
---

# `fastmcp.utilities.openapi`

## Functions

### `format_array_parameter` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/openapi.py#L41" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
format_array_parameter(values: list, parameter_name: str, is_query_parameter: bool = False) -> str | list
```


Format an array parameter according to OpenAPI specifications.

**Args:**
- `values`: List of values to format
- `parameter_name`: Name of the parameter (for error messages)
- `is_query_parameter`: If True, can return list for explode=True behavior

**Returns:**
- String (comma-separated) or list (for query params with explode=True)


### `format_deep_object_parameter` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/openapi.py#L91" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
format_deep_object_parameter(param_value: dict, parameter_name: str) -> dict[str, str]
```


Format a dictionary parameter for deepObject style serialization.

According to OpenAPI 3.0 spec, deepObject style with explode=true serializes
object properties as separate query parameters with bracket notation.

For example: `{"id": "123", "type": "user"}` becomes `param[id]=123&param[type]=user`.

**Args:**
- `param_value`: Dictionary value to format
- `parameter_name`: Name of the parameter

**Returns:**
- Dictionary with bracketed parameter names as keys


### `parse_openapi_to_http_routes` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/openapi.py#L201" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
parse_openapi_to_http_routes(openapi_dict: dict[str, Any]) -> list[HTTPRoute]
```


Parses an OpenAPI schema dictionary into a list of HTTPRoute objects
using the openapi-pydantic library.

Supports both OpenAPI 3.0.x and 3.1.x versions.


### `clean_schema_for_display` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/openapi.py#L741" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
clean_schema_for_display(schema: JsonSchema | None) -> JsonSchema | None
```


Clean up a schema dictionary for display by removing internal/complex fields.


### `generate_example_from_schema` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/openapi.py#L801" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
generate_example_from_schema(schema: JsonSchema | None) -> Any
```


Generate a simple example value from a JSON schema dictionary.
Very basic implementation focusing on types.


### `format_json_for_description` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/openapi.py#L884" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
format_json_for_description(data: Any, indent: int = 2) -> str
```


Formats Python data as a JSON string block for markdown.


### `format_description_with_responses` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/openapi.py#L893" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
format_description_with_responses(base_description: str, responses: dict[str, Any], parameters: list[ParameterInfo] | None = None, request_body: RequestBodyInfo | None = None) -> str
```


Formats the base description string with response, parameter, and request body information.

**Args:**
- `base_description`: The initial description to be formatted.
- `responses`: A dictionary of response information, keyed by status code.
- `parameters`: A list of parameter information,
  including path and query parameters. Each parameter includes details such as name,
  location, whether it is required, and a description.
- `request_body`: Information about the request body,
  including its description, whether it is required, and its content schema.

**Returns:**
- The formatted description string with additional details about responses, parameters,
- and the request body.


### `extract_output_schema_from_responses` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/openapi.py#L1419" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
extract_output_schema_from_responses(responses: dict[str, ResponseInfo], schema_definitions: dict[str, Any] | None = None, openapi_version: str | None = None) -> dict[str, Any] | None
```


Extract output schema from OpenAPI responses for use as MCP tool output schema.

This function finds the first successful response (200, 201, 202, 204) with a
JSON-compatible content type and extracts its schema. If the schema is not an
object type, it wraps it to comply with MCP requirements.

**Args:**
- `responses`: Dictionary of ResponseInfo objects keyed by status code
- `schema_definitions`: Optional schema definitions to include in the output schema
- `openapi_version`: OpenAPI version string, used to optimize nullable field handling

**Returns:**
- MCP-compliant output schema with potential wrapping, or None if no suitable schema found


## Classes

### `ParameterInfo` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/openapi.py#L124" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Represents a single parameter for an HTTP operation in our IR.


### `RequestBodyInfo` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/openapi.py#L136" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Represents the request body for an HTTP operation in our IR.


### `ResponseInfo` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/openapi.py#L146" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Represents response information in our IR.


### `HTTPRoute` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/openapi.py#L154" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Intermediate Representation for a single OpenAPI operation.


### `OpenAPIParser` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/openapi.py#L255" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Unified parser for OpenAPI schemas with generic type parameters to handle both 3.0 and 3.1.


**Methods:**

#### `parse` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/openapi.py#L619" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
parse(self) -> list[HTTPRoute]
```

Parse the OpenAPI schema into HTTP routes.




================================================
FILE: docs/python-sdk/fastmcp-utilities-tests.mdx
================================================
---
title: tests
sidebarTitle: tests
---

# `fastmcp.utilities.tests`

## Functions

### `temporary_settings` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/tests.py#L25" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
temporary_settings(**kwargs: Any)
```


Temporarily override FastMCP setting values.

**Args:**
- `**kwargs`: The settings to override, including nested settings.


### `run_server_in_process` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/tests.py#L75" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
run_server_in_process(server_fn: Callable[..., None], *args, **kwargs) -> Generator[str, None, None]
```


Context manager that runs a FastMCP server in a separate process and
returns the server URL. When the context manager is exited, the server process is killed.

**Args:**
- `server_fn`: The function that runs a FastMCP server. FastMCP servers are
  not pickleable, so we need a function that creates and runs one.
- `*args`: Arguments to pass to the server function.
- `provide_host_and_port`: Whether to provide the host and port to the server function as kwargs.
- `host`: Host to bind the server to (default\: "127.0.0.1").
- `port`: Port to bind the server to (default\: find available port).
- `**kwargs`: Keyword arguments to pass to the server function.

**Returns:**
- The server URL.


### `caplog_for_fastmcp` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/tests.py#L141" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
caplog_for_fastmcp(caplog)
```


Context manager to capture logs from FastMCP loggers even when propagation is disabled.


## Classes

### `HeadlessOAuth` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/tests.py#L152" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


OAuth provider that bypasses browser interaction for testing.

This simulates the complete OAuth flow programmatically by making HTTP requests
instead of opening a browser and running a callback server. Useful for automated testing.


**Methods:**

#### `redirect_handler` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/tests.py#L165" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
redirect_handler(self, authorization_url: str) -> None
```

Make HTTP request to authorization URL and store response for callback handler.


#### `callback_handler` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/tests.py#L171" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
callback_handler(self) -> tuple[str, str | None]
```

Parse stored response and return (auth_code, state).




================================================
FILE: docs/python-sdk/fastmcp-utilities-types.mdx
================================================
---
title: types
sidebarTitle: types
---

# `fastmcp.utilities.types`


Common types used across FastMCP.

## Functions

### `get_cached_typeadapter` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/types.py#L41" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
get_cached_typeadapter(cls: T) -> TypeAdapter[T]
```


TypeAdapters are heavy objects, and in an application context we'd typically
create them once in a global scope and reuse them as often as possible.
However, this isn't feasible for user-generated functions. Instead, we use a
cache to minimize the cost of creating them as much as possible.


### `issubclass_safe` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/types.py#L116" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
issubclass_safe(cls: type, base: type) -> bool
```


Check if cls is a subclass of base, even if cls is a type variable.


### `is_class_member_of_type` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/types.py#L126" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
is_class_member_of_type(cls: Any, base: type) -> bool
```


Check if cls is a member of base, even if cls is a type variable.

Base can be a type, a UnionType, or an Annotated type. Generic types are not
considered members (e.g. T is not a member of list\[T]).


### `find_kwarg_by_type` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/types.py#L148" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
find_kwarg_by_type(fn: Callable, kwarg_type: type) -> str | None
```


Find the name of the kwarg that is of type kwarg_type.

Includes union types that contain the kwarg_type, as well as Annotated types.


### `replace_type` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/types.py#L377" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
replace_type(type_, type_map: dict[type, type])
```


Given a (possibly generic, nested, or otherwise complex) type, replaces all
instances of old_type with new_type.

This is useful for transforming types when creating tools.

**Args:**
- `type_`: The type to replace instances of old_type with new_type.
- `old_type`: The type to replace.
- `new_type`: The type to replace old_type with.

Examples:
```python
>>> replace_type(list[int | bool], {int: str})
list[str | bool]

>>> replace_type(list[list[int]], {int: str})
list[list[str]]
```


## Classes

### `FastMCPBaseModel` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/types.py#L34" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Base model for FastMCP models.


### `Image` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/types.py#L174" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Helper class for returning images from tools.


**Methods:**

#### `to_image_content` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/types.py#L211" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
to_image_content(self, mime_type: str | None = None, annotations: Annotations | None = None) -> mcp.types.ImageContent
```

Convert to MCP ImageContent.


### `Audio` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/types.py#L233" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Helper class for returning audio from tools.


**Methods:**

#### `to_audio_content` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/types.py#L270" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
to_audio_content(self, mime_type: str | None = None, annotations: Annotations | None = None) -> mcp.types.AudioContent
```

### `File` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/types.py#L291" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


Helper class for returning audio from tools.


**Methods:**

#### `to_resource_content` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/types.py#L330" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>

```python
to_resource_content(self, mime_type: str | None = None, annotations: Annotations | None = None) -> mcp.types.EmbeddedResource
```

### `ContextSamplingFallbackProtocol` <sup><a href="https://github.com/jlowin/fastmcp/blob/main/src/fastmcp/utilities/types.py#L414" target="_blank"><Icon icon="github" style="width: 14px; height: 14px;" /></a></sup>


