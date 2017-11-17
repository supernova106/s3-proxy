# Description

Proxy server to retrieve and serve object from S3 bucket. Currently supports html and jpeg/jpg types, though this can be further customized.

## Setup

- Grab the latest binary from the releases page.

- On macOS you can install or upgrade to the latest released version with Homebrew:

```shell
brew install dep
brew upgrade dep
```

- Please refer to: [Go Dep](https://github.com/golang/dep)
- `dep ensure` to install dependencies
- `cp .env.sample .env` and update the necessary parameters.
- the app supports default AWS authentication methods, using AWS credential, environment variables as well as AWS profile and IAM role.
- for Production deployment, set `GIN_MODE=release`

## Release

- get `go releaser` for MacOS

```shell
brew install goreleaser/tap/goreleaser
```

- Next, you need to export a GITHUB_TOKEN environment variable, which should contain a GitHub token with the repo scope selected. It will be used to deploy releases to your GitHub repository. Create a token [here](https://github.com/settings/tokens/new).
- `export GITHUB_TOKEN=YOUR_TOKEN`

- GoReleaser uses the latest Git tag of your repository. Create a tag and push it to GitHub:

```shell
git tag -a v0.1.0 -m "First release"
git push origin v0.1.0
```

- Please refer to [semantic versioning](http://semver.org/) when making a tag
- If you don't want to create a tag yet, you can also create a release based on the latest commit by using the --snapshot flag.
- Now you can run GoReleaser at the root of your repository:

```shell
goreleaser
```
