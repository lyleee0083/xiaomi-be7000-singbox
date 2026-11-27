# xiaomi-be7000-singbox Repository Instructions

## Scope

- Confirm the repository root, target branch, worktree status, and applicable instructions before editing. Keep changes within the authorized task and preserve existing work.
- Track build source, build configuration and scripts, test and validation code, and essential resources and documentation. Maintained configuration and rule data are source. Keep temporary files in the task directory. Deliver build packages and validation outputs through the project artifact channel when authorized.

## GitHub

- Make fixes directly in the maintained source code, configuration, or build logic that owns the behavior, including in forks. Do not add patch files, build-time source rewrites, or redundant wrappers as substitutes for direct fixes.
- Start new upstream-based work and upstream updates from the latest commit on the applicable upstream branch. Verify the branch tip remotely and record the upstream repository, branch, and full commit SHA. Use the upstream default branch unless the task specifies another branch or revision.
- Run all repository builds and compilation exclusively in GitHub Actions; do not build or compile locally. Deliver the verified artifacts from the successful workflow run.
- Prefer the verified official `gh`. On Windows, resolve the executable without running it and verify its contents with trusted installer evidence or an official release package. Use the verified absolute path and SHA-256. Use `gh api` for operations without a dedicated command. An authorized connector may perform the same read-only operation if the CLI is unavailable or unverified and repair is outside scope. Respect access denials.
- Every GitHub repository must have an active `Protect main` ruleset for its default branch. Require pull requests and block branch deletion and force pushes. Require status checks that actually run on pull requests targeting the protected branch.
- Enable and verify `delete_branch_on_merge: true` when creating or configuring a repository. Before merging a pull request, verify this setting in its source repository, including a fork, and report any limit on automatic branch deletion.
- Before writing to a shared branch, verify the current head and check for competing writers. Follow the repository rules for reviews and checks, and verify the remote commit and source branch cleanup after merging.
- Use Simplified Chinese for README and GitHub Release titles and bodies, including generated notes, unless directed otherwise. Keep official names and identifiers.
- Start README with one project-name H1 and an overview, then applicable H2s in order 功能 → 快速开始 → 配置 → 构建与测试 → 仓库结构 → 版本发布 → 许可证.
- Use `<project> <version>` for Release titles. Follow the summary with applicable H2s in order 主要变更 → 下载与校验 → 安装与升级 → 验证状态.
- Use tagged code fences, relative README and tag-pinned Release file links. Keep upstream links, known refs, credits, licenses and copyright notices.

## Build and Verification

- See [README.md](README.md) for the SagerNet/sing-box source and imported revision. Check the source commit, submodules, and licenses when updating upstream code. Preserve project-specific changes.
- Use `go.mod`, [release/be7000.json](release/be7000.json), and `.github/workflows/` to determine the current Go version, build targets, build tags, and test entry points. Run all builds, compilation, tests, and platform verification exclusively in GitHub Actions and keep resources embedded by supported builds.
- Verify Linux ARM64 compatibility for BE7000 changes through the GitHub Actions workflow (`.github/workflows/be7000.yml`). Report GitHub Actions build and verification results separately from router acceptance. Keep releases, container publication, and deployment within task authorization.

## Versioned deliverables

- Use [release/be7000.json](release/be7000.json) for the BE7000 release version and upstream provenance. For changes affecting the versioned executable, advance the version and rebuild the Linux ARM64 musl executable. Verify that the direct filename and `sing-box version --name` report the same version. Keep upstream provenance in build metadata.

## Release asset integrity

- Deliver the executable directly with `SHA256SUMS.txt`. Publish the same verified build output files to GitHub Releases with identical filenames and SHA-256 hashes.
- Match the Release tag to the verified build commit. Check every published asset against the delivery files and checksum manifest. Report missing source, build, or Release evidence.
