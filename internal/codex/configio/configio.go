package configio

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	toml "github.com/pelletier/go-toml/v2"
)

// KnownTopLevelKeys is the set of recognized TOML keys for Codex config.
var KnownTopLevelKeys = map[string]struct{}{
	"model": {}, "model_provider": {}, "model_context_window": {}, "model_max_output_tokens": {},
	"approval_policy": {}, "sandbox_mode": {}, "sandbox_workspace_write": {}, "notify": {},
	"history": {}, "file_opener": {}, "hide_agent_reasoning": {}, "show_raw_agent_reasoning": {},
	"model_reasoning_effort": {}, "model_reasoning_summary": {}, "model_verbosity": {},
	"model_supports_reasoning_summaries": {}, "project_doc_max_bytes": {}, "tui": {},
	"shell_environment_policy": {}, "model_providers": {}, "mcp_servers": {}, "profiles": {},
	// Allow profile pointer if present though resource does not manage it
	"profile": {},
	// Agentsmith internal table
	"__agentsmith": {},
}

// ResolvePath returns the absolute path to the target config.toml for a given scope.
// scope: "home" | "project" | "custom". When scope==custom, customPath must be provided.
// workdir is used for project scope.
func ResolvePath(scope, workdir, customPath string) (string, error) {
	switch scope {
	case "home":
		codexHome := os.Getenv("CODEX_HOME")
		if codexHome == "" {
			home, err := os.UserHomeDir()
			if err != nil {
				return "", fmt.Errorf("get home dir: %w", err)
			}
			codexHome = filepath.Join(home, ".codex")
		}
		return filepath.Join(codexHome, "config.toml"), nil
	case "project":
		if workdir == "" {
			return "", fmt.Errorf("workdir not configured for project scope")
		}
		return filepath.Join(workdir, ".codex", "config.toml"), nil
	case "custom":
		if strings.TrimSpace(customPath) == "" {
			return "", fmt.Errorf("custom scope requires 'path'")
		}
		if !filepath.IsAbs(customPath) {
			// Resolve to absolute from cwd
			abs, err := filepath.Abs(customPath)
			if err != nil {
				return "", err
			}
			return abs, nil
		}
		return customPath, nil
	default:
		return "", fmt.Errorf("invalid scope: %s", scope)
	}
}

// ReadTOMLMap reads a TOML file into a generic map structure. If the file does not exist, returns an empty map and no error.
func ReadTOMLMap(path string) (map[string]any, error) {
	if path == "" {
		return map[string]any{}, nil
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return map[string]any{}, nil
	}
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var out map[string]any
	if len(bytes.TrimSpace(b)) == 0 {
		return map[string]any{}, nil
	}
	if err := toml.Unmarshal(b, &out); err != nil {
		return nil, err
	}
	if out == nil {
		out = map[string]any{}
	}
	return out, nil
}

// ValidateUnknownKeys checks if the map contains keys not in KnownTopLevelKeys.
func ValidateUnknownKeys(m map[string]any) (unknown []string) {
	for k := range m {
		if _, ok := KnownTopLevelKeys[k]; !ok {
			unknown = append(unknown, k)
		}
	}
	return
}

// MergePreserveUnknown overlays desired onto existing for keys present in desired only, preserving all other keys.
func MergePreserveUnknown(existing, desired map[string]any) map[string]any {
	out := map[string]any{}
	for k, v := range existing {
		out[k] = v
	}
	for k, v := range desired {
		out[k] = v
	}
	// ensure internal marker table exists
	ensureAgentsmithMarker(out)
	return out
}

func ensureAgentsmithMarker(m map[string]any) {
	t, _ := m["__agentsmith"].(map[string]any)
	if t == nil {
		t = map[string]any{}
	}
	t["managed"] = true
	t["updated_at"] = time.Now().UTC().Format(time.RFC3339)
	m["__agentsmith"] = t
}

// MarshalTOML renders a map into TOML bytes.
func MarshalTOML(m map[string]any) ([]byte, error) {
	buf := new(bytes.Buffer)
	enc := toml.NewEncoder(buf)
	if err := enc.Encode(m); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// AtomicWrite writes bytes to a temp file and renames it to the path. If backup is true and the file exists,
// creates a .bak copy first. Mode string is parsed in base-8 (e.g., "0600").
func AtomicWrite(path string, data []byte, mode string, backup bool, createDirs bool, headerComment string) error {
	dir := filepath.Dir(path)
	if createDirs {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return fmt.Errorf("mkdir: %w", err)
		}
	}
	// Optional backup
	if backup {
		if _, err := os.Stat(path); err == nil {
			bak := path + ".bak"
			if err := copyFile(path, bak); err != nil {
				return fmt.Errorf("backup write: %w", err)
			}
		}
	}
	// Prepend header comment if requested
	if strings.TrimSpace(headerComment) != "" {
		data = append([]byte(headerComment+"\n"), data...)
	}

	tmp := path + ".tmp"
	f, err := os.OpenFile(tmp, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o600)
	if err != nil {
		return err
	}
	if _, err := f.Write(data); err != nil {
		_ = f.Close()
		return err
	}
	if err := f.Sync(); err != nil {
		_ = f.Close()
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	if err := os.Rename(tmp, path); err != nil {
		return err
	}
	// Apply mode if provided
	if strings.TrimSpace(mode) != "" {
		// strip potential 0x or 0 prefix by parsing base 8 explicitly
		var parsed os.FileMode
		// Accept form like "0600" or "600"
		if strings.HasPrefix(mode, "0") {
			mode = strings.TrimLeft(mode, "0")
		}
		if mode == "" {
			mode = "600"
		}
		// naive parse
		var v uint64
		for _, r := range mode {
			if r < '0' || r > '7' {
				return fmt.Errorf("invalid file_mode: %s", mode)
			}
			v = v*8 + uint64(r-'0')
		}
		parsed = os.FileMode(v)
		if err := os.Chmod(path, parsed); err != nil {
			return err
		}
	}
	return nil
}

func copyFile(src, dst string) error {
	in, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	return os.WriteFile(dst, in, 0o600)
}

// SHA256Hex computes SHA256 hex of input string.
func SHA256Hex(s string) string {
	sum := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sum[:])
}
