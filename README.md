# ci-cd-example-pipeline

A GitHub Actions pipeline template you can drop into any project. Copy the four workflows, adjust two lines, and you get lint, tests, Docker builds, semantic versioning, and automated GitHub Releases.

The Go API in this repo is just the example application the pipeline runs on: you can replace it with anything.

## Workflows

```
push / PR
    └── ci.yml
        ├── lint       checks code style and common errors
        ├── test       runs tests and measures coverage
        └── comment    posts a summary on the PR with coverage %

merge to main
    └── cd.yml
        ├── build.yml    compiles and pushes a Docker image to GHCR
        └── release.yml  creates a semver tag, CHANGELOG, and GitHub Release
```

### Versioning (Conventional Commits)

| Commit prefix | Bump |
|---------------|------|
| `feat:` | minor — `v1.0.0 → v1.1.0` |
| `fix:` | patch — `v1.1.0 → v1.1.1` |
| `feat!:` | major — `v1.x.x → v2.0.0` |
| `chore:` / `docs:` / `refactor:` | patch (default) |

## Adopting the pipeline

### 1. Copy the workflows

```bash
cp -r .github/workflows/ your-project/.github/workflows/
```

### 2. Swap the language-specific steps

Open `ci.yml` and replace the `lint` and `test` steps with your language's tooling:

| Language | Lint | Test |
|----------|------|------|
| Go | `golangci/golangci-lint-action` | `go test -race -coverprofile=coverage.out ./...` |
| Node.js | `npm run lint` | `npm test -- --coverage` |
| Python | `ruff check .` | `pytest --cov=. --cov-report=term` |
| Rust | `cargo clippy` | `cargo test` |
| Java | `checkstyle` | `mvn test` |

The `build.yml` and `release.yml` workflows are language-agnostic and work as-is for any Dockerized project.

### 3. Update the Docker image name in `build.yml`

The default tag is `ghcr.io/${{ github.repository }}`, which resolves automatically from your repo name, so no change needed in most cases.

### 4. Configure GitHub

**Branch protection** (Settings → Branches → Add rule for `main`):
- ✅ Require a pull request before merging
- ✅ Require status checks: `lint`, `test`
- ✅ Require branches to be up to date before merging
- ✅ Do not allow bypassing

## Using this as a template

If this repo is marked as a GitHub Template (Settings → General → ✅ Template repository), click **Use this template** to create a clean repo with all the files and no commit history.

## Example application

The `handler/` and `main.go` files are a minimal Go HTTP API used to demonstrate the pipeline:

| Method | Route | Response |
|--------|-------|----------|
| GET | `/health` | `{"status":"ok"}` |
| POST | `/words` | `{"words":["hello","world"]}` |

Replace or delete them when adopting the pipeline for your own project.

## License

This project is licensed under the [MIT LICENSE](LICENSE).
