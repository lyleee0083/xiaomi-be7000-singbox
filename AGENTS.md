# xiaomi-be7000-singbox Repository Instructions

## Scope

- Confirm the repository root, target branch, worktree status, and applicable instructions before editing. Keep changes within the authorized task and preserve existing work.
- Track build source, build configuration and scripts, test and validation code, and essential resources and documentation. Maintained configuration and rule data are source. Keep temporary files in the task directory. Deliver build packages and validation outputs through the project artifact channel when authorized.

## GitHub

- Prefer the verified official `gh`. On Windows, resolve the executable without running it and verify its contents with trusted installer evidence or an official release package. Use the verified absolute path and SHA-256. Use `gh api` for operations without a dedicated command. An authorized connector may perform the same read-only operation if the CLI is unavailable or unverified and repair is outside scope. Respect access denials.
- Every GitHub repository must have an active `Protect main` ruleset for its default branch. Require pull requests and block branch deletion and force pushes. Require status checks that actually run on pull requests targeting the protected branch.
- Enable and verify `delete_branch_on_merge: true` when creating or configuring a repository. Before merging a pull request, verify this setting in its source repository, including a fork, and report any limit on automatic branch deletion.
- Before writing to a shared branch, verify the current head and check for competing writers. Follow the repository rules for reviews and checks, and verify the remote commit and source branch cleanup after merging.
- Maintain one H1 in the primary README, with the project name and purpose. Where applicable, keep the H2 order Features, Quick Start, Configuration, Build and Test, Repository Structure, Releases, and License. Use the established language, fenced code blocks with language identifiers, and relative links for repository files.
- Name and link the upstream repository in the README, and record the source tag or commit when known. Credit imported components in the README or linked third-party notices and preserve applicable license and copyright notices.

## Build and Verification

- See [README.md](README.md) for the SagerNet/sing-box source and imported revision. Check the source commit, submodules, and licenses when updating upstream code. Preserve project-specific changes.
- Use `go.mod`, `Makefile`, and `.github/workflows/` to determine the current Go version, build targets, build tags, and test entry points. Verify the affected platforms and keep resources embedded by supported builds.
- Check Linux ARM64 compatibility for BE7000 changes. Report local or CI builds separately from router acceptance. Keep releases, container publication, and deployment within task authorization.

## Versioned deliverables

- Use [release/be7000.json](release/be7000.json) for the BE7000 release version and upstream provenance. For changes affecting the versioned executable, advance the version and rebuild the Linux ARM64 musl executable. Verify that the direct filename and `sing-box version --name` report the same version.
- Deliver the executable directly with `SHA256SUMS.txt`. Publish the same verified build output files with identical filenames and SHA-256 hashes. Keep upstream provenance in build metadata.
- Verify that the Release tag points to the approved build's source commit and that all published assets match the build delivery and checksum manifest.
