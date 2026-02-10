# Terraform Provider: Agentsmith

The Agentsmith provider allows for the declarative management of configuration for various AI coding assistants, including Claude, Codex, and Gemini. It operates on local files within a configurable working directory.

## Usage

```hcl
terraform {
  required_providers {
    agentsmith = {
      source  = "nkbud/agentsmith"
      version = "~> 0.1"
    }
  }
}

provider "agentsmith" {
  # workdir = "/path/to/your/project"
}
```

## Development

### Build

```shell
go build -v ./...
```

### Test

```shell
# Unit tests
go test -v -cover -timeout=120s -parallel=10 ./...

# Acceptance tests
TF_ACC=1 go test -v -cover -timeout 120m ./...
```

### Generate Documentation

```shell
make generate
```

---

## Publishing to the Terraform Registry

This section documents everything needed to publish a Terraform provider to the [Terraform Registry](https://registry.terraform.io/). The project already contains all the required **code-level** artifacts. Below is a summary of what exists and what must be done in the **GitHub UI / external services**.

### Code-Level Requirements (Already in Place)

| Requirement | File | Purpose |
|---|---|---|
| **GoReleaser config** | `.goreleaser.yml` | Builds multi-platform binaries, creates ZIP archives, generates SHA256 checksums, signs with GPG, and attaches the registry manifest to each release. |
| **Release workflow** | `.github/workflows/release.yml` | Triggered on `v*` tags. Imports a GPG key, then runs GoReleaser to create a signed GitHub Release. |
| **Registry manifest** | `terraform-registry-manifest.json` | Tells the registry which protocol versions the provider supports (protocol `6.0`). |
| **Provider address** | `main.go` | `Address: "registry.terraform.io/nkbud/agentsmith"` — matches the registry namespace `nkbud/agentsmith`. |
| **MPL-2.0 license** | `LICENSE` | The Terraform Registry requires an MPL-2.0, MIT, or similarly OSI-approved license. |
| **Generated docs** | `docs/` | Created by `tfplugindocs`; the registry renders these automatically. |

### GitHub UI / External Steps Required

Follow these steps **in order** to enable publishing.

#### 1. Create a GPG Signing Key

The Terraform Registry verifies release integrity via GPG signatures. You need a key pair.

```shell
# Generate a new GPG key (use a real email; no passphrase is simpler for CI)
gpg --full-generate-key

# Export the ASCII-armored private key (for GitHub Secrets)
gpg --armor --export-secret-keys <KEY_ID>

# Export the public key (for the Terraform Registry)
gpg --armor --export <KEY_ID>
```

#### 2. Add GitHub Repository Secrets

Go to **Settings → Secrets and variables → Actions** in your GitHub repository and create:

| Secret Name | Value |
|---|---|
| `GPG_PRIVATE_KEY` | The full ASCII-armored private key block (from `gpg --armor --export-secret-keys`) |
| `PASSPHRASE` | The passphrase for the GPG key (leave empty string if none) |

These are consumed by the release workflow (`.github/workflows/release.yml`).

#### 3. Sign In to the Terraform Registry

1. Go to [registry.terraform.io](https://registry.terraform.io/).
2. Click **Sign In** → **Sign in with GitHub**.
3. Authorize the Terraform Registry GitHub App for the **nkbud** account (or the organization that owns the repository).

#### 4. Publish the Provider

1. On the Terraform Registry, click **Publish** → **Provider**.
2. Select the repository `nkbud/terraform-provider-agentsmith`.
3. Paste your **GPG public key** (from `gpg --armor --export <KEY_ID>`) when prompted. The registry uses this to verify release signatures.
4. Accept the terms of service and click **Publish**.

The registry will now watch for GitHub Releases that contain the expected artifacts (binaries, checksums, signature, manifest).

#### 5. Create a Release

Tag and push to trigger the release workflow:

```shell
git tag v0.1.0
git push origin v0.1.0
```

The release workflow will:
1. Import the GPG key from the `GPG_PRIVATE_KEY` secret.
2. Run GoReleaser, which builds binaries for all platforms, creates archives, generates `SHA256SUMS`, signs the checksums, and creates a GitHub Release with all artifacts attached.
3. The Terraform Registry detects the new release and makes it available within a few minutes.

#### 6. Verify

After the release workflow completes:
- Check the [GitHub Releases page](https://github.com/nkbud/terraform-provider-agentsmith/releases) to confirm all artifacts are present (ZIP archives, SHA256SUMS, SHA256SUMS.sig, and the manifest JSON).
- Check the [Terraform Registry](https://registry.terraform.io/providers/nkbud/agentsmith/latest) to confirm the provider version is listed.
- Test installation: `terraform init` in a configuration that references `nkbud/agentsmith`.

### Release Artifact Checklist

A valid release should contain these files (produced by GoReleaser):

```
terraform-provider-agentsmith_<VERSION>_darwin_amd64.zip
terraform-provider-agentsmith_<VERSION>_darwin_arm64.zip
terraform-provider-agentsmith_<VERSION>_freebsd_386.zip
terraform-provider-agentsmith_<VERSION>_freebsd_amd64.zip
terraform-provider-agentsmith_<VERSION>_freebsd_arm.zip
terraform-provider-agentsmith_<VERSION>_freebsd_arm64.zip
terraform-provider-agentsmith_<VERSION>_linux_386.zip
terraform-provider-agentsmith_<VERSION>_linux_amd64.zip
terraform-provider-agentsmith_<VERSION>_linux_arm.zip
terraform-provider-agentsmith_<VERSION>_linux_arm64.zip
terraform-provider-agentsmith_<VERSION>_windows_386.zip
terraform-provider-agentsmith_<VERSION>_windows_amd64.zip
terraform-provider-agentsmith_<VERSION>_windows_arm.zip
terraform-provider-agentsmith_<VERSION>_windows_arm64.zip
terraform-provider-agentsmith_<VERSION>_SHA256SUMS
terraform-provider-agentsmith_<VERSION>_SHA256SUMS.sig
terraform-provider-agentsmith_<VERSION>_manifest.json
```

### Subsequent Releases

For future versions, simply tag and push:

```shell
git tag v0.2.0
git push origin v0.2.0
```

The registry picks up new releases automatically — no additional UI steps are required after the initial setup.