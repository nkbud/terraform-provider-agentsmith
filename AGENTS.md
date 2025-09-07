# Codex Configuration Inspector — Implementation Plan

Goal: Implement a host‑machine configuration inspector that exposes Codex CLI settings (as documented in `codex/codex_config_helper.md`) via a Terraform data source in this provider.

Outcomes

- New Terraform data source `agentsmith_codex` that returns the effective Codex configuration from the local machine.
- Clear precedence rules (home vs project overrides; environment‑derived overrides) applied consistently.
- Examples and tests to validate decoding, merging, and schema shape.

Non‑Goals (initial)

- Running Codex CLI or parsing runtime CLI flags; we only inspect files and environment.
- Writing or mutating any Codex configuration.

Design Summary

- Data source: `codexDataSource` in `internal/provider/codex_data_source.go`, registered from `internal/provider/provider.go` alongside the existing Claude inspector.
- Sources:
  - `$CODEX_HOME/config.toml` (default `~/.codex/config.toml`).
  - Optional project configuration at `<workdir>/.codex/config.toml` (if present) with higher precedence than `$CODEX_HOME`.
  - Environment‑derived adjustments: `OPENAI_BASE_URL` (overrides built‑in OpenAI base URL), `CODEX_HOME` (locates config), and a minimal set of other environment flags we can reliably infer (e.g., `CODEX_SANDBOX_NETWORK_DISABLED`).
- Effective config:
  - Start with “base” config merged from files (home then project, last‑one‑wins per key).
  - If `profile` is set in config, overlay the profile table from `profiles.<name>` on top of base to compute `effective_config`.
  - Surface `active_profile` and both `base_config` and `profile_config` so users can audit precedence.
- Parser: Decode TOML into typed Go structs, falling back to generic maps for dynamic sections (e.g., `model_providers`, `mcp_servers`, `profiles`).
  - Library: `github.com/pelletier/go-toml/v2` (preferred) or `github.com/BurntSushi/toml`. We’ll add as a module dependency.
- Sensitivity: Do not print API keys. For provider entries that specify `env_key`, expose only whether the env var is set (boolean) and the env var name; never its value.
- UX parity with Claude inspector: Reuse patterns in `claude_data_source.go` for schema style, diagnostics, file discovery, and environment reading helpers.

Schema (MVP)

All attributes are `Computed: true`.

- `id` (string): Constant identifier (`"codex-config"`).
- `source_files` (list<string>): Paths that contributed to the merge (in order).
- `active_profile` (string): Name of the applied profile (if any).
- `effective_config` (nested): Canonical view after precedence and profile overlay, with these keys:
  - `model` (string)
  - `model_provider` (string)
  - `model_context_window` (number)
  - `model_max_output_tokens` (number)
  - `approval_policy` (string enum: `untrusted|on-failure|on-request|never`)
  - `sandbox_mode` (string enum: `read-only|workspace-write|danger-full-access`)
  - `sandbox_workspace_write` (nested): `writable_roots` (list<string>), `network_access` (bool), `exclude_tmpdir_env_var` (bool), `exclude_slash_tmp` (bool)
  - `notify` (list<string>)
  - `history` (nested): `persistence` (string), `max_bytes` (number, optional)
  - `file_opener` (string enum: `vscode|vscode-insiders|windsurf|cursor|none`)
  - `hide_agent_reasoning` (bool)
  - `show_raw_agent_reasoning` (bool)
  - `model_reasoning_effort` (string)
  - `model_reasoning_summary` (string)
  - `model_verbosity` (string)
  - `model_supports_reasoning_summaries` (bool)
  - `project_doc_max_bytes` (number)
  - `tui` (map<string,string> or reserved object)
  - `shell_environment_policy` (nested): `inherit` (string), `ignore_default_excludes` (bool), `exclude` (list<string>), `set` (map<string,string>), `include_only` (list<string>)
  - `model_providers` (list<nested>): flattened from TOML table; each element has `id`, `name`, `base_url`, `env_key`, `wire_api`, `query_params` (map), `http_headers` (map), `env_http_headers` (map), `request_max_retries` (number), `stream_max_retries` (number), `stream_idle_timeout_ms` (number), plus `env_key_is_set` (bool, computed from host env).
  - `mcp_servers` (list<nested>): flattened from TOML table; each element has `id`, `command`, `args` (list<string>), `env` (map<string,string>—marked Sensitive in Terraform schema).
  - `profiles` (list<nested>): flattened; each element has `name` and a subset of the same keys as above so users can audit what the profile would override.
- `environment` (nested): host signals used during computation:
  - `codex_home` (string), `openai_base_url` (string), `codex_sandbox_network_disabled` (bool)

Merging & Precedence

1. Determine `$CODEX_HOME` (env var or `~/.codex`).
2. Read `config.toml` from `$CODEX_HOME` and from `<workdir>/.codex` if present. Merge maps and scalars with “last file wins” (project overrides home).
3. Compute `active_profile` from `profile` if set; overlay `profiles.<active>` on top of the merged base for `effective_config`.
4. Apply environment‑derived overrides:
   - If `OPENAI_BASE_URL` is set and a built‑in `openai` provider exists, override its base URL field in `effective_config`.
   - Record whether `env_key` is set for each provider via `os.LookupEnv`.
   - Record `CODEX_SANDBOX_NETWORK_DISABLED` as a boolean signal; do not mutate `sandbox_mode` based on it (informational only).

Implementation Steps

1) Define schema and types (MVP)
- Create `internal/provider/codex_data_source.go` with:
  - `codexDataSource` struct (+ `Configure`, `Metadata`, `Schema`, `Read`).
  - Go model structs mirroring the schema sections above.
  - Helper for TOML decode into an internal `codexConfig` struct that uses maps for dynamic tables.
  - Helper to flatten dynamic TOML tables to `list<nested>` with an explicit `id`/`name` field.

2) File discovery & reading
- Implement `$CODEX_HOME` resolution and `<workdir>/.codex` detection, mirroring `claudeDataSource.getSettingsFilePaths` style.
- Read in order: home then project; accumulate `source_files` for visibility.

3) TOML parsing & merge
- Add `go-toml/v2` as a dependency and implement decode with strict errors to catch typos; report via diagnostics with file path context.
- Merge strategy:
  - Scalars: last value wins.
  - Tables (maps): shallow merge (last wins per key); for dynamic maps (`model_providers`, `mcp_servers`, `profiles`), merge entries by id/name.

4) Profile overlay
- If `profile` is set, locate `profiles.<name>`.
- Overlay profile values onto the merged base to compute `effective_config`.
- Expose `base_config` and `profile_config` sections if helpful; at minimum expose `active_profile` and `effective_config`.

5) Environment inspection
- Read `CODEX_HOME`, `OPENAI_BASE_URL`, `CODEX_SANDBOX_NETWORK_DISABLED`.
- For each provider, compute `env_key_is_set` by checking its declared `env_key`.
- Mark any env values as Sensitive when exposed (avoid dumping secrets).

6) Provider registration & examples
- Add `NewCodexDataSource` to `DataSources()` in `internal/provider/provider.go`.
- Create `examples/agentsmith_codex/main.tf` with a minimal usage block and outputs.

7) Tests
- Unit tests for:
  - Decoding minimal `config.toml`.
  - Merging home + project.
  - Applying profiles overlay.
  - Flattening `model_providers`/`mcp_servers`.
- Acceptance test to exercise `data "agentsmith_codex"` with a temporary `$CODEX_HOME`.

8) Docs
- README: add a short “Codex Inspector” section and link to `codex/codex_config_helper.md`.
- Data source docs (Terraform): schema reference and precedence notes.

Key Design Choices & Rationale

- Flatten dynamic TOML tables to `list<nested>`: Terraform requires static schema keys—lists with an explicit identifier (`id`) keep it predictable while preserving fidelity and diffs.
- Best‑effort environment awareness: We surface whether provider `env_key` vars are set without reading secrets. Only a small, well‑defined set of env variables are consulted for effective config to avoid surprising side‑effects.
- Iterative scope: Start with the high‑value top‑level keys and dynamic tables; expand to additional keys as needed.

Open Questions (to validate with stakeholders)

- Should project‑level overrides under `<workdir>/.codex/config.toml` be considered authoritative, or do we strictly honor `$CODEX_HOME` only? (Plan above supports both, with project overriding.)
- Do we want to expose `base_config` and `profile_config` explicitly in the schema, or keep the state slimmer with only `effective_config` + `active_profile`?
- Any additional env overrides we should recognize beyond `OPENAI_BASE_URL`?

Next Steps

- Implement the MVP schema and file reading logic.
- Hook up TOML parsing and profile overlay.
- Register the data source and add an example.
- Add tests and iterate on schema details based on feedback.

## Codex Configuration Resources — Implementation Plan

Goal: Add Terraform resources to author Codex CLI configuration files on the host machine, complementing the read‑only inspector. Resources should safely manage `$CODEX_HOME/config.toml` or a project‑scoped `<workdir>/.codex/config.toml`, mirroring the data source schema and precedence rules without leaking secrets.

Outcomes

- New Terraform resource `agentsmith_codex_config` that writes a Codex TOML config file.
- Clear file targeting via `scope` (home|project|custom) and predictable merge/ownership semantics.
- Schema mirrors the inspector’s core fields with nested blocks for providers, MCP servers, and profiles.
- Safe, atomic writes with backups, optional strict validation, and clear guardrails around sensitive data persistence.
- Examples and acceptance tests for create/update/delete and drift detection.

Design Summary

- Cardinality decision: exactly one resource per file on the host. The provider rejects multiple resources targeting the same resolved path during plan/apply and uses best‑effort file locks to avoid concurrent writes.
-
- Primary resource: `codexConfigResource` in `internal/provider/codex_config_resource.go`, registered in `internal/provider/provider.go`.
- Single‑writer model: One resource instance owns a single TOML file (home, project, or explicit path). We do not support multiple resources writing to the same file in the MVP.
- Ownership strategies:
  - `merge_strategy = preserve_unknown` (default): update declared keys; preserve unknown top‑level keys/tables.
  - `merge_strategy = replace_all`: resource owns the full file; writes canonical content; removes undeclared keys.
  - `merge_strategy = fail_on_unknown`: refuse to write if unknown keys exist in the file (safe guardrail for strict environments).
- Atomic write + backup: write to `*.tmp`, fsync+rename, optional `*.bak` backup. Default permissions `0600`.
- Profile overlay guidance: Resource writes exactly the file it targets. Effective runtime overlay still follows: base (home) → project → profile. Resource does not merge across files; users declare separate resources for home/project if needed.
- Secret handling: We never read or print provider API keys. Writing `mcp_servers.env` is allowed but gated by an explicit `allow_sensitive_env_writes = true` to reduce foot‑guns. Provider API keys continue to be referenced via `env_key` only.

Resource: `agentsmith_codex_config` (MVP)

- Meta attributes:
  - `id` (Computed): stable identifier: `sha256(path)`.
  - `scope` (Required): `"home"|"project"|"custom"`.
  - `path` (Optional): absolute or relative path to `config.toml` when `scope = "custom"`.
  - `merge_strategy` (Optional): `"preserve_unknown"|"replace_all"|"fail_on_unknown"` (default `preserve_unknown`).
  - `create_directories` (Optional bool, default true): create target dir if missing.
  - `file_mode` (Optional string, default `"0600"`): permissions applied on create.
  - `backup_on_write` (Optional bool, default true): if true, writes `config.toml.bak` before update.
  - `allow_sensitive_env_writes` (Optional bool, default false): required to persist `mcp_servers.env` values to disk.
  - `keep_file_on_destroy` (Optional bool, default true): if true, do not delete file on destroy; instead remove only managed keys (when `preserve_unknown`) or leave file intact.
  - `validate_strict` (Optional bool, default true): validate keys against known schema; surface typos as diagnostics.
  - `resolved_path` (Computed string): fully resolved target path.

- Top‑level config (all Optional; omitted keys are removed only when `replace_all`):
  - `model` (string)
  - `model_provider` (string)
  - `model_context_window` (number)
  - `model_max_output_tokens` (number)
  - `approval_policy` (string: `untrusted|on-failure|on-request|never`)
  - `sandbox_mode` (string: `read-only|workspace-write|danger-full-access`)
  - `sandbox_workspace_write` (block, max 1):
    - `writable_roots` (list(string))
    - `network_access` (bool)
    - `exclude_tmpdir_env_var` (bool)
    - `exclude_slash_tmp` (bool)
  - `notify` (list(string))
  - `history` (block, max 1): `persistence` (string), `max_bytes` (number)
  - `file_opener` (string: `vscode|vscode-insiders|windsurf|cursor|none`)
  - `hide_agent_reasoning` (bool)
  - `show_raw_agent_reasoning` (bool)
  - `model_reasoning_effort` (string)
  - `model_reasoning_summary` (string)
  - `model_verbosity` (string)
  - `model_supports_reasoning_summaries` (bool)
  - `project_doc_max_bytes` (number)
  - `tui` (map(string))
  - `shell_environment_policy` (block, max 1):
    - `inherit` (string)
    - `ignore_default_excludes` (bool)
    - `exclude` (list(string))
    - `set` (map(string))
    - `include_only` (list(string))

- Dynamic tables
  - `model_providers` (block, multiple):
    - `id` (Required string)
    - `name` (Optional string)
    - `base_url` (Optional string)
    - `env_key` (Optional string) — name of env var that carries the secret; value is never read/written.
    - `wire_api` (Optional string)
    - `query_params` (Optional map(string))
    - `http_headers` (Optional map(string))
    - `env_http_headers` (Optional map(string))
    - `request_max_retries` (Optional number)
    - `stream_max_retries` (Optional number)
    - `stream_idle_timeout_ms` (Optional number)
  - `mcp_servers` (block, multiple):
    - `id` (Required string)
    - `command` (Required string)
    - `args` (Optional list(string))
    - `env` (Optional map(string), Sensitive) — persisted only if `allow_sensitive_env_writes = true`.
  - `profiles` (block, multiple): a profile overlay with the same subset of keys as top‑level config (e.g., model, provider, sandbox, etc.), plus `name` (Required string). Values here are written under `profiles.<name>`.

CRUD Semantics

- Create:
  - Resolve target `path` from `scope` and environment (`CODEX_HOME`, default `~/.codex`).
  - Read existing file (if present). If `merge_strategy = fail_on_unknown` and unknown keys are present: error.
  - Merge desired state with existing content per strategy.
  - Validate against schema (strict if `validate_strict = true`).
  - Write atomically with optional backup; apply `file_mode` on create.
  - Set `id`, `resolved_path` in state.

- Read:
  - Parse current TOML; decode into internal model.
  - Populate state attributes from file; unknown keys are ignored unless `fail_on_unknown`.
  - Report drift using normalized TOML serialization for managed keys.

- Update:
  - Same flow as Create; compute diff on managed keys; write atomically.

- Delete:
  - If `keep_file_on_destroy = true` and `merge_strategy = preserve_unknown`: remove only managed keys; retain the rest.
  - If `replace_all` and `keep_file_on_destroy = true`: leave file intact but clear ownership marker.
  - If `keep_file_on_destroy = false`: remove file if it has a management marker and contains only managed keys; otherwise error unless forced via a future flag.

Filesystem Scope & Precedence

- `scope = "home"`: `path = $CODEX_HOME/config.toml` (default `$HOME/.codex/config.toml`).
- `scope = "project"`: `path = <workdir>/.codex/config.toml` (provider resolves from working directory used by Terraform).
- `scope = "custom"`: `path` must be provided.
- Precedence at runtime remains: home → project → profile overlay. The resource writes only its targeted file; it does not attempt to coordinate with other files.

Safety & Secrets

- Provider API secrets: configured via `env_key`. The resource never reads or writes secret values.
- MCP server `env`: treated as Sensitive. Persisted to disk only when `allow_sensitive_env_writes = true`; otherwise ignored with a warning.
- Backups and atomic writes reduce risk of corruption. Default `0600` permissions minimize exposure.

Notation & Ownership Markers

- Write a header comment at the top of the file to record ownership when `replace_all` is used:
  - `# Managed by terraform-provider-agentsmith: <resource addr>`.
- For `preserve_unknown`, maintain an internal hidden marker table:
  - `[__agentsmith] managed = true` and a set of managed keys to power safe partial deletion and drift checks.

Implementation Steps

1) Resource scaffolding
- Add `internal/provider/codex_config_resource.go` implementing `Resource`, `Schema`, `Configure`, `Create`, `Read`, `Update`, `Delete`, and `ImportState`.
- Add registration to `Resources()` in `internal/provider/provider.go`.

2) File I/O helpers
- Create `internal/codex/configio` package to centralize: resolve paths, read/validate TOML, merge models, atomic write with backup, permission handling, and ownership markers.
- Reuse TOML models from the data source; ensure a single source of truth for structs and (de)serialization.

3) Merge engine
- Implement per‑field merge for scalars and shallow map merges for dynamic tables. Respect `merge_strategy`.
- Normalize TOML for diffing to ensure stable updates irrespective of map iteration order.

4) Validation & diagnostics
- Strict schema validation (unknown keys, type mismatches) using go‑toml/v2 decode errors wrapped with path and file context.
- Cross‑field checks (e.g., if `model_provider` references a non‑existent provider, emit warning).

5) Tests
- Unit tests for:
  - Creating a new file (home/project) with defaults and permissions.
  - `preserve_unknown` merges (retain comments not guaranteed, retain unknown keys).
  - `replace_all` behavior (removes undeclared keys) and ownership markers.
  - Sensitive `mcp_servers.env` honoring `allow_sensitive_env_writes`.
  - Import existing file into resource state.
- Acceptance tests:
  - Exercise `resource "agentsmith_codex_config"` across `scope` variants using temp dirs for `$HOME`/`$CODEX_HOME`.
  - Round‑trip with `data "agentsmith_codex"` to assert effective config.

6) Examples & docs
- Add `examples/agentsmith_codex_config/home.tf` and `examples/agentsmith_codex_config/project.tf` with realistic providers, MCP servers, and a profile.
- Document safety flags (`merge_strategy`, `allow_sensitive_env_writes`, `keep_file_on_destroy`) and precedence.
- Update README with a “Codex Config Writer” section, linking to `codex/codex_config_helper.md`.

Trade‑offs & Future Work

- MVP intentionally avoids multiple resources writing to the same file to prevent race conditions and last‑writer‑wins surprises. If demand arises, we can introduce a single in‑provider file aggregator or split resources (e.g., `agentsmith_codex_model_provider`, `agentsmith_codex_profile`) with an internal lock/compose mechanism.
- Comment preservation is not guaranteed by TOML encoders; we prioritize correctness and safety over round‑tripping comments. Future work could explore a comment‑aware parser or block markers.
- Additional env‑derived toggles can be recognized as read‑time signals in the data source; the resource will continue to write only file‑backed settings.
